package jsoninterpreter

import (
	"fmt"
	"strconv"
	"errors"
)

func Get(json []byte, path ... string) ([]byte, error){
	// path null.
	if len(path) == 0 {
		return nil, errors.New("Error: Path can not be null.")
	}
	// main offset track of this search.
	offset := 0
	currentPath := path[0]
	// important chars for json.
	// 34 = "
	// 44 = ,
	// 58 = :
	// 91 = [
	// 93 = ]
	// 123 = {
	// 125 = }
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	// creating a bool array fill with false
	isJsonChar := make([]bool, 256)
	// only interested chars is true
	for _,v := range chars {
		isJsonChar[v] = true
	}
	// trim spaces of start
	for space(json[offset]) {
		offset++
	}
	// braceType determine whether or not search will be a json search or array search
	braceType := json[offset]
	// main iteration off all bytes.
	for k := 0 ; k < len(path) ; k ++ {
		// 91 = [, begining of an array search
		if braceType == 91 {
			// path value cast to integer for determine index.
			arrayIndex, err := strconv.Atoi(currentPath)
			if err != nil {
				// braceType and current path type is confilicts.
				return nil, errors.New("Error: Index Expected, got string.")
			}
			// is this search has found something.
			// done := false
			// zeroth index search.
			if arrayIndex == 0 {
				// increment offset for not catch current brace.
				offset++
				// inner iteration for brace search.
				for i := offset; i < len(json) ; i ++ {
					// curr is current byte of reading.
					curr := json[i]
					// open curly brace
					if curr == 123 {
						// change brace type of next search.
						braceType = curr
						if k != len(path) - 1{
							// if its not last path than change currentPath to next path.
							currentPath = path[k + 1]
						}
						// assign offset to brace index.
						offset = i
						// found it.
						// done = true
						// break the array search scope.
						break
					}
					// open square brace
					if curr == 91 {
						// change brace type of next search.
						braceType = curr
						if k != len(path) - 1{
							// if its not last path than change currentPath to next path.
							currentPath = path[k + 1]
						}
						// searching for zeroth index is conflicts with searching zeroth array or arrays zeroth element.
						offset = i + 1
						// found it.
						// done = true
						// break the array search scope.
						break
					}
					// doesnt have to always find a brace. it can be a value.
					if !space(curr){
						// done = true
						break
					}
				}
				// if offset == 1 {
				// 	return nil, errors.New("Error: Bad format")
				// }
			}else{
				// brace level every brace increments the level
				level := 0
				// main in quote flag for determine what is in quote and what is not
				inQuote := false
				// index found flag.
				found := false
				// index count of current element.
				indexCount := 0
				// not interested with column char in this search
				isJsonChar[58] = false
				for i := offset ; i < len(json) ; i ++ {
					// curr is current byte of reading.
					curr := json[i]
					// just interesting with json chars. other wise continue.
					if !isJsonChar[curr]{
						continue
					}
					// if current byte is quote
					if curr == 34 {
						// check befour char it might be escape char.
						if json[i - 1] == 92 {
							continue
						}
						// change inQuote flag to opposite.
						inQuote = !inQuote
						continue
					}
					if inQuote {
						continue
					}else{
						// open braces
						if curr == 91 || curr == 123{
							// if found befour done with this search
							// break array search scope
							if found {
								level++
								braceType = curr
								currentPath = path[k + 1]
								found = false
								// done = true
								break
							}
							level++
							continue
						}
						if curr == 93 || curr == 125 {
							level--
							// if level is less than 1 it mean index not in this array. 
							if level < 1 {
								return nil, errors.New("Error: Index out of range")
							}
							continue
						}
						// not found befour
						if !found {
							// same level with path
							if level == 1 {
								// curren byte is comma
								if curr == 44 {
									// inc index
									indexCount++
									if indexCount == arrayIndex {
										offset = i + 1
										if k == len(path) - 1{
											// last path and found than break
											// done = true
											break
										}
										// not last path keep going. for find next brace Type.
										found = true
										continue
									}
									continue
								}
								continue
							}
							continue
						}
						continue
					}
				}
				// check true for column char again for keep same with first decleration.
				isJsonChar[58] = true
			}
		}else{
			inQuote := false
			found := false
			start := 0
			end := 0
			level := k
			// not interested with comma to this level
			isJsonChar[44] = false
			for i := offset ; i < len(json) ; i ++ {
				curr := json[i]
				if !isJsonChar[curr]{
					continue
				}
				if curr == 34 {
					inQuote = !inQuote
					if found {
						continue
					}
					if level != k + 1 {
						continue
					}
					if inQuote {
						start = i + 1
						continue
					}
					end = i
					continue
				}
				if inQuote {
					continue
				}else{
					if curr == 91 {
						if found {
							braceType = curr
							currentPath = path[k + 1]
							break
						}
						level++
						continue
					}
					if curr == 123 {
						if found {
							k++
							level++
							currentPath = path[k]
							found = false
							continue
						}
						level++
						continue
					}
					// close brace
					if curr == 93 || curr == 125 {
						level--
						continue
					}
					// column
					if level == k + 1 {
						if curr == 58 {
							if compare(json, start, end, currentPath) {
								offset = i + 1
								found = true
								if k == len(path) - 1{
									isJsonChar[44] = true
									break
								}else{
									continue
								}
							}
							// interested with comma to this level
							isJsonChar[44] = true
							// not interested with column to this level
							isJsonChar[58] = false
							// little jump alogirthm :{} -> ,
							for j := i ;  j < len(json) ; j ++ {
								curr := json[j]
								if !isJsonChar[curr]{
									continue
								}
								// quote
								if curr == 34 {
									if json[j - 1] == 92 {
										continue
									}
									inQuote = !inQuote
									continue
								}
								if inQuote {
									continue
								}else{
									if curr == 91 || curr == 123 {
										level++
										continue
									}
									if curr == 93 || curr == 125 {
										level--
										continue
									}
									// comma
									if curr == 44 {
										if level == k + 1 {
											i = j
											break
										}
										continue
									}
									continue
								}

							}
							// not interested with comma to this level
							isJsonChar[44] = false
							// interested with column to this level
							isJsonChar[58] = true
							continue
						}
						continue
					}
				}
			}
			isJsonChar[44] = true
			if !found {
				return nil, errors.New("Error: Last key not found.")
			}
		}
	}
	if offset == 0 {
		return nil, errors.New("Error: Something went wrong... not sure.")
	}
	for space(json[offset]) {
		offset++
	}
	// starts with { [
	if json[offset] == 91 || json[offset] == 123 {
		level := 0
		inQuote := false
		for i := offset ; i < len(json) ; i ++ {
			curr := json[i]
			if !isJsonChar[curr]{
				continue
			}
			if curr == 34 {
				// escape character control
				if json[i - 1] == 92 {
					continue
				}
				inQuote = !inQuote
				continue
			}
			if inQuote {
				continue
			}else{
				if curr == 91 || curr == 123 {
					level++
				}
				if curr == 93 || curr == 125 {
					level--
					if level == 0 {
						return json[offset:i + 1], nil
					}
					continue
				}
				continue
			}
			continue
		}
	}else{
		// starts with quote
		if json[offset] == 34 {
			inQuote := false
			for i := offset ;  i < len(json) ; i ++ {
				curr := json[i]
				// quote
				if curr == 34 {
					// escape character control
					if json[i - 1] == 92 {
						continue
					}
					if inQuote {
						return json[offset + 1:i], nil
					}
					inQuote = !inQuote
					continue
				}
			}
		}else{
			// starts without quote
			for i := offset ;  i < len(json) ; i ++ {
				if isJsonChar[json[i]] {
					return json[offset:i], nil
				}
			}
		}
	}
	return nil, errors.New("Error: Something went wrong at the end... not sure.")
}

