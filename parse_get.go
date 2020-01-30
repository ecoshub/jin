package jint


func (p *parse) Get(path ...string) ([]byte, error) {
	if len(path) == 0 {
		// later
		return nil, NULL_PATH_ERROR()
	}
	curr, err := p.core.walk(path)
	if err != nil {
		return nil, err
	}
	// if len(curr.down) == 0 {
	// 	return trim(curr.value), nil
	// }else{

	// }
	// return curr.value, nil
	return trim(curr.value), nil
}
