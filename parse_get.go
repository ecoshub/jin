package jint

func (p *parse) Get(path ...string) ([]byte, error) {
	if len(path) == 0 {
		return nil, BAD_JSON_ERROR(0)
	}
	curr := p.core.walk(path)
	if curr == nil {
		return []byte{}, nil
	} else {
		return nil, KEY_NOT_FOUND_ERROR()
	}
	return curr.getVal(p.json), nil
}

func (n *node) getVal(json []byte) []byte {
	if n.valueStored {
		return n.value
	}
	if n.start == n.end {
		return []byte{}
	}
	n.value = trim(json[n.start:n.end])
	n.valueStored = true
	return n.value
}
