package jin

import (
	"strconv"
)

// Set sets the value that path has pointed.
// Path can point anything, a key-value pair, a value, an array, an object.
// Path variable can not be null,
// otherwise it will provide an error message.
func Set(json []byte, newValue []byte, path ...string) ([]byte, error) {
	if len(path) == 0 {
		return json, errNullPath()
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

// SetString is a variation of Set() func.
// SetString takes the set value as string.
func SetString(json []byte, newValue string, path ...string) ([]byte, error) {
	return Set(json, []byte(formatType(newValue)), path...)
}

// SetInt is a variation of Set() func.
// SetInt takes the set value as integer.
func SetInt(json []byte, newValue int, path ...string) ([]byte, error) {
	return Set(json, []byte(strconv.Itoa(newValue)), path...)
}

// SetFloat is a variation of Set() func.
// SetFloat takes the set value as float64.
func SetFloat(json []byte, newValue float64, path ...string) ([]byte, error) {
	return Set(json, []byte(strconv.FormatFloat(newValue, 'f', -1, 64)), path...)
}

// SetBool is a variation of Set() func.
// SetBool takes the set value as boolean.
func SetBool(json []byte, newValue bool, path ...string) ([]byte, error) {
	if newValue {
		return Set(json, []byte("true"), path...)
	}
	return Set(json, []byte("false"), path...)
}

// SetKey sets the key value of key-value pair that path has pointed.
// Path must point to an object.
// otherwise it will provide an error message.
// Path variable can not be null,
func SetKey(json []byte, newKey string, path ...string) ([]byte, error) {
	if len(newKey) == 0 {
		return json, errNullKey()
	}
	if len(path) == 0 {
		return json, errNullPath()
	}
	var err error
	var keyStart int
	var start int
	newPath := make([]string, len(path))
	copy(newPath, path[:len(path)-1])
	newPath[len(newPath)-1] = newKey
	_, _, _, err = core(json, false, newPath...)
	if err != nil {
		if ErrEqual(err, ErrCodeKeyNotFound) {
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
			return json, errBadJSON(keyStart)
		}
		return json, err
	}
	return json, errBadJSON(keyStart)
}

// Store sets or adds a new value to the body
// like a real json store event
func Store(json []byte, key string, value []byte, path ...string) ([]byte, error) {
	_, start, end, err := core(json, false, append(path, key)...)
	if err != nil {
		if ErrEqual(err, ErrCodeKeyNotFound) {
			return AddKeyValue(json, key, value, path...)
		}
		return json, err
	}
	if json[start-1] == 34 && json[end] == 34 {
		return replace(json, value, start-1, end+1), nil
	}
	return replace(json, value, start, end), nil
}

// StoreString is a variation of Store() func.
// StoreString takes the Store value as string.
func StoreString(json []byte, key string, newValue string, path ...string) ([]byte, error) {
	return Store(json, key, []byte(formatType(newValue)), path...)
}

// StoreInt is a variation of Store() func.
// StoreInt takes the Store value as integer.
func StoreInt(json []byte, key string, newValue int, path ...string) ([]byte, error) {
	return Store(json, key, []byte(strconv.Itoa(newValue)), path...)
}

// StoreFloat is a variation of Store() func.
// StoreFloat takes the Store value as float64.
func StoreFloat(json []byte, key string, newValue float64, path ...string) ([]byte, error) {
	return Store(json, key, []byte(strconv.FormatFloat(newValue, 'f', -1, 64)), path...)
}

// StoreBool is a variation of Store() func.
// StoreBool takes the Store value as boolean.
func StoreBool(json []byte, key string, newValue bool, path ...string) ([]byte, error) {
	if newValue {
		return Store(json, key, []byte("true"), path...)
	}
	return Store(json, key, []byte("false"), path...)
}
