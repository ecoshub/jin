package jint

import "strconv"

func IterateKeyValue(json []byte, callback func([]byte, []byte, error), path ... string) {
	if len(json) == 0 {
		callback(nil, nil, BAD_JSON_ERROR())
		return
	}
	offset := 0
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _,v := range chars {
		isJsonChar[v] = true
	}
	for space(json[offset]) {
		if offset > len(json) - 1{
			callback(nil, nil, BAD_JSON_ERROR())
			return
		}
		offset++
	}
	if len(path) == 0 {
		if json[offset] == 123 {
			inQuote := false
			valueStart := 0
			start := 0
			end := 0
			level := 0
			key := []byte{}
			for i := offset ; i < len(json) ; i ++ {
				curr := json[i]
				if !isJsonChar[curr]{
					continue
				}
				if curr == 34 {
					inQuote = !inQuote
					if level != 1 {
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
					if curr == 91 || curr == 123{
						level++
						continue
					}
					if curr == 93 || curr == 125 {
						level--
						if level < 1 {
							// trim spaces from beginning 
							for space(json[valueStart]) {
								if valueStart > len(json) - 1{
									callback(nil, nil, BAD_JSON_ERROR())
									return
								}
								valueStart++
							}
							if json[valueStart] == 91 || json[valueStart] == 123 {
								end := 0
								for j := i - 1;  j > start ; j -- {
									curr := json[j]
									// if curreny byte is one of close braces this means end of the value is i
									if curr == 93 || curr == 125 {
										if offset == j {
											callback(nil, nil, BAD_JSON_ERROR())
										}
										end = j + 1
										break
									}
								}
								callback(key, json[valueStart:end], nil)
								return
							}
							// if a quoted value then strip quotes
							if json[valueStart] == 34 {
								end := 0
								for j := valueStart + 1;  j < len(json) ; j ++ {
									curr := json[j]
									// find ending quote
									// quote
									if curr == 34 {
										// just interested with json chars. Other wise continue.
										if json[j - 1] == 92 {
											continue
										}
										end = j
										break
									}
								}
								callback(key, json[valueStart + 1:end], nil)
								return
							}else{							
								end := 0
								for j := valueStart + 1;  j < len(json) ; j ++ {
									curr := json[j]
									// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
									if space(curr) || curr == 44 || curr == 93 || curr == 125 {
										if offset == j {
											callback(nil, nil, BAD_JSON_ERROR())
										}
										end = j
										break
									}
								}
								callback(key, json[valueStart:end], nil)
								return
							}
						}
						continue
					}
					if level == 1 {
						if curr == 44 {
							// trim spaces from beginning 
							for space(json[valueStart]) {
								if valueStart > len(json) - 1{
									// callback(nil, BAD_JSON_ERROR())
									callback(nil, nil, BAD_JSON_ERROR())
									return
								}
								valueStart++
							}
							if json[valueStart] == 91 || json[valueStart] == 123 {
								end := 0
								for j := i - 1;  j > start ; j -- {
									curr := json[j]
									// if curreny byte is one of close braces this means end of the value is i
									if curr == 93 || curr == 125 {
										if offset == j {
											callback(nil, nil, BAD_JSON_ERROR())
										}
										end = j + 1
										break
									}
								}
								callback(key, json[valueStart:end], nil)
								continue
							}
							// if a quoted value then strip quotes
							if json[valueStart] == 34 {
								end := 0
								for j := valueStart + 1;  j < len(json) ; j ++ {
									curr := json[j]
									// find ending quote
									// quote
									if curr == 34 {
										// just interested with json chars. Other wise continue.
										if json[j - 1] == 92 {
											continue
										}
										end = j
										break
									}
								}
								callback(key, json[valueStart + 1:end], nil)
								continue
							}else{							
								end := 0
								for j := valueStart + 1;  j < len(json) ; j ++ {
									curr := json[j]
									// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
									if space(curr) || curr == 44 || curr == 93 || curr == 125 {
										if offset == j {
											callback(nil, nil, BAD_JSON_ERROR())
										}
										end = j
										break
									}
								}
								callback(key, json[valueStart:end], nil)
								continue
							}
						}
						if curr == 58{
							valueStart = i + 1
							key = json[start:end]
							// key = string(json[start:end])
							continue
						}
					}
				}
			}
		}else{
			callback(nil, nil, OBJECT_EXPECTED_ERROR())
			return
		}
	}
	currentPath := path[0]
	braceType := json[offset]
	for k := 0 ; k < len(path) ; k ++ {
		if braceType == 91 {
			arrayIndex, err := strconv.Atoi(currentPath)
			if err != nil {
				return
			}
			if arrayIndex == 0 {
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
								braceType = curr
								currentPath = path[k + 1]
								found = false
								break
							}
							level++
							continue
						}
						if curr == 93 || curr == 125 {
							if level < 0 {
								return
							}
							level--
							continue
						}
						if !found {
							if level == 1 {
								if curr == 44 {
									indexCount++
									if indexCount == arrayIndex {
										if k == len(path) - 1{
											offset = i + 1
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
				isJsonChar[58] = true
			}
		}else{
			inQuote := false
			found := false
			start := 0
			end := 0
			level := k
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
					if curr == 93 || curr == 125 {
						level--
						continue
					}
					if level == k + 1 {
						if curr == 58 {
							// comparisson
							if len(currentPath) == end - start {
								same := true
								for j := 0 ; j < len(currentPath) ; j ++ {
									if currentPath[j] != json[start + j] {
										same = false
										break
									}
								}
								if same {
									offset = i + 1
									found = true
									if k == len(path) - 1{
										isJsonChar[44] = true
										// keyStart = start
										break
									}else{
										continue
									}
								}
							}
							isJsonChar[44] = true
							isJsonChar[58] = false
							for j := i ;  j < len(json) ; j ++ {
								curr := json[j]
								if !isJsonChar[curr]{
									continue
								}
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
							isJsonChar[44] = false
							isJsonChar[58] = true
							continue
						}
						continue
					}
				}
			}
			if !found {
				return
			}
			isJsonChar[44] = true
		}
	}
	if offset == 0 {
		return
	}
	for space(json[offset]) {
		if offset > len(json) - 1{
			return
		}
		offset++
	}
	if json[offset] == 123 {
		inQuote := false
		valueStart := offset
		start := offset
		end := offset
		level := 0
		key := []byte{}
		for i := offset ; i < len(json) ; i ++ {
			curr := json[i]
			if !isJsonChar[curr]{
				continue
			}
			if curr == 34 {
				inQuote = !inQuote
				if level != 1 {
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
				if curr == 91 || curr == 123{
					level++
					continue
				}
				if curr == 93 || curr == 125 {
					level--
					if level < 1 {
						// trim spaces from beginning 
						for space(json[valueStart]) {
							if valueStart > len(json) - 1{
								callback(nil, nil, BAD_JSON_ERROR())
								return
							}
							valueStart++
						}
						if json[valueStart] == 91 || json[valueStart] == 123 {
							for j := i - 1;  j > start ; j -- {
								curr := json[j]
								// if curreny byte is one of close braces this means end of the value is i
								if curr == 93 || curr == 125 {
									if offset == j {
										callback(nil, nil, BAD_JSON_ERROR())
									}
									callback(key, json[valueStart:j + 1], nil)
									return
								}
							}
						}
						// if a quoted value then strip quote
						if json[valueStart] == 34 {
							for j := valueStart + 1;  j < len(json) ; j ++ {
								curr := json[j]
								// find ending quote
								// quote
								if curr == 34 {
									// just interested with json chars. Other wise continue.
									if json[j - 1] == 92 {
										continue
									}
									callback(key, json[valueStart + 1:j], nil)
									return
								}
							}
						}else{
							for j := valueStart + 1;  j < len(json) ; j ++ {
								curr := json[j]
								// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
								if space(curr) || curr == 44 || curr == 93 || curr == 125 {
									if offset == j {
										callback(nil, nil, BAD_JSON_ERROR())
									}
									callback(key, json[valueStart:j], nil)
									return 
								}
							}
						}
						continue
					}
					continue
				}
				if level == 1 {
					if curr == 44 {
						// trim spaces from beginning 
						for space(json[valueStart]) {
							if valueStart > len(json) - 1{
								callback(nil, nil, BAD_JSON_ERROR())
								return
							}
							valueStart++
						}
						if json[valueStart] == 91 || json[valueStart] == 123 {
							end := 0
							for j := i - 1;  j > start ; j -- {
								curr := json[j]
								// if curreny byte is one of close braces this means end of the value is i
								if curr == 93 || curr == 125 {
									if offset == j {
										callback(nil, nil, BAD_JSON_ERROR())
									}
									end = j + 1
									break
								}
							}
							callback(key, json[valueStart:end], nil)
							continue
						}
						// if a quoted value then strip quotes
						if json[valueStart] == 34 {
							end := 0
							for j := valueStart + 1;  j < len(json) ; j ++ {
								curr := json[j]
								// find ending quote
								// quote
								if curr == 34 {
									// just interested with json chars. Other wise continue.
									if json[j - 1] == 92 {
										continue
									}
									end = j
									break
								}
							}
							callback(key, json[valueStart + 1:end], nil)
							continue
						}else{							
							end := 0
							for j := valueStart + 1;  j < len(json) ; j ++ {
								curr := json[j]
								// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
								if space(curr) || curr == 44 || curr == 93 || curr == 125 {
									if offset == j {
										callback(nil, nil, BAD_JSON_ERROR())
									}
									end = j
									break
								}
							}
							callback(key, json[valueStart:end], nil)
							continue
						}
					}
					if curr == 58{
						valueStart = i + 1
						key = json[start:end]
						// key = string(json[start:end])
						continue
					}
				}
			}
		}
	}else{
		callback(nil, nil, OBJECT_EXPECTED_ERROR())
		return
	}
	callback(nil, nil, BAD_JSON_ERROR())
	return
}


func IterateArray(json []byte, callback func([]byte, error), path ... string) {
	if len(json) == 0 {
		callback(nil, BAD_JSON_ERROR())
		return
	}
	offset := 0
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _,v := range chars {
		isJsonChar[v] = true
	}
	for space(json[offset]) {
		if offset > len(json) - 1{
			callback(nil, BAD_JSON_ERROR())
			return
		}
		offset++
	}
	if len(path) == 0 {
		if json[offset] == 91 {
			start := offset + 1
			level := 0
			inQuote := false
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
						level++
						continue
					}
					if curr == 93 || curr == 125 {
						if level < 2 {
							// trim spaces from beginning 
							for space(json[start]) {
								if start > len(json) - 1{
									callback(nil, BAD_JSON_ERROR())
									return
								}
								start++
							}
							if json[start] != 91 && json[start] != 123 {
								// if a quoted value then strip quotes
								if json[start] == 34 {
									end := 0
									for j := start + 1;  j < len(json) ; j ++ {
										curr := json[j]
										// find ending quote
										// quote
										if curr == 34 {
											// just interested with json chars. Other wise continue.
											if json[j - 1] == 92 {
												continue
											}
											end = j
											break
											// return
										}
									}
									callback(json[start + 1:end], nil)
									return
								}else{							
									end := 0
									for j := start + 1;  j < len(json) ; j ++ {
										curr := json[j]
										// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
										if space(curr) || curr == 44 || curr == 93 || curr == 125 {
											if offset == j {
												callback(nil, BAD_JSON_ERROR())
												return
											}
											end = j
											break
										}
									}
									callback(json[start:end], nil)
									return
								}
							}
							end := 0
							for j := i - 1;  j > start ; j -- {
								curr := json[j]
								// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
								if curr == 93 || curr == 125 {
									if offset == j {
										callback(nil, BAD_JSON_ERROR())
										return
									}
									end = j + 1
									break
								}
							}
							callback(json[start:end], nil)
							return
						}
						level--
						continue
					}
					if level == 1 {
						if curr == 44 {
							// trim spaces from beginning 
							for space(json[start]) {
								if start > len(json) - 1{
									callback(nil, BAD_JSON_ERROR())
									return
								}
								start++
							}
							if json[start] != 91 && json[start] != 123 {
								// if a quoted value then strip quotes
								if json[start] == 34 {
									end := 0
									for j := start + 1;  j < len(json) ; j ++ {
										curr := json[j]
										// find ending quote
										// quote
										if curr == 34 {
											// just interested with json chars. Other wise continue.
											if json[j - 1] == 92 {
												continue
											}
											end = j
											break
											// return
										}
									}
									callback(json[start + 1:end], nil)
									start = i + 1
									continue
								}else{							
									end := 0
									for j := start + 1;  j < len(json) ; j ++ {
										curr := json[j]
										// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
										if space(curr) || curr == 44 || curr == 93 || curr == 125 {
											if offset == j {
												callback(nil, BAD_JSON_ERROR())
											}
											end = j
											break
										}
									}
									callback(json[start:end], nil)
									start = i + 1
									continue
								}
							}
							end := 0
							for j := i - 1;  j > start ; j -- {
								curr := json[j]
								// if curreny byte is one of close braces this means end of the value is i
								if curr == 93 || curr == 125 {
									if offset == j {
										callback(nil, BAD_JSON_ERROR())
									}
									end = j + 1
									break
								}
							}
							callback(json[start:end], nil)
							start = i + 1
							continue
						}
						continue
					}
					continue
				}
			}
			isJsonChar[58] = true
		}else{
			callback(nil, ARRAY_EXPECTED_ERROR())
			return
		}
	}
	currentPath := path[0]
	braceType := json[offset]
	for k := 0 ; k < len(path) ; k ++ {
		if braceType == 91 {
			arrayIndex, err := strconv.Atoi(currentPath)
			if err != nil {
				callback(nil, INDEX_EXPECTED_ERROR())
				return
			}
			if arrayIndex == 0 {
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
								braceType = curr
								currentPath = path[k + 1]
								found = false
								break
							}
							level++
							continue
						}
						if curr == 93 || curr == 125 {
							if level < 0 {
								callback(nil, INDEX_OUT_OF_RANGE_ERROR())
								return
							}
							level--
							continue
						}
						if !found {
							if level == 1 {
								if curr == 44 {
									indexCount++
									if indexCount == arrayIndex {
										if k == len(path) - 1{
											offset = i + 1
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
				isJsonChar[58] = true
			}
		}else{
			inQuote := false
			found := false
			start := 0
			end := 0
			level := k
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
					if curr == 93 || curr == 125 {
						level--
						continue
					}
					if level == k + 1 {
						if curr == 58 {
							if len(currentPath) == end - start {
								same := true
								for j := 0 ; j < len(currentPath) ; j ++ {
									if currentPath[j] != json[start + j] {
										same = false
										break
									}
								}
								if same {
									offset = i + 1
									found = true
									if k == len(path) - 1{
										isJsonChar[44] = true
										break
									}else{
										continue
									}
								}
							}
							isJsonChar[44] = true
							isJsonChar[58] = false
							for j := i ;  j < len(json) ; j ++ {
								curr := json[j]
								if !isJsonChar[curr]{
									continue
								}
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
							isJsonChar[44] = false
							isJsonChar[58] = true
							continue
						}
						continue
					}
				}
			}
			if !found {
				callback(nil, KEY_NOT_FOUND_ERROR())
				return
			}
			isJsonChar[44] = true
		}
	}
	if offset == 0 {
		callback(nil, BAD_JSON_ERROR())
		return
	}
	for space(json[offset]) {
		if offset > len(json) - 1{
			callback(nil, BAD_JSON_ERROR())
			return
		}
		offset++
	}
	if json[offset] == 91 {
		start := offset + 1
		level := 0
		inQuote := false
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
					level++
					continue
				}
				if curr == 93 || curr == 125 {
					if level < 2 {
						// trim spaces from beginning 
						for space(json[start]) {
							if start > len(json) - 1{
								callback(nil, BAD_JSON_ERROR())
								return
							}
							start++
						}
						if json[start] != 91 && json[start] != 123 {
							// if a quoted value then strip quotes
							if json[start] == 34 {
								end := 0
								for j := start + 1;  j < len(json) ; j ++ {
									curr := json[j]
									// find ending quote
									// quote
									if curr == 34 {
										// just interested with json chars. Other wise continue.
										if json[j - 1] == 92 {
											continue
										}
										end = j
										break
										// return
									}
								}
								callback(json[start + 1:end], nil)
								return
							}else{							
								end := 0
								for j := start + 1;  j < len(json) ; j ++ {
									curr := json[j]
									// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
									if space(curr) || curr == 44 || curr == 93 || curr == 125 {
										if offset == j {
											callback(nil, BAD_JSON_ERROR())
											return
										}
										end = j
										break
									}
								}
								callback(json[start:end], nil)
								return
							}
						}
						end := 0
						for j := i - 1;  j > start ; j -- {
							curr := json[j]
							// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
							if curr == 93 || curr == 125 {
								if offset == j {
									callback(nil, BAD_JSON_ERROR())
									return
								}
								end = j + 1
								break
							}
						}
						callback(json[start:end], nil)
						return
					}
					level--
					continue
				}
				if level == 1 {
					if curr == 44 {
						// trim spaces from beginning 
						for space(json[start]) {
							if start > len(json) - 1{
								callback(nil, BAD_JSON_ERROR())
								return
							}
							start++
						}
						if json[start] != 91 && json[start] != 123 {
							// if a quoted value then strip quotes
							if json[start] == 34 {
								end := 0
								for j := start + 1;  j < len(json) ; j ++ {
									curr := json[j]
									// find ending quote
									// quote
									if curr == 34 {
										// just interested with json chars. Other wise continue.
										if json[j - 1] == 92 {
											continue
										}
										end = j
										break
										// return
									}
								}
								callback(json[start + 1:end], nil)
								start = i + 1
								continue
							}else{							
								end := 0
								for j := start + 1;  j < len(json) ; j ++ {
									curr := json[j]
									// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
									if space(curr) || curr == 44 || curr == 93 || curr == 125 {
										if offset == j {
											callback(nil, BAD_JSON_ERROR())
										}
										end = j
										break
									}
								}
								callback(json[start:end], nil)
								start = i + 1
								continue
							}
						}
						end := 0
						for j := i - 1;  j > start ; j -- {
							curr := json[j]
							// if curreny byte is one of close braces this means end of the value is i
							if curr == 93 || curr == 125 {
								if offset == j {
									callback(nil, BAD_JSON_ERROR())
								}
								end = j + 1
								break
							}
						}
						callback(json[start:end], nil)
						start = i + 1
						continue
					}
					continue
				}
				continue
			}
		}
		isJsonChar[58] = true
	}else{
		callback(nil, OBJECT_EXPECTED_ERROR())
		return
	}
	callback(nil, BAD_JSON_ERROR())
	return
}