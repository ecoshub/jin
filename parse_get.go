package jin

import "strconv"

func (p *parse) Get(path ...string) ([]byte, error) {
	if len(path) == 0 {
		return p.json, nil
	}
	curr, err := p.core.walk(path)
	if err != nil {
		return nil, err
	}
	return cleanValue(curr.value), nil
}

func (p *parse) GetString(path ...string) (string, error) {
	val, err := p.Get(path...)
	if err != nil {
		return "", err
	}
	return string(val), err
}

func (p *parse) GetInt(path ...string) (int, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return -1, err
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return -1, ERROR_PARSE_INT(val)
	}
	return intVal, nil
}

func (p *parse) GetFloat(path ...string) (float64, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return -1, err
	}
	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return -1, ERROR_PARSE_FLOAT(val)
	}
	return floatVal, nil
}

func (p *parse) GetBool(path ...string) (bool, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return false, err
	}
	if val == "true" {
		return true, nil
	}
	if val == "false" {
		return false, nil
	}
	return false, ERROR_PARSE_BOOL(val)
}

func (p *parse) GetStringArray(path ...string) ([]string, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, ERROR_PARSE_STRING_ARRAY(val)
	}
	if val[0] == 91 && val[lena-1] == 93 {
		arr := ParseArray(val)
		if arr == nil {
			return nil, ERROR_PARSE_STRING_ARRAY(val)
		}
		return arr, nil
	}
	return nil, ERROR_PARSE_STRING_ARRAY(val)
}

func (p *parse) GetIntArray(path ...string) ([]int, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, ERROR_PARSE_INT_ARRAY(val)
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
								return nil, ERROR_PARSE_INT(cleanValueString(element))
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
							return nil, ERROR_PARSE_INT(cleanValueString(element))
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
	return nil, ERROR_PARSE_INT_ARRAY(val)
}

func (p *parse) GetFloatArray(path ...string) ([]float64, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, ERROR_PARSE_FLOAT_ARRAY(val)
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
								return nil, ERROR_PARSE_FLOAT(cleanValueString(element))
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
							return nil, ERROR_PARSE_FLOAT(cleanValueString(element))
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
	return nil, ERROR_PARSE_FLOAT_ARRAY(val)
}

func (p *parse) GetBoolArray(path ...string) ([]bool, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, ERROR_PARSE_BOOL_ARRAY(val)
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
								return nil, ERROR_PARSE_BOOL(cleanValueString(element))
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
							return nil, ERROR_PARSE_BOOL(element)
						}
						start = i + 1
						continue
					}
				}
			}
		}
		return newArray, nil
	}
	return nil, ERROR_PARSE_BOOL_ARRAY(val)
}
