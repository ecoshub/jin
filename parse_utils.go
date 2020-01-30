package jint

import "fmt"

type node struct {
	// type 0 is value
	// type 1 is key&Value
	// type 2 is object
	// type 3 is array
	typ         int
	label       string
	value       []byte
	up          *node
	down        []*node
}

type parse struct {
	core *node
	json []byte
}

func CreateNode(up *node) *node {
	Node := node{up: up, down: []*node{}}
	return &Node
}

func Parse(json []byte) (*parse, error) {
	core := CreateNode(nil)
	err := pCore(json, core)
	if err != nil {
		return nil, err
	}
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

func (n *node) walk(path []string) (*node, error) {
	lenp := len(path)
	curr := n
	if len(path) == 0 {
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

func trim(json []byte) []byte {
	if len(json) < 1 {
		return nil
	}
	start := 0
	for i := start; i < len(json); i++ {
		if space(json[i]) {
			start++
		} else {
			break
		}
	}
	end := len(json) - 1
	for i := end; i > start; i-- {
		if space(json[i]) {
			end--
		} else {
			break
		}
	}
	if json[start] == 34 {
		start++
	}
	if json[end] == 34 {
		end--
	}
	return json[start : end+1]
}

// func (n * node) getIndex() int {
// 	for i, v := range n.up.down {
// 		if v.label == n.label {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func (n * node) setOffsetUp(off int){
// 	index := n.getIndex() + 1
// 	n = n.up
// 	n.end += off
// 	start := false
// 	for i, d := range n.down{
// 		if i == index {
// 			start = true
// 		}
// 		if start {
// 			d.start += off
// 			d.end += off
// 			d.hasValue = false
// 			if len(d.down) > 0 {
// 				d.setOffset(0, off)
// 			}
// 		}
// 	}
// 	if n.up != nil {
// 		n.setOffsetUp(off)
// 	}
// }

// func (n * node) setOffset(startIndex, off int){
// 	start := false
// 	for i, d := range n.down{
// 		if i == startIndex {
// 			start = true
// 		}
// 		if start {
// 			d.start += off
// 			d.end += off
// 			d.hasValue = false
// 			if len(d.down) > 0 {
// 				d.setOffset(0, off)
// 			}
// 		}
// 	}
// }


func (p *parse) Three(withValues bool) {
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
