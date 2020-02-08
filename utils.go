package jin

import "strconv"

type sequance struct {
	list   []int
	index  int
	length int
}

func makeSeq(length int) *sequance {
	s := sequance{list: make([]int, length), index: 0, length: length}
	return &s
}

func (s *sequance) push(element int) {
	if s.index > s.length-1 {
		newList := make([]int, s.length+4)
		copy(newList, s.list)
		s.list = newList
		s.length = s.length + 4
	}
	s.list[s.index] = element
	s.index++
}

func (s *sequance) pop() int {
	if s.index > -1 {
		s.index--
		return s.list[s.index]
	}
	return 0
}

func (s *sequance) last() int {
	return s.list[s.index-1]
}

func (s *sequance) getlist() []int {
	return s.list[:s.index]
}

func (s *sequance) inc() {
	s.list[s.index-1]++
}

func replace(json, newValue []byte, start, end int) []byte {
	newJSON := make([]byte, 0, len(json)-end+start+len(newValue))
	newJSON = append(newJSON, json[:start]...)
	newJSON = append(newJSON, newValue...)
	newJSON = append(newJSON, json[end:]...)
	return newJSON
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

func createTabs(n int) []byte {
	res := make([]byte, n)
	for i := range res {
		res[i] = 9
	}
	return res
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

// make private after
func stripQuotes(str string) string {
	if len(str) > 1 {
		if str[0] == 34 && str[len(str)-1] == 34 {
			str = str[1 : len(str)-1]
		}
	}
	return str
}

func stripQuotesByte(str []byte) []byte {
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

func trim(str []byte) []byte {
	start := 0
	lens := len(str) - 1
	for space(str[start]) {
		if start > lens {
			break
		}
		start++
	}
	end := lens
	for space(str[end]) {
		if end < 1 {
			break
		}
		end--
	}
	return str[start : end+1]
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
	start := 0
	lens := len(str)
	end := lens - 1
	for i := start; i < lens-1; i++ {
		if !space(str[i]) {
			break
		}
		start++
	}
	for i := end; i > start+1; i-- {
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
