package jint

func Delete(json []byte, path ...string) ([]byte, error) {
	lenp := len(path)
	if lenp == 0 {
		return json, NULL_PATH_ERROR()
	}
	var currBrace byte
	if lenp > 1 {
		_, valueStart, _, err := core(json, false, path[:lenp-1]...)
		if err != nil {
			return json, err
		}
		currBrace = json[valueStart]
	}
	if lenp == 1 {
		var offset int
		for space(json[offset]) {
			if offset > len(json)-1 {
				return nil, BAD_JSON_ERROR(offset)
			}
			offset++
		}
		currBrace = json[offset]
	}
	var start int
	var end int
	var err error
	if currBrace == 91 {
		_, start, end, err = core(json, false, path...)
		if err != nil {
			return json, err
		}
		if json[end] == 34 {
			end++
			start--
		}
	}
	if currBrace == 123 {
		start, _, end, err = core(json, false, path...)
		if err != nil {
			return json, err
		}
		if json[end] == 34 {
			end++
		}
		start--
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
