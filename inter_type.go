package jin

import "strconv"

const (
	TypeArray   string = "array"
	TypeObject  string = "object"
	TypeString  string = "string"
	TypeBoolean string = "boolean"
	TypeNull    string = "null"
	TypeNumber  string = "number"
)

// IsObject is a type control function.
// If path points to an object it will return true, otherwise it will return false.
// In this instance 'object' means everything that has starts and ends with curly brace.
func IsObject(json []byte, path ...string) (bool, error) {
	state, _, err := typeControlCore(json, []byte{91, 123}, true, path...)
	if err != nil {
		return false, err
	}
	return state, nil
}

// IsArray is a type control function.
// If path points to an array it will return true, otherwise it will return false.
// In this instance 'array' means everything that has starts and ends with square brace.
func IsArray(json []byte, path ...string) (bool, error) {
	state, _, err := typeControlCore(json, []byte{91}, true, path...)
	if err != nil {
		return false, err
	}
	return state, nil
}

// IsValue is a type control function.
// If path points to an value it will return true, otherwise it will return false.
// In this instance 'value' means everything that has not starts and ends with any brace.
func IsValue(json []byte, path ...string) (bool, error) {
	state, _, err := typeControlCore(json, []byte{91, 123}, false, path...)
	if err != nil {
		return false, err
	}
	return state, nil
}

// GetType returns the types of value that path has point.
// possible return types 'array' | 'object' | 'string' | 'boolean' | 'null' | 'number'
func GetType(json []byte, path ...string) (string, error) {
	_, start, err := typeControlCore(json, []byte{}, false, path...)
	if err != nil {
		return "", err
	}
	switch json[start] {
	case '[':
		return TypeArray, nil
	case '{':
		return TypeObject, nil
	case '"':
		return TypeString, nil
	}
	val, err := GetString(json, path...)
	if err != nil {
		return "", err
	}
	switch val {
	case "true", "false":
		return TypeBoolean, nil
	case "null":
		return TypeNull, nil
	}
	_, err = strconv.ParseInt(val, 10, 64)
	if err == nil {
		return TypeNumber, nil
	}
	_, err = strconv.ParseFloat(val, 64)
	if err == nil {
		return TypeNumber, nil
	}
	return TypeString, nil
}

// IsEmpty is a control function.
// If path points to an value it will return 'value' string
// If path points to an array that has zero element in it,
// then it will return true, otherwise it will return false.
func IsEmpty(json []byte, path ...string) (bool, error) {
	var start int
	var end int = len(json) - 1
	var err error
	if len(path) == 0 {
		for space(json[start]) {
			if start > len(json)-1 {
				return false, errBadJSON(start)
			}
			start++
			continue
		}
		for space(json[end]) {
			if end < 1 {
				return false, errBadJSON(end)
			}
			end--
			continue
		}
	} else {
		_, start, end, err = core(json, false, path...)
		if err != nil {
			return false, err
		}
		end--
	}
	braceStart := json[start]
	braceEnd := json[end]
	if braceStart == 91 || braceStart == 123 {
		if braceStart+2 != braceEnd {
			return false, errBadJSON(end)
		}
		for i := start + 1; i < end-1; i++ {
			if !space(json[i]) {
				return false, nil
			}
		}
	} else {
		return false, errObjectExpected()
	}
	return true, nil
}

func typeControlCore(json []byte, control []byte, equal bool, path ...string) (bool, int, error) {
	var start int
	var err error
	if len(path) == 0 {
		for space(json[start]) {
			if start > len(json)-1 {
				return false, -1, errBadJSON(start)
			}
			start++
			continue
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
			}
			return false, start, nil
		}
	}
	if equal {
		return false, start, nil
	}
	return true, start, nil
}
