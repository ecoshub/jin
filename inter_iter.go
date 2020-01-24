package jint

import "strconv"

func IterateArray(json []byte, callback func([]byte, error) bool, path ... string) error{
	if len(json) == 0 {
		callback(nil, BAD_JSON_ERROR(0))
		return BAD_JSON_ERROR(0)
	}
	offset := 0
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _,v := range chars {
		isJsonChar[v] = true
	}
	for space(json[offset]) {
		if offset > len(json) - 1{
			callback(nil, BAD_JSON_ERROR(0))
			return BAD_JSON_ERROR(offset)
		}else{
			offset++
			continue
		}
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
					for k := i - 1 ; k > 0 ; k -- {
						if json[k] != 92 {
							if (i - 1 - k) % 2 == 0 {
								inQuote = !inQuote
								break
							}else{
								break
							}
						}
						continue
					}
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
						if level < 1 {
							return INDEX_OUT_OF_RANGE_ERROR()
						}
						if level < 2 {
							// trim spaces from beginning 
							for space(json[start]) {
								if start > len(json) - 1{
									callback(nil, BAD_JSON_ERROR(start))
									return BAD_JSON_ERROR(start)
								}
								start++
							}
							if json[start] != 91 && json[start] != 123 {
								// if a quoted value then strip quotes
								if json[start] == 34 {
									end := 0
									for j := start + 1;  j < len(json) ; j ++ {
										curr := json[j]
										if  curr == 92 {
											continue
										}
										// find ending quote
										// quote
										if curr == 34 {
											// just interested with json chars. Other wise continue.
											end = j
											break
										}
									}
									callback(json[start + 1:end], nil)
									return nil
								}else{							
									end := 0
									for j := start + 1;  j < len(json) ; j ++ {
										curr := json[j]
										// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
										if space(curr) || curr == 44 || curr == 93 || curr == 125 {
											if offset == j {
												callback(nil, BAD_JSON_ERROR(j))
												return BAD_JSON_ERROR(j)
											}
											end = j
											break
										}
									}
									callback(json[start:end], nil)
									return nil
								}
							}
							end := 0
							for j := i - 1;  j > start ; j -- {
								curr := json[j]
								// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
								if curr == 93 || curr == 125 {
									if offset == j {
										callback(nil, BAD_JSON_ERROR(j))
										return BAD_JSON_ERROR(j)
									}
									end = j + 1
									break
								}
							}
							callback(json[start:end], nil)
							return nil
						}
						level--
						continue
					}
					if level == 1 {
						if curr == 44 {
							// trim spaces from beginning 
							for space(json[start]) {
								if start > len(json) - 1{
									callback(nil, BAD_JSON_ERROR(start))
									return BAD_JSON_ERROR(start)
								}
								start++
							}
							if json[start] != 91 && json[start] != 123 {
								// if a quoted value then strip quotes
								if json[start] == 34 {
									end := 0
									for j := start + 1;  j < len(json) ; j ++ {
										curr := json[j]
										if curr == 92 {
											continue
										}
										// find ending quote
										// quote
										if curr == 34 {
											// just interested with json chars. Other wise continue.
											end = j
											break
										}
									}
									if !callback(json[start + 1:end], nil) {
										return nil
									}
									start = i + 1
									continue
								}else{							
									end := 0
									for j := start + 1;  j < len(json) ; j ++ {
										curr := json[j]
										// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
										if space(curr) || curr == 44 || curr == 93 || curr == 125 {
											if offset == j {
												if !callback(nil, BAD_JSON_ERROR(j)) {
													return nil
												}
											}
											end = j
											break
										}
									}
									if !callback(json[start:end], nil) {
										return nil
									}
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
										if !callback(nil, BAD_JSON_ERROR(j)) {
											return nil
										}
									}
									end = j + 1
									break
								}
							}
							if !callback(json[start:end], nil) {
								return nil
							}
							start = i + 1
							continue
						}
						continue
					}
					continue
				}
			}
			callback(nil, BAD_JSON_ERROR(offset))
			return BAD_JSON_ERROR(offset)
		}else{
			callback(nil, ARRAY_EXPECTED_ERROR())
			return ARRAY_EXPECTED_ERROR()
		}
	}
	currentPath := path[0]
	braceType := json[offset]
	for k := 0 ; k < len(path) ; k ++ {
		if braceType == 91 {
			arrayIndex, err := strconv.Atoi(currentPath)
			if err != nil {
				callback(nil, INDEX_EXPECTED_ERROR())
				return INDEX_EXPECTED_ERROR()
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
						for k := i - 1 ; k > 0 ; k -- {
							if json[k] != 92 {
								if (i - 1 - k) % 2 == 0 {
									inQuote = !inQuote
									break
								}else{
									break
								}
							}
							continue
						}
						continue
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
								break
							}
							level++
							continue
						}
						if curr == 93 || curr == 125 {
							if level < 1 {
								return INDEX_OUT_OF_RANGE_ERROR()
							}
							level--
							continue
						}
						if !found {
							if level == 1 {
								if curr == 44 {
									indexCount++
									if indexCount == arrayIndex {
										found = true
										if k == len(path) - 1{
											offset = i + 1
											break
										}
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
				if !found {
					callback(nil, INDEX_OUT_OF_RANGE_ERROR())
					return INDEX_OUT_OF_RANGE_ERROR()
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
						if level < 1 {
							return INDEX_OUT_OF_RANGE_ERROR()
						}
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
									for k := i - 1 ; k > 0 ; k -- {
										if json[k] != 92 {
											if (i - 1 - k) % 2 == 0 {
												inQuote = !inQuote
												break
											}else{
												break
											}
										}
										continue
									}
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
				return KEY_NOT_FOUND_ERROR()
			}
			isJsonChar[44] = true
		}
	}
	if offset == 0 {
		callback(nil, BAD_JSON_ERROR(0))
		return BAD_JSON_ERROR(0)
	}
	for space(json[offset]) {
		if offset > len(json) - 1{
			callback(nil, BAD_JSON_ERROR(offset))
			return BAD_JSON_ERROR(offset)
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
				for k := i - 1 ; k > 0 ; k -- {
					if json[k] != 92 {
						if (i - 1 - k) % 2 == 0 {
							inQuote = !inQuote
							break
						}else{
							break
						}
					}
					continue
				}
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
								callback(nil, BAD_JSON_ERROR(start))
								return BAD_JSON_ERROR(start)
							}
							start++
						}
						if json[start] != 91 && json[start] != 123 {
							// if a quoted value then strip quotes
							if json[start] == 34 {
								end := 0
								for j := start + 1;  j < len(json) ; j ++ {
									curr := json[j]
									if curr == 92 {
										continue
									}
									// find ending quote
									// quote
									if curr == 34 {
										// just interested with json chars. Other wise continue.
										end = j
										break
									}
								}
								callback(json[start + 1:end], nil)
								return nil
							}else{			
								end := 0
								for j := start + 1;  j < len(json) ; j ++ {
									curr := json[j]
									// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
									if space(curr) || curr == 44 || curr == 93 || curr == 125 {
										if offset == j {
											callback(nil, BAD_JSON_ERROR(j))
											return BAD_JSON_ERROR(j)
										}
										end = j
										break
									}
								}
								callback(json[start:end], nil)
								return nil
							}
						}
						end := 0
						for j := i - 1;  j > start ; j -- {
							curr := json[j]
							// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
							if curr == 93 || curr == 125 {
								if offset == j {
									callback(nil, BAD_JSON_ERROR(j))
									return BAD_JSON_ERROR(j)
								}
								end = j + 1
								break
							}
						}
						callback(json[start:end], nil)
						return nil
					}
					level--
					continue
				}
				if level == 1 {
					if curr == 44 {
						// trim spaces from beginning 
						for space(json[start]) {
							if start > len(json) - 1{
								callback(nil, BAD_JSON_ERROR(start))
								return BAD_JSON_ERROR(start)
							}
							start++
						}
						if json[start] != 91 && json[start] != 123 {
							// if a quoted value then strip quotes
							if json[start] == 34 {
								end := 0
								for j := start + 1;  j < len(json) ; j ++ {
									curr := json[j]
									if curr == 92 {
										continue
									}
									// find ending quote
									// quote
									if curr == 34 {
										// just interested with json chars. Other wise continue.
										end = j
										break
									}
								}
								if !callback(json[start + 1:end], nil) {
									return nil
								}
								start = i + 1
								continue
							}else{
								end := 0
								for j := start + 1;  j < len(json) ; j ++ {
									curr := json[j]
									// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
									if space(curr) || curr == 44 || curr == 93 || curr == 125 {
										if offset == j {
											if !callback(nil, BAD_JSON_ERROR(j)) {
												return nil
											}
										}
										end = j
										break
									}
								}
								if !callback(json[start:end], nil) {
									return nil
								}
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
									if !callback(nil, BAD_JSON_ERROR(j)) {
										return nil
									}
								}
								end = j + 1
								break
							}
						}
						if !callback(json[start:end], nil) {
							return nil
						}
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
		return OBJECT_EXPECTED_ERROR()
	}
	callback(nil, BAD_JSON_ERROR(-1))
	return BAD_JSON_ERROR(-1)
}

