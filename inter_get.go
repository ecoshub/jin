package jin

import (
	"strconv"
)

// Get returns the value that path has pointed.
// It stripes quotation marks from string values.
// Path can point anything, a key-value pair, a value, an array, an object.
// Path variable can not be null,
// otherwise it will provide an error message.
func Get(json []byte, path ...string) ([]byte, error) {
	if len(path) == 0 {
		return json, nil
	}
	_, start, end, err := core(json, false, path...)
	if err != nil {
		return nil, err
	}
	return json[start:end], err
}

// GetString is a variation of Get() func.
// GetString returns the value that path has pointed as string.
func GetString(json []byte, path ...string) (string, error) {
	val, err := Get(json, path...)
	if err != nil {
		return "", err
	}
	return string(val), err
}

// GetInt is a variation of Get() func.
// GetInt returns the value that path has pointed as integer.
// returns an error message if the value to be returned cannot be converted to an integer
func GetInt(json []byte, path ...string) (int, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return -1, err
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return -1, errIntegerParse(val)
	}
	return intVal, nil
}

// GetFloat is a variation of Get() func.
// GetFloat returns the value that path has pointed as float.
// returns an error message if the value to be returned cannot be converted to an float
func GetFloat(json []byte, path ...string) (float64, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return -1, err
	}
	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return -1, errFloatParse(val)
	}
	return floatVal, nil
}

// GetBool is a variation of Get() func.
// GetBool returns the value that path has pointed as boolean.
// returns an error message if the value to be returned cannot be converted to an boolean
func GetBool(json []byte, path ...string) (bool, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return false, err
	}
	if val == "true" {
		return true, nil
	}
	if val == "false" {
		return false, nil
	}
	return false, errBoolParse(val)
}

// GetStringArray is a variation of Get() func.
// GetStringArray returns the value that path has pointed as string slice.
// returns an error message if the value to be returned cannot be converted to an string slice.
func GetStringArray(json []byte, path ...string) ([]string, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, errStringArrayParse(val)
	}
	if val[0] == 91 && val[lena-1] == 93 {
		arr := ParseArray(val)
		if arr == nil {
			return nil, errStringArrayParse(val)
		}
		return arr, nil
	}
	return nil, errStringArrayParse(val)
}

// GetIntArray is a variation of Get() func.
// GetIntArray returns the value that path has pointed as integer slice.
// returns an error message if the value to be returned cannot be converted to an integer slice.
func GetIntArray(json []byte, path ...string) ([]int, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, errIntegerArrayParse(val)
	}
	if val[0] == 91 && val[lena-1] == 93 {
		newArray := make([]int, 0, 16)
		start := 1
		inQuote := false
		level := 0
		for i := 0; i < len(val); i++ {
			curr := val[i]
			if curr == 92 {
				i++
				continue
			}
			if curr == 34 {
				inQuote = !inQuote
				continue
			}
			if inQuote {
				continue
			} else {
				if curr == 91 || curr == 123 {
					level++
				}
				if curr == 93 || curr == 125 {
					level--
					if curr == 93 {
						if level == 0 {
							element := val[start:i]
							num, err := strconv.Atoi(cleanValueString(element))
							if err != nil {
								return nil, errIntegerParse(cleanValueString(element))
							}
							newArray = append(newArray, num)
							break
						}
					}
				}
				if level == 1 {
					if curr == 44 {
						element := val[start:i]
						num, err := strconv.Atoi(cleanValueString(element))
						if err != nil {
							return nil, errIntegerParse(cleanValueString(element))
						}
						newArray = append(newArray, num)
						start = i + 1
						continue
					}
				}
			}
		}
		return newArray, nil
	}
	return nil, errIntegerArrayParse(val)
}

