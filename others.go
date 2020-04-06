package jin

// Length function gives the length of any array or object
// that path has pointed.
func Length(json []byte, path ...string) (int, error) {
	if string(json) == "[]" || string(json) == "{}" {
		return 0, nil
	}
	start, end, err := getStartEnd(json, path...)
	if err != nil {
		return -1, err
	}
	length := 0
	level := 0
	inQuote := false
	chars := []byte{34, 44, 91, 93, 123, 125}
	isJSONChar := make([]bool, 256)
	for _, v := range chars {
		isJSONChar[v] = true
	}
	for i := start; i < end; i++ {
		curr := json[i]
		if !isJSONChar[curr] {
			continue
		}
		if curr == 34 {
			for n := i - 1; n > -1; n-- {
				if json[n] != 92 {
					if (i-n)%2 != 0 {
						inQuote = !inQuote
						break
					} else {
						break
					}
				}
				continue
			}
			continue
		}
		if inQuote {
			continue
		}
		if curr == 123 || curr == 91 {
			level++
			continue
		}
		if curr == 125 || curr == 93 {
			level--
			continue
		}
		if curr == 44 && level == 1 {
			length++
		}
	}
	if length == 0 {
		for i := start + 1; i < end-1; i++ {
			if !space(json[i]) {
				return 1, nil
			}
		}
		return 0, nil
	}
	return length + 1, nil
}

// Flatten is tool for formatting json strings.
// flattens indent formations.
func Flatten(json []byte) []byte {
	newJSON := make([]byte, 0, len(json))
	inQuote := false
	for i := 0; i < len(json); i++ {
		curr := json[i]
		if curr == 92 {
			newJSON = append(newJSON, curr)
			if i+1 < len(json) {
				newJSON = append(newJSON, json[i+1])
			}
			i++
			continue
		}
		if curr == 34 {
			newJSON = append(newJSON, curr)
			inQuote = !inQuote
			continue
		}
		if inQuote {
			newJSON = append(newJSON, curr)
			continue
		} else {
			if !space(curr) {
				newJSON = append(newJSON, curr)
				continue
			}
		}
	}
	return newJSON
}

// Indent is tool for formatting JSON strings.
// Adds Indentation to JSON string.
// It uses tab indentation.
func Indent(json []byte) []byte {
	json = Flatten(json)
	newJSON := make([]byte, 0, len(json))
	inQuote := false
	level := 0
	for i := 0; i < len(json); i++ {
		curr := json[i]
		if curr == 34 {
			newJSON = append(newJSON, curr)
			if json[i-1] == 92 {
				continue
			}
			inQuote = !inQuote
			continue
		}
		if inQuote {
			newJSON = append(newJSON, curr)
			continue
		} else {
			if !space(curr) {
				if curr == 91 {
					level++
					// add curr
					newJSON = append(newJSON, curr)
					// NL
					newJSON = append(newJSON, 10)
					// tab
					newJSON = append(newJSON, createTabs(level)...)
					continue
				}
				if curr == 93 {
					level--
					// NL
					newJSON = append(newJSON, 10)
					// tab
					newJSON = append(newJSON, createTabs(level)...)
					// add curr
					newJSON = append(newJSON, curr)
					continue
				}
				if curr == 123 {
					level++
					// add curr
					newJSON = append(newJSON, curr)
					// NL
					newJSON = append(newJSON, 10)
					// tab
					newJSON = append(newJSON, createTabs(level)...)
					continue
				}
				if curr == 125 {
					level--
					// NL
					newJSON = append(newJSON, 10)
					// tab
					newJSON = append(newJSON, createTabs(level)...)
					// add curr
					newJSON = append(newJSON, curr)
					continue
				}
				if curr == 44 {
					newJSON = append(newJSON, curr)
					// NL
					newJSON = append(newJSON, 10)
					// tab
					newJSON = append(newJSON, createTabs(level)...)
					continue
				}
				if curr == 58 {
					newJSON = append(newJSON, curr)
					// space
					newJSON = append(newJSON, 32)
					continue
				}
				newJSON = append(newJSON, curr)
				continue
			}
		}
	}
	return newJSON
}

// ParseArray is a parse function for converting string type arrays to string slices
func ParseArray(arr string) []string {
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
							val = cleanValueString(val)
							newArray = append(newArray, val)
							break
						}
					}
				}
				if level == 1 {
					if curr == 44 {
						val := arr[start:i]
						val = cleanValueString(val)
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
