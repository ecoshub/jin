package jint

import "strconv"
import "fmt"

type Scheme struct {
	keys []string
}

func MakeScheme(keys...string) *Scheme {
	return &Scheme{keys: keys}
}

func (s *Scheme) MakeJson(values...interface{}) []byte {
	strValues := make([]string, len(values))
	for i,v := range values {
		strValues[i] = fmt.Sprintf("%v", v)
	}
	return MakeJson(s.keys, strValues)
}

func MakeArray(values []string) []byte {
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

func MakeJson(keys, values []string) []byte {
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

func formatType(val string) string {
	if len(val) > 0 {
		if isBool(val) {
			return val
		}
		if isInt(val) {
			if val[0] == 48 && len(val) > 1 {
				return `"` + val + `"`
			}
			return val
		}
		if isFloat(val) {
			return val
		}
		return `"` + val + `"`
	}
	return `""`
}

func isBool(val string) bool {
	return val == "true" || val == "false"
}

func isFloat(val string) bool {
	_, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return false
	}
	return true
}

func isInt(val string) bool {
	_, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return false
	}
	return true
}
