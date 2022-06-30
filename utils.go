package jin

import (
	"strconv"
	"unsafe"
)

func replace(json, newValue []byte, start, end int) []byte {
	newJSON := make([]byte, 0, len(json)-end+start+len(newValue))
	newJSON = append(newJSON, json[:start]...)
	newJSON = append(newJSON, newValue...)
	newJSON = append(newJSON, json[end:]...)
	return newJSON
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

func createTabs(n int) []byte {
	res := make([]byte, n)
	for i := range res {
		res[i] = 9
	}
	return res
}

// make private after
func stripQuotes(str string) string {
	if len(str) > 1 {
		if str[0] == 34 && str[len(str)-1] == 34 {
			str = str[1 : len(str)-1]
		}
	}
	return str
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
		if val == "null" {
			return val
		}
		start := val[0]
		end := val[len(val)-1]
		if (start == 34 && end == 34) || (start == 91 && end == 93) || (start == 123 && end == 125) {
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
	return err == nil
}

func isInt(val string) bool {
	_, err := strconv.ParseInt(val, 10, 32)
	return err == nil
}

func cleanValueString(str string) string {
	start := 0
	lens := len(str)
	for space(str[start]) && start < lens-1 {
		start++
	}
	end := lens - 1
	for space(str[end]) && end > 1 {
		end--
	}
	if str[start] == 34 && str[end] == 34 {
		start++
		end--
	}
	return str[start : end+1]
}

func cleanValue(str []byte) []byte {
	if len(str) == 0 {
		return str
	}
	start := 0
	lens := len(str)
	end := lens - 1
	for i := start; i < lens-1; i++ {
		if !space(str[i]) {
			break
		}
		start++
	}
	for i := end; i > start; i-- {
		if !space(str[i]) {
			break
		}
		end--
	}
	if str[start] == 34 && str[end] == 34 {
		start++
		end--
	}
	return str[start : end+1]
}

func byteArrayToString(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}

func getStartEnd(json []byte, path ...string) (int, int, error) {
	lenj := len(json)
	if lenj < 2 {
		return -1, -1, errBadJSON(0)
	}
	var err error
	var start int
	var end int
	lenp := len(path)
	if lenp != 0 {
		_, start, end, err = core(json, false, path...)
		if err != nil {
			return -1, -1, err
		}
	} else {
		for space(json[start]) {
			if start > len(json)-1 {
				return -1, -1, errBadJSON(start)
			}
			start++
			continue
		}
		end = lenj - 1
		for space(json[end]) {
			if end < start {
				return -1, -1, errBadJSON(start)
			}
			end--
			continue
		}
	}
	return start, end, nil
}
