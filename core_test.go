package jin

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const (
	testsDir       string = "tests"
	testFileName   string = "test-json.json"
	pathsFileName  string = "test-json-paths.json"
	valuesFileName string = "test-json-values.json"
	failMessage    string = "test failed trying this function"
)

var (
	testFiles          []string
	testFileDir        string = "test" + sep() + testFileName
	pathsFileNameDir   string = "test" + sep() + pathsFileName
	valuesFileNameDir  string = "test" + sep() + valuesFileName
	errorSkipPath      error  = errors.New("skipped, no paths file found")
	errorSkipValue     error  = errors.New("skipped, no values file found")
	errorSkipJSON      error  = errors.New("skipped, no json file found")
	errorTriggerFailed error  = errors.New("skipped, node trigger failed")
	errorEmptyPath     error  = errors.New("skipped, empty path")
	errorNullValue     error  = errors.New("skipped, null value")
	errorNullArray     error  = errors.New("skipped, null array")
)

func init() {
	testFiles = dir(getCurrentDir() + sep() + testsDir)
}

func errorMessage(where string) error {
	return fmt.Errorf("%v: '%v'", failMessage, where)
}

func triggerNode(state string, fileName string) error {
	writeFile(testFileDir, readFile(testsDir+fileName))
	str, err := executeNode("node", "test/test-case-creator.js", state)
	if err != nil {
		return fmt.Errorf("err:%v inner:%v err:%v", errorTriggerFailed, err, str)
	}
	return nil
}

func getComponents(file string) ([]string, []string, []byte, error) {
	json := readFile(testsDir + file)
	if json == nil {
		return nil, nil, nil, errorSkipJSON
	}

	pathFile := string(readFile(pathsFileNameDir))
	if pathFile == "" {
		return nil, nil, nil, errorSkipPath
	}
	paths := strings.Split(pathFile, "\n")

	valueFile := string(readFile(valuesFileNameDir))
	if valueFile == "" {
		return nil, nil, nil, errorSkipValue
	}
	values := strings.Split(valueFile, "\n")
	return paths, values, json, nil
}

func formatValue(value []byte) string {
	if len(value) > 1 {
		if (value[0] == 91 && value[len(value)-1] == 93) ||
			(value[0] == 123 && value[len(value)-1] == 125) {
			return string(Flatten(value))
		}
	}
	return string(value)
}

func coreTestFunction(t *testing.T, state string, mainTest func(json []byte, path []string, expected string) ([]byte, string, string, error)) {
	flatTest := false
	for _, file := range testFiles {
		t.Logf("file: %v", testsDir+file)
	start:
		err := triggerNode(state, file)
		if err != nil {
			t.Errorf("%v\n", err)
			continue
		}
		paths, values, json, err := getComponents(file)
		if flatTest {
			json = Flatten(json)
		}
		if err != nil {
			t.Logf("%v\n", err)
			continue
		}
		for i := 0; i < len(paths); i++ {
			path := ParseArray(paths[i])
			expected := stripQuotes(values[i])
			// this is the core of test
			value, expected, sticker, err := mainTest(json, path, expected)
			// ---
			got := formatValue(value)
			if err != nil {
				t.Logf(err.Error())
				t.Error(errorMessage("coreTest/" + sticker))
				t.Logf("path:%v\n", path)
				t.Logf("got:>%v<\n", got)
				t.Logf("expected:>%v<\n", expected)
				return
			}
			if got != expected {
				t.Error(errorMessage("coreTest/" + sticker))
				t.Logf("path:%v\n", path)
				t.Logf("got:>%v<\n", got)
				t.Logf("expected:>%v<\n", expected)
				return
			}
		}
		if !flatTest {
			flatTest = true
			goto start
		}
		flatTest = false
	}
}

func TestNode(t *testing.T) {
	str, err := executeNode("node", "test/test-node.js")
	t.Logf(str)
	if err != nil {
		t.Logf("err:%v inner:%v", errorTriggerFailed, str)
		return
	}
	if str == "node is running well" {
		t.Logf("node is running well")
	} else {
		t.Logf("err:%v inner:%v", errorTriggerFailed, str)
		return
	}
}

func TestInterperterGet(t *testing.T) {
	coreTestFunction(t, "all", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Get"
		value, err := Get(json, path...)
		return value, expected, sticker, err
	})
}

