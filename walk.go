package jin

import "strconv"

func (j *JO) Walk(callback func(k string, v []byte, p []string) (bool, error)) error {
	path := []string{}
	return walk(j.body, path, callback)
}

func Walk(json []byte, callback func(k string, v []byte, p []string) (bool, error)) error {
	path := []string{}
	return walk(json, path, callback)
}

func walk(json []byte, path []string, callback func(k string, v []byte, p []string) (bool, error)) error {
	t, err := GetType(json, path...)
	if err != nil {
		return err
	}
	switch t {
	case TypeObject:
		err := IterateKeyValue(json, func(keyBytes, _ []byte) (bool, error) {
			err = walk(json, append(path, string(keyBytes)), callback)
			if err != nil {
				return false, err
			}
			return true, nil
		}, path...)
		if err != nil {
			return err
		}
	case TypeArray:
		i := 0
		err := IterateArray(json, func(_ []byte) (bool, error) {
			err = walk(json, append(path, strconv.Itoa(i)), callback)
			if err != nil {
				return false, err
			}
			i++
			return true, nil
		}, path...)
		if err != nil {
			return err
		}
	default:
		key := ""
		if len(path) != 0 {
			key = path[len(path)-1]
		}
		val, err := Get(json, path...)
		if err != nil {
			return err
		}
		callback(key, val, path)
	}
	return nil
}
