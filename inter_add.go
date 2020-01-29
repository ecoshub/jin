package jint

import "strconv"

func AddKeyValue(json []byte, key string, value []byte, path ...string) ([]byte, error) {
	if len(json) < 2 {
		return json, BAD_JSON_ERROR(0)
	}
	var start int
	var end int
	var err error
	if len(path) == 0 {
		for i := 0; i < len(json); i++ {
			if !space(json[i]) {
				if json[i] == 123 {
					start = i
					if i == len(json)-1 {
						return json, BAD_JSON_ERROR(i)
					}
					break
				} else {
					return json, OBJECT_EXPECTED_ERROR()
				}
			}
		}
		for i := len(json) - 1; i > -1; i-- {
			if !space(json[i]) {
				if json[i] == 125 {
					end = i + 1
					if i == 0 {
						return json, BAD_JSON_ERROR(i)
					}
					break
				} else {
					return json, OBJECT_EXPECTED_ERROR()
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
		} else {
			path = append(path, key)
			// key already exist control
			_, _, _, err = core(json, false, path...)
			if err != nil {
				if err.Error() == KEY_NOT_FOUND_ERROR().Error() {
					val := []byte(`,"` + key + `":` + string(value))
					json = replace(json, val, end-1, end-1)
					return json, nil
				}
				return json, err
			}
			return json, KEY_ALREADY_EXIST_ERROR()
		}
	} else {
		return json, OBJECT_EXPECTED_ERROR()
	}
	return json, BAD_JSON_ERROR(-1)
}

func Add(json []byte, value []byte, path ...string) ([]byte, error) {
	var start int
	var end int
	var err error
	if len(json) < 2 {
		return json, BAD_JSON_ERROR(0)
	}
	if len(path) == 0 {
		for i := 0; i < len(json); i++ {
			if !space(json[i]) {
				if json[i] == 91 {
					start = i
					if i == len(json)-1 {
						return json, BAD_JSON_ERROR(i)
					}
					break
				} else {
					return json, ARRAY_EXPECTED_ERROR()
				}
			}
		}
		for i := len(json) - 1; i > -1; i-- {
			if !space(json[i]) {
				if json[i] == 93 {
					end = i + 1
					if i == 0 {
						return json, BAD_JSON_ERROR(i)
					}
					break
				} else {
					return json, ARRAY_EXPECTED_ERROR()
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
		} else {
			val := make([]byte, len(value)+1)
			val[0] = 44
			copy(val[1:], value)
			json = replace(json, val, end-1, end-1)
			return json, nil
		}
	} else {
		return json, ARRAY_EXPECTED_ERROR()
	}
	return json, BAD_JSON_ERROR(-1)
}

func AddKeyValueString(json []byte, key, value string, path ...string) ([]byte, error) {
	if value[0] != 34 && value[len(value)-1] != 34 {
		value = `"` + value + `"`
	}
	return AddKeyValue(json, key, []byte(value), path...)
}

func AddKeyValueInt(json []byte, key string, value int, path ...string) ([]byte, error) {
	return AddKeyValue(json, key, []byte(strconv.Itoa(value)), path...)
}

func AddKeyValueFloat(json []byte, key string, value float64, path ...string) ([]byte, error) {
	return AddKeyValue(json, key, []byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

func AddKeyValueBool(json []byte, key string, value bool, path ...string) ([]byte, error) {
	if value {
		return AddKeyValue(json, key, []byte("true"), path...)
	}
	return AddKeyValue(json, key, []byte("false"), path...)
}

func AddString(json []byte, value string, path ...string) ([]byte, error) {
	if value[0] != 34 && value[len(value)-1] != 34 {
		return Add(json, []byte(`"`+value+`"`), path...)
	}
	return Add(json, []byte(value), path...)
}

func AddInt(json []byte, value int, path ...string) ([]byte, error) {
	return Add(json, []byte(strconv.Itoa(value)), path...)
}

func AddFloat(json []byte, value float64, path ...string) ([]byte, error) {
	return Add(json, []byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

func AddBool(json []byte, value bool, path ...string) ([]byte, error) {
	if value {
		return Add(json, []byte("true"), path...)
	}
	return Add(json, []byte("false"), path...)
}

func Insert(json []byte, index int, value []byte, path ...string) ([]byte, error) {
	// lenpath == 0 and empty array control needed
	var start int
	var err error
	if len(path) == 0 {
		for i := 0; i < len(json); i++ {
			if !space(json[i]) {
				start = i
				break
			}
		}
	} else {
		_, start, _, err = core(json, true, path...)
		if err != nil {
			return json, err
		}
	}
	if json[start] != 91 {
		return json, ARRAY_EXPECTED_ERROR()
	}
	indexStr := strconv.Itoa(index)
	path = append(path, indexStr)
	_, start, _, err = core(json, true, path...)
	if err != nil {
		return json, err
	}
	val := make([]byte, len(value)+1)
	copy(val, value)
	val[len(val)-1] = 44
	if json[start-1] == 34 {
		json = replace(json, val, start-1, start-1)
		return json, nil
	}
	json = replace(json, val, start, start)
	return json, nil
}

func InsertString(json []byte, index int, value string, path ...string) ([]byte, error) {
	if value[0] != 34 && value[len(value)-1] != 34 {
		return Insert(json, index, []byte(`"`+value+`"`), path...)
	}
	return Insert(json, index, []byte(value), path...)
}

func InsertInt(json []byte, index, value int, path ...string) ([]byte, error) {
	return Insert(json, index, []byte(strconv.Itoa(value)), path...)
}

func InsertFloat(json []byte, index int, value float64, path ...string) ([]byte, error) {
	return Insert(json, index, []byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

func InsertBool(json []byte, index int, value bool, path ...string) ([]byte, error) {
	if value {
		return Insert(json, index, []byte("true"), path...)
	}
	return Insert(json, index, []byte("false"), path...)
}
