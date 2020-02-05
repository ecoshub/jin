package jin

import (
	"fmt"
	"strconv"
)

type node struct {
	label string
	value []byte
	up    *node
	down  []*node
}

type Parser struct {
	core *node
	json []byte
}

func createNode(up *node) *node {
	Node := node{up: up}
	return &Node
}

func Parse(json []byte) (*Parser, error) {
	core := &node{up: nil}
	err := pCore(json, core)
	if err != nil {
		return nil, err
	}
	if core.down == nil {
		return nil, badJSONError(0)
	}
	core = core.down[0]
	pars := Parser{core: core, json: json}
	return &pars, nil
}

func (n *node) insert(up *node, index int) error {
	lend := len(up.down)
	if lend != 0 {
		if lend-1 < index {
			return indexOutOfRangeError()
		}
		for i := index; i < lend; i++ {
			up.down[i].label = strconv.Itoa(i + 1)
		}
		n.label = strconv.Itoa(index)
		n.up = up
		return nil
	}
	return indexOutOfRangeError()
}

func (n *node) deAttach() {
	newDown := make([]*node, 0, len(n.up.down)-1)
	for _, d := range n.up.down {
		if d.label != n.label {
			newDown = append(newDown, d)
		}
	}
	n.up.down = newDown
}

func (n *node) walk(path []string) (*node, error) {
	for _, p := range path {
		for _, d := range n.down {
			if d.label == p {
				n = d
				goto cont
			}
		}
		return nil, keyNotFoundError()
	cont:
		continue
	}
	return n, nil
}

func (n *node) getIndex() int {
	for i, v := range n.up.down {
		if v.label == n.label {
			return i
		}
	}
	return -1
}

func (p *Parser) Tree(withValues bool) string {
	str := ""
	p.core.createTree(p.json, 0, withValues, &str)
	return str
}

func (n *node) createTree(json []byte, depth int, withValues bool, str *string) {
	for _, d := range n.down {
		tab := ""
		for i := 0; i < depth-1; i++ {
			tab += fmt.Sprintf("\t")
		}
		if withValues {
			if depth != 0 {
				*str += fmt.Sprintf("\t%v %-6v : %v\n", tab+string(9492)+" ", d.label, string(d.value))
			} else {
				*str += fmt.Sprintf("%v%v %-6v : %v\n", tab, string(9472), d.label, string(d.value))
			}
		} else {
			if depth != 0 {
				*str += fmt.Sprintf("\t%v %v\n", tab+string(9492)+" ", d.label)
			} else {
				*str += fmt.Sprintf("%v%v %v\n", tab, string(9472), d.label)
			}
		}
		if len(d.down) > 0 {
			d.createTree(json, depth+1, withValues, str)
		}
	}
}
