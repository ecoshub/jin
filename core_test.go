package jint

import (
	"testing"
	"strings"
	"strconv"
	test "jint/test"
)

var (
	json []byte
	paths [][]string
	values []string
)

func InitValues(t *testing.T){
	json = test.ReadFile("test/test-json.json")
	newPaths := strings.Split(string(test.ReadFile("test/test-json-paths.json")), "\n")
	newValues := strings.Split(string(test.ReadFile("test/test-json-values.json")), "\n")
	paths = make([][]string, 0, len(newPaths))
	values = make([]string, 0, len(newValues))
	if len(newPaths) == 0 {
		t.Errorf("Paths length is zero.\n")
		return
	}
	if len(newValues) == 0 {
		t.Errorf("Values length is zero.\n")
		return
	}
	for _,val := range newValues {
		values = append(values, val)
	}
	for _,val := range newPaths {
		paths = append(paths, ParseArray(val))
	}
}

func TestGetInit(t *testing.T){
	str, err := test.ExecuteNode("get")
	if err != nil {
		t.Errorf("Init Error E:%v, S:%v\n", err, str)
		return
	}
	InitValues(t)
}

func TestGet(t *testing.T){
	for i, _ := range paths {
		_, start, end, err := Core(json, paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Get), path:%v err:%v\n", paths[i], err)
			return
		}
		// t.Logf("val:>%v<\n", string(value))
		value := json[start:end]
		if json[start - 1] != 34  {
			value = Flatten(value)
		}
		if string(value) != StripQuotes(values[i]) {
			t.Errorf("Fail, not same answer path:%v\n, got:\t\t>%v<\n, expected:\t>%v<  i:%v\n", paths[i], string(value), StripQuotes(values[i]), i)
			return
		}
	}
}

func TestSetInit(t *testing.T){
	str, err := test.ExecuteNode("set")
	if err != nil {
		t.Errorf("Init Error E:%v , S:%v\n", err, str)
		return
	}
	InitValues(t)
}

func TestSet(t *testing.T){

	for i, _ := range paths {
		value, err := Set(json, []byte(`"test-string"`), paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Set), path:%v err:%v\n", paths[i], err)
			return
		}
		value, err = Get(value, paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Get), path:%v err:%v\n", paths[i], err)
			return
		}
		if string(value) != StripQuotes(values[i]) {
			t.Errorf("Fail, not same answer path:%v\n, got:\t\t>%v<\n, expected:\t>%v<\n", paths[i], string(value), values[i] )
			return
		}
	}
}

func TestSetKeyInit(t *testing.T){
	str, err := test.ExecuteNode("get")
	if err != nil {
		t.Errorf("Init Error E:%v , S:%v\n", err, str)
		return
	}
	InitValues(t)
}

func TestSetKey(t *testing.T){
	for i, _ := range paths {
		keyStart, _, _, err1 := Core(json, paths[i]...)
		if err1 != nil {
			t.Errorf("Total Fail(Core), path:%v\n", paths[i])
			return	
		}
		newJson, err2 := SetKey(json, "test-key", paths[i]...)
		// it is a number
		if keyStart == -1 {
			if err2 == nil {
				t.Errorf("It is an element of an array cannot be set a new key %v", paths[i])
				return
			}
		}else{
			if err2 != nil {
				t.Errorf("It is a key it can be set a new key %v", paths[i])
				return
			}
			newPath := make([]string, len(paths[i]))
			copy(newPath, paths[i][:len(paths[i]) - 1])
			newPath[len(newPath) - 1] = "test-key"
			_, start, end, err := Core(newJson, newPath...)
			if err != nil {
				t.Errorf("Total Fail(Get), path:%v err:%v\n", paths[i], err)
				return
			}
			value := newJson[start:end]
			if newJson[start - 1] != 34  {
				value = Flatten(value)
			}
			if string(value) != StripQuotes(values[i]) {
				t.Errorf("Fail, not same answer path:%v\n, got:\t\t>%v<\n, expected:\t>%v<\n", newPath, string(value), values[i] )
				return
			}
		}
	}
}

func TestAddKVInit(t *testing.T){
	str, err := test.ExecuteNode("addkv")
	if err != nil {
		t.Errorf("Init Error E:%v , S:%v\n", err, str)
		return
	}
	InitValues(t)
}

func TestAddKV(t *testing.T){
	for i, _ := range paths {
		value, err := AddKeyValue(json, "test-key", []byte(`"test-value"`), paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Set), path:%v err:%v\n", paths[i], err)
			return
		}
		value, err = Get(value, paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Get), path:%v err:%v\n", paths[i], err)
			return
		}
		if string(Flatten(value)) != StripQuotes(values[i]) {
			t.Errorf("Fail, not same answer path:%v\n, got:\t\t>%v<\n, expected:\t>%v<\n", paths[i], string(value), values[i] )
			return
		}
	}
}

