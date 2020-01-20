package jint

import "strconv"

func Set(json []byte, newValue []byte, path ... string) ([]byte, error){
	_, start, end, err := Core(json, path...)
	if err != nil {
		return json, err
	}
	return replace(json, newValue, start - 1, end + 1), nil
}

func SetString(json []byte, newValue string, path ... string) ([]byte, error){
	if newValue[0] != 34 && newValue[len(newValue) - 1] != 34 {
		return Set(json, []byte(`"` + newValue + `"`), path...)
	}
	return Set(json, []byte(newValue), path...)
}

func SetInt(json []byte, newValue int, path ... string) ([]byte, error){
	return Set(json, []byte(strconv.Itoa(newValue)), path...)
}

func SetFloat(json []byte, newValue float64, path ... string) ([]byte, error){
	return Set(json, []byte(strconv.FormatFloat(newValue, 'e', -1, 64)), path...)
}

func SetBool(json []byte, newValue bool, path ... string) ([]byte, error){
	if newValue {
		return Set(json, []byte("true"), path...)
	}
	return Set(json, []byte("false"), path...)
}

func SetKey(json []byte, newKey string, path ... string) ([]byte, error){
	if len(newKey) < 1 {
		return json, NULL_NEW_VALUE_ERROR()
	}
	newPath := make([]string, len(path))
	if len(path) == 0 {
		newPath = []string{newKey}
	}else{
		copy(newPath, path[:len(path) - 1])
		newPath[len(newPath) - 1] = newKey
		keyStart, _, _, err := Core(json, path...)
		if err != nil {
			return json, err
		}
		_, _, _, err = Core(json, newPath...)
		if err != nil {
			// key exist error code is 08
			if err.Error() == KEY_NOT_FOUND_ERROR().Error() {
				if keyStart != -1 {
					for i := keyStart ; i < len(json) ; i++ {
						curr := json[i]
						if curr == 34 {
							return replace(json, []byte(newKey), keyStart, i), nil
						}
					}
				}
			}
		}
	}
	return json, KEY_ALREADY_EXIST_ERROR()
}