package jin

import "strconv"

// AddKeyValue adds a key-value pair to an object.
// Path variable must point to an object,
// otherwise it will provide an error message.
func AddKeyValue(json []byte, key string, value []byte, path ...string) ([]byte, error) {
	if value == nil {
		return AddKeyValue(json, key, []byte("null"), path...)
	}
	var start int
	var end int
	var err error
	if len(json) < 2 {
		return json, errBadJSON(0)
	}
	if len(path) == 0 {
		for i := 0; i < len(json); i++ {
			if !space(json[i]) {
				if json[i] == 123 {
					start = i
					if i == len(json)-1 {
						return json, errBadJSON(i)
					}
					break
				} else {
					return json, errObjectExpected()
				}
			}
		}
		for i := len(json) - 1; i > -1; i-- {
			if !space(json[i]) {
				if json[i] == 125 {
					end = i + 1
					if i == 0 {
						return json, errBadJSON(i)
					}
					break
				} else {
					return json, errObjectExpected()
				}
			}
		}
	} else {
		_, start, end, err = core(json, false, path...)
		if err != nil {
			return json, err
		}
	}
	if json[start] == 123 && json[end-1] == 125 {
		empty := true
		for i := start + 1; i < end-1; i++ {
			if !space(json[i]) {
				empty = false
			}
		}
		if empty {
			val := []byte(`"` + key + `":` + string(value))
			json = replace(json, val, end-1, end-1)
			return json, nil
		}
		path = append(path, key)
		_, _, _, err = core(json, false, path...)
		if err != nil {
			if ErrEqual(err, ErrCodeKeyNotFound) {
				val := []byte(`,"` + key + `":` + string(value))
				json = replace(json, val, end-1, end-1)
				return json, nil
			}
			return json, err
		}
		return json, errKeyAlreadyExist(key)
	}
	return json, errObjectExpected()
}

// Add adds a value to an array.
// Path variable must point to an array,
// otherwise it will provide an error message.
func Add(json []byte, value []byte, path ...string) ([]byte, error) {
	var start int
	var end int
	var err error
	if len(json) < 2 {
		return json, errBadJSON(0)
	}
	if len(path) == 0 {
		for i := 0; i < len(json); i++ {
			if !space(json[i]) {
				if json[i] == 91 {
					start = i
					if i == len(json)-1 {
						return json, errBadJSON(i)
					}
					break
				} else {
					return json, errArrayExpected()
				}
			}
		}
		for i := len(json) - 1; i > -1; i-- {
			if !space(json[i]) {
				if json[i] == 93 {
					end = i + 1
					if i == 0 {
						return json, errBadJSON(i)
					}
					break
				} else {
					return json, errArrayExpected()
				}
			}
		}
	} else {
		_, start, end, err = core(json, false, path...)
		if err != nil {
			return json, err
		}
	}
	if json[start] == 91 && json[end-1] == 93 {
		empty := true
		for i := start + 1; i < end-1; i++ {
			if !space(json[i]) {
				empty = false
			}
		}
		if empty {
			json = replace(json, value, end-1, end-1)
			return json, nil
		}
		val := make([]byte, len(value)+1)
		val[0] = 44
		copy(val[1:], value)
		json = replace(json, val, end-1, end-1)
		return json, nil
	}
	return json, errArrayExpected()
}

// Insert inserts a value to an array.
// Path variable must point to an array,
// otherwise it will provide an error message.
func Insert(json []byte, index int, value []byte, path ...string) ([]byte, error) {
	var start int
	var end int
	var err error
	if len(path) == 0 {
		for i := 0; i < len(json); i++ {
			if !space(json[i]) {
				if json[i] == 91 {
					start = i
					if i == len(json)-1 {
						return json, errBadJSON(i)
					}
					break
				} else {
					return json, errArrayExpected()
				}
			}
		}
		for i := len(json) - 1; i > -1; i-- {
			if !space(json[i]) {
				if json[i] == 93 {
					end = i + 1
					if i == 0 {
						return json, errBadJSON(i)
					}
					break
				} else {
					return json, errArrayExpected()
				}
			}
		}
	} else {
		_, start, end, err = core(json, false, path...)
		if err != nil {
			return json, err
		}
	}
	if json[start] != 91 || json[end-1] != 93 {
		return json, errArrayExpected()
	}
	_, start, end, err = core(json, false, append(path, strconv.Itoa(index))...)
	if err != nil {
		return json, err
	}
	if json[start-1] == 34 {
		start--
	}
	if json[end] == 34 {
		end++
	}
	var startEdge int
	var endEdge int
	for i := start - 1; i > 0; i-- {
		if !space(json[i]) {
			startEdge = i
			break
		}
	}
	for i := end; i < len(json); i++ {
		if !space(json[i]) {
			endEdge = i
			break
		}
	}
	if (json[startEdge] == 91 || json[startEdge] == 123) && json[startEdge]+2 == json[endEdge] {
		val := make([]byte, 0, len(value)+1)
		val = append(val, value...)
		val = append(val, 44)
		json = replace(json, val, start, start)
		return json, nil
	}
	if json[endEdge] == 44 {
		val := make([]byte, 0, len(value)+1)
		val = append(val, value...)
		val = append(val, 44)
		json = replace(json, val, start, start)
		return json, nil
	}
	if json[startEdge] == 44 {
		val := make([]byte, 0, len(value)+1)
		val = append(val, 44)
		val = append(val, value...)
		json = replace(json, val, start-1, start-1)
		return json, nil
	}
	return json, errBadJSON(start)
}

