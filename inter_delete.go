package jint

func Delete(json []byte, path ...string) ([]byte, error) {
	lenp := len(path)
	if lenp == 0 {
		return json, NULL_PATH_ERROR()
	}
	var currBrace byte
	var start int
	var err error
	if lenp == 1 {
		for space(json[start]) {
			if start > len(json)-1 {
				return nil, BAD_JSON_ERROR(start)
			}
			start++
		}
		currBrace = json[start]
	} else {
		_, start, _, err = core(json, true, path[:lenp-1]...)
		if err != nil {
			return json, err
		}
		currBrace = json[start]
	}
	var end int
	switch currBrace {
	case 91:
		_, start, end, err = core(json, false, path...)
		if err != nil {
			return json, err
		}
		if json[start-1] == 34 && json[end] == 34 {
			end++
			start--
		}
	case 123:
		start, _, end, err = core(json, false, path...)
		if err != nil {
			return json, err
		}
		if json[end] == 34 {
			end++
		}
		start--
	default:
		return nil, BAD_JSON_ERROR(-1)
	}
	for i := end; i < len(json); i++ {
		curr := json[i]
		if !space(curr) {
			if curr == currBrace+2 {
				for j := start - 1; j > -1; j-- {
					curr = json[j]
					if !space(curr) {
						if curr == 44 {
							return replace(json, []byte{}, j, end), nil
						}
						if curr == currBrace {
							return replace(json, []byte{}, j+1, end), nil
						}
						break
					}
				}
				break
			}
			if curr == 44 {
				return replace(json, []byte{}, start, i+1), nil
			}
			break
		}
	}
	return json, BAD_JSON_ERROR(-1)
}
