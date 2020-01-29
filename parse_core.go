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
	hasValue bool
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
	core.label = "0"
	mapping(json, core)
	pars := parse{core: core, json: json}
	// pars.Three(true)
	return &pars
}

func ParseString(json string) *parse {
	arr := []byte(json)
	core := CreateNode(nil)
	core.label = "*"
	mapping(arr, core)
	pars := parse{core: core, json: arr}
	return &pars
}

func (n *node) getVal(json []byte) []byte {
	if n.hasValue {
		return n.value
	}
	start := n.start
	for i := start; i < n.end; i++ {
		if space(json[i]) {
			start++
		} else {
			break
		}
	}
	if json[start] == 34 {
		start++
	}
	end := n.end
	for i := end - 1; i > start; i-- {
		if space(json[i]) {
			end--
		} else {
			break
		}
	}
	if json[end-1] == 34 {
		end--
	}
	n.value = json[start:end]
	n.hasValue = true
	return n.value
}

func (n *node) setVal(val []byte) {
	n.value = val
	n.hasValue = true
}

func (n *node) getIndex() int {
	for i, v := range n.up.down {
		if v.label == n.label {
			return i
		}
	}
	return -1
}

func (n *node) add(label string) *node {
	if n.down == nil {
		new := CreateNode(n)
		new.label = label
		n.attach(new)
		return new
	} else {
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
	fmt.Println("WARNING add()")
	return nil
}

func (n *node) attach(other *node) {
	other.up = n
	n.down = append(n.down, other)
}

func (n *node) deAttach(label string) {
	newDownList := make([]*node, len(n.down)-1)
	count := 0
	for _, v := range n.down {
		if v.label != label {
			newDownList[count] = v
			count++
		}
	}
	n.down = newDownList
}

func (n *node) insert(other *node, index int) {
	lend := len(n.down)
	newDown := make([]*node, lend+1)
	count := 0
	for i := 0; i < lend; i++ {
		if i == index {
			newDown[count] = other
			count++
			newDown[count] = n.down[i]
			count++
		} else {
			newDown[count] = n.down[i]
			count++
		}
	}
	n.down = newDown
	other.up = n
}

func (n *node) setOffset(startIndex, off int) {
	start := false
	for i, d := range n.down {
		if i == startIndex {
			start = true
		}
		if start {
			d.start += off
			d.end += off
			d.hasValue = false
			if len(d.down) > 0 {
				d.setOffset(0, off)
			}
		}
	}
}

func (n *node) setOffsetUp(off int) {
	index := n.getIndex() + 1
	n = n.up
	n.end += off
	start := false
	for i, d := range n.down {
		if i == index {
			start = true
		}
		if start {
			d.start += off
			d.end += off
			d.hasValue = false
			if len(d.down) > 0 {
				d.setOffset(0, off)
			}
		}
	}
	if n.up != nil {
		n.setOffsetUp(off)
	}
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
	val := curr.getVal(p.json)
	if val[0] == byte('"') {
		val = val[1 : len(val)-1]
	}
	return val, true
}

func (p *parse) GetString(path ...string) (string, bool) {
	val, done := p.Get(path...)
	return string(val), done
}

func (p *parse) GetInt(path ...string) (int, bool) {
	val, done := p.GetString(path...)
	num, err := strconv.Atoi(val)
	if err != nil {
		return -1, false
	}
	return num, done
}

func (p *parse) GetBool(path ...string) (bool, bool) {
	val, done := p.GetString(path...)
	if !done {
		return false, false
	}
	if val == "true" {
		return true, true
	}
	if val == "false" {
		return false, true
	}
	return false, false
}

func (p *parse) GetFloat(path ...string) (float64, bool) {
	val, done := p.GetString(path...)
	if !done {
		return 0.0, false
	}
	num, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return -1, false
	}
	return num, done
}

