package jin

import "strconv"

// AddKeyValue adds a key-value pair to an object.
// Path variable must point to an object,
// otherwise it will provide an error message.
func (p *Parser) AddKeyValue(key string, newVal []byte, path ...string) error {
	lenp := len(path)
	lenv := len(key)
	var curr *node
	var err error
	if lenv == 0 {
		return errNullKey()
	}
	json, err := p.Get(path...)
	if err != nil {
		return err
	}
	curr = p.core
	if lenp == 0 {
		for _, d := range curr.down {
			if d.label == key {
				return errKeyAlreadyExist(key)
			}
		}
		if len(json) >= 2 {
			if json[0] == 123 && json[len(json)-1] == 125 {
				newKV := []byte(`,"` + key + `":` + string(newVal))
				if lenv >= 2 {
					if newVal[0] == 91 || newVal[0] == 123 {
						newNode := createNode(nil)
						pCore(newVal, newNode)
						newNode.label = key
						newNode.value = newVal
						if len(p.json) == 2 {
							p.json = replace(p.json, newKV[1:], len(p.json)-1, len(p.json)-1)
						} else {
							p.json = replace(p.json, newKV, len(p.json)-1, len(p.json)-1)
						}
						curr.down = append(curr.down, newNode)
						newNode.up = curr
						return nil
					}
				}
				p.json = replace(p.json, newKV, len(p.json)-1, len(p.json)-1)
				newNode := createNode(nil)
				newNode.label = key
				newNode.value = newVal
				curr.down = append(curr.down, newNode)
				newNode.up = curr
				return nil
			}
			return errObjectExpected()
		}
		return errBadJSON(0)
	}
	curr, err = p.core.walk(path)
	if err != nil {
		return err
	}
	for _, d := range curr.up.down {
		if d.label == key {
			return errKeyAlreadyExist(key)
		}
	}
	if len(json) >= 2 {
		if json[0] == 123 && json[len(json)-1] == 125 {
			if lenv >= 2 {
				if newVal[0] == 91 || newVal[0] == 123 {
					newNode := createNode(nil)
					pCore(newVal, newNode)
					newNode.label = key
					newNode.value = newVal
					curr.down = append(curr.down, newNode)
					newNode.up = curr
					p.json, err = AddKeyValue(p.json, key, newVal, path...)
					if err != nil {
						return err
					}
					for i := 0; i < lenp; i++ {
						newNode = newNode.up
						newNode.value, err = Get(p.json, path[:lenp-i]...)
						if err != nil {
							return err
						}
					}
					return nil
				}
			}
			newNode := createNode(nil)
			newNode.label = key
			newNode.value = newVal
			curr.down = append(curr.down, newNode)
			newNode.up = curr
			p.json, err = AddKeyValue(p.json, key, newVal, path...)
			if err != nil {
				return err
			}
			for i := 0; i < lenp; i++ {
				newNode = newNode.up
				newNode.value, err = Get(p.json, path[:lenp-i]...)
				if err != nil {
					return err
				}
			}
			return nil
		}
		return errObjectExpected()
	}
	return errBadJSON(0)
}

