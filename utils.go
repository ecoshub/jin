package jint

import "strconv"

type scheme struct {
	keys []string
}

func Scheme(keys []string) *scheme {
	return &scheme{keys: keys}
}

func (s *scheme) MakeJson(values []string) []byte {
	return MakeJson(s.keys, values)
}

func replace(json, newValue []byte, start, end int) []byte {
	newJson := make([]byte, 0, len(json)-end+start+len(newValue))
	newJson = append(newJson, json[:start]...)
	newJson = append(newJson, newValue...)
	newJson = append(newJson, json[end:]...)
	return newJson
}

func trimSpace(str string, start, eoe int) string {
	for space(str[start]) {
		start++
	}
	end := start
	for !space(str[end]) && end < eoe {
		end++
	}
	return str[start:end]
}

func compare(json []byte, start, end int, key string) bool {
	if len(key) != end-start {
		return false
	}
	for i := 0; i < len(key); i++ {
		if key[i] != json[start+i] {
			return false
		}
	}
	return true
}

func space(curr byte) bool {
	// space
	if curr == 32 {
		return true
	}
	// tab
	if curr == 9 {
		return true
	}
	// new line NL
	if curr == 10 {
		return true
	}
	// return CR
	if curr == 13 {
		return true
	}
	return false
}

func Flatten(json []byte) []byte {
	newJson := make([]byte, 0, len(json))
	inQuote := false
	for i := 0; i < len(json); i++ {
		curr := json[i]
		if curr == 92 {
			newJson = append(newJson, curr)
			if i+1 < len(json) {
				newJson = append(newJson, json[i+1])
			}
			i++
			continue
		}
		if curr == 34 {
			newJson = append(newJson, curr)
			inQuote = !inQuote
			continue
		}
		if inQuote {
			newJson = append(newJson, curr)
			continue
		} else {
			if !space(curr) {
				newJson = append(newJson, curr)
				continue
			}
		}
	}
	return newJson
}

func createTabs(n int) []byte {
	res := make([]byte, n)
	for i, _ := range res {
		res[i] = 9
	}
	return res
}

func Format(json []byte) []byte {
	json = Flatten(json)
	newJson := make([]byte, 0, len(json))
	inQuote := false
	level := 0
	for i := 0; i < len(json); i++ {
		curr := json[i]
		if curr == 34 {
			newJson = append(newJson, curr)
			if json[i-1] == 92 {
				continue
			}
			inQuote = !inQuote
			continue
		}
		if inQuote {
			newJson = append(newJson, curr)
			continue
		} else {
			if !space(curr) {
				if curr == 91 {
					level++
					// add curr
					newJson = append(newJson, curr)
					// NL
					newJson = append(newJson, 10)
					// tab
					newJson = append(newJson, createTabs(level)...)
					continue
				}
				if curr == 93 {
					level--
					// NL
					newJson = append(newJson, 10)
					// tab
					newJson = append(newJson, createTabs(level)...)
					// add curr
					newJson = append(newJson, curr)
					continue
				}
				if curr == 123 {
					level++
					// add curr
					newJson = append(newJson, curr)
					// NL
					newJson = append(newJson, 10)
					// tab
					newJson = append(newJson, createTabs(level)...)
					continue
				}
				if curr == 125 {
					level--
					// NL
					newJson = append(newJson, 10)
					// tab
					newJson = append(newJson, createTabs(level)...)
					// add curr
					newJson = append(newJson, curr)
					continue
				}
				if curr == 44 {
					newJson = append(newJson, curr)
					// NL
					newJson = append(newJson, 10)
					// tab
					newJson = append(newJson, createTabs(level)...)
					continue
				}
				if curr == 58 {
					newJson = append(newJson, curr)
					// space
					newJson = append(newJson, 32)
					continue
				}
				newJson = append(newJson, curr)
				continue
			}
		}
	}
	return newJson
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

func ParseArray(arr string) []string {
	arr = string(Flatten([]byte(arr)))
	if len(arr) < 2 {
		return []string{}
	}
	if arr[0] == 91 && arr[len(arr)-1] == 93 {
		if len(arr) == 2 {
			return []string{}
		}
		newArray := make([]string, 0, 16)
		start := 1
		inQuote := false
		level := 0
		for i := 0; i < len(arr); i++ {
			curr := arr[i]
			if curr == 92 {
				i++
				continue
			}
			if curr == 34 {
				inQuote = !inQuote
				continue
			}
			if inQuote {
				continue
			} else {
				if curr == 91 || curr == 123 {
					level++
				}
				if curr == 93 || curr == 125 {
					level--
					if curr == 93 {
						if level == 0 {
							val := arr[start:i]
							val = StripQuotes(val)
							newArray = append(newArray, val)
							start = i + 1
							break
						}
					}
				}
				if level == 1 {
					if curr == 44 {
						val := arr[start:i]
						val = StripQuotes(val)
						newArray = append(newArray, val)
						start = i + 1
						continue
					}
				}
			}
		}
		return newArray
	}
	return nil
}

// make private after
func StripQuotes(str string) string {
	if len(str) > 1 {
		if str[0] == 34 && str[len(str)-1] == 34 {
			str = str[1 : len(str)-1]
		}
	}
	return str
}

func StripQuotesByte(str []byte) []byte {
	if len(str) > 1 {
		if str[0] == 34 && str[len(str)-1] == 34 {
			str = str[1 : len(str)-1]
		}
	}
	return str
}
