package jint

import "strconv"

func AddKeyValue(json []byte, key string, value []byte, path ... string) ([]byte, error){
	_, _, valueEnd , err := Core(json, path...)
	if err != nil {
		return json, err
	}
	// close curly brace
	if json[valueEnd - 1] == 125 {
		val := []byte(`,"` + key + `":` + string(value))
		json = replace(json, val, valueEnd - 1, valueEnd - 1)
		return json, nil
	}
	return json, OBJECT_EXPECTED_ERROR()
}

func AddValue(json []byte, value []byte, path ... string) ([]byte, error){
	_, _, valueEnd , err := Core(json, path...)
	if err != nil {
		return json, err
	}
	// close square brace
	if json[valueEnd - 1] == 93 {
		val := make([]byte, len(value) + 1)
		val[0] = 44
		copy(val[1:], value)
		json = replace(json, val, valueEnd - 1, valueEnd - 1)
		return json, nil
	}
	return json, OBJECT_EXPECTED_ERROR()
}

func AddKeyValueString(json []byte, key, value string, path ... string) ([]byte, error){
	if value[0] != 34 && value[len(value) - 1] != 34 {
		value = `"` + value + `"`
	}
	return AddKeyValue(json, key, []byte(value), path...)
}

func AddKeyValueInt(json []byte, key string, value int, path ... string) ([]byte, error){
	return AddKeyValue(json, key, []byte(strconv.Itoa(value)), path...)
}

func AddKeyValueFloat(json []byte, key string, value float64, path ... string) ([]byte, error){
	return AddKeyValue(json, key, []byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

func AddKeyValueBool(json []byte, key string, value bool, path ... string) ([]byte, error){
	if value {
		return AddKeyValue(json, key, []byte("true"), path...)
	}
	return AddKeyValue(json, key, []byte("false"), path...)
}

func AddValueString(json []byte, value string, path ... string) ([]byte, error){
	if value[0] != 34 && value[len(value) - 1] != 34 {
		return AddValue(json, []byte(`"` + value + `"`), path...)
	}
	return AddValue(json, []byte(value), path...)
}

func AddValueInt(json []byte, value int, path ... string) ([]byte, error){
	return AddValue(json, []byte(strconv.Itoa(value)), path...)
}

func AddValueFloat(json []byte, value float64, path ... string) ([]byte, error){
	return AddValue(json, []byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

func AddValueBool(json []byte, value bool, path ... string) ([]byte, error){
	if value {
		return AddValue(json, []byte("true"), path...)
	}
	return AddValue(json, []byte("false"), path...)
}

func InsertValue(json []byte, index int, value []byte, path ... string) ([]byte, error){
	_, valueStart, _, err := Core(json, path...)
	if err != nil {
		return json, err
	}
	if json[valueStart] != 91 {
		return json, ARRAY_EXPECTED_ERROR()
	}
	if index < 0 {
		return json, INDEX_OUT_OF_RANGE_ERROR()
	}
	indexStr := strconv.Itoa(index)
	path = append(path, indexStr)
	_, valueStart, _, err = Core(json, path...)
	if err != nil {
		return json, err
	}
	val := make([]byte, len(value) + 1)
	copy(val, value)
	val[len(val) - 1] = 44
	if json[valueStart - 1] == 34 {
		json = replace(json, val, valueStart - 1, valueStart - 1)
		return json, nil
	}
	json = replace(json, val, valueStart, valueStart)
	return json, nil
}

func InsertValueString(json []byte, index int, value string, path ... string) ([]byte, error){
	if value[0] != 34 && value[len(value) - 1] != 34 {
		return InsertValue(json, index, []byte(`"` + value + `"`), path...)
	}
	return InsertValue(json, index, []byte(value), path...)
}

func InsertValueInt(json []byte, index, value int, path ... string) ([]byte, error){
	return InsertValue(json, index, []byte(strconv.Itoa(value)), path...)
}

func InsertValueFloat(json []byte, index int, value float64, path ... string) ([]byte, error){
	return InsertValue(json, index, []byte(strconv.FormatFloat(value, 'e', -1, 64)), path...)
}

func InsertValueBool(json []byte, index int, value bool, path ... string) ([]byte, error){
	if value {
		return InsertValue(json, index, []byte("true"), path...)
	}
	return InsertValue(json, index, []byte("false"), path...)
}