func IterateKeyValue(json []byte, callback func([]byte, []byte, error) bool, path ... string) error{
	if len(json) == 0 {
		callback(nil, nil, BAD_JSON_ERROR(0))
		return BAD_JSON_ERROR(0)
	}
	offset := 0
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _,v := range chars {
		isJsonChar[v] = true
	}
	for space(json[offset]) {
		if offset > len(json) - 1{
			callback(nil, nil, BAD_JSON_ERROR(offset))
			return BAD_JSON_ERROR(offset)
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
						if level < 1 {
							callback(nil, nil, INDEX_OUT_OF_RANGE_ERROR())
							return BAD_JSON_ERROR(valueStart)
						}
						if level < 2 {
							// trim spaces from beginning 
							for space(json[valueStart]) {
								if valueStart > len(json) - 1{
									callback(nil, nil, BAD_JSON_ERROR(valueStart))
									return BAD_JSON_ERROR(valueStart)
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
											callback(nil, nil, BAD_JSON_ERROR(j))
										}
										end = j + 1
										break
									}
								}
								callback(key, json[valueStart:end], nil)
								return nil
							}
							// if a quoted value then strip quotes
							if json[valueStart] == 34 {
								end := 0
								for j := valueStart + 1;  j < len(json) ; j ++ {
									curr := json[j]
									if curr == 92 {
										continue
									}
									// find ending quote
									// quote
									if curr == 34 {
										// just interested with json chars. Other wise continue.
										end = j
										break
									}
								}
								callback(key, json[valueStart + 1:end], nil)
								return nil
							}else{							
								end := 0
								for j := valueStart + 1;  j < len(json) ; j ++ {
									curr := json[j]
									// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
									if space(curr) || curr == 44 || curr == 93 || curr == 125 {
										if offset == j {
											callback(nil, nil, BAD_JSON_ERROR(j))
										}
										end = j
										break
									}
								}
								callback(key, json[valueStart:end], nil)
								return nil
							}
						}
						level--
						continue
					}
					if level == 1 {
						if curr == 44 {
							// trim spaces from beginning 
							for space(json[valueStart]) {
								if valueStart > len(json) - 1{
									callback(nil, nil, BAD_JSON_ERROR(valueStart))
									return BAD_JSON_ERROR(valueStart)
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
											if !callback(nil, nil, BAD_JSON_ERROR(j)) {
												return nil
											}
										}
										end = j + 1
										break
									}
								}
								if !callback(key, json[valueStart:end], nil) {
									return nil
								}
								continue
							}
							// if a quoted value then strip quotes
							if json[valueStart] == 34 {
								end := 0
								for j := valueStart + 1;  j < len(json) ; j ++ {
									curr := json[j]
									if curr == 92 {
										continue
									}
									// find ending quote
									// quote
									if curr == 34 {
										// just interested with json chars. Other wise continue.
										end = j
										break
									}
								}
								if !callback(key, json[valueStart + 1:end], nil) {
									return nil
								}
								continue
							}else{							
								end := 0
								for j := valueStart + 1;  j < len(json) ; j ++ {
									curr := json[j]
									// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
									if space(curr) || curr == 44 || curr == 93 || curr == 125 {
										if offset == j {
											callback(nil, nil, BAD_JSON_ERROR(j))
										}
										end = j
										break
									}
								}
								if !callback(key, json[valueStart:end], nil) {
									return nil
								}
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
			callback(nil, nil, BAD_JSON_ERROR(offset))
			return BAD_JSON_ERROR(offset)
		}else{
			callback(nil, nil, OBJECT_EXPECTED_ERROR())
			return OBJECT_EXPECTED_ERROR()
		}
	}
	currentPath := path[0]
	braceType := json[offset]
	for k := 0 ; k < len(path) ; k ++ {
		if braceType == 91 {
			arrayIndex, err := strconv.Atoi(currentPath)
			if err != nil {
				return INDEX_EXPECTED_ERROR()
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
						for k := i - 1 ; k > 0 ; k -- {
							if json[k] != 92 {
								if (i - 1 - k) % 2 == 0 {
									inQuote = !inQuote
									break
								}else{
									break
								}
							}
							continue
						}
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
								break
							}
							level++
							continue
						}
						if curr == 93 || curr == 125 {
							if level < 0 {
								return INDEX_OUT_OF_RANGE_ERROR()
							}
							level--
							continue
						}
						if !found {
							if level == 1 {
								if curr == 44 {
									indexCount++
									if indexCount == arrayIndex {
										found = true
										if k == len(path) - 1{
											offset = i + 1
											break
										}
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
				if !found {
					callback(nil, nil, INDEX_OUT_OF_RANGE_ERROR())
					return INDEX_OUT_OF_RANGE_ERROR()
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
									for k := i - 1 ; k > 0 ; k -- {
										if json[k] != 92 {
											if (i - 1 - k) % 2 == 0 {
												inQuote = !inQuote
												break
											}else{
												break
											}
										}
										continue
									}
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
				return KEY_NOT_FOUND_ERROR()
			}
			isJsonChar[44] = true
		}
	}
	if offset == 0 {
		return BAD_JSON_ERROR(0)
	}
	for space(json[offset]) {
		if offset > len(json) - 1{
			return BAD_JSON_ERROR(offset)
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
								callback(nil, nil, BAD_JSON_ERROR(valueStart))
								return BAD_JSON_ERROR(valueStart)
							}
							valueStart++
						}
						if json[valueStart] == 91 || json[valueStart] == 123 {
							for j := i - 1;  j > start ; j -- {
								curr := json[j]
								// if curreny byte is one of close braces this means end of the value is i
								if curr == 93 || curr == 125 {
									if offset == j {
										callback(nil, nil, BAD_JSON_ERROR(j))
										return BAD_JSON_ERROR(j)
									}
									callback(key, json[valueStart:j + 1], nil)
									return nil
								}
							}
						}
						// if a quoted value then strip quote
						if json[valueStart] == 34 {
							for j := valueStart + 1;  j < len(json) ; j ++ {
								curr := json[j]
								if curr == 92 {
									continue
								}
								// find ending quote
								// quote
								if curr == 34 {
									// just interested with json chars. Other wise continue.
									callback(key, json[valueStart + 1:j], nil)
									return nil
								}
							}
						}else{
							for j := valueStart + 1;  j < len(json) ; j ++ {
								curr := json[j]
								// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
								if space(curr) || curr == 44 || curr == 93 || curr == 125 {
									if offset == j {
										callback(nil, nil, BAD_JSON_ERROR(j))
										return BAD_JSON_ERROR(j)
									}
									callback(key, json[valueStart:j], nil)
									return  nil
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
								callback(nil, nil, BAD_JSON_ERROR(valueStart))
								return BAD_JSON_ERROR(valueStart)
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
										if !callback(nil, nil, BAD_JSON_ERROR(j)) {
											return nil
										}
									}
									end = j + 1
									break
								}
							}
							if !callback(key, json[valueStart:end], nil) {
								return nil
							}
							continue
						}
						// if a quoted value then strip quotes
						if json[valueStart] == 34 {
							end := 0
							for j := valueStart + 1;  j < len(json) ; j ++ {
								curr := json[j]
								if curr == 92 {
									continue
								}
								// find ending quote
								// quote
								if curr == 34 {
									// just interested with json chars. Other wise continue.
									end = j
									break
								}
							}
							if !callback(key, json[valueStart + 1:end], nil) {
								return nil
							}
							continue
						}else{							
							end := 0
							for j := valueStart + 1;  j < len(json) ; j ++ {
								curr := json[j]
								// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
								if space(curr) || curr == 44 || curr == 93 || curr == 125 {
									if offset == j {
										if !callback(nil, nil, BAD_JSON_ERROR(j)) {
											return nil
										}
									}
									end = j
									break
								}
							}
							if !callback(key, json[valueStart:end], nil) {
								return nil
							}
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
		return OBJECT_EXPECTED_ERROR()
	}
	callback(nil, nil, BAD_JSON_ERROR(-1))
	return BAD_JSON_ERROR(-1)
}
