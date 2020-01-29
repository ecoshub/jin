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
	val, done := Get(json, path...)
	return string(val), done
}

func GetInt(json []byte, path ...string) (int, error) {
	value, err := GetString(json, path...)
	if err != nil {
		return -1, err
	}
	intVal, err := strconv.Atoi(value)
	if err != nil {
		return -1, PARSE_INT_ERROR(value)
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

// get array functions not safe in random json files.

func GetStringArray(json []byte, path ...string) ([]string, error) {
	val, err := GetString(json, path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, PARSE_STRING_ARRAY_ERROR(val)
	}
	if val[0] == '[' && val[lena-1] == ']' {
		newArray := make([]string, 0, 16)
		start := 1
		inQuote := false
		for i := 1; i < lena-1; i++ {
			curr := val[i]
			if curr == 34 || curr == 44 {
				if curr == 34 {
					// escape character control
					if val[i-1] == 92 {
						continue
					}
					inQuote = !inQuote
					continue
				}
				if inQuote {
					continue
				} else {
					if curr == 44 {
						newArray = append(newArray, trimSpace(val, start, i))
						start = i + 1
					}
				}
			}
		}
		newArray = append(newArray, trimSpace(val, start, lena-1))
		return newArray, nil
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
	if val[0] == '[' && val[lena-1] == ']' {
		newArray := make([]int, 0, 16)
		start := 1
		inQuote := false
		for i := 1; i < lena-1; i++ {
			curr := val[i]
			if curr == 34 || curr == 44 {
				if curr == 34 {
					// escape character control
					if val[i-1] == 92 {
						continue
					}
					inQuote = !inQuote
					continue
				}
				if inQuote {
					continue
				} else {
					if curr == 44 {
						num, err := strconv.Atoi(trimSpace(val, start, i))
						if err != nil {
							return nil, PARSE_INT_ERROR(trimSpace(val, start, i))
						}
						newArray = append(newArray, num)
						start = i + 1
					}
				}
			}
		}

		num, err := strconv.Atoi(trimSpace(val, start, lena-1))
		if err != nil {
			return nil, PARSE_INT_ERROR(trimSpace(val, start, lena-1))
		}
		newArray = append(newArray, num)
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
	if val[0] == '[' && val[lena-1] == ']' {
		newArray := make([]float64, 0, 16)
		start := 1
		inQuote := false
		for i := 1; i < lena-1; i++ {
			curr := val[i]
			if curr == 34 || curr == 44 {
				if curr == 34 {
					// escape character control
					if val[i-1] == 92 {
						continue
					}
					inQuote = !inQuote
					continue
				}
				if inQuote {
					continue
				} else {
					if curr == 44 {
						num, err := strconv.ParseFloat(trimSpace(val, start, i), 64)
						if err != nil {
							return nil, PARSE_FLOAT_ERROR(trimSpace(val, start, i))
						}
						newArray = append(newArray, num)
						start = i + 1
					}
				}
			}
		}

		num, err := strconv.ParseFloat(trimSpace(val, start, lena-1), 64)
		if err != nil {
			return nil, PARSE_FLOAT_ERROR(trimSpace(val, start, lena-1))
		}
		newArray = append(newArray, num)
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
	if val[0] == '[' && val[lena-1] == ']' {
		newArray := make([]bool, 0, 16)
		start := 1
		inQuote := false
		for i := 1; i < lena-1; i++ {
			curr := val[i]
			if curr == 34 || curr == 44 {
				if curr == 34 {
					// escape character control
					if val[i-1] == 92 {
						continue
					}
					inQuote = !inQuote
					continue
				}
				if inQuote {
					continue
				} else {
					if curr == 44 {
						val := trimSpace(val, start, i)
						if val == "true" || val == "false" {
							if val == "true" {
								newArray = append(newArray, true)
								start = i + 1
							} else {
								newArray = append(newArray, false)
								start = i + 1
							}
						} else {
							return nil, PARSE_BOOL_ERROR(val)
						}
					}
				}
			}
		}
		val := trimSpace(val, start, lena-2)
		if val == "true" || val == "false" {
			if val == "true" {
				newArray = append(newArray, true)
			} else {
				newArray = append(newArray, false)
			}
		} else {
			return nil, PARSE_BOOL_ERROR(val)
		}
		return newArray, nil
	} else {
		return nil, PARSE_BOOL_ARRAY_ERROR(val)
	}
}
