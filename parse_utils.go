package jint

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

type parse struct {
	core *node
	json []byte
}

func CreateNode(up *node) *node {
	Node := node{up: up, down: []*node{}}
	if up != nil {
		up.down = append(up.down, &Node)
	}
	return &Node
}

func Parse(json []byte) (*parse, error) {
	core := CreateNode(nil)
	err := pCore(json, core)
	if err != nil {
		return nil, err
	}
	if core.down == nil {
		return nil, BAD_JSON_ERROR(0)
	}
	core = core.down[0]
	pars := parse{core: core, json: json}
	return &pars, nil
}

func (n *node) link(label string) *node {
	if n.down == nil {
		new := CreateNode(n)
		new.label = label
		n.attach(new)
		return new
	}
	for _, d := range n.down {
		if d.label == label {
			return d
		}
	}
	new := CreateNode(n)
	new.label = label
	n.attach(new)
	return new
}

func (n *node) attach(other *node) {
	other.up = n
	n.down = append(n.down, other)
}

func (n *node) insert(up *node, index int) error {
	lend := len(up.down)
	if lend != 0 {
		if lend-1 < index {
			return INDEX_OUT_OF_RANGE_ERROR()
		}
		for i := index; i < lend; i++ {
			up.down[i].label = strconv.Itoa(i + 1)
		}
		n.label = strconv.Itoa(index)
		n.up = up
		return nil
	}
	return INDEX_OUT_OF_RANGE_ERROR()
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
	lenp := len(path)
	curr := n
	if lenp == 0 {
		return curr, nil
	}
	for i := 0; i < lenp; i++ {
		exists := false
		for _, down := range curr.down {
			if down.label == path[i] {
				curr = down
				exists = true
			}
		}
		if !exists {
			return nil, KEY_NOT_FOUND_ERROR()
		}
	}
	return curr, nil
}

func (n *node) getIndex() int {
	for i, v := range n.up.down {
		if v.label == n.label {
			return i
		}
	}
	return -1
}

func (p *parse) Tree(withValues bool) {
	p.core.recursivePrint(p.json, 0, withValues)
}

func (n *node) recursivePrint(json []byte, depth int, withValues bool) {
	for _, d := range n.down {
		str := ""
		for i := 0; i < depth-1; i++ {
			str += fmt.Sprintf("\t")
		}
		if withValues {
			if depth != 0 {
				fmt.Printf("\t%v %-6v : %v\n", str+string(9492)+" ", d.label, string(d.value))
			} else {
				fmt.Printf("%v%v %-6v : %v\n", str, string(9472), d.label, string(d.value))
			}
		} else {
			if depth != 0 {
				fmt.Printf("\t%v %v\n", str+string(9492)+" ", d.label)
			} else {
				fmt.Printf("%v%v %v\n", str, string(9472), d.label)
			}
		}
		if len(d.down) > 0 {
			d.recursivePrint(json, depth+1, withValues)
		}
	}
}