func GetString(json []byte, path ... string) (string, error){
	val, done := Get(json, path...)
	return string(val), done
}

func GetInt(json []byte, path ... string) (int, error){
	val, err := GetString(json, path...)
	if err != nil {
		return -1, err
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return -1, fmt.Errorf("Cast Error: value '%v' can not cast to int.", val)
	}
	return intVal, nil
}

func GetFloat(json []byte, path ... string) (float64, error){
	val, err := GetString(json, path...)
	if err != nil {
		return -1, err
	}
	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return -1, fmt.Errorf("Cast Error: value '%v' can not cast to float64.", val)
	}
	return floatVal, nil
}

func GetBool(json []byte, path ... string) (bool, error){
	val, err := GetString(json, path...)
	if err != nil {
		return false, err
	}
	if val == "true" {
		return true, nil
	}
	if val == "false" {
		return false, nil
	}
	return false, fmt.Errorf("Cast Error: value '%v' can not cast to bool.", val)
}

func GetStringArray(json []byte, path ... string) ([]string, error){
	val, err := GetString(json, path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, fmt.Errorf("Cast Error: value '%v' can not cast to []string.", val)
	}
	if val[0] == '[' && val[lena - 1] == ']' {
		newArray := make([]string, 0, 16)
		start := 1
		inQuote := false
		for i := 1 ; i < lena - 1 ; i ++ {
			curr := val[i]
			if curr == 34 || curr == 44 {
				if curr == 34 {
					// escape character control
					if val[i - 1] == 92 {
						continue
					}
					inQuote = !inQuote
					continue
				}
				if inQuote {
					continue
				}else{
					if curr == 44 {
						newArray = append(newArray, trimSpace(val, start, i))
						start = i + 1
					}
				}
			}
		}
		newArray = append(newArray, trimSpace(val, start, lena - 2))
		return newArray, nil
	}else{
		return nil, fmt.Errorf("Cast Error: value '%v' can not cast to []string.", val)
	}
}

