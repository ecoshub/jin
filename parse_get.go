package jin

import "strconv"

// ToJSON returns underlying json value as byte array
func (p *Parser) ToJSON() []byte {
	return p.json
}

// Get returns the value that path has pointed.
// It stripes quotation marks from string values.
// Path can point anything, a key-value pair, a value, an array, an object.
// Path variable can not be null,
// otherwise it will provide an error message.
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

// GetNew returns the value that path has pointed.
// It stripes quotation marks from string values.
// Path can point anything, a key-value pair, a value, an array, an object.
// Path variable can not be null,
// otherwise it will provide an error message.
func (p *Parser) GetNew(path ...string) []byte {
	var err error
	curr := p.core
	if len(path) != 0 {
		curr, err = p.core.walk(path)
		if err != nil {
			return nil
		}
	}
	if len(curr.down) == 0 {
		return cleanValue(curr.value)
	}
	temp := make([]byte, 0, 128)
	temp = curr.dive(temp)
	return cleanValue(temp)
}

// GetString is a variation of Get() func.
// GetString returns the value that path has pointed as string.
func (p *Parser) GetString(path ...string) (string, error) {
	val, err := p.Get(path...)
	if err != nil {
		return "", err
	}
	return string(val), err
}

// GetInt is a variation of Get() func.
// GetInt returns the value that path has pointed as integer.
// returns an error message if the value to be returned cannot be converted to an integer
func (p *Parser) GetInt(path ...string) (int, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return -1, err
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return -1, ErrIntegerParse(val)
	}
	return intVal, nil
}

// GetFloat is a variation of Get() func.
// GetFloat returns the value that path has pointed as float.
// returns an error message if the value to be returned cannot be converted to an float
func (p *Parser) GetFloat(path ...string) (float64, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return -1, err
	}
	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return -1, ErrFloatParse(val)
	}
	return floatVal, nil
}

// GetBool is a variation of Get() func.
// GetBool returns the value that path has pointed as boolean.
// returns an error message if the value to be returned cannot be converted to an boolean
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
	return false, ErrBoolParse(val)
}

// GetStringArray is a variation of Get() func.
// GetStringArray returns the value that path has pointed as string slice.
// returns an error message if the value to be returned cannot be converted to an string slice.
func (p *Parser) GetStringArray(path ...string) ([]string, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, ErrStringArrayParse(val)
	}
	if val[0] == 91 && val[lena-1] == 93 {
		arr := ParseArray(val)
		if arr == nil {
			return nil, ErrStringArrayParse(val)
		}
		return arr, nil
	}
	return nil, ErrStringArrayParse(val)
}

// GetIntArray is a variation of Get() func.
// GetIntArray returns the value that path has pointed as integer slice.
// returns an error message if the value to be returned cannot be converted to an integer slice.
func (p *Parser) GetIntArray(path ...string) ([]int, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, ErrIntegerArrayParse(val)
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
								return nil, ErrIntegerParse(cleanValueString(element))
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
							return nil, ErrIntegerParse(cleanValueString(element))
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
	return nil, ErrIntegerArrayParse(val)
}

// GetFloatArray is a variation of Get() func.
// GetFloatArray returns the value that path has pointed as float slice.
// returns an error message if the value to be returned cannot be converted to an float slice.
func (p *Parser) GetFloatArray(path ...string) ([]float64, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, ErrFloatArrayParse(val)
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
								return nil, ErrFloatParse(cleanValueString(element))
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
							return nil, ErrFloatParse(cleanValueString(element))
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
	return nil, ErrFloatArrayParse(val)
}

// GetBoolArray is a variation of Get() func.
// GetBoolArray returns the value that path has pointed as boolean slice.
// returns an error message if the value to be returned cannot be converted to an boolean slice.
func (p *Parser) GetBoolArray(path ...string) ([]bool, error) {
	val, err := p.GetString(path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, ErrBoolArrayParse(val)
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
								return nil, ErrBoolParse(cleanValueString(element))
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
							return nil, ErrBoolParse(element)
						}
						start = i + 1
						continue
					}
				}
			}
		}
		return newArray, nil
	}
	return nil, ErrBoolArrayParse(val)
}