func TestAddInit(t *testing.T){
	str, err := test.ExecuteNode("add")
	if err != nil {
		t.Errorf("Init Error E:%v , S:%v\n", err, str)
		return
	}
	InitValues(t)
}

func TestAdd(t *testing.T){
	for i, _ := range paths {
		// t.Logf("path:>%v %T %v<\n",paths[i], paths[i], len(paths[i]))
		value, err := Add(json, []byte(`"test-value"`), paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Set), path:%v err:%v\n", paths[i], err)
			return
		}
		value, err = Get(value, paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Get), path:%v err:%v\n", paths[i], err)
			return
		}
		if string(Flatten(value)) != StripQuotes(values[i]) {
			t.Errorf("Fail, not same answer path:%v\n, got:\t\t>%v<\n, expected:\t>%v<\n", paths[i], string(Flatten(value)), values[i] )
			return
		}
	}
}

func TestInsertInit(t *testing.T){
	str, err := test.ExecuteNode("insert")
	if err != nil {
		t.Errorf("Init Error E:%v , S:%v\n", err, str)
		return
	}
	InitValues(t)
}

func TestInsert(t *testing.T){
	var err error
	var value []byte
	for i, _ := range paths {
		json, err = Insert(json, 0, []byte(`"test-value"`), paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Set), path:%v err:%v\n", paths[i], err)
			return
		}
		value, err = Get(json, paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Get), path:%v err:%v\n", paths[i], err)
			return
		}
		if string(Flatten(value)) != StripQuotes(values[i]) {
			t.Errorf("Fail, not same answer path:%v\n, got:\t\t>%v<\n, expected:\t>%v<\n", paths[i], string(value), values[i] )
			return
		}
	}
}

func TestDeleteKVInit(t *testing.T){
	str, err := test.ExecuteNode("deleteKV")
	if err != nil {
		t.Errorf("Init Error E:%v , S:%v\n", err, str)
		return
	}
	InitValues(t)
}

func TestDeleteKV(t *testing.T){
	for i, _ := range paths {
		value, err := AddKeyValue(json, "test-key", []byte(`"test-value"`), paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Set), path:%v err:%v\n", paths[i], err)
			return
		}
		newPath := make([]string, len(paths[i]))
		copy(newPath, paths[i])
		newPath = append(newPath, "test-key")
		value, err = Delete(value, newPath...)
		if err != nil {
			t.Errorf("Total Fail(Set), path:%v err:%v\n", paths[i], err)
			return
		}
		value, err = Get(value, paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Get), path:%v err:%v\n", paths[i], err)
			return
		}
		if string(Flatten(value)) != StripQuotes(values[i]) {
			t.Errorf("Fail, not same answer path:%v\n, got:\t\t>%v<\n, expected:\t>%v<\n", paths[i], string(value), values[i] )
			return
		}
	}
}

func TestDeleteVInit(t *testing.T){
	str, err := test.ExecuteNode("deleteV")
	if err != nil {
		t.Errorf("Init Error E:%v , S:%v\n", err, str)
		return
	}
	InitValues(t)
}

func TestDeleteV(t *testing.T){
	for i, _ := range paths {
		value, err := Get(json, paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Get), path:%v err:%v\n", paths[i], err)
			return
		}
		arr := ParseArray(string(value))
		value, err = Add(json, []byte(`"test-value"`), paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Set), path:%v err:%v\n", paths[i], err)
			return
		}
		newPath := make([]string, len(paths[i]))
		copy(newPath, paths[i])
		newPath = append(newPath, strconv.Itoa(len(arr)))
		value, err = Delete(value, newPath...)
		if err != nil {
			t.Errorf("Total Fail(Set), path:%v err:%v\n", paths[i], err)
			return
		}
		value, err = Get(value, paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Get), path:%v err:%v\n", paths[i], err)
			return
		}
		if string(Flatten(value)) != StripQuotes(values[i]) {
			t.Errorf("Fail, not same answer path:%v\n, got:\t\t>%v<\n, expected:\t>%v<\n", paths[i], string(value), values[i] )
			return
		}
	}
}

func TestArrayIterInit(t *testing.T){
	str, err := test.ExecuteNode("arrayiter")
	if err != nil {
		t.Errorf("Init Error E:%v , S:%v\n", err, str)
		return
	}
	InitValues(t)
}

func TestArrayIter(t *testing.T){
	for _, path := range paths {
		count := 0
		err := IterateArray(json, func(value []byte, err error) bool {
			newPath := make([]string, len(path))
			copy(newPath, path)
			newPath = append(newPath, strconv.Itoa(count))
			value2, err := Get(json, newPath...)
			if string(value) != string(value2) {
				t.Errorf("Fail, not same answer path:%v\n, got:\t\t>%v<\n, expected:\t>%v<\n",  string(value2), string(value), newPath)
				return false
			}else{
				count++
				return true
			}
		}, path...)
		if err != nil {
			t.Errorf("Total Fail(ArrayIter), path:%v err:%v\n", path, err)
			return
		}
	}
}
