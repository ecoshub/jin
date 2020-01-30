package jint

import "strconv"

func pCore(json []byte, core *node) error {
	if len(json) == 0 {
		return BAD_JSON_ERROR(0)
	}
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _, v := range chars {
		isJsonChar[v] = true
	}
	offset := 0
	for space(json[offset]) {
		if offset > len(json)-1 {
			return BAD_JSON_ERROR(offset)
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
	brace := make([]int, 0, 16)
	indexLevel := make([]int, 0, 16)
	path := make([]string, 0, 16)
	if json[offset] == 123 || json[offset] == 91 {
		indexLevel = append(indexLevel, 0)
		brace = append(brace, offset)
		level++
		lastIndex = offset + 1
	} else {
		return BAD_JSON_ERROR(offset)
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
			} else {
				inQuote = !inQuote
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
					core.value = json[lastIndex:i]
					core.typ = 0
					core = core.up
				case 91, 44:
					core = core.link(strconv.Itoa(indexLevel[len(indexLevel)-1]))
					core.value = json[lastIndex:i]
					core.typ = 0
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
					core.value = json[lastIndex:i]
					core.typ = 0
					core = core.up
				}
				if len(path) == 0 {
					return nil
				}
				core = core.up
				core = core.link(path[len(path)-1])
				core.value = json[brace[len(brace)-1]:i + 1]
				if len(core.value) != 0 {
					if core.value[0] == 91 {
						core.typ = 3
					}
					if core.value[0] == 123 {
						core.typ = 2
					}
				}else{
					return BAD_JSON_ERROR(i)
				}
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
					core.value = json[lastIndex:i]
					core.typ = 1
					core = core.up
				}
				if len(path) == 0 {
					return nil
				}
				core = core.up
				core = core.link(path[len(path)-1])
				core.value = json[brace[len(brace)-1]:i + 1]
				if len(core.value) != 0 {
					if core.value[0] == 91 {
						core.typ = 3
					}
					if core.value[0] == 123 {
						core.typ = 2
					}
				}else{
					return BAD_JSON_ERROR(i)
				}
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
	return nil
}
