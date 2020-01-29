package jint

import (
	"fmt"
	"strconv"
)

type node struct {
	label    string
	value    []byte
	start    int
	end      int
	valueStored bool
	up       *node
	down     []*node
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
	mapping(json, core)
	pars := parse{core: core, json: json}
	return &pars
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

func (p *parse) Get(path ...string) ([]byte, bool) {
	if len(path) == 0 {
		return nil, false
	}
	curr := p.core.walk(path)
	if curr == nil {
		return []byte{}, false
	}
	return curr.getVal(p.json), true
}

func mapping(json []byte, core *node) {
	if len(json) == 0 {
		return
	}
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _, v := range chars {
		isJsonChar[v] = true
	}
	offset := 0
	for space(json[offset]) {
		if offset > len(json)-1 {
			return
		} else {
			offset++
			continue
		}
	}
	inQuote := false
	last := json[offset]
	lastIndex := -1
	level := 0
	key := ""
	brace := make([]int, 0, 32)
	indexLevel := make([]int, 0, 32)
	path := make([]string, 0, 32)
	if json[offset] == 123 || json[offset] == 91 {
		indexLevel = append(indexLevel, 0)
		brace = append(brace, offset)
		level++
		lastIndex = offset + 1
	} else {
		return
	}
	for i := offset + 1; i < len(json); i++ {
		curr := json[i]
		if !isJsonChar[curr] {
			continue
		}
		if curr == 34 {
			if inQuote {
				for n := i - 1; n > -1; n-- {
					if json[n] != 92 {
						if (i-1-n)%2 == 0 {
							inQuote = !inQuote
							break
						} else {
							break
						}
					}
					continue
				}
				continue
			}else{
				inQuote = !inQuote
				continue
			}
		}
		if inQuote {
			continue
		} else {
			switch curr {
			case 58:
				switch last {
				// { , -> :
				case 123, 44:
					key = string(trim(json[lastIndex:i]))
					break
				}
				break
			case 44:
				switch last {
				// [ : , -> ,
				// middle value area
				case 58:
					core = core.link(key)
					core.start = lastIndex
					core.end = i
					core = core.up
				case 91, 44:
					core = core.link(strconv.Itoa(indexLevel[len(indexLevel)-1]))
					core.start = lastIndex
					core.end = i
					core = core.up
				}
				indexLevel[len(indexLevel)-1]++
				break
			case 93:
				// , -> ]
				// last value area
				switch last {
				case 44, 91:
					core = core.link(strconv.Itoa(indexLevel[len(indexLevel)-1]))
					core.start = lastIndex
					core.end = i
					core = core.up
				}
				if len(path) == 0 {
					return
				}
				core = core.up
				core = core.link(path[len(path)-1])
				core.start = brace[len(brace)-1]
				core.end = i + 1
				core = core.up

				path = path[:len(path)-1]
				indexLevel = indexLevel[:len(indexLevel)-1]
				brace = brace[:len(brace)-1]
				level--
				break
			case 125:
				// : -> }
				// last value area
				switch last {
				case 58:
					core = core.link(key)
					core.start = lastIndex
					core.end = i
					core = core.up
				}
				if len(path) == 0 {
					return
				}
				core = core.up
				core = core.link(path[len(path)-1])
				core.start = brace[len(brace)-1]
				core.end = i + 1
				core = core.up

				path = path[:len(path)-1]
				indexLevel = indexLevel[:len(indexLevel)-1]
				brace = brace[:len(brace)-1]
				level--
				break
			case 91:
				switch last {
				// : -> [
				case 58:
					core = core.link(key)
					path = append(path, key)
				default:
					core = core.link(strconv.Itoa(indexLevel[len(indexLevel)-1]))
					path = append(path, strconv.Itoa(indexLevel[len(indexLevel)-1]))
				}
				indexLevel = append(indexLevel, 0)
				brace = append(brace, i)
				level++
				break
			case 123:
				switch last {
				// : -> {
				case 58:
					core = core.link(key)
					path = append(path, key)
				default:
					core = core.link(strconv.Itoa(indexLevel[len(indexLevel)-1]))
					path = append(path, strconv.Itoa(indexLevel[len(indexLevel)-1]))
				}
				indexLevel = append(indexLevel, 0)
				brace = append(brace, i)
				level++
				break
			}
			last = curr
			lastIndex = i + 1
			continue
		}
	}
}

func trim(json []byte) []byte {
	if len(json) < 1 {
		return nil
	}
	start := 0
	for i := start; i < len(json); i++ {
		if space(json[i]){
			start++
		} else {
			break
		}
	}
	end := len(json) - 1
	for i := end; i > start; i-- {
		if space(json[i]){
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