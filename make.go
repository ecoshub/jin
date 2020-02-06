package jin

import "strconv"
import "fmt"

/*
Scheme is a tool for creating non-nested JSONs.
It provides a struct for saveing a JSON scheme for later usage.
Do not access or manipulate this struct.
Please use methods provided for.

This Sections provides examples for all Scheme.Method.

	person := jin.MakeScheme("name","lastname","age")
	eco := person.MakeJson("eco","hub","28")
	fmt.Println(string(eco))
	// Output: {"name":"eco","lastname":"hub","age":28}

	person.Add("ip")
	person.Add("location")
	sheldon := person.MakeJson("Sheldon","Bloom","42","192.168.1.105", "USA")
	fmt.Println(string(sheldon))
	// Output: {"name":"Sheldon","lastname":"Bloom","age":42,"ip":"192.168.1.105","location":"USA"}

	fmt.Println(person.GetCurrentKeys())
	// Output: [name lastname age ip location]
	fmt.Println(person.GetOriginalKeys())
	// Output: [name lastname age]

	person.Remove("location")
	john := person.MakeJson("John","Wiki","28","192.168.1.102")
	fmt.Println(string(john))
	// Output: {"name":"John","lastname":"Wiki","age":28,"ip":"192.168.1.102"}

	// restores original form of scheme
	person.Restore()
	ted := person.MakeJson("ted","stinson","38")
	fmt.Println(string(ted))
	// Output: {"name":"ted","lastname":"stinson","age":38}

	person.Save()
	fmt.Println(person.GetCurrentKeys())
	// Output: [name lastname age ip location]
	fmt.Println(person.GetOriginalKeys())
	// Output: [name lastname age]*/
type Scheme struct {
	originalKeys []string
	keys         []string
}

// MakeScheme is constructor method for creating Scheme's.
// It only need keys for making json.
// Then trigger Schemes with 'MakeJson' method to create JSON.
func MakeScheme(keys ...string) *Scheme {
	return &Scheme{keys: keys, originalKeys: keys}
}

// MakeJson is main creation method for creating JSON's from Schemes.
func (s *Scheme) MakeJson(values ...interface{}) []byte {
	return MakeJson(s.keys, values)
}

// Add adds a new key value to the current scheme.
// If given key is already exists it returns false, otherwise returns true.
func (s *Scheme) Add(key string) bool {
	for _, k := range s.keys {
		if k == key {
			return false
		}
	}
	s.keys = append(s.keys, key)
	return true
}

// Remove removes the key value to the current scheme.
// If given key is not exists it returns false, otherwise returns true.
// More information on Scheme.Add methods example.
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

// Save saves current keys for protect them temporary changes.
// More information on Scheme.Add methods example.
func (s *Scheme) Save() {
	s.originalKeys = s.keys
}

// Restore Schemes original form.
// More information on Scheme.Add methods example.
func (s *Scheme) Restore() {
	s.keys = s.originalKeys
}

// GetOriginalKeys is a simple get function for get Schemes original keys.
// More information on Scheme.Add methods example.
func (s *Scheme) GetOriginalKeys() []string {
	return s.originalKeys
}

// GetCurrentKeys is a simple get function for get Schemes current keys.
// More information on Scheme.Add methods example.
func (s *Scheme) GetCurrentKeys() []string {
	return s.keys
}

// MakeEmptyArray simply creates "[]" this as byte slice.
func MakeEmptyArray() []byte {
	return []byte{91, 93}
}

// MakeArray creates an array formation from given values and returns them as byte slice.
// Do not use any slice/array for paramter.
// It will accept this kind types but won't be able to make valid representation for use!
func MakeArray(elements ...interface{}) []byte {
	if elements == nil {
		return []byte{91, 93}
	}
	js := make([]byte, 0, 128)
	js = append(js, 91)
	for _, e := range elements {
		js = append(js, []byte(formatType(fmt.Sprintf("%v", e)))...)
		js = append(js, 44)
	}
	js = js[:len(js)-1]
	js = append(js, 93)
	return js
}

// MakeArrayString is a variation of MakeArray() func.
// Parameter type must be slice of string.
// For more information look MakeArray() function.
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

// MakeArrayInt is a variation of MakeArray() func.
// Parameter type must be slice of integer.
// For more information look MakeArray() function.
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

// MakeArrayBool is a variation of MakeArray() func.
// Parameter type must be slice of boolean.
// For more information look MakeArray() function.
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

// MakeArrayFloat is a variation of MakeArray() func.
// Parameter type must be slice of float64.
// For more information look MakeArray() function.
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

// MakeArrayBytes is a variation of MakeArray() func.
// Parameter type must be slice of byte.
// For more information look MakeArray() function.
func MakeArrayBytes(values ...[]byte) []byte {
	if values == nil {
		return []byte{91, 93}
	}
	js := make([]byte, 0, 128)
	js = append(js, 91)
	for _, v := range values {
		js = append(js, v...)
		js = append(js, 44)
	}
	js = js[:len(js)-1]
	js = append(js, 93)
	return js
}

// MakeEmptyJson simply creates "{}" this as byte slice.
func MakeEmptyJson() []byte {
	return []byte{123, 125}
}

// MakeJson creates an JSON formation from given key and value slices, and returns them as byte slice.
// Do not use any slice/array for 'values' variable paramter.
// It will accept this kind types but won't be able to make valid representation for use!
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

// MakeJsonString creates an JSON formation from given key and value string slices, and returns them as byte slice.
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

// MakeJsonWithMap creates an JSON formation from given string-string-map, and returns them as byte slice.
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