func TestInterperterSet(t *testing.T) {
	coreTestFunction(t, "all", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Set"
		if len(path) == 0 {
			t.Logf("warning: %v, func: %v, path: %v", errorEmptyPath.Error(), sticker, path)
			return []byte(expected), expected, sticker, nil
		}
		testVal := []byte(`test-value`)
		json, err := Set(json, testVal, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		value, err := Get(json, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, string(testVal), sticker, nil
	})
}

func TestInterperterSetKey(t *testing.T) {
	coreTestFunction(t, "object-values", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "SetKey"
		newKey := "test-key"
		json, err := SetKey(json, newKey, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		value, err := Get(json, append(path[:len(path)-1], newKey)...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, expected, sticker, nil
	})
}

func TestInterperterAddKeyValue(t *testing.T) {
	coreTestFunction(t, "object", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "AddKeyValue"
		newKey := "test-key"
		newVal := []byte("test-value")
		value, err := Get(json, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/GetControl", err
		}
		if string(value) == "null" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullValue.Error(), sticker, path)
			return []byte(expected), expected, sticker, nil
		}
		json, err = AddKeyValue(json, newKey, []byte("test-value"), path...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		value, err = Get(json, append(path, newKey)...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, string(newVal), sticker, nil
	})
}

func TestInterperterAdd(t *testing.T) {
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Add"
		newVal := []byte("test-value")
		json, err := Add(json, newVal, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		array := ParseArray(expected)
		value, err := Get(json, append(path, strconv.Itoa(len(array)))...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, string(newVal), sticker, nil
	})
}

func TestInterperterInsert(t *testing.T) {
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Insert"
		if expected == "[]" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullArray.Error(), sticker, path)
			return []byte(expected), expected, sticker, nil
		}
		newVal := []byte("test-value")
		json, err := Insert(json, 0, newVal, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		value, err := Get(json, append(path, "0")...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, string(newVal), sticker, nil
	})
}

func TestInterperterDelete(t *testing.T) {
	sticker := "Delete"
	newKey := "test-key"
	newVal := []byte("test-value")
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		array := ParseArray(expected)
		json, err := Add(json, newVal, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/Add", err
		}
		json, err = Delete(json, append(path, strconv.Itoa(len(array)))...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		value, err := Get(json, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, expected, sticker, nil
	})
	coreTestFunction(t, "object", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		if string(expected) == "null" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullValue.Error(), sticker, path)
			return []byte(expected), expected, sticker, nil
		}
		json, err := AddKeyValue(json, newKey, newVal, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/AddKeyValue", err
		}
		json, err = Delete(json, append(path, newKey)...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		value, err := Get(json, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, expected, sticker, nil
	})
}

func TestInterperterIterateArray(t *testing.T) {
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "IterateArray"
		if expected == "[]" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullArray.Error(), sticker, path)
			return []byte(expected), expected, sticker, nil
		}
		array := ParseArray(expected)
		count := 0
		done := true
		err := IterateArray(json, func(value []byte) (bool, error) {
			got := formatValue(value)
			expected := stripQuotes(array[count])
			if got != expected {
				t.Logf("path:%v\n", path)
				t.Logf("got:%v\n", got)
				t.Logf("expected:%v\n", expected)
				done = false
				return false, nil
			}
			count++
			return true, nil
		}, path...)
		if count != len(array) {
			t.Logf("error. iteration count and real arrays count is different.\n")
			done = false
		}
		if err != nil {
			t.Logf("error. %v\n", err)
			done = false
		}
		if !done {
			return nil, "*expected*", sticker, errorMessage("TestInterperter/" + sticker)
		}
		return nil, "", sticker, nil
	})
}

func TestInterperterIterateKeyValue(t *testing.T) {
	coreTestFunction(t, "object", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "IterateKeyValue"
		if string(expected) == "null" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullValue.Error(), sticker, path)
			return []byte(expected), expected, sticker, nil
		}
		done := true
		err := IterateKeyValue(json, func(key, value []byte) (bool, error) {
			value2, err := Get(json, append(path, string(key))...)
			if err != nil {
				done = false
				return false, nil
			}
			got := formatValue(value)
			expected := formatValue(value2)
			if got != expected {
				t.Logf("path:%v\n", path)
				t.Logf("got:%v\n", got)
				t.Logf("expected:%v\n", expected)
				done = false
				return false, nil
			}
			return true, nil
		}, path...)
		if err != nil {
			t.Logf("error. %v\n", err)
			done = false
		}
		if !done {
			return nil, "*expected*", sticker, errorMessage("IterateKeyValue/" + sticker)
		}
		return nil, "", sticker, nil
	})
}

func TestInterpreterGetKeys(t *testing.T) {
	coreTestFunction(t, "keys", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Interpreter.GetKeys"
		keys, err := GetKeys(json, path...)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		expKeys := ParseArray(expected)
		if !stringArrayEqual(keys, expKeys) {
			return []byte("some element"), "some element", sticker, errors.New("not equal")
		}
		return []byte(""), "", sticker, nil
	})
}