// AddKeyValueString is a variation of AddKeyValue() func.
// Type of new value must be a string.
func AddKeyValueString(json []byte, key, value string, path ...string) ([]byte, error) {
	if len(key) == 0 {
		return nil, errNullKey()
	}
	return AddKeyValue(json, key, []byte(formatType(value)), path...)
}

// AddKeyValueInt is a variation of AddKeyValue() func.
// Type of new value must be an integer.
func AddKeyValueInt(json []byte, key string, value int, path ...string) ([]byte, error) {
	if len(key) == 0 {
		return json, errNullKey()
	}
	return AddKeyValue(json, key, []byte(strconv.Itoa(value)), path...)
}

// AddKeyValueUint is a variation of AddKeyValue() func.
// Type of new value must be an integer.
func AddKeyValueUint(json []byte, key string, value uint, path ...string) ([]byte, error) {
	if len(key) == 0 {
		return json, errNullKey()
	}
	return AddKeyValue(json, key, []byte(strconv.FormatUint(uint64(value), 64)), path...)
}

// AddKeyValueInt32 is a variation of AddKeyValue() func.
// Type of new value must be an int32.
func AddKeyValueInt32(json []byte, key string, value int32, path ...string) ([]byte, error) {
	if len(key) == 0 {
		return json, errNullKey()
	}
	return AddKeyValue(json, key, []byte(strconv.FormatInt(int64(value), 32)), path...)
}

// AddKeyValueInt64 is a variation of AddKeyValue() func.
// Type of new value must be an int64.
func AddKeyValueInt64(json []byte, key string, value int64, path ...string) ([]byte, error) {
	if len(key) == 0 {
		return json, errNullKey()
	}
	return AddKeyValue(json, key, []byte(strconv.FormatInt(value, 64)), path...)
}

// AddKeyValueUint32 is a variation of AddKeyValue() func.
// Type of new value must be an uint32.
func AddKeyValueUint32(json []byte, key string, value uint32, path ...string) ([]byte, error) {
	if len(key) == 0 {
		return json, errNullKey()
	}
	return AddKeyValue(json, key, []byte(strconv.FormatUint(uint64(value), 32)), path...)
}

// AddKeyValueUint64 is a variation of AddKeyValue() func.
// Type of new value must be an uint64.
func AddKeyValueUint64(json []byte, key string, value uint64, path ...string) ([]byte, error) {
	if len(key) == 0 {
		return json, errNullKey()
	}
	return AddKeyValue(json, key, []byte(strconv.FormatUint(value, 64)), path...)
}

// AddKeyValueFloat is a variation of AddKeyValue() func.
// Type of new value must be a float64.
func AddKeyValueFloat(json []byte, key string, value float64, path ...string) ([]byte, error) {
	if len(key) == 0 {
		return json, errNullKey()
	}
	return AddKeyValue(json, key, []byte(strconv.FormatFloat(value, 'f', -1, 64)), path...)
}

// AddKeyValueBool is a variation of AddKeyValue() func.
// Type of new value must be a boolean.
func AddKeyValueBool(json []byte, key string, value bool, path ...string) ([]byte, error) {
	if len(key) == 0 {
		return json, errNullKey()
	}
	if value {
		return AddKeyValue(json, key, []byte("true"), path...)
	}
	return AddKeyValue(json, key, []byte("false"), path...)
}

// AddString is a variation of Add() func.
// Type of new value must be an string.
func AddString(json []byte, value string, path ...string) ([]byte, error) {
	return Add(json, []byte(formatType(value)), path...)
}

// AddInt is a variation of Add() func.
// Type of new value must be an integer.
func AddInt(json []byte, value int, path ...string) ([]byte, error) {
	return Add(json, []byte(strconv.Itoa(value)), path...)
}

