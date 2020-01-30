package jint

// func ( p * parse) Set(newVal []byte, path ... string) error{
// 	if len(path) == 0 {
// 		return BAD_JSON_ERROR(0)
// 	}
// 	curr, err := p.core.walk(path)
// 	if err != nil {
// 		return err
// 	}
// 	val := curr.getVal(p.json)
// 	oldType := 0 //value
// 	newType := 0 // value
// 	if val[0] ==123 || val[0] == 91{
// 		oldType = 1 // json
// 	}
// 	if newVal[0] ==123 || newVal[0] == 91{
// 		newType = 1 // json
// 	}
// 	newJson := make([]byte, 0, len(p.json) - len(val) + len(newVal))
// 	newJson = append(newJson, p.json[:curr.start]...)
// 	newJson = append(newJson, newVal...) 
// 	newJson = append(newJson, p.json[curr.start + len(val):]...)
// 	p.json = newJson
// 	newVal = Flatten(newVal)
// 	delt := 0
// 	switch oldType{
// 	case 0:
// 		switch newType{
// 		case 0:
// 			// value to value
// 			curr.setVal(newVal)
// 			delt = len(newVal) - len(val)
// 		case 1:
// 			// value to json/array
// 			newCore := CreateNode(nil)
// 			pCore(newVal, newCore)
// 			if newCore.down[0] != nil {
// 				newCore = newCore.down[0]
// 			}
// 			newCore.label = curr.label
// 			newCore.up = curr.up

// 			index := curr.getIndex()
// 			curr.up.down[index] = newCore
// 			delt = len(newVal) - len(val)
// 			newCore.start += curr.start
// 			newCore.end += curr.start
// 			for _,d := range newCore.down {
// 				d.start += curr.start
// 				d.end += curr.start
// 				d.hasValue = false
// 			}
// 			newCore.up.setOffset(index + 1, delt)
// 		}
// 	case 1:
// 		switch newType{
// 		case 0:
// 			// json/array to value
// 			off := curr.end - curr.start
// 			delt = len(newVal) - off
// 			curr.down = []*node{}
// 			curr.end = curr.start + len(newVal)
// 			curr.hasValue = false
// 		case 1:
// 			// json/array to json/array
// 			delt = len(newVal) - len(val)
// 			index := curr.getIndex()
// 			newCore := CreateNode(nil)
// 			pCore(newVal, newCore)
// 			if newCore.down[0] != nil {
// 				newCore = newCore.down[0]
// 			}
// 			newCore.label = curr.label
// 			curr.up.down[index] = newCore
// 			newCore.up = curr.up
// 			newCore.up.setOffset(index, curr.start)

// 		}
// 	}
// 	curr.setOffsetUp(delt)
// 	return nil
// }

// func (n * node) setVal(val []byte){
// 	n.value = val
// 	n.hasValue = true
// }