func TestInterpreterGetValues(t *testing.T) {
	coreTestFunction(t, "values", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Interpreter.GetValues"
		json = Flatten(json)
		values, err := GetValues(json, path...)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		expValues := ParseArray(expected)
		if !stringArrayEqual(values, expValues) {
			return []byte("some element"), "some element", sticker, errors.New("not equal")
		}
		return []byte(""), "", sticker, nil
	})
}

func TestInterpreterGetKeysValues(t *testing.T) {
	coreTestFunction(t, "keys", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Interpreter.GetKeysValues"
		json = Flatten(json)
		expValues, err := GetValues(json, path...)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		expKeys, err := GetKeys(json, path...)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		keys, values, err := GetKeysValues(json, path...)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		if !stringArrayEqual(keys, expKeys) || !stringArrayEqual(values, expValues) {
			return []byte("some element"), "some element", sticker, errors.New("not equal")
		}
		return []byte(""), "", sticker, nil
	})
}

func TestInterpreterGetLength(t *testing.T) {
	coreTestFunction(t, "length", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Interpreter.Length"
		length, err := Length(json, path...)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		if strconv.Itoa(length) != expected {
			return []byte(strconv.Itoa(length)), expected, sticker, errors.New("not equal")
		}
		return []byte(""), "", sticker, nil
	})
}

func TestParserGet(t *testing.T) {
	coreTestFunction(t, "all", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Parser.Get"
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		value, err := prs.Get(path...)
		return value, expected, sticker, err
	})
}

func TestParserSet(t *testing.T) {
	coreTestFunction(t, "all", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Parser.Set"
		if len(path) == 0 {
			t.Logf("warning: %v, func: %v, path: %v", errorEmptyPath.Error(), sticker, path)
			return []byte(expected), expected, sticker, nil
		}
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		testVal := []byte(`test-value`)
		err = prs.Set(testVal, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		value, err := prs.Get(path...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, string(testVal), sticker, nil
	})
}

func TestParserSetKey(t *testing.T) {
	coreTestFunction(t, "object-values", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Parser.SetKey"
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		newKey := "test-key"
		err = prs.SetKey(newKey, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		value, err := prs.Get(append(path[:len(path)-1], newKey)...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, expected, sticker, nil
	})
}

func TestParserAddKeyValue(t *testing.T) {
	coreTestFunction(t, "object", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Parser.AddKeyValue"
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		newKey := "test-key"
		newVal := []byte("test-value")
		value, err := prs.Get(path...)
		if err != nil {
			return nil, "*expected*", sticker + "/GetControl", err
		}
		if string(value) == "null" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullValue.Error(), sticker, path)
			return []byte(expected), expected, sticker, nil
		}
		err = prs.AddKeyValue(newKey, []byte("test-value"), path...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		value, err = prs.Get(append(path, newKey)...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, string(newVal), sticker, nil
	})
}

func TestParserAdd(t *testing.T) {
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Parser.Add"
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		newVal := []byte("test-value")
		err = prs.Add(newVal, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		array := ParseArray(expected)
		value, err := prs.Get(append(path, strconv.Itoa(len(array)))...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, string(newVal), sticker, nil
	})
}

func TestParserInsert(t *testing.T) {
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		sticker := "Parser.Insert"
		if expected == "[]" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullArray.Error(), sticker, path)
			return []byte(expected), expected, sticker, nil
		}
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		newVal := []byte("test-value")
		err = prs.Insert(0, newVal, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		value, err := prs.Get(append(path, "0")...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, string(newVal), sticker, nil
	})
}

func TestParserDelete(t *testing.T) {
	sticker := "Parser.Delete"
	newKey := "test-key"
	newVal := []byte("test-value")
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		array := ParseArray(expected)
		err = prs.Add(newVal, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/Add", err
		}
		err = prs.Delete(append(path, strconv.Itoa(len(array)))...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		value, err := prs.Get(path...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, expected, sticker, nil
	})
	coreTestFunction(t, "object", func(json []byte, path []string, expected string) ([]byte, string, string, error) {
		if string(expected) == "null" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullValue.Error(), sticker, path)
			return []byte(expected), expected, sticker, nil
		}
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, expected, sticker, err
		}
		err = prs.AddKeyValue(newKey, newVal, path...)
		if err != nil {
			return nil, "*expected*", sticker + "/AddKeyValue", err
		}
		err = prs.Delete(append(path, newKey)...)
		if err != nil {
			return nil, "*expected*", sticker + "/" + sticker, err
		}
		value, err := prs.Get(path...)
		if err != nil {
			return nil, "*expected*", sticker + "/Get", err
		}
		return value, expected, sticker, nil
	})
}
