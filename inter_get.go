package jint

import "strconv"

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

func GetString(json []byte, path ...string) (string, error) {
	val, err := Get(json, path...)
	if err != nil {
		return "", err
	}
	return string(val), err
}

func GetInt(json []byte, path ...string) (int, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return -1, err
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return -1, PARSE_INT_ERROR(val)
	}
	return intVal, nil
}

func GetFloat(json []byte, path ...string) (float64, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return -1, err
	}
	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return -1, PARSE_FLOAT_ERROR(val)
	}
	return floatVal, nil
}

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
	return false, PARSE_BOOL_ERROR(val)
}

func GetStringArray(json []byte, path ...string) ([]string, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, PARSE_STRING_ARRAY_ERROR(val)
	}
	if val[0] == 91 && val[lena-1] == 93 {
		arr := ParseArray(val)
		if arr == nil {
			return nil, PARSE_STRING_ARRAY_ERROR(val)
		}
		return arr, nil
	} else {
		return nil, PARSE_STRING_ARRAY_ERROR(val)
	}
}

func GetIntArray(json []byte, path ...string) ([]int, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, PARSE_INT_ARRAY_ERROR(val)
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
							num, err := strconv.Atoi(trimAndStripQuote(element))
							if err != nil {
								return nil, PARSE_INT_ERROR(trimAndStripQuote(element))
							}
							newArray = append(newArray, num)
							start = i + 1
							break
						}
					}
				}
				if level == 1 {
					if curr == 44 {
						element := val[start:i]
						num, err := strconv.Atoi(trimAndStripQuote(element))
						if err != nil {
							return nil, PARSE_INT_ERROR(trimAndStripQuote(element))
						}
						newArray = append(newArray, num)
						start = i + 1
						continue
					}
				}
			}
		}
		return newArray, nil
	} else {
		return nil, PARSE_INT_ARRAY_ERROR(val)
	}
}

func GetFloatArray(json []byte, path ...string) ([]float64, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, PARSE_FLOAT_ARRAY_ERROR(val)
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
							num, err := strconv.ParseFloat(trimAndStripQuote(element), 64)
							if err != nil {
								return nil, PARSE_FLOAT_ERROR(trimAndStripQuote(element))
							}
							newArray = append(newArray, num)
							start = i + 1
							break
						}
					}
				}
				if level == 1 {
					if curr == 44 {
						element := val[start:i]
						num, err := strconv.ParseFloat(trimAndStripQuote(element), 64)
						if err != nil {
							return nil, PARSE_FLOAT_ERROR(trimAndStripQuote(element))
						}
						newArray = append(newArray, num)
						start = i + 1
						continue
					}
				}
			}
		}
		return newArray, nil
	} else {
		return nil, PARSE_FLOAT_ARRAY_ERROR(val)
	}
}

func GetBoolArray(json []byte, path ...string) ([]bool, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, PARSE_BOOL_ARRAY_ERROR(val)
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
							element = trimAndStripQuote(element)
							if element == "true" || element == "false" {
								if element == "true" {
									newArray = append(newArray, true)
								} else {
									newArray = append(newArray, false)
								}
							} else {
								return nil, PARSE_BOOL_ERROR(trimAndStripQuote(element))
							}
							start = i + 1
							break
						}
					}
				}
				if level == 1 {
					if curr == 44 {
						element := val[start:i]
						element = trimAndStripQuote(element)
						if element == "true" || element == "false" {
							if element == "true" {
								newArray = append(newArray, true)
							} else {
								newArray = append(newArray, false)
							}
						} else {
							return nil, PARSE_BOOL_ERROR(element)
						}
						start = i + 1
						continue
					}
				}
			}
		}
		return newArray, nil
	} else {
		return nil, PARSE_BOOL_ARRAY_ERROR(val)
	}
}