func (p *parse) Set(newVal []byte, path ...string) bool {
	if len(path) == 0 {
		return false
	}
	curr := p.core.down[0].walk(path)
	if curr == nil {
		return false
	}
	val := curr.getVal(p.json)
	oldType := 0 //value
	newType := 0 // value
	if val[0] == byte('{') || val[0] == byte('[') {
		oldType = 1 // json
	}
	if newVal[0] == byte('{') || newVal[0] == byte('[') {
		newType = 1 // json
	}
	newJson := make([]byte, 0, len(p.json)-len(val)+len(newVal))
	newJson = append(newJson, p.json[:curr.start]...)
	newJson = append(newJson, newVal...)
	newJson = append(newJson, p.json[curr.start+len(val):]...)
	p.json = newJson
	newVal = Flatten(newVal)
	delt := 0
	switch oldType {
	case 0:
		switch newType {
		case 0:
			// value to value
			curr.setVal(newVal)
			delt = len(newVal) - len(val)
		case 1:
			// value to json/array
			newCore := CreateNode(nil)
			mapping(newVal, newCore)
			if newCore.down[0] != nil {
				newCore = newCore.down[0]
			}
			newCore.label = curr.label
			newCore.up = curr.up

			index := curr.getIndex()
			curr.up.down[index] = newCore
			delt = len(newVal) - len(val)
			newCore.start += curr.start
			newCore.end += curr.start
			for _, d := range newCore.down {
				d.start += curr.start
				d.end += curr.start
				d.hasValue = false
			}
			newCore.up.setOffset(index+1, delt)
		}
	case 1:
		switch newType {
		case 0:
			// json/array to value
			off := curr.end - curr.start
			delt = len(newVal) - off
			curr.down = []*node{}
			curr.end = curr.start + len(newVal)
			curr.hasValue = false
		case 1:
			// json/array to json/array
			delt = len(newVal) - len(val)
			index := curr.getIndex()
			newCore := CreateNode(nil)
			mapping(newVal, newCore)
			if newCore.down[0] != nil {
				newCore = newCore.down[0]
			}
			newCore.label = curr.label
			curr.up.down[index] = newCore
			newCore.up = curr.up
			newCore.up.setOffset(index, curr.start)

		}
	}
	curr.setOffsetUp(delt)
	return true
}

func (p *parse) SetString(newVal string, path ...string) bool {
	arr := []byte(newVal)
	return p.Set(arr, path...)
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
		}
		if inQuote {
			continue
		} else {
			switch curr {
			case 58:
				switch last {
				// { , -> :
				case 123, 44:
					key = trimSpaceString(json[lastIndex:i])
					break
				}
				break
			case 44:
				switch last {
				// [ : , -> ,
				// middle value area
				case 58:
					core = core.add(key)
					core.start = lastIndex
					core.end = i
					core = core.up
				case 91, 44:
					core = core.add(strconv.Itoa(indexLevel[len(indexLevel)-1]))
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
					core = core.add(strconv.Itoa(indexLevel[len(indexLevel)-1]))
					core.start = lastIndex
					core.end = i
					core = core.up
				}
				if len(path) == 0 {
					return
				}
				core = core.up
				core = core.add(path[len(path)-1])
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
					core = core.add(key)
					core.start = lastIndex
					core.end = i
					core = core.up
				}
				if len(path) == 0 {
					return
				}
				core = core.up
				core = core.add(path[len(path)-1])
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
					core = core.add(key)
					path = append(path, key)
				default:
					core = core.add(strconv.Itoa(indexLevel[len(indexLevel)-1]))
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
					core = core.add(key)
					path = append(path, key)
				default:
					core = core.add(strconv.Itoa(indexLevel[len(indexLevel)-1]))
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

func trimSpaceString(json []byte) string {
	if len(json) < 1 {
		return ""
	}
	start := 0
	for i := start; i < len(json); i++ {
		if space(json[i]) || json[i] == 34 {
			start++
		} else {
			break
		}
	}
	end := len(json) - 1
	for i := end; i > start; i-- {
		if space(json[i]) || json[i] == 34 {
			end--
		} else {
			break
		}
	}
	return string(json[start : end+1])
}
