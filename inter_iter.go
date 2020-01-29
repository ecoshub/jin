package jint

func IterateArray(json []byte, callback func([]byte) bool, path ...string) error {
	var start int
	var end int
	var err error
	if len(path) == 0 {
		for space(json[start]) {
			if start > len(json)-1 {
				return BAD_JSON_ERROR(start)
			} else {
				start++
				continue
			}
		}
	} else {
		_, start, _, err = core(json, true, path...)
		if err != nil {
			return err
		}
	}
	chars := []byte{34, 44, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _, v := range chars {
		isJsonChar[v] = true
	}
	if json[start] == 91 {
		start++
		inQuote := false
		level := 0
		for i := start; i < len(json); i++ {
			curr := json[i]
			if !isJsonChar[curr] {
				continue
			}
			if curr == 34 {
				for k := i - 1; k > 0; k-- {
					if json[k] != 92 {
						if (i-1-k)%2 == 0 {
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
					if level < 1 {
						return nil
					}
					if curr == 93 {
						if level == 0 {
							for j := start; j < i; j++ {
								if !space(json[j]) {
									start = j
									break
								}
							}
							for j := i - 1; j > start; j-- {
								if !space(json[j]) {
									end = j
									break
								}
							}
							if json[start] == 34 && json[end] == 34 {
								callback(json[start+1 : end])
								return nil
							} else {
								callback(json[start : end+1])
								return nil
							}
						}
					}
					level--
					continue
				}
				if level == 0 {
					if curr == 44 {
						end = i - 1
						for j := start; j < i; j++ {
							if !space(json[j]) {
								start = j
								break
							}
						}
						for j := end; j > start; j-- {
							if !space(json[j]) {
								end = j
								break
							}
						}
						if json[start] == 34 && json[end] == 34 {
							if !callback(json[start+1 : end]) {
								return nil
							}
						} else {
							if !callback(json[start : end+1]) {
								return nil
							}
						}
						start = i + 1
						continue
					}
				}
			}
		}
		return nil
	} else {
		return ARRAY_EXPECTED_ERROR()
	}
	return BAD_JSON_ERROR(-1)
}

func IterateKeyValue(json []byte, callback func([]byte, []byte) bool, path ...string) error {
	var start int
	var err error
	var end int
	if len(path) == 0 {
		for space(json[start]) {
			if start > len(json)-1 {
				return BAD_JSON_ERROR(start)
			} else {
				start++
				continue
			}
		}
	} else {
		_, start, _, err = core(json, true, path...)
		if err != nil {
			return err
		}
	}
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _, v := range chars {
		isJsonChar[v] = true
	}
	if json[start] == 123 {
		start++
		keyStart := 0
		keyEnd := 0
		inQuote := false
		level := 0
		var key []byte
		for i := start; i < len(json); i++ {
			curr := json[i]
			if !isJsonChar[curr] {
				continue
			}
			if curr == 34 {
				for k := i - 1; k > 0; k-- {
					if json[k] != 92 {
						if (i-1-k)%2 == 0 {
							inQuote = !inQuote
							break
						} else {
							goto cont
						}
					}
					continue
				}
				if inQuote {
					keyStart = i
					continue
				}
				keyEnd = i
			cont:
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
					if level < 1 {
						return nil
					}
					if curr == 125 {
						if level == 0 {
							for j := start; j < i; j++ {
								if !space(json[j]) {
									start = j
									break
								}
							}
							for j := i - 1; j > start; j-- {
								if !space(json[j]) {
									end = j
									break
								}
							}
							if json[start] == 34 && json[end] == 34 {
								callback(key, json[start+1:end])
								return nil
							} else {
								callback(key, json[start:end+1])
								return nil
							}
						}
					}
					level--
					continue
				}
				if level == 0 {
					if curr == 44 {
						end = i - 1
						for j := start; j < i; j++ {
							if !space(json[j]) {
								start = j
								break
							}
						}
						for j := i - 1; j > start; j-- {
							if !space(json[j]) {
								end = j
								break
							}
						}
						if json[start] == 34 && json[end] == 34 {
							if !callback(key, json[start+1:end]) {
								return nil
							}
						} else {
							if !callback(key, json[start:end+1]) {
								return nil
							}
						}
						continue
					}
					if curr == 58 {
						key = json[keyStart+1 : keyEnd]
						start = i + 1
						continue
					}
				}
			}
		}
		return nil
	} else {
		return OBJECT_EXPECTED_ERROR()
	}
	return BAD_JSON_ERROR(-1)
}
