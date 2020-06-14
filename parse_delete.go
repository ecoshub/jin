package jin

// Delete can delete any key-value pair, array value, array, object.
// Path value must be provided,
// otherwise it will provide an error message.
func (p *Parser) Delete(path ...string) error {
	var err error
	var curr *node
	lenp := len(path)
	if lenp == 0 {
		return nullPathError()
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
