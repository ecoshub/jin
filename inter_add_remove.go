package jsoninterpreter

import (
	"fmt"
	"strconv"
	"errors"
)

func AddKeyValue(json []byte, key, value string, path ... string) ([]byte, error){
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _,v := range chars {
		isJsonChar[v] = true
	}
	if len(path) == 0 {
		if json[0] == 123 {
			level := 0
			inQuote := false
			for i := 0 ; i < len(json) ; i ++ {
				curr := json[i]
				if !isJsonChar[curr]{
					continue
				}
				if curr == 34 {
					if json[i - 1] == 92 {
						continue
					}
					inQuote = !inQuote
					continue
				}
				if inQuote {
					continue
				}else{
					if curr == 91 || curr == 123 {
						level++
					}
					if curr == 93 || curr == 125 {
						level--
						if level == 0 {
							json = replace(json, []byte(fmt.Sprintf(`,"%v":"%v"`, key, value)),i,i)
							return json, nil
						}
						continue
					}
					continue
				}
				continue
			}
		}else{
			return json, errors.New("Error: last path must be pointed at an object not to an array")
		}
	}
	path = append(path, key)
	offset := 0
	currentPath := path[0]
	for space(json[offset]) {
		offset++
	}
	lastOffset := 0
	braceType := json[offset]
	for k := 0 ; k < len(path) ; k ++ {
		if braceType == 91 {
			arrayIndex, err := strconv.Atoi(currentPath)
			if err != nil {
				return json, errors.New("Error: Index Expected, got string.")
			}
			done := false
			if arrayIndex == 0 {
				offset++
				for i := offset; i < len(json) ; i ++ {
					curr := json[i]
					if curr == 123 {
						braceType = curr
						if k != len(path) - 1{
							currentPath = path[k + 1]
							lastOffset = offset
						}else{
							return json, errors.New("Error: last path must be pointed at an object not to an array")
						}
						offset = i
						done = true
						break
					}
					if curr == 91 {
						braceType = curr
						if k != len(path) - 1{
							currentPath = path[k + 1]
							lastOffset = offset
						}else{
							return json, errors.New("Error: last path must be pointed at an object nto ot an array")
						}
						offset = i + 1
						done = true
						break
					}
					if !space(curr){
						done = true
						break
					}
				}
			}else{
				level := 0
				inQuote := false
				found := false
				indexCount := 0
				isJsonChar[58] = false
				for i := offset ; i < len(json) ; i ++ {
					curr := json[i]
					if !isJsonChar[curr]{
						continue
					}
					if curr == 34 {
						if json[i - 1] == 92 {
							continue
						}
						inQuote = !inQuote
						continue
					}
					if inQuote {
						continue
					}else{
						if curr == 91 || curr == 123{
							if found {
								level++
								braceType = curr
								currentPath = path[k + 1]
								found = false
								done = true
								break
							}
							level++
							continue
						}
						if curr == 93 || curr == 125 {
							level--
							if level < 1 {
								done = false
							}
							continue
						}
						if !found {
							if level == 1 {
								if curr == 44 {
									indexCount++
									if indexCount == arrayIndex {
										offset = i + 1
										if k == len(path) - 1{
											return json, errors.New("Error: last path must be pointed at an object not to an array")
										}
										lastOffset = offset
										found = true
										continue
									}
									continue
								}
								continue
							}
							continue
						}
						continue
					}
				}
				isJsonChar[58] = true
			}
			if !done {
				return json, errors.New("Error: Index out of range")
			}
		}else{
			inQuote := false
			found := false
			start := 0
			end := 0
			level := k
			isJsonChar[44] = false
			for i := offset ; i < len(json) ; i ++ {
				curr := json[i]
				if !isJsonChar[curr]{
					continue
				}
				if curr == 34 {
					inQuote = !inQuote
					if found {
						continue
					}
					if level != k + 1 {
						continue
					}
					if inQuote {
						start = i + 1
						continue
					}
					end = i
					continue
				}
				if inQuote {
					continue
				}else{
					if curr == 91 {
						if found {
							braceType = curr
							currentPath = path[k + 1]
							break
						}
						level++
						continue
					}
					if curr == 123 {
						if found {
							k++
							level++
							currentPath = path[k]
							found = false
							continue
						}
						level++
						continue
					}
					if curr == 93 || curr == 125 {
						level--
						continue
					}
					if level == k + 1 {
						if curr == 58 {
							if compare(json, start, end, currentPath) {
								offset = i + 1
								found = true
								if k == len(path) - 1{
									isJsonChar[44] = true
									break
								}else{
									lastOffset = offset
									continue
								}
							}
							isJsonChar[44] = true
							isJsonChar[58] = false
							for j := i ;  j < len(json) ; j ++ {
								curr := json[j]
								if !isJsonChar[curr]{
									continue
								}
								if curr == 34 {
									if json[j - 1] == 92 {
										continue
									}
									inQuote = !inQuote
									continue
								}
								if inQuote {
									continue
								}else{
									if curr == 91 || curr == 123 {
										level++
										continue
									}
									if curr == 93 || curr == 125 {
										level--
										continue
									}
									if curr == 44 {
										if level == k + 1 {
											i = j
											break
										}
										continue
									}
									continue
								}

							}
							isJsonChar[44] = false
							isJsonChar[58] = true
							continue
						}
						continue
					}
				}
			}
			isJsonChar[44] = true
			if k == len(path) - 1 {
				if found {
					return json, errors.New("Error: key already exist")
				}
			}else{
				if !found {
					return json, errors.New("Error: key not found.")
				}
			}
		}
	}
	if lastOffset == 0 {
		return json, errors.New("Error: Something went wrong... not sure, maybe bad JSON format...")
	}
	for space(json[lastOffset]) {
		lastOffset++
	}
	if json[lastOffset] == 123 {
		level := 0
		inQuote := false
		for i := lastOffset ; i < len(json) ; i ++ {
			curr := json[i]
			if !isJsonChar[curr]{
				continue
			}
			if curr == 34 {
				if json[i - 1] == 92 {
					continue
				}
				inQuote = !inQuote
				continue
			}
			if inQuote {
				continue
			}else{
				if curr == 91 || curr == 123 {
					level++
				}
				if curr == 93 || curr == 125 {
					level--
					if level == 0 {
						// control for comma maybe needed.
						json = replace(json, []byte(fmt.Sprintf(`,"%v":"%v"`, key, value)),i,i)
						return json, nil
					}
					continue
				}
				continue
			}
			continue
		}
	}else{
		return json, errors.New("Error: Error: last path must be pointed at an object not to an array")
	}
	return json, errors.New("Error: Something went wrong... not sure, maybe bad JSON format...")
}