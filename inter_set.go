package jsoninterpreter

import "strconv"


func Set(json []byte, newValue []byte, path ... string) ([]byte, error){
	_, start, end, err := Core(json, path...)
	if err != nil {
		return json, err
	}
	return replace(json, newValue, start, end), nil
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
	newPath := make([]string, len(path))
	copy(newPath, path[:len(path) - 1])
	newPath[len(newPath) - 1] = newKey
	_, _, _, err := Core(json, newPath...)
	if err != nil {
		// key exist error code is 07
		if err.Error()[11:13] == "07" {
			keyStart, _, _, err := Core(json, path...)
			if err != nil {
				return json, err
			}
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
	return json, KEY_ALREADY_EXIST_ERROR()
}