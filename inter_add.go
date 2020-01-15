package jsoninterpreter

import (
	"fmt"
	"strconv"
	"errors"
)

func dummy(){fmt.Println()}

func AddKeyValue(json []byte, key string, value []byte, path ... string) ([]byte, error){
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _,v := range chars {
		isJsonChar[v] = true
	}
	offset := 0
	for space(json[offset]) {
		offset++
	}

	if len(path) == 0 {
		if json[offset] == 123 {
			lenj := len(json)
			for i := offset ; i < lenj; i ++ {
				curr := json[lenj - i - 1]
				if !isJsonChar[curr]{
					continue
				}
				if !space(curr){
					if curr == 125 {
						return replace(json, []byte(`,"` + key + `":` + string(value)),lenj - i - 1,lenj - i - 1), nil
					}else{
						return json, errors.New("Error: Bad JSON format")
					}
					continue
				}
			}
		}else{
			return json, errors.New("Error: Last path must be pointed at an object.")
		}
	}
	path = append(path, key)
	currentPath := path[0]
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
						return replace(json, []byte(`,"` + key + `":` + string(value)),i,i), nil
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
	return json, errors.New("Error: Something went wrong... not sure, maybe bad JSON format...")
}

func AddKeyValueString(json []byte, key, value string, path ... string) ([]byte, error){
	if value[0] != 34 && value[len(value) - 1] != 34 {
		value = `"` + value + `"`
	}
	return AddKeyValue(json, key, []byte(value), path...)
}

func AddKeyValueInt(json []byte, key string, value int, path ... string) ([]byte, error){
	return AddKeyValue(json, key, []byte(strconv.Itoa(value)), path...)
}

func AddKeyValueFloat(json []byte, key string, value float64, path ... string) ([]byte, error){
	return AddKeyValue(json, key, []byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

func AddKeyValueBool(json []byte, key string, value bool, path ... string) ([]byte, error){
	if value {
		return AddKeyValue(json, key, []byte("true"), path...)
	}
	return AddKeyValue(json, key, []byte("false"), path...)
}

func AddValue(json []byte, value []byte, path ... string) ([]byte, error){
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _,v := range chars {
		isJsonChar[v] = true
	}
	offset := 0
	for space(json[offset]) {
		offset++
	}
	if len(path) == 0 {
		if json[offset] == 91 {
			lenj := len(json)
			for i := offset ; i < lenj; i ++ {
				curr := json[lenj - i - 1]
				if !isJsonChar[curr]{
					continue
				}
				if !space(curr){
					if curr == 93 {
						return replace(json, []byte("," + string(value)),lenj - i - 1,lenj - i - 1), nil
					}else{
						return json, errors.New("Error: Bad JSON format")
					}
					continue
				}
			}
		}else{
			return json, errors.New("Error: Last path must be pointed at an array.")
		}
	}
	currentPath := path[0]
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
						}
						offset = i
						done = true
						break
					}
					if curr == 91 {
						braceType = curr
						if k != len(path) - 1{
							currentPath = path[k + 1]
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
											done = true
											break
										}
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
			if !found {
				return json, errors.New("Error: key not found.")
			}
		}
	}
	if offset == 0 {
		return json, errors.New("Error: Something went wrong... not sure, maybe bad JSON format...")
	}
	for space(json[offset]) {
		offset++
	}
	if json[offset] == 91 {
		level := 0
		inQuote := false
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
				if curr == 91 || curr == 123 {
					level++
				}
				if curr == 93 || curr == 125 {
					level--
					if level == 0 {
						// control for comma maybe needed.
						return replace(json, []byte("," + string(value)),i,i), nil
					}
					continue
				}
				continue
			}
			continue
		}
	}else{
		return json, errors.New("Error: Last path must be pointed at an array.")
	}
	return json, errors.New("Error: Something went wrong... not sure, maybe bad JSON format...")
}

func AddValueString(json []byte, value string, path ... string) ([]byte, error){
	if value[0] != 34 && value[len(value) - 1] != 34 {
		return AddValue(json, []byte(`"` + value + `"`), path...)
	}
	return AddValue(json, []byte(value), path...)
}

func AddValueInt(json []byte, value int, path ... string) ([]byte, error){
	return AddValue(json, []byte(strconv.Itoa(value)), path...)
}

func AddValueFloat(json []byte, value float64, path ... string) ([]byte, error){
	return AddValue(json, []byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

func AddValueBool(json []byte, value bool, path ... string) ([]byte, error){
	if value {
		return AddValue(json, []byte("true"), path...)
	}
	return AddValue(json, []byte("false"), path...)
}

func InsertValue(json []byte, value []byte, index int, path ... string) ([]byte, error){
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _,v := range chars {
		isJsonChar[v] = true
	}
	offset := 0
	for space(json[offset]) {
		offset++
	}
	if len(path) == 0 {
		if json[offset] == 91 {
			done := false
			if index == 0 {
				return replace(json, []byte(string(value) + ","),offset + 1,offset + 1), nil
			}else{
				level := 0
				inQuote := false
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
						if level == 1 {
							if curr == 44 {
								indexCount++
								if indexCount == index {
									offset = i + 1
									return replace(json, []byte("," + string(value)),i,i), nil
								}
								continue
							}
							continue
						}
						continue
					}
				}
				isJsonChar[58] = true
				if !done {
					return json, errors.New("Error: Index out of range")
				}
			}
		}else{
			return json, errors.New("Error: Last path must be pointed at an array.")
		}
	}
	currentPath := path[0]
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
						}
						offset = i
						done = true
						break
					}
					if curr == 91 {
						braceType = curr
						if k != len(path) - 1{
							currentPath = path[k + 1]
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
											done = true
											break
										}
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
			if !found {
				return json, errors.New("Error: key not found.")
			}
		}
	}
	if offset == 0 {
		return json, errors.New("Error: Something went wrong... not sure, maybe bad JSON format...")
	}
	for space(json[offset]) {
		offset++
	}
	if json[offset] == 91 {
		if index == 0 {
			return replace(json, []byte(string(value) + ","),offset + 1,offset + 1), nil
		}else{
			done := false
			level := 0
			inQuote := false
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
						level++
						continue
					}
					if curr == 93 || curr == 125 {
						level--
						if level < 1 {
							return json, errors.New("Error: Index out of range")
						}
						continue
					}
					if level == 1 {
						if curr == 44 {
							indexCount++
							if indexCount == index {
								offset = i + 1
								return replace(json, []byte("," + string(value)),i,i), nil
							}
							continue
						}
						continue
					}
					continue
				}
			}
			// isJsonChar[58] = true
			if !done {
				return json, errors.New("Error: Index out of range")
			}
		}
	}else{
		return json, errors.New("Error: Last path must be pointed at an array.")
	}
	return json, errors.New("Error: Something went wrong... not sure, maybe bad JSON format...")
}