package jsoninterpreter

import "strconv"

func Set(json []byte, newValue []byte, path ... string) ([]byte, error){
	if len(path) == 0 {
		return nil, NULL_PATH_ERROR()
	}
	if len(newValue) == 0 {
		return nil, NULL_NEW_VALUE_ERROR()
	}
	offset := 0
	currentPath := path[0]
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _,v := range chars {
		isJsonChar[v] = true
	}
	for space(json[offset]) {
		offset++
	}
	braceType := json[offset]

	for k := 0 ; k < len(path) ; k ++ {
		if braceType == 91 {
			arrayNumber, err := strconv.Atoi(currentPath)
			if err != nil {
				return json, INDEX_EXPECTED_ERROR()
			}
			if arrayNumber == 0 {
				offset++
				for i := offset; i < len(json) ; i ++ {
					curr := json[i]
					if curr == 123 || curr == 91{
						braceType = curr
						if k != len(path) - 1{
							currentPath = path[k + 1]
						}
						offset = i
						break
					}
					if !space(curr){
						break
					}
				}
			}else{
				level := 0
				inQuote := false
				found := false
				indexCount := 0
				// not interested with column to this level
				isJsonChar[58] = false
				for i := offset ; i < len(json) ; i ++ {
					curr := json[i]
					if !isJsonChar[curr]{
						continue
					}
					if curr == 34 {
						if json[i - 1] == 92 {
							continue
						}
						inQuote = !inQuote
						continue
					}
					if inQuote {
						continue
					}else{
						if curr == 91 || curr == 123{
							if found {
								offset = i
								level++
								braceType = curr
								currentPath = path[k + 1]
								found = false
								break
							}
							level++
							continue
						}
						if curr == 93 || curr == 125 {
							level--
							if level < 1 {
								return nil, INDEX_OUT_OF_RANGE_ERROR()
							}
							continue
						}
						if !found {
							if level == 1 {
								if curr == 44 {
									indexCount++
									if indexCount == arrayNumber {
										offset = i + 1
										if k == len(path) - 1{
											break
										}
										found = true
										continue
									}
									continue
								}
								continue
							}
							continue
						}
						continue
					}
				}
				// interested with column to this level
				isJsonChar[58] = true
			}
		}else{
			inQuote := false
			found := false
			start := 0
			end := 0
			level := k
			// not interested with comma to this level
			isJsonChar[44] = false
			for i := offset ; i < len(json) ; i ++ {
				curr := json[i]
				if !isJsonChar[curr]{
					continue
				}
				if curr == 34 {
					inQuote = !inQuote
					if found {
						continue
					}
					if level != k + 1 {
						continue
					}
					if inQuote {
						start = i + 1
						continue
					}
					end = i
					continue
				}
				if inQuote {
					continue
				}else{
					if curr == 91 {
						if found {
							braceType = curr
							currentPath = path[k + 1]
							break
						}
						level++
						continue
					}
					if curr == 123 {
						if found {
							k++
							level++
							currentPath = path[k]
							found = false
							continue
						}
						level++
						continue
					}
					// close brace
					if curr == 93 || curr == 125 {
						level--
						continue
					}
					// column
					if level == k + 1 {
						if curr == 58 {
							if compare(json, start, end, currentPath) {
								offset = i + 1
								found = true
								if k == len(path) - 1{
									isJsonChar[44] = true
									break
								}else{
									continue
								}
							}
							// interested with comma to this level
							isJsonChar[44] = true
							// not interested with column to this level
							isJsonChar[58] = false
							// little jump algorithm :{} -> ,
							for j := i ;  j < len(json) ; j ++ {
								curr := json[j]
								if !isJsonChar[curr]{
									continue
								}
								// quote
								if curr == 34 {
									if json[j - 1] == 92 {
										continue
									}
									inQuote = !inQuote
									continue
								}
								if inQuote {
									continue
								}else{
									if curr == 91 || curr == 123 {
										level++
										continue
									}
									if curr == 93 || curr == 125 {
										level--
										continue
									}
									// comma
									if curr == 44 {
										if level == k + 1 {
											i = j
											break
										}
										continue
									}
									continue
								}

							}
							// not interested with comma to this level
							isJsonChar[44] = false
							// interested with column to this level
							isJsonChar[58] = true
							continue
						}
						continue
					}
				}
			}
			isJsonChar[44] = true
			if !found {
				return json,KEY_NOT_FOUND_ERROR()
			}
		}
	}
	if offset == 0 {
		return json, BAD_JSON_ERROR()
	}
	for space(json[offset]) {
		offset++
	}
	// starts with { [
	if json[offset] == 91 || json[offset] == 123 {
		level := 0
		inQuote := false
		for i := offset ; i < len(json) ; i ++ {
			curr := json[i]
			if !isJsonChar[curr]{
				continue
			}
			if curr == 34 {
				// escape character control
				if json[i - 1] == 92 {
					continue
				}
				inQuote = !inQuote
				continue
			}
			if inQuote {
				continue
			}else{
				if curr == 91 || curr == 123 {
					level++
				}
				if curr == 93 || curr == 125 {
					level--
					if level == 0 {
						return replace(json, newValue, offset, i + 1), nil
					}
					continue
				}
				continue
			}
			continue
		}
	}else{
		// starts with quote
		if json[offset] == 34 {
			inQuote := false
			for i := offset ;  i < len(json) ; i ++ {
				curr := json[i]
				// quote
				if curr == 34 {
					// escape character control
					if json[i - 1] == 92 {
						continue
					}
					if inQuote {
						return replace(json, newValue, offset, i + 1), nil
					}
					inQuote = !inQuote
					continue
				}
			}
		}else{
			// starts without quote
			for i := offset ;  i < len(json) ; i ++ {
				if isJsonChar[json[i]] {
					// strip others and return value.
					if offset == i {
						return json[offset:i], EMPTY_ARRAY_ERROR()
					}
					return replace(json, newValue, offset, i), nil
				}
			}
		}
	}
	return nil, BAD_JSON_ERROR()
}

