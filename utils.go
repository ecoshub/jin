package jint

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
