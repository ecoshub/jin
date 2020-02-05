package jin

import "strconv"

func Set(json []byte, newValue []byte, path ...string) ([]byte, error) {
	if len(path) == 0 {
		return json, nullPathError()
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
		return json, nullNewValueError()
	}
	if len(path) == 0 {
		return json, nullPathError()
	}
	var err error
	var keyStart int
	var start int
	newPath := make([]string, len(path))
	copy(newPath, path[:len(path)-1])
	newPath[len(newPath)-1] = newKey
	_, _, _, err = core(json, false, newPath...)
	if err != nil {
		if err.Error() == keyNotFoundError().Error() {
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
			return json, badJSONError(keyStart)
		}
		return json, keyExpectedError()
	}
	return json, keyAlreadyExistsError()
}
