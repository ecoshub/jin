package jint

func Delete(json []byte, path ... string) ([]byte, error) {
	lenp := len(path)
	if lenp == 0 {
		return json, NULL_PATH_ERROR()
	}
	var currBrace byte
	if lenp > 1 {
		_, valueStart, _, err := Core(json, path[:lenp - 1]...)
		if err != nil {
			return json, err
		}
		currBrace = json[valueStart]
	}
	if lenp == 1 {
		var offset int
		for space(json[offset]) {
			if offset > len(json) - 1{
				return nil, BAD_JSON_ERROR() 
			}
			offset++
		}
		currBrace = json[offset]
	}
	var start int
	var valueEnd int
	var err error
	if currBrace == 91 {
		_, start, valueEnd, err = Core(json, path...)
		if err != nil {
			return json, err
		}
	}
	if currBrace == 123 {
		start, _, valueEnd, err = Core(json, path...)
		if err != nil {
			return json, err
		}
		start--
	}
	if json[valueEnd] == 34 {
		valueEnd++
		start--
	}
	for i := valueEnd; i < len(json) ; i++ {
		curr := json[i]
		if !space(curr){
			if curr == currBrace + 2 {
				// back comma search
				for j := start - 1; j > -1 ; j -- {
					curr = json[j]
					if !space(curr){
						if json[j] == 44 {
							return replace(json, []byte{}, j, valueEnd), nil
						}else{
							return json, BAD_JSON_ERROR()
						}
					}
				}
				break
			}
			if curr == 44 {
				return replace(json, []byte{}, start, i + 1), nil
			}
			return json, BAD_JSON_ERROR()
		}
	}
	return nil, BAD_JSON_ERROR() 
}