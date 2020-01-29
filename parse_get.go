package jint

func (p *parse) Get(path ...string) ([]byte, error) {
	if len(path) == 0 {
		return nil, NULL_PATH_ERROR()
	}
	curr, err := p.core.walk(path)
	if err != nil {
		return nil, err
	}
	return curr.getVal(p.json), nil
}

func (n *node) getVal(json []byte) []byte {
	if n.hasValue {
		return n.value
	}
	if n.start == n.end {
		return []byte{}
	}
	n.value = trim(json[n.start:n.end])
	n.hasValue = true
	return n.value
}
