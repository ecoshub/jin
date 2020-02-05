package jin

import "strconv"

func Set(json []byte, newValue []byte, path ...string) ([]byte, error) {
	if len(path) == 0 {
		return json, ERROR_NULL_PATH()
	}
	_, start, end, err := core(json, false, path...)
	if err != nil {
		return json, err
	}
	if json[start-1] == 34 && json[end] == 34 {
		return replace(json, newValue, start-1, end+1), nil
	}
	return replace(json, newValue, start, end), nil
}

func SetString(json []byte, newValue string, path ...string) ([]byte, error) {
	if newValue[0] != 34 && newValue[len(newValue)-1] != 34 {
		return Set(json, []byte(`"`+newValue+`"`), path...)
	}
	return Set(json, []byte(newValue), path...)
}

func SetInt(json []byte, newValue int, path ...string) ([]byte, error) {
	return Set(json, []byte(strconv.Itoa(newValue)), path...)
}

func SetFloat(json []byte, newValue float64, path ...string) ([]byte, error) {
	return Set(json, []byte(strconv.FormatFloat(newValue, 'e', -1, 64)), path...)
}

func SetBool(json []byte, newValue bool, path ...string) ([]byte, error) {
	if newValue {
		return Set(json, []byte("true"), path...)
	}
	return Set(json, []byte("false"), path...)
}

func SetKey(json []byte, newKey string, path ...string) ([]byte, error) {
	if len(newKey) == 0 {
		return json, ERROR_NULL_NEW_VALUE()
	}
	if len(path) == 0 {
		return json, ERROR_NULL_PATH()
	}
	var err error
	var keyStart int
	var start int
	newPath := make([]string, len(path))
	copy(newPath, path[:len(path)-1])
	newPath[len(newPath)-1] = newKey
	_, _, _, err = core(json, false, newPath...)
	if err != nil {
		if err.Error() == ERROR_KEY_NOT_FOUND().Error() {
			keyStart, start, _, err = core(json, false, path...)
			if err != nil {
				return json, err
			}
			for i := keyStart; i < start; i++ {
				curr := json[i]
				if curr == 92 {
					i++
				}
				if curr == 34 {
					return replace(json, []byte(newKey), keyStart, i), nil
				}
			}
			return json, ERROR_BAD_JSON(keyStart)
		}
		return json, ERROR_KEY_EXPECTED()
	}
	return json, ERROR_KEY_ALREADY_EXISTS()
}
