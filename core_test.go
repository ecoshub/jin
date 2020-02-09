package jin

import (
	"errors"
	"fmt"
	"test"
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
	sep               = test.Sep()
	testFiles         []string
	testFileDir       string = "test" + sep + testFileName
	pathsFileNameDir  string = "test" + sep + pathsFileName
	valuesFileNameDir string = "test" + sep + valuesFileName
	// tempDir              string =
	errorSkipPath      error = errors.New("skipped, no paths file found")
	errorSkipValue     error = errors.New("skipped, no values file found")
	errorSkipJSON      error = errors.New("skipped, no json file found")
	errorTriggerFailed error = errors.New("skipped, node trigger failed")
	errorEmptyPath     error = errors.New("skipped, empty path")
	errorNullValue     error = errors.New("skipped, null value")
	errorNullArray     error = errors.New("skipped, null array")
)

func init() {
	testFiles = test.Dir(test.GetCurrentDir() + sep + testsDir)
}

func errorMessage(where string) error {
	return fmt.Errorf("%v: '%v'", failMessage, where)
}

func triggerNode(state string, fileName string) error {
	test.WriteFile(testFileDir, test.ReadFile(testsDir+fileName))
	str, err := test.ExecuteNode(state)
	if err != nil {
		return fmt.Errorf("err:%v inner:%v", errorTriggerFailed, str)
	}
	return nil
}

func getComponents(file string) ([]string, []string, []byte, error) {
	json := test.ReadFile(testsDir + file)
	if json == nil {
		return nil, nil, nil, errorSkipJSON
	}

	pathFile := string(test.ReadFile(pathsFileNameDir))
	if pathFile == "" {
		return nil, nil, nil, errorSkipPath
	}
	paths := strings.Split(pathFile, "\n")

	valueFile := string(test.ReadFile(valuesFileNameDir))
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

func coreTestFunction(t *testing.T, state string, mainTest func(json []byte, path []string, expected string) ([]byte, error, string, string)) {
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
			value, err, expected, sticker := mainTest(json, path, expected)
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

func TestInterperterGet(t *testing.T) {
	coreTestFunction(t, "all", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "Get"
		value, err := Get(json, path...)
		return value, err, expected, sticker
	})
}

func TestInterperterSet(t *testing.T) {
	coreTestFunction(t, "all", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "Set"
		if len(path) == 0 {
			t.Logf("warning: %v, func: %v, path: %v", errorEmptyPath.Error(), sticker, path)
			return []byte(expected), nil, expected, sticker
		}
		testVal := []byte(`test-value`)
		json, err := Set(json, testVal, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		value, err := Get(json, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, string(testVal), sticker
	})
}

func TestInterperterSetKey(t *testing.T) {
	coreTestFunction(t, "object-values", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "SetKey"
		newKey := "test-key"
		json, err := SetKey(json, newKey, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		value, err := Get(json, append(path[:len(path)-1], newKey)...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, expected, sticker
	})
}

func TestInterperterAddKeyValue(t *testing.T) {
	coreTestFunction(t, "object", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "AddKeyValue"
		newKey := "test-key"
		newVal := []byte("test-value")
		value, err := Get(json, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/GetControl"
		}
		if string(value) == "null" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullValue.Error(), sticker, path)
			return []byte(expected), nil, expected, sticker
		}
		json, err = AddKeyValue(json, newKey, []byte("test-value"), path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		value, err = Get(json, append(path, newKey)...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, string(newVal), sticker
	})
}

func TestInterperterAdd(t *testing.T) {
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "Add"
		newVal := []byte("test-value")
		json, err := Add(json, newVal, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		array := ParseArray(expected)
		value, err := Get(json, append(path, strconv.Itoa(len(array)))...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, string(newVal), sticker
	})
}

func TestInterperterInsert(t *testing.T) {
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "Insert"
		if expected == "[]" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullArray.Error(), sticker, path)
			return []byte(expected), nil, expected, sticker
		}
		newVal := []byte("test-value")
		json, err := Insert(json, 0, newVal, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		value, err := Get(json, append(path, "0")...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, string(newVal), sticker
	})
}

func TestInterperterDelete(t *testing.T) {
	sticker := "Delete"
	newKey := "test-key"
	newVal := []byte("test-value")
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		array := ParseArray(expected)
		json, err := Add(json, newVal, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Add"
		}
		json, err = Delete(json, append(path, strconv.Itoa(len(array)))...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		value, err := Get(json, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, expected, sticker
	})
	coreTestFunction(t, "object", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		if string(expected) == "null" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullValue.Error(), sticker, path)
			return []byte(expected), nil, expected, sticker
		}
		json, err := AddKeyValue(json, newKey, newVal, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/AddKeyValue"
		}
		json, err = Delete(json, append(path, newKey)...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		value, err := Get(json, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, expected, sticker
	})
}

