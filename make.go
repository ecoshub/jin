package jint

import "strconv"
import "fmt"

type Scheme struct {
	originalKeys []string
	keys         []string
}

func MakeScheme(keys ...string) *Scheme {
	return &Scheme{keys: keys, originalKeys: keys}
}

func (s *Scheme) MakeJson(values ...interface{}) []byte {
	return MakeJson(s.keys, values)
}

func (s *Scheme) Add(key string) bool {
	for _, k := range s.keys {
		if k == key {
			return false
		}
	}
	s.keys = append(s.keys, key)
	return true
}

func (s *Scheme) Remove(key string) bool {
	newKeys := make([]string, 0, len(s.keys))
	result := false
	for _, k := range s.keys {
		if k != key {
			newKeys = append(newKeys, k)
		} else {
			result = true
		}
	}
	s.keys = newKeys
	return result
}

func (s *Scheme) Save() {
	s.originalKeys = s.keys
}

func (s *Scheme) Restore() {
	s.keys = s.originalKeys
}

func (s *Scheme) Print() {
	fmt.Println("Current  Keys : ", s.keys)
	fmt.Println("Original Keys : ", s.originalKeys)
}

func MakeArray(elements ...[]byte) []byte {
	arr := make([]byte, 0, 128)
	arr = append(arr, 91)
	for _, el := range elements {
		arr = append(arr, el...)
		arr = append(arr, 44)
	}
	arr = arr[:len(arr)-1]
	arr = append(arr, 93)
	return arr
}

func MakeArrayString(values []string) []byte {
	if values == nil {
		return []byte(`[]`)
	}
	js := make([]byte, 0, 128)
	js = append(js, 91)
	for _, v := range values {
		js = append(js, []byte(formatType(v))...)
		js = append(js, 44)
	}
	js = js[:len(js)-1]
	js = append(js, 93)
	return js
}

func MakeEmptyArray() []byte {
	return []byte{91, 93}
}

func MakeArrayInt(values []int) []byte {
	if values == nil {
		return []byte{91, 93}
	}
	js := make([]byte, 0, 128)
	js = append(js, 91)
	for _, v := range values {
		js = append(js, []byte(strconv.Itoa(v))...)
		js = append(js, 44)
	}
	js = js[:len(js)-1]
	js = append(js, 93)
	return js
}

func MakeArrayBool(values []bool) []byte {
	if values == nil {
		return []byte{91, 93}
	}
	js := make([]byte, 0, 128)
	js = append(js, 91)
	for _, v := range values {
		if v == true {
			js = append(js, []byte("true")...)
		} else {
			js = append(js, []byte("false")...)
		}
		js = append(js, 44)
	}
	js = js[:len(js)-1]
	js = append(js, 93)
	return js
}

func MakeArrayFloat(values []float64) []byte {
	if values == nil {
		return []byte{91, 93}
	}
	js := make([]byte, 0, 128)
	js = append(js, 91)
	for _, v := range values {
		num := strconv.FormatFloat(v, 'e', -1, 64)
		start := 0
		for i := 0; i < len(num); i++ {
			if num[i] == 'e' {
				start = i
			}
		}
		exp, _ := strconv.Atoi(num[start+2:])
		if exp == 0 {
			js = append(js, []byte(num[:start])...)
		} else {
			js = append(js, []byte(num)...)
		}
		js = append(js, 44)
	}
	js = js[:len(js)-1]
	js = append(js, 93)
	return js
}

func MakeJsonWithMap(json map[string]string) []byte {
	if json == nil {
		return []byte{123, 125}
	}
	js := make([]byte, 0, 128)
	js = append(js, 123)
	for k, v := range json {
		js = append(js, 34)
		js = append(js, []byte(k)...)
		js = append(js, 34)
		js = append(js, 58)
		js = append(js, []byte(formatType(v))...)
		js = append(js, 44)
	}
	js = js[:len(js)-1]
	js = append(js, 125)
	return js
}

func MakeJson(keys []string, values []interface{}) []byte {
	if len(keys) != len(values) {
		return nil
	}
	if keys == nil {
		return []byte{123, 125}
	}
	js := make([]byte, 0, 128)
	js = append(js, 123)
	for i, k := range keys {
		js = append(js, 34)
		js = append(js, []byte(k)...)
		js = append(js, 34)
		js = append(js, 58)
		js = append(js, []byte(formatType(fmt.Sprintf("%v", values[i])))...)
		js = append(js, 44)
	}
	js = js[:len(js)-1]
	js = append(js, 125)
	return js
}

func MakeJsonString(keys, values []string) []byte {
	if len(keys) != len(values) {
		return nil
	}
	if keys == nil {
		return []byte{123, 125}
	}
	js := make([]byte, 0, 128)
	js = append(js, 123)
	for i, k := range keys {
		js = append(js, 34)
		js = append(js, []byte(k)...)
		js = append(js, 34)
		js = append(js, 58)
		js = append(js, []byte(formatType(values[i]))...)
		js = append(js, 44)
	}
	js = js[:len(js)-1]
	js = append(js, 125)
	return js
}

func MakeEmptyJson() []byte {
	return []byte{123, 125}
}
