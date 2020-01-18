package jint


func replace(json, newValue []byte, start, end int) []byte {
	newJson := make([]byte, 0, len(json) - end + start + len(newValue))
	newJson = append(newJson, json[:start]...)
	newJson = append(newJson, newValue...)
	newJson = append(newJson, json[end:]...)
	return newJson
}

func trimSpace(str string, start, eoe int) string {
	for space(str[start]){
		start++
	}
	end := start
	for !space(str[end]) && end < eoe {
		end++
	}
	return str[start:end]
}

func compare(json []byte, start, end int , key string) bool{
	if len(key) != end - start {
		return false
	}
	for i := 0 ; i < len(key) ; i ++ {
		if key[i] != json[start + i] {
			return false
		}
	}
	return true
}

func space(curr byte) bool{
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