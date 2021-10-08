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

// Parser is provides a struct for saving a JSON as nodes.
// Do not access or manipulate this struct.
// Please use methods provided for.
type Parser struct {
	core *node
	json []byte
}

func createNode(up *node) *node {
	Node := node{up: up}
	return &Node
}

func packKeyValue(label string, value []byte) []byte {
	byteOutput := make([]byte, 0, 16)
	byteOutput = append(byteOutput, 34)
	byteOutput = append(byteOutput, []byte(label)...)
	byteOutput = append(byteOutput, 34)
	byteOutput = append(byteOutput, 58)
	byteOutput = append(byteOutput, value...)
	return byteOutput
}

func packValue(value []byte) []byte {
	byteOutput := make([]byte, 0, 16)
	byteOutput = append(byteOutput, value...)
	return byteOutput
}

func (n *node) dive(bytes []byte) []byte {
	temp := make([]byte, 0, 32)
	temp = append(temp, n.value[0])
	for _, d := range n.down {
		val := d.value
		if len(d.down) != 0 {
			val = d.dive(bytes)
		}
		if d.up.value[0] == 123 {
			temp = append(temp, packKeyValue(d.label, val)...)
		}
		if d.up.value[0] == 91 {
			temp = append(temp, packValue(val)...)
		}
		temp = append(temp, 44)
	}
	if len(temp) == 0 {
		return nil
	}
	temp = temp[:len(temp)-1]
	temp = append(temp, n.value[0]+2)
	bytes = append(bytes, temp...)
	return bytes
}

// Parse is constructor method for creating Parsers.
func Parse(json []byte) (*Parser, error) {
	core := &node{up: nil}
	err := pCore(json, core)
	if err != nil {
		return nil, err
	}
	if core.down == nil {
		return nil, errBadJSON(0)
	}
	core = core.down[0]
	pars := Parser{core: core, json: json}
	return &pars, nil
}

func (n *node) insert(up *node, index int) error {
	lend := len(up.down)
	if lend != 0 {
		if lend-1 < index {
			return errIndexOutOfRange()
		}
		for i := index; i < lend; i++ {
			up.down[i].label = strconv.Itoa(i + 1)
		}
		n.label = strconv.Itoa(index)
		n.up = up
		return nil
	}
	return errIndexOutOfRange()
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
		return nil, errKeyNotFound(p)
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

// Tree is a simple tool for visualizing JSONs as semi-tree formation.
// This function returns that form as string.
func (p *Parser) Tree() string {
	str := ""
	p.core.createTree(p.json, 0, false, &str)
	return str
}

// TreeFull is same function with Tree, except it returns tree with values.
func (p *Parser) TreeFull() string {
	str := ""
	p.core.createTree(p.json, 0, true, &str)
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
				*str += fmt.Sprintf("\t%v %-6v : %v\n", tab+string(rune(9492))+" ", d.label, string(d.value))
			} else {
				*str += fmt.Sprintf("%v%v %-6v : %v\n", tab, string(rune(9472)), d.label, string(d.value))
			}
		} else {
			if depth != 0 {
				*str += fmt.Sprintf("\t%v %v\n", tab+string(rune(9492))+" ", d.label)
			} else {
				*str += fmt.Sprintf("%v%v %v\n", tab, string(rune(9472)), d.label)
			}
		}
		if len(d.down) > 0 {
			d.createTree(json, depth+1, withValues, str)
		}
	}
}
