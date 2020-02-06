package jin

import "strconv"

func pCore(json []byte, core *node) error {
	lenj := len(json)
	if lenj < 2 {
		return nil
	}
	inQuote := false
	braceList := makeSeq(4)
	indexList := makeSeq(4)
	indexList.Push(0)
	var start int
	var end int
	var key []byte
	var valStart int
	var last byte
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJSONChar := make([]bool, 256)
	for _, v := range chars {
		isJSONChar[v] = true
	}
	for space(json[start]) {
		if start > len(json)-1 {
			return nil
		}
		start++
	}
	for i := start; i < lenj; i++ {
		curr := json[i]
		if !isJSONChar[curr] {
			continue
		}
		if curr == 34 {
			if inQuote {
				for n := i - 1; n > -1; n-- {
					if json[n] != 92 {
						if (i-n)&1 != 0 {
							inQuote = !inQuote
							break
						} else {
							goto cont
						}
					}
					continue
				}
			} else {
				inQuote = true
				start = i + 1
				continue
			}
			if inQuote {
				start = i + 1
				continue
			}
			end = i
		cont:
			continue
		}
		if inQuote {
			continue
		} else {
			switch curr {
			case 91, 123:
				switch last {
				case 58:
					newNode := &node{up: core, label: string(key)}
					core.down = append(core.down, newNode)
					core = newNode
				default:
					newNode := &node{up: core, label: strconv.Itoa(indexList.Last())}
					core.down = append(core.down, newNode)
					core = newNode
				}
				indexList.Push(0)
				braceList.Push(i)
				valStart = i + 1
				last = curr
				continue
			case 93, 125:
				switch last {
				case 58:
					newNode := &node{up: core, label: string(key), value: json[valStart:i]}
					core.down = append(core.down, newNode)
				case 44:
					newNode := &node{up: core, label: strconv.Itoa(indexList.Last()), value: json[valStart:i]}
					core.down = append(core.down, newNode)
					valStart = i + 1
				}
				core.value = json[braceList.Pop() : i+1]
				indexList.Pop()
				core = core.up
				last = curr
				continue
			case 58:
				key = json[start:end]
				valStart = i + 1
				last = 58
				continue
			case 44:
				switch last {
				case 58:
					newNode := &node{up: core, label: string(key), value: json[valStart:i]}
					core.down = append(core.down, newNode)
				case 44, 91:
					newNode := &node{up: core, label: strconv.Itoa(indexList.Last()), value: json[valStart:i]}
					core.down = append(core.down, newNode)
				}
				indexList.Inc()
				valStart = i + 1
				last = 44
				continue
			}
		}
	}
	return nil
}
