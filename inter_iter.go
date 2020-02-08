package jin

// import "fmt"
// IterateArray is a callback function that can iterate any array and return value as byte slice.
// It stripes quotation marks from string values befour return.
// Path value can be left blank for access main JSON.
func IterateArray(json []byte, callback func([]byte) bool, path ...string) error {
	var start int
	var err error
	if len(path) == 0 {
		for space(json[start]) {
			if start > len(json)-2 {
				return badJSONError(start)
			}
			start++
			continue
		}
	} else {
		_, start, _, err = core(json, true, path...)
		if err != nil {
			return err
		}
	}
	chars := []byte{34, 44, 91, 93, 123, 125}
	isJSONChar := make([]bool, 256)
	for _, v := range chars {
		isJSONChar[v] = true
	}
	if json[start] != 91 {
		return arrayExpectedError()
	}
	start++
	inQuote := false
	level := 0
	for i := start; i < len(json); i++ {
		curr := json[i]
		if !isJSONChar[curr] {
			continue
		}
		if curr == 34 {
			if inQuote {
				for k := i - 1; k > 0; k-- {
					if json[k] == 92 {
						continue
					}
					if (i-k)%2 != 0 {
						inQuote = !inQuote
						goto cont
					}
					break
				}
			} else {
				inQuote = !inQuote
				continue
			}
		cont:
			continue
		}
		if inQuote {
			continue
		} else {
			if curr == 91 || curr == 123 {
				level++
				continue
			}
			if curr == 93 || curr == 125 {
				if level < 0 {
					return nil
				}
				if curr == 93 {
					if level == 0 {
						callback(cleanValue(json[start:i]))
						return nil
					}
				}
				level--
				continue
			}
			if level == 0 {
				if curr == 44 {
					if !callback(cleanValue(json[start:i])) {
						return nil
					}
					start = i + 1
					continue
				}
			}
		}
	}
	return badJSONError(start)
}

// IterateKeyValue is a callback function that can iterate any object and return key-value pair as byte slices.
// It stripes quotation marks from string values befour return.
// Path value can be left blank for access main JSON.
func IterateKeyValue(json []byte, callback func([]byte, []byte) bool, path ...string) error {
	var start int
	var err error
	var end int
	if len(path) == 0 {
		for space(json[start]) {
			if start > len(json)-1 {
				return badJSONError(start)
			}
			start++
			continue
		}
	} else {
		_, start, _, err = core(json, true, path...)
		if err != nil {
			return err
		}
	}
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJSONChar := make([]bool, 256)
	for _, v := range chars {
		isJSONChar[v] = true
	}
	if json[start] == 123 {
		keyStart := 0
		keyEnd := 0
		inQuote := false
		level := 0
		var key []byte
		for i := start + 1; i < len(json); i++ {
			curr := json[i]
			if !isJSONChar[curr] {
				continue
			}
			if curr == 34 {
				if inQuote {
					for n := i - 1; n > -1; n-- {
						if json[n] != 92 {
							if (i-n)%2 != 0 {
								inQuote = !inQuote
								break
							} else {
								goto cont
							}
						}
						continue
					}
				} else {
					inQuote = !inQuote
				}
				if inQuote {
					keyStart = i
					continue
				}
				keyEnd = i
			cont:
				continue
			}
			if inQuote {
				continue
			} else {
				if curr == 91 || curr == 123 {
					level++
					continue
				}
				if curr == 93 || curr == 125 {
					if level < 1 {
						return nil
					}
					if curr == 125 {
						if level == 0 {
							for j := start; j < i; j++ {
								if !space(json[j]) {
									start = j
									break
								}
							}
							for j := i - 1; j > start; j-- {
								if !space(json[j]) {
									end = j
									break
								}
							}
							if json[start] == 34 && json[end] == 34 {
								callback(key, json[start+1:end])
								return nil
							}
							callback(key, json[start:end+1])
							return nil
						}
					}
					level--
					continue
				}
				if level == 0 {
					if curr == 44 {
						end = i - 1
						for j := start; j < i; j++ {
							if !space(json[j]) {
								start = j
								break
							}
						}
						for j := i - 1; j > start; j-- {
							if !space(json[j]) {
								end = j
								break
							}
						}
						if json[start] == 34 && json[end] == 34 {
							if !callback(key, json[start+1:end]) {
								return nil
							}
						} else {
							if !callback(key, json[start:end+1]) {
								return nil
							}
						}
						continue
					}
					if curr == 58 {
						key = json[keyStart+1 : keyEnd]
						start = i + 1
						continue
					}
				}
			}
		}
		return nil
	}
	return objectExpectedError()
}
