package jin

// Delete() can delete any key-value pair, array value, array, object.
// Path value must be provided,
// otherwise it will provide an error message.
func Delete(json []byte, path ...string) ([]byte, error) {
	lenp := len(path)
	if lenp == 0 {
		return json, nullPathError()
	}
	ks, s, e, err := core(json, false, path...)
	if err != nil {
		return json, err
	}
	start := 0
	if ks == -1 {
		start = s
	} else {
		start = ks - 1
	}
	if json[start-1] == 34 {
		start--
	}
	if json[e] == 34 {
		e++
	}
	var startEdge int
	var endEdge int
	for i := start - 1; i > 0; i-- {
		if !space(json[i]) {
			startEdge = i
			break
		}
	}
	for i := e; i < len(json); i++ {
		if !space(json[i]) {
			endEdge = i
			break
		}
	}
	if (json[startEdge] == 91 || json[startEdge] == 123) && json[startEdge]+2 == json[endEdge] {
		json = replace(json, []byte{}, start, e)
		return json, nil
	}
	if json[endEdge] == 44 {
		json = replace(json, []byte{}, start, endEdge+1)
		return json, nil
	}
	if json[startEdge] == 44 {
		json = replace(json, []byte{}, startEdge, e)
		return json, nil
	}
	return nil, badJSONError(start)
}