func GetIntArray(json []byte, path ... string) ([]int, error){
	val, err := GetString(json, path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, fmt.Errorf("Cast Error: value '%v' can not cast to []int.", val)
	}
	if val[0] == '[' && val[lena - 1] == ']' {
		newArray := make([]int, 0, 16)
		start := 1
		inQuote := false
		for i := 1 ; i < lena - 1 ; i ++ {
			curr := val[i]
			if curr == 34 || curr == 44 {
				if curr == 34 {
					// escape character control
					if val[i - 1] == 92 {
						continue
					}
					inQuote = !inQuote
					continue
				}
				if inQuote {
					continue
				}else{
					if curr == 44 {
						num, err := strconv.Atoi(trimSpace(val, start, i))
						if err != nil {
							return nil,  fmt.Errorf("Cast Error: value '%v' can not cast to int.", trimSpace(val, start, i))
						}
						newArray = append(newArray, num)
						start = i + 1
					}
				}
			}
		}

		num, err := strconv.Atoi(trimSpace(val, start, lena - 2))
		if err != nil {
			return nil,  fmt.Errorf("Cast Error: value '%v' can not cast to int.", trimSpace(val, start, lena - 2))
		}
		newArray = append(newArray, num)
		return newArray, nil
	}else{
		return nil, fmt.Errorf("Cast Error: value '%v' can not cast to []int.", val)
	}
}

