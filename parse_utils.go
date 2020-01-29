package jint

import "fmt"

type node struct {
	label       string
	value       []byte
	start       int
	end         int
	valueStored bool
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

func Parse(json []byte) *parse {
	core := CreateNode(nil)
	pCore(json, core)
	pars := parse{core: core, json: json}
	return &pars
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

func (n *node) walk(path []string) *node {
	lenp := len(path)
	curr := n
	if len(path) == 0 {
		return curr
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
			return nil
		}
	}
	return curr
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
				fmt.Printf("\t%v %-6v : %v\n", str+string(9492)+" ", d.label, string(d.getVal(json)))
			} else {
				fmt.Printf("%v%v %-6v : %v\n", str, string(9472), d.label, string(d.getVal(json)))
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
