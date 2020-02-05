package jin

import "strconv"

func AddKeyValue(json []byte, key string, value []byte, path ...string) ([]byte, error) {
	var start int
	var end int
	var err error
	if len(path) == 0 {
		for i := 0; i < len(json); i++ {
			if !space(json[i]) {
				if json[i] == 123 {
					start = i
					if i == len(json)-1 {
						return json, error_bad_json(i)
					}
					break
				} else {
					return json, error_object_expected()
				}
			}
		}
		for i := len(json) - 1; i > -1; i-- {
			if !space(json[i]) {
				if json[i] == 125 {
					end = i + 1
					if i == 0 {
						return json, error_bad_json(i)
					}
					break
				} else {
					return json, error_object_expected()
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
			if err.Error() == error_key_not_found().Error() {
				val := []byte(`,"` + key + `":` + string(value))
				json = replace(json, val, end-1, end-1)
				return json, nil
			}
			return json, err
		}
		return json, error_key_already_exists()
	}
	return json, error_object_expected()
}

func Add(json []byte, value []byte, path ...string) ([]byte, error) {
	var start int
	var end int
	var err error
	if len(json) < 2 {
		return json, error_bad_json(0)
	}
	if len(path) == 0 {
		for i := 0; i < len(json); i++ {
			if !space(json[i]) {
				if json[i] == 91 {
					start = i
					if i == len(json)-1 {
						return json, error_bad_json(i)
					}
					break
				} else {
					return json, error_array_expected()
				}
			}
		}
		for i := len(json) - 1; i > -1; i-- {
			if !space(json[i]) {
				if json[i] == 93 {
					end = i + 1
					if i == 0 {
						return json, error_bad_json(i)
					}
					break
				} else {
					return json, error_array_expected()
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
		return json, error_array_expected()
	}
	return json, error_bad_json(-1)
}

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
						return json, error_bad_json(i)
					}
					break
				} else {
					return json, error_array_expected()
				}
			}
		}
		for i := len(json) - 1; i > -1; i-- {
			if !space(json[i]) {
				if json[i] == 93 {
					end = i + 1
					if i == 0 {
						return json, error_bad_json(i)
					}
					break
				} else {
					return json, error_array_expected()
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
		return json, error_array_expected()
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
	return nil, error_bad_json(start)
}

func AddKeyValueString(json []byte, key, value string, path ...string) ([]byte, error) {
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