func GetFloatArray(json []byte, path ... string) ([]float64, error){
	val, err := GetString(json, path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, fmt.Errorf("Cast Error: value '%v' can not cast to []float64.", val)
	}
	if val[0] == '[' && val[lena - 1] == ']' {
		newArray := make([]float64, 0, 16)
		start := 1
		inQuote := false
		for i := 1 ; i < lena - 1 ; i ++ {
			curr := val[i]
			if curr == 34 || curr == 44 {
				if curr == 34 {
					// escape character control
					if val[i - 1] == 92 {
						continue
					}
					inQuote = !inQuote
					continue
				}
				if inQuote {
					continue
				}else{
					if curr == 44 {
						num, err := strconv.ParseFloat(trimSpace(val, start, i), 64)
						if err != nil {
							return nil,  fmt.Errorf("Cast Error: value '%v' can not cast to float64.", trimSpace(val, start, i))
						}
						newArray = append(newArray, num)
						start = i + 1
					}
				}
			}
		}

		num, err := strconv.ParseFloat(trimSpace(val, start, lena - 2), 64)
		if err != nil {
			return nil,  fmt.Errorf("Cast Error: value '%v' can not cast to float64.", trimSpace(val, start, lena - 2))
		}
		newArray = append(newArray, num)
		return newArray, nil
	}else{
		return nil, fmt.Errorf("Cast Error: value '%v' can not cast to []float64.", val)
	}
}

func GetBoolArray(json []byte, path ... string) ([]bool, error){
	val, err := GetString(json, path...)
	if err != nil {
		return nil, err
	}
	lena := len(val)
	if lena < 2 {
		return nil, fmt.Errorf("Cast Error: value '%v' can not cast to []bool.", val)
	}
	if val[0] == '[' && val[lena - 1] == ']' {
		newArray := make([]bool, 0, 16)
		start := 1
		inQuote := false
		for i := 1 ; i < lena - 1 ; i ++ {
			curr := val[i]
			if curr == 34 || curr == 44 {
				if curr == 34 {
					// escape character control
					if val[i - 1] == 92 {
						continue
					}
					inQuote = !inQuote
					continue
				}
				if inQuote {
					continue
				}else{
					if curr == 44 {
						val := trimSpace(val, start, i)
						if val == "true" || val == "false" {
							if val == "true"{
								newArray = append(newArray, true)
								start = i + 1
							}else{
								newArray = append(newArray, false)
								start = i + 1
							}
						}else{
							return nil,  fmt.Errorf("Cast Error: value '%v' can not cast to bool.", val)
						}
					}
				}
			}
		}
		val := trimSpace(val, start, lena - 2)
		if val == "true" || val == "false" {
			if val == "true"{
				newArray = append(newArray, true)
			}else{
				newArray = append(newArray, false)
			}
		}else{
			return nil,  fmt.Errorf("Cast Error: value '%v' can not cast to bool.", val)
		}
		return newArray, nil
	}else{
		return nil, fmt.Errorf("Cast Error: value '%v' can not cast to []bool.", val)
	}
}


