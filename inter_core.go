package jsoninterpreter

import "strconv"

// Only this function commented, other Get() and Set() functions based on same logic. 
func Core(json []byte, path ... string) (int, int, int, error){
	// null path control.
	if len(path) == 0 {
		return -1, -1, -1, NULL_PATH_ERROR()
	}
	// null json control.
	if len(json) == 0 {
		return -1, -1, -1, BAD_JSON_ERROR() 
	}
	// main offset track of this search.
	offset := 0
	currentPath := path[0]
	// important chars for json interpretation.
	// 34 = "
	// 44 = ,
	// 58 = :
	// 91 = [
	// 93 = ]
	// 123 = {
	// 125 = }
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	// creating a bool array fill with false
	isJsonChar := make([]bool, 256)
	// only interested chars is true
	for _,v := range chars {
		isJsonChar[v] = true
	}
	// trim spaces of start
	for space(json[offset]) {
		// json length overflow control
		if offset > len(json) - 1{
			return -1, -1, -1, BAD_JSON_ERROR() 
		}
		offset++
	}
	// braceType determine whether or not search will be a key search or index search
	braceType := json[offset]
	// if last path pointing at a key, that placeholder stores key start point.
	keyStart := -1
	// main iteration off all bytes.
	for k := 0 ; k < len(path) ; k ++ {
		// 91 = [, beginning of an array search
		if braceType == 91 {
			// ARRAY SEACH SCOPE
			// path value cast to integer for determine index.
			arrayIndex, err := strconv.Atoi(currentPath)
			if err != nil {
				// braceType and current path type is conflicts.
				return -1, -1, -1, INDEX_EXPECTED_ERROR()
			}
			// zeroth index search.
			if arrayIndex == 0 {
				// Increment offset for not catch current brace.
				offset++
				// Inner iteration for brace search.
				for i := offset; i < len(json) ; i ++ {
					// curr is current byte of reading.
					curr := json[i]
					// Open brace
					if curr == 123 || curr == 91{
						// change brace type of next search.
						braceType = curr
						if k != len(path) - 1{
							// If its not last path than change currentPath to next path.
							currentPath = path[k + 1]
						}
						// Assign offset to brace index.
						offset = i
						// Break the array search scope.
						break
					}
					// Doesn't have to always find a brace. It can be a value.
					if !space(curr){
						break
					}
				}
			}else{
				// Brace level every brace increments the level
				level := 0
				// main in quote flag for determine what is in quote and what is not
				inQuote := false
				// index found flag.
				found := false
				// Index count of current element.
				indexCount := 0
				// Not interested with column char in this search
				isJsonChar[58] = false
				for i := offset ; i < len(json) ; i ++ {
					// curr is current byte of reading.
					curr := json[i]
					// Just interested with json chars. Other wise continue.
					if !isJsonChar[curr]{
						continue
					}
					// If current byte is quote
					if curr == 34 {
						// check before char it might be escape char.
						if json[i - 1] == 92 {
							continue
						}
						// Change inQuote flag to opposite.
						inQuote = !inQuote
						continue
					}
					if inQuote {
						continue
					}else{
						// Open braces
						if curr == 91 || curr == 123{
							// if found before done with this search
							// break array search scope
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
							// if level is less than 1 it mean index not in this array. 
							if level < 0 {
								return -1, -1, -1, INDEX_OUT_OF_RANGE_ERROR()
								// done = false
							}
							level--
							continue
						}
						// Not found before
						if !found {
							// same level with path
							if level == 1 {
								// current byte is comma
								if curr == 44 {
									// Inc index
									indexCount++
									if indexCount == arrayIndex {
										if k == len(path) - 1{
											// last path found, break
											offset = i + 1
											break
										}
										// not last path keep going. For find next brace Type.
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
				// Check true for column char again for keep same with first declaration.
				isJsonChar[58] = true
			}
		}else{
			// KEY SEACH SCOPE
			// main in quote flag for determine what is in quote and what is not.
			inQuote := false
			// Key found flag.
			found := false
			// Key start index.
			start := 0
			// Key end index.
			end := 0
			// Current level.
			level := k
			// Not interested with comma in this search
			isJsonChar[44] = false
			for i := offset ; i < len(json) ; i ++ {
				// curr is current byte of reading.
				curr := json[i]
				// Just interested with json chars. Other wise continue.
				if !isJsonChar[curr]{
					continue
				}
				// If current byte is quote
				if curr == 34 {
					// change inQuote flag to opposite.
					inQuote = !inQuote
					// If key found no need to determine start and end points.
					if found {
						continue
					}
					// If level not same as path level no need to determine start and end points.
					if level != k + 1 {
						continue
					}
					// If starting new quote that means key starts here
					if inQuote {
						start = i + 1
						continue
					}
					// if quote ends that means key ends here
					end = i
					continue
				}
				if inQuote {
					continue
				}else{
					// open square brace
					if curr == 91 {
						// if found and new brace is square brace than 
						// next search is array search break loop and
						// update the current path 
						if found {
							braceType = curr
							currentPath = path[k + 1]
							break
						}
						level++
						continue
					}
					if curr == 123 {
						// if found and new brace is curly brace than 
						// next search is key search continue with this loop and
						// update the current path 
						// close found flag for next search.
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
					// Close brace
					if curr == 93 || curr == 125 {
						level--
						continue
					}
					// same level with path
					if level == k + 1 {
						// column
						if curr == 58 {
							// comp between current path and key
							// length comp between current path and key
							if len(currentPath) == end - start {
								same := true
								// compare all elements
								for j := 0 ; j < len(currentPath) ; j ++ {
									if currentPath[j] != json[start + j] {
										same = false
										break
									}
								}
								if same {
									offset = i + 1
									found = true
									// if it is the last path element break
									// and include comma element to json chars.
									if k == len(path) - 1{
										isJsonChar[44] = true
										keyStart = start
										break
									}else{
										continue
									}
								}
							}
							// Include comma element to json chars for jump function
							isJsonChar[44] = true
							// exclude column element to json chars for jump function
							isJsonChar[58] = false
							// jump function start :{} -> ,
							// it is fast travel from column to comma
							// first we need keys 
							// for this purpose skipping values. 
							// Only need value if key is correct
							for j := i ;  j < len(json) ; j ++ {
								// curr is current byte of reading.
								curr := json[j]
								// Just interested with json chars. Other wise continue.
								if !isJsonChar[curr]{
									continue
								}
								// Quote
								if curr == 34 {
									// check before char it might be escape char.
									if json[j - 1] == 92 {
										continue
									}
									// Change inQuote flag to opposite.
									inQuote = !inQuote
									continue
								}
								if inQuote {
									continue
								}else{
									// This brace conditions for level trace
									// it is necessary to keep level value correct
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
										// level same with path
										if level == k + 1 {
											// jump i to j
											i = j
											break
										}
										continue
									}
									continue
								}

							}
							// exclude comma element to json chars, jump func is ending.
							isJsonChar[44] = false
							// Include column element to json chars, jump func is ending.
							isJsonChar[58] = true
							continue
						}
						continue
					}
				}
			}
			// key not found return error
			if !found {
				return -1, -1, -1, KEY_NOT_FOUND_ERROR()
			}
			// Include comma element to json chars to restore original.
			isJsonChar[44] = true
		}
	}
	// this means not search operation has take place
	// it must be some kinda error or bad format
	if offset == 0 {
		return -1, -1, -1, BAD_JSON_ERROR()
	}
	// skip spaces from top.
	for space(json[offset]) {
		// json length overflow control
		if offset > len(json) - 1{
			return -1, -1, -1, BAD_JSON_ERROR()
		}
		offset++
	}
	// If value starts with open braces
	if json[offset] == 91 || json[offset] == 123 {
		// main level indicator.
		level := 0
		// Quote check flag
		inQuote := false
		for i := offset ; i < len(json) ; i ++ {
			// curr is current byte of reading.
			curr := json[i]
			// Just interested with json chars. Other wise continue.
			if !isJsonChar[curr]{
				continue
			}
			if curr == 34 {
				// // check before char it might be escape char.
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
						// Close brace found in same level with start.
						// Return all of it.
						return keyStart, offset, i + 1, nil
					}
					continue
				}
				continue
			}
			continue
		}
	}else{
		// If value starts with quote
		if json[offset] == 34 {
			for i := offset + 1;  i < len(json) ; i ++ {
				curr := json[i]
				// find ending quote
				// quote
				if curr == 34 {
					// just interested with json chars. Other wise continue.
					if json[i - 1] == 92 {
						continue
					}
					return keyStart, offset + 1, i, nil
				}
			}
		}else{
			for i := offset ;  i < len(json) ; i ++ {
				curr := json[i]
				// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
				if space(curr) || curr == 44 || curr == 93 || curr == 125 {
					if offset == i {
						return -1, -1, -1, EMPTY_ARRAY_ERROR()
					}
					return keyStart, offset, i, nil
				}
			}
		}
	}
	// This means not search operation has take place
	// not any formatting operation has take place
	// it must be some kinda bad JSON format
	return -1, -1, -1,  BAD_JSON_ERROR()
}

// // Only this function commented, other Get() and Set() functions based on same logic. 
// func Offset(json []byte, path ... string) (int, int, error){
// 	// null path control.
// 	if len(path) == 0 {
// 		return -1, -1, NULL_PATH_ERROR()
// 	}
// 	// null json control.
// 	if len(json) == 0 {
// 		return -1, -1, BAD_JSON_ERROR() 
// 	}
// 	// main offset track of this search.
// 	offset := 0
// 	currentPath := path[0]
// 	// important chars for json interpretation.
// 	// 34 = "
// 	// 44 = ,
// 	// 58 = :
// 	// 91 = [
// 	// 93 = ]
// 	// 123 = {
// 	// 125 = }
// 	chars := []byte{34, 44, 58, 91, 93, 123, 125}
// 	// creating a bool array fill with false
// 	isJsonChar := make([]bool, 256)
// 	// only interested chars is true
// 	for _,v := range chars {
// 		isJsonChar[v] = true
// 	}
// 	// trim spaces of start
// 	for space(json[offset]) {
// 		// json length overflow control
// 		if offset > len(json) - 1{
// 			return -1, -1, BAD_JSON_ERROR() 
// 		}
// 		offset++
// 	}
// 	// braceType determine whether or not search will be a key search or index search
// 	braceType := json[offset]
// 	// if last path pointing at a key, that placeholder stores key start point.
// 	keyStart := -1
// 	// main iteration off all bytes.
// 	for k := 0 ; k < len(path) ; k ++ {
// 		// 91 = [, beginning of an array search
// 		if braceType == 91 {
// 			// ARRAY SEACH SCOPE
// 			// path value cast to integer for determine index.
// 			arrayIndex, err := strconv.Atoi(currentPath)
// 			if err != nil {
// 				// braceType and current path type is conflicts.
// 				return -1, -1, INDEX_EXPECTED_ERROR()
// 			}
// 			// zeroth index search.
// 			if arrayIndex == 0 {
// 				// Increment offset for not catch current brace.
// 				offset++
// 				// Inner iteration for brace search.
// 				for i := offset; i < len(json) ; i ++ {
// 					// curr is current byte of reading.
// 					curr := json[i]
// 					// Open brace
// 					if curr == 123 || curr == 91{
// 						// change brace type of next search.
// 						braceType = curr
// 						if k != len(path) - 1{
// 							// If its not last path than change currentPath to next path.
// 							currentPath = path[k + 1]
// 						}
// 						// Assign offset to brace index.
// 						offset = i
// 						// Break the array search scope.
// 						break
// 					}
// 					// Doesn't have to always find a brace. It can be a value.
// 					if !space(curr){
// 						break
// 					}
// 				}
// 			}else{
// 				// Brace level every brace increments the level
// 				level := 0
// 				// main in quote flag for determine what is in quote and what is not
// 				inQuote := false
// 				// index found flag.
// 				found := false
// 				// Index count of current element.
// 				indexCount := 0
// 				// Not interested with column char in this search
// 				isJsonChar[58] = false
// 				for i := offset ; i < len(json) ; i ++ {
// 					// curr is current byte of reading.
// 					curr := json[i]
// 					// Just interested with json chars. Other wise continue.
// 					if !isJsonChar[curr]{
// 						continue
// 					}
// 					// If current byte is quote
// 					if curr == 34 {
// 						// check before char it might be escape char.
// 						if json[i - 1] == 92 {
// 							continue
// 						}
// 						// Change inQuote flag to opposite.
// 						inQuote = !inQuote
// 						continue
// 					}
// 					if inQuote {
// 						continue
// 					}else{
// 						// Open braces
// 						if curr == 91 || curr == 123{
// 							// if found before done with this search
// 							// break array search scope
// 							if found {
// 								offset = i
// 								braceType = curr
// 								currentPath = path[k + 1]
// 								found = false
// 								break
// 							}
// 							level++
// 							continue
// 						}
// 						if curr == 93 || curr == 125 {
// 							// if level is less than 1 it mean index not in this array. 
// 							if level < 0 {
// 								return -1, -1, INDEX_OUT_OF_RANGE_ERROR()
// 								// done = false
// 							}
// 							level--
// 							continue
// 						}
// 						// Not found before
// 						if !found {
// 							// same level with path
// 							if level == 1 {
// 								// current byte is comma
// 								if curr == 44 {
// 									// Inc index
// 									indexCount++
// 									if indexCount == arrayIndex {
// 										if k == len(path) - 1{
// 											// last path found, break
// 											offset = i + 1
// 											break
// 										}
// 										// not last path keep going. For find next brace Type.
// 										found = true
// 										continue
// 									}
// 									continue
// 								}
// 								continue
// 							}
// 							continue
// 						}
// 						continue
// 					}
// 				}
// 				// Check true for column char again for keep same with first declaration.
// 				isJsonChar[58] = true
// 			}
// 		}else{
// 			// KEY SEACH SCOPE
// 			// main in quote flag for determine what is in quote and what is not.
// 			inQuote := false
// 			// Key found flag.
// 			found := false
// 			// Key start index.
// 			start := 0
// 			// Key end index.
// 			end := 0
// 			// Current level.
// 			level := k
// 			// Not interested with comma in this search
// 			isJsonChar[44] = false
// 			for i := offset ; i < len(json) ; i ++ {
// 				// curr is current byte of reading.
// 				curr := json[i]
// 				// Just interested with json chars. Other wise continue.
// 				if !isJsonChar[curr]{
// 					continue
// 				}
// 				// If current byte is quote
// 				if curr == 34 {
// 					// change inQuote flag to opposite.
// 					inQuote = !inQuote
// 					// If key found no need to determine start and end points.
// 					if found {
// 						continue
// 					}
// 					// If level not same as path level no need to determine start and end points.
// 					if level != k + 1 {
// 						continue
// 					}
// 					// If starting new quote that means key starts here
// 					if inQuote {
// 						start = i + 1
// 						continue
// 					}
// 					// if quote ends that means key ends here
// 					end = i
// 					continue
// 				}
// 				if inQuote {
// 					continue
// 				}else{
// 					// open square brace
// 					if curr == 91 {
// 						// if found and new brace is square brace than 
// 						// next search is array search break loop and
// 						// update the current path 
// 						if found {
// 							braceType = curr
// 							currentPath = path[k + 1]
// 							break
// 						}
// 						level++
// 						continue
// 					}
// 					if curr == 123 {
// 						// if found and new brace is curly brace than 
// 						// next search is key search continue with this loop and
// 						// update the current path 
// 						// close found flag for next search.
// 						if found {
// 							k++
// 							level++
// 							currentPath = path[k]
// 							found = false
// 							continue
// 						}
// 						level++
// 						continue
// 					}
// 					// Close brace
// 					if curr == 93 || curr == 125 {
// 						level--
// 						continue
// 					}
// 					// same level with path
// 					if level == k + 1 {
// 						// column
// 						if curr == 58 {
// 							// comp between current path and key
// 							// length comp between current path and key
// 							if len(currentPath) == end - start {
// 								same := true
// 								// compare all elements
// 								for j := 0 ; j < len(currentPath) ; j ++ {
// 									if currentPath[j] != json[start + j] {
// 										same = false
// 										break
// 									}
// 								}
// 								if same {
// 									offset = i + 1
// 									found = true
// 									// if it is the last path element break
// 									// and include comma element to json chars.
// 									if k == len(path) - 1{
// 										isJsonChar[44] = true
// 										keyStart = start
// 										break
// 									}else{
// 										continue
// 									}
// 								}
// 							}
// 							// Include comma element to json chars for jump function
// 							isJsonChar[44] = true
// 							// exclude column element to json chars for jump function
// 							isJsonChar[58] = false
// 							// jump function start :{} -> ,
// 							// it is fast travel from column to comma
// 							// first we need keys 
// 							// for this purpose skipping values. 
// 							// Only need value if key is correct
// 							for j := i ;  j < len(json) ; j ++ {
// 								// curr is current byte of reading.
// 								curr := json[j]
// 								// Just interested with json chars. Other wise continue.
// 								if !isJsonChar[curr]{
// 									continue
// 								}
// 								// Quote
// 								if curr == 34 {
// 									// check before char it might be escape char.
// 									if json[j - 1] == 92 {
// 										continue
// 									}
// 									// Change inQuote flag to opposite.
// 									inQuote = !inQuote
// 									continue
// 								}
// 								if inQuote {
// 									continue
// 								}else{
// 									// This brace conditions for level trace
// 									// it is necessary to keep level value correct
// 									if curr == 91 || curr == 123 {
// 										level++
// 										continue
// 									}
// 									if curr == 93 || curr == 125 {
// 										level--
// 										continue
// 									}
// 									// comma
// 									if curr == 44 {
// 										// level same with path
// 										if level == k + 1 {
// 											// jump i to j
// 											i = j
// 											break
// 										}
// 										continue
// 									}
// 									continue
// 								}

// 							}
// 							// exclude comma element to json chars, jump func is ending.
// 							isJsonChar[44] = false
// 							// Include column element to json chars, jump func is ending.
// 							isJsonChar[58] = true
// 							continue
// 						}
// 						continue
// 					}
// 				}
// 			}
// 			// key not found return error
// 			if !found {
// 				return -1, -1, KEY_NOT_FOUND_ERROR()
// 			}
// 			// Include comma element to json chars to restore original.
// 			isJsonChar[44] = true
// 		}
// 	}
// 	// this means not search operation has take place
// 	// it must be some kinda error or bad format
// 	if offset == 0 {
// 		return -1, -1, BAD_JSON_ERROR()
// 	}
// 	// skip spaces from top.
// 	for space(json[offset]) {
// 		// json length overflow control
// 		if offset > len(json) - 1{
// 			return -1, -1, BAD_JSON_ERROR()
// 		}
// 		offset++
// 	}
// 	// This means not search operation has take place
// 	// not any formatting operation has take place
// 	// it must be some kinda bad JSON format
// 	return keyStart, offset, nil
// }

// func Core(json []byte, path...string) (int,int,int, error){
// 	chars := []byte{34, 44, 58, 91, 93, 123, 125}
// 	// creating a bool array fill with false
// 	isJsonChar := make([]bool, 256)
// 	// only interested chars is true
// 	for _,v := range chars {
// 		isJsonChar[v] = true
// 	}
// 	keyStart, offset, err := Offset(json, path...)
// 	if err != nil {
// 		return -1, -1, -1, err
// 	}
// 	// If value starts with open braces
// 	if json[offset] == 91 || json[offset] == 123 {
// 		// main level indicator.
// 		level := 0
// 		// Quote check flag
// 		inQuote := false
// 		for i := offset ; i < len(json) ; i ++ {
// 			// curr is current byte of reading.
// 			curr := json[i]
// 			// Just interested with json chars. Other wise continue.
// 			if !isJsonChar[curr]{
// 				continue
// 			}
// 			if curr == 34 {
// 				// // check before char it might be escape char.
// 				if json[i - 1] == 92 {
// 					continue
// 				}
// 				inQuote = !inQuote
// 				continue
// 			}
// 			if inQuote {
// 				continue
// 			}else{
// 				if curr == 91 || curr == 123 {
// 					level++
// 				}
// 				if curr == 93 || curr == 125 {
// 					level--
// 					if level == 0 {
// 						// Close brace found in same level with start.
// 						// Return all of it.
// 						return keyStart, offset, i + 1, nil
// 					}
// 					continue
// 				}
// 				continue
// 			}
// 			continue
// 		}
// 	}else{
// 		// If value starts with quote
// 		if json[offset] == 34 {
// 			inQuote := false
// 			for i := offset ;  i < len(json) ; i ++ {
// 				curr := json[i]
// 				// quote
// 				// find ending quote
// 				if curr == 34 {
// 					// just interested with json chars. Other wise continue.
// 					if json[i - 1] == 92 {
// 						continue
// 					}
// 					if inQuote {
// 						// Strip quotes and return.
// 						return keyStart, offset + 1, i, nil
// 					}
// 					inQuote = !inQuote
// 					continue
// 				}
// 			}
// 		}else{
// 			for i := offset ;  i < len(json) ; i ++ {
// 				curr := json[i]
// 				// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
// 				if space(curr) || curr == 44 || curr == 93 || curr == 125 {
// 					if offset == i {
// 						return -1, -1, -1, EMPTY_ARRAY_ERROR()
// 					}
// 					return keyStart, offset, i, nil
// 				}
// 			}
// 		}
// 	}
// 	// This means not search operation has take place
// 	// not any formatting operation has take place
// 	// it must be some kinda bad JSON format
// 	return -1, -1, -1, BAD_JSON_ERROR()
// }