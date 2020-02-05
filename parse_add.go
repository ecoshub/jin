package jin

import "strconv"

func (p *parse) AddKeyValue(key string, newVal []byte, path ...string) error {
	lenp := len(path)
	lenv := len(key)
	var curr *node
	var err error
	if lenv == 0 {
		return ERROR_NULL_KEY()
	}
	json, err := p.Get(path...)
	if err != nil {
		return err
	}
	curr = p.core
	if lenp == 0 {
		for _, d := range curr.down {
			if d.label == key {
				return ERROR_KEY_ALREADY_EXISTS()
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
			return ERROR_OBJECT_EXPECTED()
		}
		return ERROR_BAD_JSON(0)
	}
	curr, err = p.core.walk(path)
	if err != nil {
		return err
	}
	for _, d := range curr.up.down {
		if d.label == key {
			return ERROR_KEY_ALREADY_EXISTS()
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
		return ERROR_OBJECT_EXPECTED()
	}
	return ERROR_BAD_JSON(0)
}

func (p *parse) Add(newVal []byte, path ...string) error {
	lenp := len(path)
	lenv := len(newVal)
	var curr *node
	var err error
	if lenv == 0 {
		return ERROR_NULL_KEY()
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
			return ERROR_ARRAY_EXPECTED()
		}
		return ERROR_BAD_JSON(0)
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
		return ERROR_ARRAY_EXPECTED()
	}
	return ERROR_BAD_JSON(0)
}

func (p *parse) Insert(newIndex int, newVal []byte, path ...string) error {
	lenp := len(path)
	lenv := len(newVal)
	var curr *node
	var err error
	if lenv == 0 {
		return ERROR_NULL_KEY()
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
			return ERROR_ARRAY_EXPECTED()
		}
		return ERROR_BAD_JSON(0)
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
		return ERROR_ARRAY_EXPECTED()
	}
	return ERROR_BAD_JSON(0)
}

func (p *parse) AddKeyValueString(key, value string, path ...string) error {
	return p.AddKeyValue(key, []byte(value), path...)
}

func (p *parse) AddKeyValueInt(key string, value int, path ...string) error {
	return p.AddKeyValue(key, []byte(strconv.Itoa(value)), path...)
}

func (p *parse) AddKeyValueFloat(key string, value float64, path ...string) error {
	return p.AddKeyValue(key, []byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

func (p *parse) AddKeyValueBool(key string, value bool, path ...string) error {
	if value {
		return p.AddKeyValue(key, []byte("true"), path...)
	}
	return p.AddKeyValue(key, []byte("false"), path...)
}

func (p *parse) AddString(value string, path ...string) error {
	if value[0] != 34 && value[len(value)-1] != 34 {
		return p.Add([]byte(`"`+value+`"`), path...)
	}
	return p.Add([]byte(value), path...)
}

func (p *parse) AddInt(value int, path ...string) error {
	return p.Add([]byte(strconv.Itoa(value)), path...)
}

func (p *parse) AddFloat(value float64, path ...string) error {
	return p.Add([]byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

func (p *parse) AddBool(value bool, path ...string) error {
	if value {
		return p.Add([]byte("true"), path...)
	}
	return p.Add([]byte("false"), path...)
}

func (p *parse) InsertString(index int, value string, path ...string) error {
	if value[0] != 34 && value[len(value)-1] != 34 {
		return p.Insert(index, []byte(`"`+value+`"`), path...)
	}
	return p.Insert(index, []byte(value), path...)
}

func (p *parse) InsertInt(index, value int, path ...string) error {
	return p.Insert(index, []byte(strconv.Itoa(value)), path...)
}

func (p *parse) InsertFloat(index int, value float64, path ...string) error {
	return p.Insert(index, []byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

func (p *parse) InsertBool(index int, value bool, path ...string) error {
	if value {
		return p.Insert(index, []byte("true"), path...)
	}
	return p.Insert(index, []byte("false"), path...)
}
