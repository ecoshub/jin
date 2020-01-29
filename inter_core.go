package jint

import "strconv"
// import "fmt"

// Only this function commented, other Get() and Set() functions based on same logic.
// Do not use with zero length path! no control for that
// Not for public usage
func core(json []byte, justStart bool, path ...string) (int, int, int, error) {
	// null json control.
	if len(json) == 0 {
		return -1, -1, -1, BAD_JSON_ERROR(0)
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
	for _, v := range chars {
		isJsonChar[v] = true
	}
	// trim spaces of start
	for space(json[offset]) {
		// json length overflow control
		if offset > len(json)-1 {
			return -1, -1, -1, BAD_JSON_ERROR(offset)
		} else {
			offset++
			continue
		}
	}
	// braceType determine whether or not search will be a key search or index search
	braceType := json[offset]
	// if last path pointing at a key, that placeholder stores key start point.
	keyStart := -1
	// main iteration off all bytes.
	for k := 0; k < len(path); k++ {
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
				// Inner iteration for brace type search.
				for i := offset; i < len(json); i++ {
					// curr is current byte of reading.
					curr := json[i]
					if !space(curr) {
						// Open brace
						if curr == 123 || curr == 91 {
							// change brace type of next search.
							braceType = curr
							if k != len(path)-1 {
								// If its not last path than change currentPath to next path.
								currentPath = path[k+1]
							}
							// Assign offset to brace index.
							offset = i
							break
						} else {
							if k != len(path)-1 {
								return -1, -1, -1, INDEX_OUT_OF_RANGE_ERROR()
							}
							break
						}
						// Doesn't have to always find a brace. It can be a value. always break after one non-space char.
						break
					}
				}
			} else {
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
				for i := offset; i < len(json); i++ {
					// curr is current byte of reading.
					curr := json[i]
					// Just interested with json chars. Other wise continue.
					if !isJsonChar[curr] {
						continue
					}
					// If current byte is quote
					if curr == 34 {
						for n := i - 1; n > -1; n-- {
							if json[n] != 92 {
								if (i-1-n)%2 == 0 {
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
					} else {
						// Not found before
						if !found {
							// same level with path
							if level == 1 {
								// current byte is comma
								if curr == 44 {
									// Inc index
									indexCount++
									if indexCount == arrayIndex {
										found = true
										for j := i + 1; j < len(json); j++ {
											curr := json[j]
											if !space(curr) {
												if curr != 91 && curr != 123 {
													if k != len(path)-1 {
														return -1, -1, -1, INDEX_OUT_OF_RANGE_ERROR()
													}
													break
												}
												break
											}
										}
										if k == len(path)-1 {
											// last path found, break
											offset = i + 1
											break
										} else {
											continue
										}
										// keep going for find next brace Type.
									}
									continue
								}
							}
						}
						// Open braces
						if curr == 91 || curr == 123 {
							// if found before done with this search
							// break array search scope
							if found {
								offset = i
								braceType = curr
								currentPath = path[k+1]
								break
							} else {
								level++
								continue
							}
						}
						if curr == 93 || curr == 125 {
							// if level is less than 1 it mean index not in this array.
							if level < 2 {
								return -1, -1, -1, INDEX_OUT_OF_RANGE_ERROR()
							} else {
								level--
								continue
							}
						}
					}
				}
				if !found {
					return -1, -1, -1, INDEX_OUT_OF_RANGE_ERROR()
				}
				// Check true for column char again for keep same with first declaration.
				isJsonChar[58] = true
			}
		} else {
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
			for i := offset; i < len(json); i++ {
				// curr is current byte of reading.
				curr := json[i]
				// Just interested with json chars. Other wise continue.
				if !isJsonChar[curr] {
					continue
				}
				// If current byte is quote
				if curr == 34 {
					// If key found no need to determine start and end points.
					if found {
						continue
					}
					// If level not same as path level no need to determine start and end points.
					if level != k+1 {
						continue
					}
					// escape char ccontrol algorithm
					for n := i - 1; n > -1; n-- {
						if json[n] != 92 {
							if (i-1-n)%2 == 0 {
								inQuote = !inQuote
								break
							} else {
								goto cont
							}
						}
						continue
					}
					// If starting new quote that means key starts here
					if inQuote {
						start = i + 1
						continue
					}
					// if quote ends that means key ends here
					end = i
				cont:
					continue
				}
				if inQuote {
					continue
				} else {
					// same level with path
					if level == k+1 {
						// column
						if curr == 58 {
							// comp between current path and key
							// length comp between current path and key
							if len(currentPath) == end-start {
								same := true
								// compare all elements
								for j := 0; j < len(currentPath); j++ {
									if currentPath[j] != json[start+j] {
										same = false
										break
									} else {
										continue
									}
								}
								if same {
									offset = i + 1
									found = true
									// if it is the last path element break
									// and include comma character to json chars.
									if k == len(path)-1 {
										keyStart = start
										break
									} else {
										continue
									}
								}
							}
							// Include comma character to json chars for jump function
							isJsonChar[44] = true
							// exclude column character to json chars for jump function
							isJsonChar[58] = false
							// exclude space character to json chars for jump function
							// jump function start :{} -> ,
							// it is fast travel from column to comma
							// first we need keys
							// for this purpose skipping values.
							// Only need value if key is correct
							for j := i; j < len(json); j++ {
								// curr is current byte of reading.
								curr := json[j]
								// Just interested with json chars. Other wise continue.
								if !isJsonChar[curr] {
									continue
								}
								// Quote
								if curr == 34 {
									// check before char it might be escape char.
									// escape char ccontrol algorithm
									for n := j - 1; n > -1; n-- {
										if json[n] != 92 {
											if (j-1-n)%2 == 0 {
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
								} else {
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
										if level == k+1 {
											// jump i to j
											i = j
											break
										} else {
											continue
										}
									}
									continue
								}

							}
							// exclude comma character to json chars, jump func is ending.
							isJsonChar[44] = false
							// Include column character to json chars, jump func is ending.
							isJsonChar[58] = true
							// Include space character to json chars, jump func is ending.
							continue
						}
					}
					// open square brace
					if curr == 91 {
						// if found and new brace is square brace than
						// next search is array search break loop and
						// update the current path
						if found {
							offset = i
							braceType = curr
							currentPath = path[k+1]
							break
						} else {
							level++
							continue
						}
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
						} else {
							level++
							continue
						}
					}
					// Close brace
					if curr == 93 || curr == 125 {
						level--
						continue
					}
				}
			}
			// key not found return error
			if !found {
				return -1, -1, -1, KEY_NOT_FOUND_ERROR()
			}
			// Include comma character to json chars to restore original.
			isJsonChar[44] = true
			// Include space character to json chars to restore original.
		}
	}
	// this means not search operation has take place
	// it must be some kinda error or bad format
	if offset == 0 {
		return -1, -1, -1, BAD_JSON_ERROR(0)
	}
	// skip spaces from top.
	for space(json[offset]) {
		// json length overflow control
		if offset > len(json)-1 {
			return -1, -1, -1, BAD_JSON_ERROR(offset)
		} else {
			offset++
			continue
		}
	}
	if justStart {
		return keyStart, offset, 0, nil
	}
	// If value starts with open braces
	if json[offset] == 91 || json[offset] == 123 {
		// main level indicator.
		level := 0
		// Quote check flag
		inQuote := false
		for i := offset; i < len(json); i++ {
			// curr is current byte of reading.
			curr := json[i]
			// Just interested with json chars. Other wise continue.
			if !isJsonChar[curr] {
				continue
			}
			if curr == 34 {	
				// check before char it might be escape char.
				// escape char ccontrol algorithm
				for n := i - 1; n > -1; n-- {
					if json[n] != 92 {
						if (i-1-n)%2 == 0 {
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
			} else {
				if curr == 91 || curr == 123 {
					level++
					continue
				}
				if curr == 93 || curr == 125 {
					if level == 1 {
						return keyStart, offset, i + 1, nil
					} else {
						level--
						continue
					}
				}
				continue
			}
			continue
		}
	} else {
		// If value starts with quote
		if json[offset] == 34 {
			for i := offset + 1; i < len(json); i++ {
				curr := json[i]
				if curr == 92 {
					i++
					continue
				} else {
					// find ending quote
					if curr == 34 {
						// just interested with json chars. Other wise continue.
						return keyStart, offset + 1, i, nil
					}
				}
			}
		} else {
			for i := offset; i < len(json); i++ {
				curr := json[i]
				// if curreny byte is space or one of these ',' ']' '}' this means end of the value is i
				if space(curr) || curr == 44 || curr == 93 || curr == 125 {
					if offset == i {
						return -1, -1, -1, EMPTY_ARRAY_ERROR()
					} else {
						return keyStart, offset, i, nil
					}
				}
			}
		}
	}
	// This means not search operation has take place
	// not any formatting operation has take place
	// it must be some kinda bad JSON format
	return -1, -1, -1, BAD_JSON_ERROR(offset)
}