// GetFloatArray is a variation of Get() func.
// GetFloatArray returns the value that path has pointed as float slice.
// returns an error message if the value to be returned cannot be converted to an float slice.
func GetFloatArray(json []byte, path ...string) ([]float64, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, errFloatArrayParse(val)
	}
	if val[0] == 91 && val[lena-1] == 93 {
		newArray := make([]float64, 0, 16)
		start := 1
		inQuote := false
		level := 0
		for i := 0; i < len(val); i++ {
			curr := val[i]
			if curr == 92 {
				i++
				continue
			}
			if curr == 34 {
				inQuote = !inQuote
				continue
			}
			if inQuote {
				continue
			} else {
				if curr == 91 || curr == 123 {
					level++
				}
				if curr == 93 || curr == 125 {
					level--
					if curr == 93 {
						if level == 0 {
							element := val[start:i]
							num, err := strconv.ParseFloat(cleanValueString(element), 64)
							if err != nil {
								return nil, errFloatParse(cleanValueString(element))
							}
							newArray = append(newArray, num)
							break
						}
					}
				}
				if level == 1 {
					if curr == 44 {
						element := val[start:i]
						num, err := strconv.ParseFloat(cleanValueString(element), 64)
						if err != nil {
							return nil, errFloatParse(cleanValueString(element))
						}
						newArray = append(newArray, num)
						start = i + 1
						continue
					}
				}
			}
		}
		return newArray, nil
	}
	return nil, errFloatArrayParse(val)
}

// GetBoolArray is a variation of Get() func.
// GetBoolArray returns the value that path has pointed as boolean slice.
// returns an error message if the value to be returned cannot be converted to an boolean slice.
func GetBoolArray(json []byte, path ...string) ([]bool, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, errBoolArrayParse(val)
	}
	if val[0] == 91 && val[lena-1] == 93 {
		newArray := make([]bool, 0, 16)
		start := 1
		inQuote := false
		level := 0
		for i := 0; i < len(val); i++ {
			curr := val[i]
			if curr == 92 {
				i++
				continue
			}
			if curr == 34 {
				inQuote = !inQuote
				continue
			}
			if inQuote {
				continue
			} else {
				if curr == 91 || curr == 123 {
					level++
				}
				if curr == 93 || curr == 125 {
					level--
					if curr == 93 {
						if level == 0 {
							element := val[start:i]
							element = cleanValueString(element)
							if element == "true" || element == "false" {
								if element == "true" {
									newArray = append(newArray, true)
								} else {
									newArray = append(newArray, false)
								}
							} else {
								return nil, errBoolParse(cleanValueString(element))
							}
							break
						}
					}
				}
				if level == 1 {
					if curr == 44 {
						element := val[start:i]
						element = cleanValueString(element)
						if element == "true" || element == "false" {
							if element == "true" {
								newArray = append(newArray, true)
							} else {
								newArray = append(newArray, false)
							}
						} else {
							return nil, errBoolParse(element)
						}
						start = i + 1
						continue
					}
				}
			}
		}
		return newArray, nil
	}
	return nil, errBoolArrayParse(val)
}

// GetKeys not tested yet
// Gets all keys that path has pointed.
func GetKeys(json []byte, path ...string) ([]string, error) {
	var keys []string
	if string(json) == "{}" {
		return nil, errEmpty()
	}
	var start int
	var err error
	if len(path) == 0 {
		for space(json[start]) {
			if start > len(json)-1 {
				return nil, errBadJSON(start)
			}
			start++
			continue
		}
	} else {
		_, start, _, err = core(json, true, path...)
		if err != nil {
			return nil, err
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
		for i := start; i < len(json); i++ {
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
					if level == 1 {
						break
					}
					level--
					continue
				}
				if curr == 58 {
					if level == 1 {
						key := json[keyStart+1 : keyEnd]
						keys = append(keys, string(key))
					}
					continue
				}
			}
		}
		return keys, nil
	}
	return nil, errObjectExpected()
}

