package jint

import (
	"testing"
	"strings"
	test "jint/test"
)

var json []byte
var paths [][]string
var values []string

func InitValues(){
	json = test.ReadFile("test/test-json.json")
	newPaths := strings.Split(string(test.ReadFile("test/test-json-paths.json")), "\n")
	newValues := strings.Split(string(test.ReadFile("test/test-json-values.json")), "\n")
	paths = make([][]string, 0, len(newPaths))
	values = make([]string, 0, len(newValues))
	for i,val := range newValues {
		if newPaths[i] != "" && newValues[i] != ""{
			paths = append(paths, ParseArray(newPaths[i]))
			values = append(values, val)
		}
	}
}

func TestGetInit(t *testing.T){
	str, err := test.ExecuteNode("get")
	if err != nil {
		t.Errorf("Init Error E:%v , S:%v\n", err, str)
		return
	}
	InitValues()
}

func TestGet(t *testing.T){
	if len(paths) != len(values) {
		t.Errorf("Paths and Values length not equal. %v %v \n", len(paths), len(values))
		return
	}
	if len(paths) == 0 {
		t.Errorf("Paths and Values length is zero.\n")
		return
	}
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
	InitValues()
}

func TestSet(t *testing.T){
	if len(paths) != len(values) {
		t.Errorf("Paths and Values length not equal. %v %v \n", len(paths), len(values))
		return
	}
	if len(paths) == 0 {
		t.Errorf("Paths and Values length is zero.\n")
		return
	}
	for i, _ := range paths {
		_, start, _, err := Core(json, paths[i]...)
		if err != nil {
			t.Errorf("Total Fail(Core), path:%v err:%v\n", paths[i], err)
			return
		}
		if json[start] != 91 && json[start] != 123 {
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
}

func TestSetKeyInit(t *testing.T){
	str, err := test.ExecuteNode("get")
	if err != nil {
		t.Errorf("Init Error E:%v , S:%v\n", err, str)
		return
	}
	InitValues()
}

func TestSetKey(t *testing.T){
	if len(paths) != len(values) {
		t.Errorf("Paths and Values length not equal. %v %v \n", len(paths), len(values))
		return
		
	}
	if len(paths) == 0 {
		t.Errorf("Paths and Values length is zero.\n")
		return
		
	}
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












func TestEnd(t *testing.T){
	str, err := test.ExecuteNode("get")
	if err != nil {
		t.Errorf("Init Error E:%v , S:%v\n", err, str)
		return
	}
	InitValues()
}