func SetValue(json []byte, newValue []byte, path ... string) ([]byte, error){
	if len(path) == 0 {
		return nil, errors.New("Error: Path can not be null.")
	}
	if len(newValue) == 0 {
		return nil, errors.New("Error: New Value can not be null.")
	}
	offset := 0
	currentPath := path[0]
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _,v := range chars {
		isJsonChar[v] = true
	}
	for space(json[offset]) {
		offset++
	}
	braceType := json[offset]

	for k := 0 ; k < len(path) ; k ++ {
		if braceType == 91 {
			arrayNumber, err := strconv.Atoi(currentPath)
			if err != nil {
				return json, errors.New("Error: Index Expected.")
			}
			done := false
			if arrayNumber == 0 {
				offset++
				for i := offset; i < len(json) ; i ++ {
					curr := json[i]
					if curr == 123 {
						braceType = curr
						if k != len(path) - 1{
							currentPath = path[k + 1]
						}
						offset = i
						done = true
						break
					}
					if curr == 91 {
						braceType = curr
						if k != len(path) - 1{
							currentPath = path[k + 1]
						}
						offset = i + 1
						done = true
						break
					}
					if !space(curr){
						done = true
						break
					}
				}
			}else{
				level := 0
				inQuote := false
				found := false
				indexCount := 0
				// not interested with column to this level
				isJsonChar[58] = false
				for i := offset ; i < len(json) ; i ++ {
					curr := json[i]
					if !isJsonChar[curr]{
						continue
					}
					if curr == 34 {
						if json[i - 1] == 92 {
							continue
						}
						inQuote = !inQuote
						continue
					}
					if inQuote {
						continue
					}else{
						if curr == 91 || curr == 123{
							if found {
								level++
								braceType = curr
								currentPath = path[k + 1]
								found = false
								done = true
								break
							}
							level++
							continue
						}
						if curr == 93 || curr == 125 {
							level--
							if level < 1 {
								done = false
								break
							}
							continue
						}
						if !found {
							if level == 1 {
								if curr == 44 {
									indexCount++
									if indexCount == arrayNumber {
										offset = i + 1
										if k == len(path) - 1{
											done = true
											break
										}
										found = true
										continue
									}
									continue
								}
								continue
							}
							continue
						}
						continue
					}
				}
				// interested with column to this level
				isJsonChar[58] = true
			}
			if !done {
				return json, errors.New("Error: Index out of range")
			}
		}else{
			inQuote := false
			found := false
			start := 0
			end := 0
			level := k
			// not interested with comma to this level
			isJsonChar[44] = false
			for i := offset ; i < len(json) ; i ++ {
				curr := json[i]
				if !isJsonChar[curr]{
					continue
				}
				if curr == 34 {
					inQuote = !inQuote
					if found {
						continue
					}
					if level != k + 1 {
						continue
					}
					if inQuote {
						start = i + 1
						continue
					}
					end = i
					continue
				}
				if inQuote {
					continue
				}else{
					if curr == 91 {
						if found {
							braceType = curr
							currentPath = path[k + 1]
							break
						}
						level++
						continue
					}
					if curr == 123 {
						if found {
							k++
							level++
							currentPath = path[k]
							found = false
							continue
						}
						level++
						continue
					}
					// close brace
					if curr == 93 || curr == 125 {
						level--
						continue
					}
					// column
					if level == k + 1 {
						if curr == 58 {
							if compare(json, start, end, currentPath) {
								offset = i + 1
								found = true
								if k == len(path) - 1{
									isJsonChar[44] = true
									break
								}else{
									continue
								}
							}
							// interested with comma to this level
							isJsonChar[44] = true
							// not interested with column to this level
							isJsonChar[58] = false
							// little jump alogirthm :{} -> ,
							for j := i ;  j < len(json) ; j ++ {
								curr := json[j]
								if !isJsonChar[curr]{
									continue
								}
								// quote
								if curr == 34 {
									if json[j - 1] == 92 {
										continue
									}
									inQuote = !inQuote
									continue
								}
								if inQuote {
									continue
								}else{
									if curr == 91 || curr == 123 {
										level++
										continue
									}
									if curr == 93 || curr == 125 {
										level--
										continue
									}
									// comma
									if curr == 44 {
										if level == k + 1 {
											i = j
											break
										}
										continue
									}
									continue
								}

							}
							// not interested with comma to this level
							isJsonChar[44] = false
							// interested with column to this level
							isJsonChar[58] = true
							continue
						}
						continue
					}
				}
			}
			isJsonChar[44] = true
			if !found {
				return json, errors.New("Error: Last key not found.")
			}
		}
	}
	if offset == 0 {
		return json, errors.New("Error: Non")
	}
	for space(json[offset]) {
		offset++
	}
	// starts with { [
	if json[offset] == 91 || json[offset] == 123 {
		level := 0
		inQuote := false
		for i := offset ; i < len(json) ; i ++ {
			curr := json[i]
			if !isJsonChar[curr]{
				continue
			}
			if curr == 34 {
				// escape character control
				if json[i - 1] == 92 {
					continue
				}
				inQuote = !inQuote
				continue
			}
			if inQuote {
				continue
			}else{
				if curr == 91 || curr == 123 {
					level++
				}
				if curr == 93 || curr == 125 {
					level--
					if level == 0 {
						return replace(json, newValue, offset, i + 1), nil
					}
					continue
				}
				continue
			}
			continue
		}
	}else{
		// starts with quote
		if json[offset] == 34 {
			inQuote := false
			for i := offset ;  i < len(json) ; i ++ {
				curr := json[i]
				// quote
				if curr == 34 {
					// escape character control
					if json[i - 1] == 92 {
						continue
					}
					if inQuote {
						return replace(json, newValue, offset + 1, i), nil
					}
					inQuote = !inQuote
					continue
				}
			}
		}else{
			// starts without quote
			for i := offset ;  i < len(json) ; i ++ {
				if isJsonChar[json[i]] {
					return replace(json, newValue, offset, i), nil
				}
			}
		}
	}
	return nil, errors.New("Error: Non 2")
}