func SetString(json []byte, newValue string, path ... string) ([]byte, error){
	if newValue[0] != 34 && newValue[len(newValue) - 1] != 34 {
		return Set(json, []byte(`"` + newValue + `"`), path...)
	}
	return Set(json, []byte(newValue), path...)
}

func SetInt(json []byte, newValue int, path ... string) ([]byte, error){
	return Set(json, []byte(strconv.Itoa(newValue)), path...)
}

func SetFloat(json []byte, newValue float64, path ... string) ([]byte, error){
	return Set(json, []byte(strconv.FormatFloat(newValue, 'e', -1, 64)), path...)
}

func SetBool(json []byte, newValue bool, path ... string) ([]byte, error){
	if newValue {
		return Set(json, []byte("true"), path...)
	}
	return Set(json, []byte("false"), path...)
}

func SetKey(json []byte, newValue []byte, path ... string) ([]byte, error){
	if len(path) == 0 {
		return json, NULL_PATH_ERROR()
	}
	if len(newValue) == 0 {
		return json, NULL_NEW_VALUE_ERROR()
	}
	for _, v := range newValue {
		if v  == 34 {
			return json, BAD_KEY_ERROR()
		}
	}
	offset := 0
	currentPath := path[0]
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _,v := range chars {
		isJsonChar[v] = true
	}
	for space(json[offset]) {
		offset++
	}
	braceType := json[offset]

	for k := 0 ; k < len(path) ; k ++ {
		if braceType == 91 {
			arrayNumber, err := strconv.Atoi(currentPath)
			if err != nil {
				return json, INDEX_EXPECTED_ERROR()
			}
			done := false
			if arrayNumber == 0 {
				offset++
				for i := offset; i < len(json) ; i ++ {
					curr := json[i]
					if curr == 123 {
						braceType = curr
						if k != len(path) - 1{
							currentPath = path[k + 1]
						}
						offset = i
						done = true
						break
					}
					if curr == 91 {
						braceType = curr
						if k != len(path) - 1{
							currentPath = path[k + 1]
						}
						offset = i + 1
						done = true
						break
					}
					if !space(curr){
						done = true
						break
					}
				}
			}else{
				level := 0
				inQuote := false
				found := false
				indexCount := 0
				// not interested with column to this level
				isJsonChar[58] = false
				for i := offset ; i < len(json) ; i ++ {
					curr := json[i]
					if !isJsonChar[curr]{
						continue
					}
					if curr == 34 {
						if json[i - 1] == 92 {
							continue
						}
						inQuote = !inQuote
						continue
					}
					if inQuote {
						continue
					}else{
						if curr == 91 || curr == 123{
							if found {
								level++
								braceType = curr
								currentPath = path[k + 1]
								found = false
								done = true
								break
							}
							level++
							continue
						}
						if curr == 93 || curr == 125 {
							level--
							if level < 1 {
								done = false
								break
							}
							continue
						}
						if !found {
							if level == 1 {
								if curr == 44 {
									indexCount++
									if indexCount == arrayNumber {
										offset = i + 1
										if k == len(path) - 1{
											done = true
											return json, KEY_EXPECTED_ERROR()
										}
										found = true
										continue
									}
									continue
								}
								continue
							}
							continue
						}
						continue
					}
				}
				// interested with column to this level
				isJsonChar[58] = true
			}
			if !done {
				return json, INDEX_OUT_OF_RANGE_ERROR()
			}
		}else{
			inQuote := false
			found := false
			start := 0
			end := 0
			level := k
			// not interested with comma to this level
			isJsonChar[44] = false
			for i := offset ; i < len(json) ; i ++ {
				curr := json[i]
				if !isJsonChar[curr]{
					continue
				}
				if curr == 34 {
					inQuote = !inQuote
					if found {
						continue
					}
					if level != k + 1 {
						continue
					}
					if inQuote {
						start = i + 1
						continue
					}
					end = i
					continue
				}
				if inQuote {
					continue
				}else{
					if curr == 91 {
						if found {
							braceType = curr
							currentPath = path[k + 1]
							break
						}
						level++
						continue
					}
					if curr == 123 {
						if found {
							k++
							level++
							currentPath = path[k]
							found = false
							continue
						}
						level++
						continue
					}
					// close brace
					if curr == 93 || curr == 125 {
						level--
						continue
					}
					// column
					if level == k + 1 {
						if curr == 58 {
							if compare(json, start, end, currentPath) {
								offset = i + 1
								found = true
								if k == len(path) - 1{
									isJsonChar[44] = true
									return replace(json, newValue, start, end), nil
									break
								}else{
									continue
								}
							}
							// interested with comma to this level
							isJsonChar[44] = true
							// not interested with column to this level
							isJsonChar[58] = false
							// little jump algorithm :{} -> ,
							for j := i ;  j < len(json) ; j ++ {
								curr := json[j]
								if !isJsonChar[curr]{
									continue
								}
								// quote
								if curr == 34 {
									if json[j - 1] == 92 {
										continue
									}
									inQuote = !inQuote
									continue
								}
								if inQuote {
									continue
								}else{
									if curr == 91 || curr == 123 {
										level++
										continue
									}
									if curr == 93 || curr == 125 {
										level--
										continue
									}
									// comma
									if curr == 44 {
										if level == k + 1 {
											i = j
											break
										}
										continue
									}
									continue
								}

							}
							// not interested with comma to this level
							isJsonChar[44] = false
							// interested with column to this level
							isJsonChar[58] = true
							continue
						}
						continue
					}
				}
			}
			isJsonChar[44] = true
			if !found {
				return json, KEY_NOT_FOUND_ERROR()
			}
		}
	}
	return json, BAD_JSON_ERROR()
}

func SetStringKey(json []byte, newValue string, path ... string) ([]byte, error){
	return SetKey(json, []byte(newValue), path...)
}