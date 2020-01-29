package jint

func IsObject(json []byte, path ...string) (bool, error) {
	state, _, err := typeControlcore(json, []byte{91, 123}, true, path...)
	if err != nil {
		return false, err
	}
	return state, nil
}

func IsArray(json []byte, path ...string) (bool, error) {
	state, _, err := typeControlcore(json, []byte{91}, true, path...)
	if err != nil {
		return false, err
	}
	return state, nil
}

func IsValue(json []byte, path ...string) (bool, error) {
	state, _, err := typeControlcore(json, []byte{91, 123}, false, path...)
	if err != nil {
		return false, err
	}
	return state, nil
}

func GetType(json []byte, path ...string) (string, error) {
	_, start, err := typeControlcore(json, []byte{}, false, path...)
	if err != nil {
		return "ERROR", err
	}
	switch json[start] {
	case 91:
		return "array", nil
	case 123:
		return "object", nil
	default:
		return "value", nil
	}
	return "ERROR", nil
}

func IsEmpty(json []byte, path ...string) (bool, error) {
	var start int
	var end int = len(json) - 1
	var err error
	if len(path) == 0 {
		for space(json[start]) {
			if start > len(json)-1 {
				return false, BAD_JSON_ERROR(start)
			} else {
				start++
				continue
			}
		}
		for space(json[end]) {
			if end < 1 {
				return false, BAD_JSON_ERROR(end)
			} else {
				end--
				continue
			}
		}
	} else {
		_, start, end, err = core(json, false, path...)
		if err != nil {
			return false, err
		}
	}
	braceStart := json[start]
	braceEnd := json[end-1]
	if braceStart == 91 || braceStart == 123 {
		if braceStart+2 != braceEnd {
			return false, BAD_JSON_ERROR(end)
		}
		for i := start + 1; i < end-1; i++ {
			if !space(json[i]) {
				return false, nil
			}
		}
	} else {
		return false, OBJECT_EXPECTED_ERROR()
	}
	return true, nil
}

func typeControlcore(json []byte, control []byte, equal bool, path ...string) (bool, int, error) {
	var start int
	var err error
	if len(path) == 0 {
		for space(json[start]) {
			if start > len(json)-1 {
				return false, -1, BAD_JSON_ERROR(start)
			} else {
				start++
				continue
			}
		}
	} else {
		_, start, _, err = core(json, true, path...)
		if err != nil {
			return false, start, err
		}
	}
	for _, v := range control {
		if json[start] == v {
			if equal {
				return true, start, nil
			} else {
				return false, start, nil
			}
		}
	}
	if equal {
		return false, start, nil
	} else {
		return true, start, nil
	}
	return false, -1, nil
}