// AddInt32 is a variation of Add() func.
// Type of new value must be an integer.
func AddInt32(json []byte, value int32, path ...string) ([]byte, error) {
	return Add(json, []byte(strconv.FormatInt(int64(value), 32)), path...)
}

// AddInt64 is a variation of Add() func.
// Type of new value must be an integer.
func AddInt64(json []byte, value int64, path ...string) ([]byte, error) {
	return Add(json, []byte(strconv.FormatInt(value, 64)), path...)
}

// AddUint is a variation of Add() func.
// Type of new value must be an integer.
func AddUint(json []byte, value uint, path ...string) ([]byte, error) {
	return Add(json, []byte(strconv.FormatUint(uint64(value), 64)), path...)
}

// AddUint32 is a variation of Add() func.
// Type of new value must be an integer.
func AddUint32(json []byte, value uint32, path ...string) ([]byte, error) {
	return Add(json, []byte(strconv.FormatUint(uint64(value), 32)), path...)
}

// AddUint64 is a variation of Add() func.
// Type of new value must be an integer.
func AddUint64(json []byte, value uint64, path ...string) ([]byte, error) {
	return Add(json, []byte(strconv.FormatUint(value, 64)), path...)
}

// AddFloat is a variation of Add() func.
// Type of new value must be an float64.
func AddFloat(json []byte, value float64, path ...string) ([]byte, error) {
	return Add(json, []byte(strconv.FormatFloat(value, 'f', -1, 64)), path...)
}

// AddBool is a variation of Add() func.
// Type of new value must be an boolean.
func AddBool(json []byte, value bool, path ...string) ([]byte, error) {
	if value {
		return Add(json, []byte("true"), path...)
	}
	return Add(json, []byte("false"), path...)
}

// InsertString is a variation of Insert() func.
// Type of new value must be an string.
func InsertString(json []byte, index int, value string, path ...string) ([]byte, error) {
	if index < 0 {
		return nil, errIndexOutOfRange()
	}
	return Insert(json, index, []byte(formatType(value)), path...)
}

// InsertInt is a variation of Insert() func.
// Type of new value must be an integer.
func InsertInt(json []byte, index, value int, path ...string) ([]byte, error) {
	if index < 0 {
		return json, errIndexOutOfRange()
	}
	return Insert(json, index, []byte(strconv.Itoa(value)), path...)
}

// InsertInt32 is a variation of Insert() func.
// Type of new value must be an integer.
func InsertInt32(json []byte, index int, value int32, path ...string) ([]byte, error) {
	if index < 0 {
		return json, errIndexOutOfRange()
	}
	return Insert(json, index, []byte(strconv.FormatInt(int64(value), 32)), path...)
}

// InsertInt64 is a variation of Insert() func.
// Type of new value must be an integer.
func InsertInt64(json []byte, index int, value int64, path ...string) ([]byte, error) {
	if index < 0 {
		return json, errIndexOutOfRange()
	}
	return Insert(json, index, []byte(strconv.FormatInt(value, 64)), path...)
}

// InsertUint is a variation of Insert() func.
// Type of new value must be an integer.
func InsertUint(json []byte, index int, value uint, path ...string) ([]byte, error) {
	if index < 0 {
		return json, errIndexOutOfRange()
	}
	return Insert(json, index, []byte(strconv.FormatUint(uint64(value), 64)), path...)
}

// InsertUint32 is a variation of Insert() func.
// Type of new value must be an integer.
func InsertUint32(json []byte, index int, value uint32, path ...string) ([]byte, error) {
	if index < 0 {
		return json, errIndexOutOfRange()
	}
	return Insert(json, index, []byte(strconv.FormatUint(uint64(value), 32)), path...)
}

// InsertUint64 is a variation of Insert() func.
// Type of new value must be an integer.
func InsertUint64(json []byte, index int, value uint64, path ...string) ([]byte, error) {
	if index < 0 {
		return json, errIndexOutOfRange()
	}
	return Insert(json, index, []byte(strconv.FormatUint(value, 64)), path...)
}

// InsertFloat is a variation of Insert() func.
// Type of new value must be an float64.
func InsertFloat(json []byte, index int, value float64, path ...string) ([]byte, error) {
	if index < 0 {
		return json, errIndexOutOfRange()
	}
	return Insert(json, index, []byte(strconv.FormatFloat(value, 'f', -1, 64)), path...)
}

// InsertBool is a variation of Insert() func.
// Type of new value must be an boolean.
func InsertBool(json []byte, index int, value bool, path ...string) ([]byte, error) {
	if index < 0 {
		return json, errIndexOutOfRange()
	}
	if value {
		return Insert(json, index, []byte("true"), path...)
	}
	return Insert(json, index, []byte("false"), path...)
}
