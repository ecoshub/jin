package jin

import "strconv"

// Set sets the value that path has pointed.
// Path can point anything, a key-value pair, a value, an array, an object.
// Path variable can not be null,
// otherwise it will provide an error message.
func (p *Parser) Set(newVal []byte, path ...string) error {
	lenp := len(path)
	lenv := len(newVal)
	var curr *node
	var err error
	if lenp == 0 {
		return nullPathError()
	}
	if lenv == 0 {
		return nullNewValueError()
	}
	curr, err = p.core.walk(path)
	if err != nil {
		return err
	}
	if lenv >= 2 {
		if newVal[0] == 91 || newVal[0] == 123 {
			newCore := createNode(nil)
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

// SetString is a variation of Set() func.
// SetString takes the set value as string.
func (p *Parser) SetString(newValue string, path ...string) error {
	if newValue[0] != 34 && newValue[len(newValue)-1] != 34 {
		return p.Set([]byte(`"`+newValue+`"`), path...)
	}
	return p.Set([]byte(newValue), path...)
}

// SetInt is a variation of Set() func.
// SetInt takes the set value as integer.
func (p *Parser) SetInt(newValue int, path ...string) error {
	return p.Set([]byte(strconv.Itoa(newValue)), path...)
}

// SetFloat is a variation of Set() func.
// SetFloat takes the set value as float64.
func (p *Parser) SetFloat(newValue float64, path ...string) error {
	return p.Set([]byte(strconv.FormatFloat(newValue, 'e', -1, 64)), path...)
}

// SetBool is a variation of Set() func.
// SetBool takes the set value as boolean.
func (p *Parser) SetBool(newValue bool, path ...string) error {
	if newValue {
		return p.Set([]byte("true"), path...)
	}
	return p.Set([]byte("false"), path...)
}

// SetKey sets the key value of key-value pair that path has pointed.
// Path must point to an object.
// otherwise it will provide an error message.
// Path variable can not be null,
func (p *Parser) SetKey(newKey string, path ...string) error {
	lenp := len(path)
	lenv := len(newKey)
	var curr *node
	var err error
	if lenp == 0 {
		return nullPathError()
	}
	if lenv == 0 {
		return nullKeyError()
	}
	curr, err = p.core.walk(path)
	if err != nil {
		return err
	}
	for _, d := range curr.up.down {
		if d.label == newKey {
			return keyAlreadyExistsError(newKey)
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
