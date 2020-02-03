package jint

import "strconv"

func (p *parse) Set(newVal []byte, path ...string) error {
	lenp := len(path)
	lenv := len(newVal)
	var curr *node
	var err error
	if lenp == 0 {
		return NULL_PATH_ERROR()
	}
	if lenv == 0 {
		return NULL_NEW_VALUE_ERROR()
	}
	curr, err = p.core.walk(path)
	if err != nil {
		return err
	}
	if lenv >= 2 {
		if newVal[0] == 91 || newVal[0] == 123 {
			newCore := CreateNode(nil)
			pCore(newVal, newCore)
			newCore.label = curr.label
			newCore.up = curr.up
			index := curr.getIndex()
			curr.up.down[index] = newCore
			newCore.value = newVal
			p.json, _ = Set(p.json, newVal, path...)
			for i := 0; i < lenp-1; i++ {
				newCore = newCore.up
				newCore.value, err = Get(p.json, path[:lenp-1-i]...)
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
	curr.value = newVal
	p.json, _ = Set(p.json, newVal, path...)
	for i := 0; i < lenp-1; i++ {
		curr = curr.up
		curr.value, err = Get(p.json, path[:lenp-1-i]...)
		if err != nil {
			return err
		}

	}
	return nil
}

func (p *parse) SetString(newValue string, path ...string) error {
	if newValue[0] != 34 && newValue[len(newValue)-1] != 34 {
		return p.Set([]byte(`"`+newValue+`"`), path...)
	}
	return p.Set([]byte(newValue), path...)
}

func (p *parse) SetInt(newValue int, path ...string) error {
	return p.Set([]byte(strconv.Itoa(newValue)), path...)
}

func (p *parse) SetFloat(newValue float64, path ...string) error {
	return p.Set([]byte(strconv.FormatFloat(newValue, 'e', -1, 64)), path...)
}

func (p *parse) SetBool(newValue bool, path ...string) error {
	if newValue {
		return p.Set([]byte("true"), path...)
	}
	return p.Set([]byte("false"), path...)
}

func (p *parse) SetKey(newKey string, path ...string) error {
	lenp := len(path)
	lenv := len(newKey)
	var curr *node
	var err error
	if lenp == 0 {
		return NULL_PATH_ERROR()
	}
	if lenv == 0 {
		return NULL_KEY_ERROR()
	}
	curr, err = p.core.walk(path)
	if err != nil {
		return err
	}
	for _, d := range curr.up.down {
		if d.label == newKey {
			return KEY_ALREADY_EXISTS_ERROR()
		}
	}
	curr.label = newKey
	p.json, _ = SetKey(p.json, newKey, path...)
	for i := 0; i < lenp-1; i++ {
		curr = curr.up
		curr.value, _ = Get(p.json, path[:lenp-1-i]...)
		if err != nil {
			return err
		}
	}
	return nil
}