// GetValues not tested yet
// Gets all values that path has pointed.
func GetValues(json []byte, path ...string) ([]string, error) {
	var values []string
	if string(json) == "{}" {
		return nil, errEmpty()
	}
	var start int
	var err error
	if len(path) == 0 {
		for space(json[start]) {
			if start > len(json)-1 {
				return nil, errBadJSON(start)
			}
			start++
			continue
		}
	} else {
		_, start, _, err = core(json, true, path...)
		if err != nil {
			return nil, err
		}
	}
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJSONChar := make([]bool, 256)
	for _, v := range chars {
		isJSONChar[v] = true
	}
	if json[start] == 123 {
		inQuote := false
		level := 0
		for i := start; i < len(json); i++ {
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
					if level == 1 {
						value := cleanValueString(string(json[start:i]))
						values = append(values, value)
						break
					}
					level--
					continue
				}
				if curr == 58 {
					if level == 1 {
						start = i + 1
					}
					continue
				}
				if curr == 44 {
					if level == 1 {
						value := cleanValueString(string(json[start:i]))
						values = append(values, value)
						start = i + 1
					}
				}
			}
		}
		return values, nil
	}
	return nil, errObjectExpected()
}

// GetKeysValues not tested yet
// Gets all keys and values that path has pointed.
func GetKeysValues(json []byte, path ...string) ([]string, []string, error) {
	var values []string
	var keys []string
	if string(json) == "{}" {
		return nil, nil, errEmpty()
	}
	var start int
	var err error
	if len(path) == 0 {
		for space(json[start]) {
			if start > len(json)-1 {
				return nil, nil, errBadJSON(start)
			}
			start++
			continue
		}
	} else {
		_, start, _, err = core(json, true, path...)
		if err != nil {
			return nil, nil, err
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
		for i := start; i < len(json); i++ {
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
					if level == 1 {
						value := cleanValueString(string(json[start:i]))
						values = append(values, value)
						break
					}
					level--
					continue
				}
				if curr == 58 {
					if level == 1 {
						key := json[keyStart+1 : keyEnd]
						keys = append(keys, string(key))
						start = i + 1
					}
					continue
				}
				if curr == 44 {
					if level == 1 {
						value := cleanValueString(string(json[start:i]))
						values = append(values, value)
						start = i + 1
					}
				}
			}
		}
		return keys, values, nil
	}
	return nil, nil, errObjectExpected()
}

// GetMap Gets all keys and values pair with string to string map.
func GetMap(json []byte, path ...string) (map[string]string, error) {
	if string(json) == "[]" || string(json) == "{}" {
		return nil, nil
	}
	start, end, err := getStartEnd(json, path...)
	if err != nil {
		return nil, err
	}
	mainMap := make(map[string]string)
	if json[start] == 123 && json[end] == 125 {
		err = IterateKeyValue(json, func(key, val []byte) (bool, error) {
			mainMap[string(key)] = string(val)
			return true, nil
		}, path...)
		if err != nil {
			return nil, err
		}
		return mainMap, nil
	}
	if json[start] == 91 && json[end] == 93 {
		count := 0
		err = IterateArray(json, func(val []byte) (bool, error) {
			mainMap[strconv.Itoa(count)] = string(val)
			count++
			return true, nil
		}, path...)
		if err != nil {
			return nil, err
		}
		return mainMap, nil
	}
	return nil, errBadJSON(start)
}

// GetAll not tested yet
// GetAll gets all values of path+key has pointed.
func GetAll(json []byte, keys []string, path ...string) ([]string, error) {
	var err error
	var val string
	results := make([]string, 0, len(keys))
	for _, k := range keys {
		val, err = GetString(json, append(path, k)...)
		if err != nil {
			return nil, err
		}
		results = append(results, val)
	}
	return results, nil
}

// GetAllMap not tested yet
// GetAllMap gets all values of path+key has pointed as string to string map.
func GetAllMap(json []byte, keys []string, path ...string) (map[string]string, error) {
	var err error
	var val string
	results := make(map[string]string)
	for _, k := range keys {
		val, err = GetString(json, append(path, k)...)
		if err != nil {
			return nil, err
		}
		results[k] = val
	}
	return results, nil
}