func TestInterperterIterateArray(t *testing.T) {
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "IterateArray"
		if expected == "[]" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullArray.Error(), sticker, path)
			return []byte(expected), nil, expected, sticker
		}
		array := ParseArray(expected)
		count := 0
		done := true
		err := IterateArray(json, func(value []byte) bool {
			got := formatValue(value)
			expected := stripQuotes(array[count])
			if got != expected {
				t.Logf("path:%v\n", path)
				t.Logf("got:%v\n", got)
				t.Logf("expected:%v\n", expected)
				done = false
				return false
			}
			count++
			return true
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
			return nil, errorMessage("TestInterperter/" + sticker), "*expected*", sticker
		}
		return nil, nil, "", sticker
	})
}

func TestInterperterIterateKeyValue(t *testing.T) {
	coreTestFunction(t, "object", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "IterateKeyValue"
		if string(expected) == "null" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullValue.Error(), sticker, path)
			return []byte(expected), nil, expected, sticker
		}
		done := true
		err := IterateKeyValue(json, func(key, value []byte) bool {
			value2, err := Get(json, append(path, string(key))...)
			if err != nil {
				done = false
				return false
			}
			got := formatValue(value)
			expected := formatValue(value2)
			if got != expected {
				t.Logf("path:%v\n", path)
				t.Logf("got:%v\n", got)
				t.Logf("expected:%v\n", expected)
				done = false
				return false
			}
			return true
		}, path...)
		if err != nil {
			t.Logf("error. %v\n", err)
			done = false
		}
		if !done {
			return nil, errorMessage("IterateKeyValue/" + sticker), "*expected*", sticker
		}
		return nil, nil, "", sticker
	})
}

func TestParserGet(t *testing.T) {
	coreTestFunction(t, "all", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "Parser.Get"
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, err, expected, sticker
		}
		value, err := prs.Get(path...)
		return value, err, expected, sticker
	})
}

func TestParserSet(t *testing.T) {
	coreTestFunction(t, "all", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "Parser.Set"
		if len(path) == 0 {
			t.Logf("warning: %v, func: %v, path: %v", errorEmptyPath.Error(), sticker, path)
			return []byte(expected), nil, expected, sticker
		}
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, err, expected, sticker
		}
		testVal := []byte(`test-value`)
		err = prs.Set(testVal, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		value, err := prs.Get(path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, string(testVal), sticker
	})
}

func TestParserSetKey(t *testing.T) {
	coreTestFunction(t, "object-values", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "Parser.SetKey"
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, err, expected, sticker
		}
		newKey := "test-key"
		err = prs.SetKey(newKey, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		value, err := prs.Get(append(path[:len(path)-1], newKey)...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, expected, sticker
	})
}

func TestParserAddKeyValue(t *testing.T) {
	coreTestFunction(t, "object", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "Parser.AddKeyValue"
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, err, expected, sticker
		}
		newKey := "test-key"
		newVal := []byte("test-value")
		value, err := prs.Get(path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/GetControl"
		}
		if string(value) == "null" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullValue.Error(), sticker, path)
			return []byte(expected), nil, expected, sticker
		}
		err = prs.AddKeyValue(newKey, []byte("test-value"), path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		value, err = prs.Get(append(path, newKey)...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, string(newVal), sticker
	})
}

func TestParserAdd(t *testing.T) {
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "Parser.Add"
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, err, expected, sticker
		}
		newVal := []byte("test-value")
		err = prs.Add(newVal, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		array := ParseArray(expected)
		value, err := prs.Get(append(path, strconv.Itoa(len(array)))...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, string(newVal), sticker
	})
}

func TestParserInsert(t *testing.T) {
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		sticker := "Parser.Insert"
		if expected == "[]" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullArray.Error(), sticker, path)
			return []byte(expected), nil, expected, sticker
		}
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, err, expected, sticker
		}
		newVal := []byte("test-value")
		err = prs.Insert(0, newVal, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		value, err := prs.Get(append(path, "0")...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, string(newVal), sticker
	})
}

func TestParserDelete(t *testing.T) {
	sticker := "Parser.Delete"
	newKey := "test-key"
	newVal := []byte("test-value")
	coreTestFunction(t, "array", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, err, expected, sticker
		}
		array := ParseArray(expected)
		err = prs.Add(newVal, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Add"
		}
		err = prs.Delete(append(path, strconv.Itoa(len(array)))...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		value, err := prs.Get(path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, expected, sticker
	})
	coreTestFunction(t, "object", func(json []byte, path []string, expected string) ([]byte, error, string, string) {
		if string(expected) == "null" {
			t.Logf("warning: %v, func: %v, path: %v", errorNullValue.Error(), sticker, path)
			return []byte(expected), nil, expected, sticker
		}
		prs, err := Parse(json)
		if err != nil {
			t.Logf("error. %v\n", err)
			return nil, err, expected, sticker
		}
		err = prs.AddKeyValue(newKey, newVal, path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/AddKeyValue"
		}
		err = prs.Delete(append(path, newKey)...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/" + sticker
		}
		value, err := prs.Get(path...)
		if err != nil {
			return nil, err, "*expected*", sticker + "/Get"
		}
		return value, nil, expected, sticker
	})
}
