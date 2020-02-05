package jin

import "strconv"

func (p *Parser) Get(path ...string) ([]byte, error) {
	if len(path) == 0 {
		return p.json, nil
	}
	curr, err := p.core.walk(path)
	if err != nil {
		return nil, err
	}
	return cleanValue(curr.value), nil
}

func (p *Parser) GetString(path ...string) (string, error) {
	val, err := p.Get(path...)
	if err != nil {
		return "", err
	}
	return string(val), err
}

func (p *Parser) GetInt(path ...string) (int, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return -1, err
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return -1, error_parse_int(val)
	}
	return intVal, nil
}

func (p *Parser) GetFloat(path ...string) (float64, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return -1, err
	}
	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return -1, error_parse_float(val)
	}
	return floatVal, nil
}

func (p *Parser) GetBool(path ...string) (bool, error) {
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
	return false, error_parse_bool(val)
}

func (p *Parser) GetStringArray(path ...string) ([]string, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, error_parse_string_array(val)
	}
	if val[0] == 91 && val[lena-1] == 93 {
		arr := ParseArray(val)
		if arr == nil {
			return nil, error_parse_string_array(val)
		}
		return arr, nil
	}
	return nil, error_parse_string_array(val)
}

func (p *Parser) GetIntArray(path ...string) ([]int, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, error_parse_int_array(val)
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
								return nil, error_parse_int(cleanValueString(element))
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
							return nil, error_parse_int(cleanValueString(element))
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
	return nil, error_parse_int_array(val)
}

func (p *Parser) GetFloatArray(path ...string) ([]float64, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, error_parse_float_array(val)
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
								return nil, error_parse_float(cleanValueString(element))
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
							return nil, error_parse_float(cleanValueString(element))
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
	return nil, error_parse_float_array(val)
}

func (p *Parser) GetBoolArray(path ...string) ([]bool, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, error_parse_bool_array(val)
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
								return nil, error_parse_bool(cleanValueString(element))
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
							return nil, error_parse_bool(element)
						}
						start = i + 1
						continue
					}
				}
			}
		}
		return newArray, nil
	}
	return nil, error_parse_bool_array(val)
}