// Add adds a value to an array.
// Path variable must point to an array,
// otherwise it will provide an error message.
func (p *Parser) Add(newVal []byte, path ...string) error {
	lenp := len(path)
	lenv := len(newVal)
	var curr *node
	var err error
	if lenv == 0 {
		return errNullKey()
	}
	json, err := p.Get(path...)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if lenp == 0 {
		curr = p.core
		if len(json) >= 2 {
			if json[0] == 91 && json[len(json)-1] == 93 {
				newValue := []byte(`,` + string(newVal))
				if lenv >= 2 {
					if newVal[0] == 91 || newVal[0] == 123 {
						newNode := createNode(nil)
						pCore(newVal, newNode)
						index := len(curr.down)
						newNode.label = strconv.Itoa(index)
						newNode.value = newVal
						if len(p.json) == 2 {
							p.json = replace(p.json, newValue[1:], len(p.json)-1, len(p.json)-1)
						} else {
							p.json = replace(p.json, newValue, len(p.json)-1, len(p.json)-1)
						}
						curr.down = append(curr.down, newNode)
						newNode.up = curr
						return nil
					}
				}
				p.json = replace(p.json, newValue, len(p.json)-1, len(p.json)-1)
				newNode := createNode(nil)
				index := len(curr.down)
				newNode.label = strconv.Itoa(index)
				newNode.value = newVal
				curr.down = append(curr.down, newNode)
				newNode.up = curr
				return nil
			}
			return errArrayExpected()
		}
		return errBadJSON(0)
	}
	curr, err = p.core.walk(path)
	if err != nil {
		return err
	}
	if len(json) >= 2 {
		if json[0] == 91 && json[len(json)-1] == 93 {
			if lenv >= 2 {
				if newVal[0] == 91 || newVal[0] == 123 {
					newNode := createNode(nil)
					pCore(newVal, newNode)
					index := len(curr.down)
					newNode.label = strconv.Itoa(index)
					newNode.value = newVal
					curr.down = append(curr.down, newNode)
					newNode.up = curr
					p.json, err = Add(p.json, newVal, path...)
					if err != nil {
						return err
					}
					for i := 0; i < lenp-1; i++ {
						newNode = newNode.up
						newNode.value, err = Get(p.json, path[:lenp-i]...)
						if err != nil {
							return err
						}
					}
					return nil
				}
			}
			newNode := createNode(nil)
			index := len(curr.down)
			newNode.label = strconv.Itoa(index)
			newNode.value = newVal
			curr.down = append(curr.down, newNode)
			newNode.up = curr
			p.json, err = Add(p.json, newVal, path...)
			if err != nil {
				return err
			}
			for i := 0; i < lenp; i++ {
				newNode = newNode.up
				newNode.value, err = Get(p.json, path[:lenp-i]...)
				if err != nil {
					return err
				}
			}
			return nil
		}
		return errArrayExpected()
	}
	return errBadJSON(0)
}

// Insert inserts a value to an array.
// Path variable must point to an array,
// otherwise it will provide an error message.
func (p *Parser) Insert(newIndex int, newVal []byte, path ...string) error {
	lenp := len(path)
	lenv := len(newVal)
	var curr *node
	var err error
	if lenv == 0 {
		return errNullKey()
	}
	json, err := p.Get(path...)
	if err != nil {
		return err
	}
	curr = p.core
	if lenp == 0 {
		if len(json) >= 2 {
			if json[0] == 91 && json[len(json)-1] == 93 {
				if lenv >= 2 {
					if newVal[0] == 91 || newVal[0] == 123 {
						newNode := createNode(nil)
						pCore(newVal, newNode)
						err = newNode.insert(curr, newIndex)
						if err != nil {
							return err
						}
						p.json, err = Insert(p.json, newIndex, newVal, path...)
						if err != nil {
							return err
						}
						err = newNode.insert(curr, newIndex)
						if err != nil {
							return err
						}
						curr.down = append(curr.down, newNode)
						newNode.up = curr
						return nil
					}
				}
				newNode := createNode(nil)
				err = newNode.insert(curr, newIndex)
				if err != nil {
					return err
				}
				newNode.value = newVal
				curr.down = append(curr.down, newNode)
				p.json, err = Insert(p.json, newIndex, newVal, path...)
				if err != nil {
					return err
				}
				err = newNode.insert(curr, newIndex)
				if err != nil {
					return err
				}
				newNode.up = curr
				return nil
			}
			return errArrayExpected()
		}
		return errBadJSON(0)
	}
	curr, err = p.core.walk(path)
	if err != nil {
		return err
	}
	if len(json) >= 2 {
		if json[0] == 91 && json[len(json)-1] == 93 {
			if lenv >= 2 {
				if newVal[0] == 91 || newVal[0] == 123 {
					newNode := createNode(nil)
					pCore(newVal, newNode)
					err = newNode.insert(curr, newIndex)
					if err != nil {
						return err
					}
					newNode.value = newVal
					curr.down = append(curr.down, newNode)
					newNode.up = curr
					p.json, err = Insert(p.json, newIndex, newVal, path...)
					if err != nil {
						return err
					}
					for i := 0; i < lenp; i++ {
						newNode = newNode.up
						newNode.value, err = Get(p.json, path[:lenp-i]...)
						if err != nil {
							return err
						}
					}
					return nil
				}
			}
			newNode := createNode(nil)
			err = newNode.insert(curr, newIndex)
			if err != nil {
				return err
			}
			newNode.value = newVal
			curr.down = append(curr.down, newNode)
			newNode.up = curr
			p.json, err = Insert(p.json, newIndex, newVal, path...)
			if err != nil {
				return err
			}
			for i := 0; i < lenp; i++ {
				newNode = newNode.up
				newNode.value, err = Get(p.json, path[:lenp-i]...)
				if err != nil {
					return err
				}
			}
			return nil
		}
		return errArrayExpected()
	}
	return errBadJSON(0)
}

