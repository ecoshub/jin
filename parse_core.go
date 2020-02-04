package jint

import "strconv"

func pCore(json []byte, core *node) error{
	if len(json) < 2 {
		return nil
	}
	inQuote := false
	braceList := MakeSeq(4)
	indexList := MakeSeq(4)
	indexList.Push(0)
	var start int
	var end int
	var key []byte
	var valStart int
	var last byte	
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _, v := range chars {
		isJsonChar[v] = true
	}
	for space(json[start]) {
		if start > len(json)-1 {
			return nil
		} else {
			start++
			continue
		}
	}
	for i := start ; i < len(json) ; i ++ {
		curr := json[i]
		if !isJsonChar[curr] {
			continue
		}
		if curr == 34 {
			if inQuote {
				for n := i - 1; n > -1; n-- {
					if json[n] != 92 {
						if (i-n) & 1 != 0 {
							inQuote = !inQuote
							break
						} else {
							goto cont
						}
					}
					continue
				}
			}else{
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
		}else{
			switch curr {
			case 91, 123:
				switch last {
				case 58:
					current := CreateNode(core)
					current.label = string(key)
					core = current
				default:
					current := CreateNode(core)
					current.label = strconv.Itoa(indexList.Last())
					core = current
				}
				indexList.Push(0)
				braceList.Push(i)
				valStart = i + 1
				last = curr
				continue
			case 93, 125:
				switch last {
				case 58:
					current := CreateNode(core)
					current.label = string(key)
					current.value = json[valStart:i]
					core = current
					core = core.up
				case 44:
					current := CreateNode(core)
					current.label = strconv.Itoa(indexList.Last())
					current.value = json[valStart:i]
					core = current
					core = core.up
					valStart = i + 1
				}
				core.value = json[braceList.Pop():i+1]
				indexList.Pop()
				core = core.up
				last = curr
				continue
			case 58:
				key = json[start:end]
				valStart = i + 1
				last = curr
				continue
			case 44:
				switch last {
				case 58:
					current := CreateNode(core)
					current.label = string(key)
					current.value = json[valStart:i]
					core = current
					core = core.up
				case 44, 91:
					current := CreateNode(core)
					current.label = strconv.Itoa(indexList.Last())
					current.value = json[valStart:i]
					core = current
					core = core.up
				}
				indexList.Inc()
				valStart = i + 1
				last = curr
				continue
			}
		}
	}
	return nil
}