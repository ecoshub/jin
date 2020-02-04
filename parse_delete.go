package jin

func (p *parse) Delete(path ...string) error {
	var err error
	var curr *node
	lenp := len(path)
	if lenp == 0 {
		return NULL_PATH_ERROR()
	}
	curr, err = p.core.walk(path)
	if err != nil {
		return err
	}
	curr.deAttach()
	p.json, err = Delete(p.json, path...)
	if err != nil {
		return err
	}
	for i := 0; i < lenp-1; i++ {
		curr = curr.up
		curr.value, err = Get(p.json, path[:lenp-i-1]...)
		if err != nil {
			return err
		}
	}
	return nil
}