// AddKeyValueString is a variation of AddKeyValue() func.
// Type of new value must be a string.
func (p *Parser) AddKeyValueString(key, value string, path ...string) error {
	if len(value) == 0 {
		return errNullNewValue()
	}
	if len(key) == 0 {
		return errNullKey()
	}
	return p.AddKeyValue(key, []byte(formatType(value)), path...)
}

// AddKeyValueInt is a variation of AddKeyValue() func.
// Type of new value must be an integer.
func (p *Parser) AddKeyValueInt(key string, value int, path ...string) error {
	if len(key) == 0 {
		return errNullKey()
	}
	return p.AddKeyValue(key, []byte(strconv.Itoa(value)), path...)
}

// AddKeyValueFloat is a variation of AddKeyValue() func.
// Type of new value must be a float64.
func (p *Parser) AddKeyValueFloat(key string, value float64, path ...string) error {
	if len(key) == 0 {
		return errNullKey()
	}
	return p.AddKeyValue(key, []byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

// AddKeyValueBool is a variation of AddKeyValue() func.
// Type of new value must be a boolean.
func (p *Parser) AddKeyValueBool(key string, value bool, path ...string) error {
	if len(key) == 0 {
		return errNullKey()
	}
	if value {
		return p.AddKeyValue(key, []byte("true"), path...)
	}
	return p.AddKeyValue(key, []byte("false"), path...)
}

// AddString is a variation of Add() func.
// Type of new value must be an string.
func (p *Parser) AddString(value string, path ...string) error {
	if len(value) == 0 {
		return errNullNewValue()
	}
	return p.Add([]byte(formatType(value)), path...)
}

// AddInt is a variation of Add() func.
// Type of new value must be an integer.
func (p *Parser) AddInt(value int, path ...string) error {
	return p.Add([]byte(strconv.Itoa(value)), path...)
}

// AddFloat is a variation of Add() func.
// Type of new value must be an float64.
func (p *Parser) AddFloat(value float64, path ...string) error {
	return p.Add([]byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

// AddBool is a variation of Add() func.
// Type of new value must be an boolean.
func (p *Parser) AddBool(value bool, path ...string) error {
	if value {
		return p.Add([]byte("true"), path...)
	}
	return p.Add([]byte("false"), path...)
}

// InsertString is a variation of Insert() func.
// Type of new value must be an string.
func (p *Parser) InsertString(index int, value string, path ...string) error {
	if len(value) == 0 {
		return errNullNewValue()
	}
	if index < 0 {
		return errIndexOutOfRange()
	}
	return p.Insert(index, []byte(formatType(value)), path...)
}

// InsertInt is a variation of Insert() func.
// Type of new value must be an integer.
func (p *Parser) InsertInt(index, value int, path ...string) error {
	if index < 0 {
		return errIndexOutOfRange()
	}
	return p.Insert(index, []byte(strconv.Itoa(value)), path...)
}

// InsertFloat is a variation of Insert() func.
// Type of new value must be an float64.
func (p *Parser) InsertFloat(index int, value float64, path ...string) error {
	if index < 0 {
		return errIndexOutOfRange()
	}
	return p.Insert(index, []byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

// InsertBool is a variation of Insert() func.
// Type of new value must be an boolean.
func (p *Parser) InsertBool(index int, value bool, path ...string) error {
	if index < 0 {
		return errIndexOutOfRange()
	}
	if value {
		return p.Insert(index, []byte("true"), path...)
	}
	return p.Insert(index, []byte("false"), path...)
}
