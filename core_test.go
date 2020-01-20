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

	for i,val := range newValues {
		if val != "" && newPaths[i] != "" {
			paths = append(paths, ParseArray(newPaths[i]))
			values = append(values, val)
		}
	}
}



func TestGetInit(t *testing.T){
	str, err := test.ExecuteNode("get")
	if err != nil {
		t.Errorf("Init Error E:%v , S:%v\n", err, str)
	}
	InitValues()
}

func TestGet(t *testing.T){
	if len(paths) != len(values) {
		t.Errorf("Paths and Values length not equal. %v %v \n", len(paths), len(values))
	}
	if len(paths) == 0 {
		t.Errorf("Paths and Values length is zero.\n")
	}
	for i, _ := range paths {
		value, done := Get(json, paths[i]...)
		if done != nil {
			t.Errorf("Total Fail(Get), path:%v\n", paths[i])
		}
		if value[0] == 91 || value[0] == 123 {
			value = Flatten(value)
		}
		if string(value) != StripQuotes(values[i]) {
			t.Errorf("Fail, not same answer path:%v\n, got:\t\t>%v<\n, expected:\t>%v<\n", paths[i], string(value), values[i] )
		}
	}
}


func TestSetInit(t *testing.T){
	str, err := test.ExecuteNode("set")
	if err != nil {
		t.Errorf("Init Error E:%v , S:%v\n", err, str)
	}
	InitValues()
}

func TestSet(t *testing.T){
	if len(paths) != len(values) {
		t.Errorf("Paths and Values length not equal. %v %v \n", len(paths), len(values))
	}
	if len(paths) == 0 {
		t.Errorf("Paths and Values length is zero.\n")
	}
	for i, _ := range paths {
		value, done := Get(json, paths[i]...)
		if done != nil {
			t.Errorf("Total Fail(Get), path:%v\n", paths[i])
			return
		}
		if value[0] == 91 || value[0] == 123 {
			value, done = Set(json, []byte(`"test-string"`), paths[i]...)
			if done != nil {
				t.Errorf("Total Fail(Set), path:%v\n", paths[i])
				return
			}
		}
		value, done = Get(value, paths[i]...)
		if done != nil {
			t.Errorf("Total Fail(Get), path:%v\n", paths[i])
			return
		}
		if string(value) != StripQuotes(values[i]) {
			t.Errorf("Fail, not same answer path:%v\n, got:\t\t>%v<\n, expected:\t>%v<\n", paths[i], string(value), values[i] )
		}
	}
}