func SetStringValue(json []byte, newValue string, path ... string) ([]byte, error){
	return SetValue(json, []byte(newValue), path...)
}

func SetIntValue(json []byte, newValue int, path ... string) ([]byte, error){
	return SetValue(json, []byte(strconv.Itoa(newValue)), path...)
}

func SetFloatValue(json []byte, newValue float64, path ... string) ([]byte, error){
	return SetValue(json, []byte(strconv.FormatFloat(newValue, 'e', -1, 64)), path...)
}

func SetBoolValue(json []byte, newValue bool, path ... string) ([]byte, error){
	if newValue {
		return SetValue(json, []byte("true"), path...)
	}
	return SetValue(json, []byte("false"), path...)
}


func SetKey(json []byte, newValue []byte, path ... string) ([]byte, error){
	if len(path) == 0 {
		return json, errors.New("Error: Path can not be null.")
	}
	if len(newValue) == 0 {
		return json, errors.New("Error: New Value can not be null.")
	}
	for _, v := range newValue {
		if v  == 34 {
			return json, errors.New("Error: Key can not contain quote symbol.")
		}
	}
	offset := 0
	currentPath := path[0]
	chars := []byte{34, 44, 58, 91, 93, 123, 125}
	isJsonChar := make([]bool, 256)
	for _,v := range chars {
		isJsonChar[v] = true
	}
	for space(json[offset]) {
		offset++
	}
	braceType := json[offset]

	for k := 0 ; k < len(path) ; k ++ {
		if braceType == 91 {
			arrayNumber, err := strconv.Atoi(currentPath)
			if err != nil {
				return json, errors.New("Error: Index Expected.")
			}
			done := false
			if arrayNumber == 0 {
				offset++
				for i := offset; i < len(json) ; i ++ {
					curr := json[i]
					if curr == 123 {
						braceType = curr
						if k != len(path) - 1{
							currentPath = path[k + 1]
						}
						offset = i
						done = true
						break
					}
					if curr == 91 {
						braceType = curr
						if k != len(path) - 1{
							currentPath = path[k + 1]
						}
						offset = i + 1
						done = true
						break
					}
					if !space(curr){
						done = true
						break
					}
				}
			}else{
				level := 0
				inQuote := false
				found := false
				indexCount := 0
				// not interested with column to this level
				isJsonChar[58] = false
				for i := offset ; i < len(json) ; i ++ {
					curr := json[i]
					if !isJsonChar[curr]{
						continue
					}
					if curr == 34 {
						if json[i - 1] == 92 {
							continue
						}
						inQuote = !inQuote
						continue
					}
					if inQuote {
						continue
					}else{
						if curr == 91 || curr == 123{
							if found {
								level++
								braceType = curr
								currentPath = path[k + 1]
								found = false
								done = true
								break
							}
							level++
							continue
						}
						if curr == 93 || curr == 125 {
							level--
							if level < 1 {
								done = false
								break
							}
							continue
						}
						if !found {
							if level == 1 {
								if curr == 44 {
									indexCount++
									if indexCount == arrayNumber {
										offset = i + 1
										if k == len(path) - 1{
											done = true
											return json, errors.New("Error: Last value must be a key value,  not an array index.")
										}
										found = true
										continue
									}
									continue
								}
								continue
							}
							continue
						}
						continue
					}
				}
				// interested with column to this level
				isJsonChar[58] = true
			}
			if !done {
				return json, errors.New("Error: Index out of range")
			}
		}else{
			inQuote := false
			found := false
			start := 0
			end := 0
			level := k
			// not interested with comma to this level
			isJsonChar[44] = false
			for i := offset ; i < len(json) ; i ++ {
				curr := json[i]
				if !isJsonChar[curr]{
					continue
				}
				if curr == 34 {
					inQuote = !inQuote
					if found {
						continue
					}
					if level != k + 1 {
						continue
					}
					if inQuote {
						start = i + 1
						continue
					}
					end = i
					continue
				}
				if inQuote {
					continue
				}else{
					if curr == 91 {
						if found {
							braceType = curr
							currentPath = path[k + 1]
							break
						}
						level++
						continue
					}
					if curr == 123 {
						if found {
							k++
							level++
							currentPath = path[k]
							found = false
							continue
						}
						level++
						continue
					}
					// close brace
					if curr == 93 || curr == 125 {
						level--
						continue
					}
					// column
					if level == k + 1 {
						if curr == 58 {
							if compare(json, start, end, currentPath) {
								offset = i + 1
								found = true
								if k == len(path) - 1{
									isJsonChar[44] = true
									return replace(json, newValue, start, end), nil
									break
								}else{
									continue
								}
							}
							// interested with comma to this level
							isJsonChar[44] = true
							// not interested with column to this level
							isJsonChar[58] = false
							// little jump alogirthm :{} -> ,
							for j := i ;  j < len(json) ; j ++ {
								curr := json[j]
								if !isJsonChar[curr]{
									continue
								}
								// quote
								if curr == 34 {
									if json[j - 1] == 92 {
										continue
									}
									inQuote = !inQuote
									continue
								}
								if inQuote {
									continue
								}else{
									if curr == 91 || curr == 123 {
										level++
										continue
									}
									if curr == 93 || curr == 125 {
										level--
										continue
									}
									// comma
									if curr == 44 {
										if level == k + 1 {
											i = j
											break
										}
										continue
									}
									continue
								}

							}
							// not interested with comma to this level
							isJsonChar[44] = false
							// interested with column to this level
							isJsonChar[58] = true
							continue
						}
						continue
					}
				}
			}
			isJsonChar[44] = true
			if !found {
				return json, errors.New("Error: Last key not found.")
			}
		}
	}
	return json, errors.New("Error: Something went wrong... not sure.")
}

func SetStringKey(json []byte, newValue string, path ... string) ([]byte, error){
	return SetKey(json, []byte(newValue), path...)
}

func replace(json, newValue []byte, start, end int) []byte {
	newJson := make([]byte, 0, len(json) - end + start + len(newValue))
	newJson = append(newJson, json[:start]...)
	newJson = append(newJson, newValue...)
	newJson = append(newJson, json[end:]...)
	return newJson
}

func trimSpace(str string, start, eoe int) string {
	for space(str[start]){
		start++
	}
	end := start
	for !space(str[end]) && end < eoe {
		end++
	}
	return str[start:end]
}

func compare(json []byte, start, end int , key string) bool{
	if len(key) != end - start {
		return false
	}
	for i := 0 ; i < len(key) ; i ++ {
		if key[i] != json[start + i] {
			return false
		}
	}
	return true
}

func space(curr byte) bool{
	// space
	if curr == 32 {
		return true
	}
	// tab
	if curr == 9 {
		return true
	}
	// new line NL
	if curr == 10 {
		return true
	}
	// return CR
	if curr == 13 {
		return true
	}
	return false
}
