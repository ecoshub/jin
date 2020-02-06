package jin

import "fmt"

func Example() {

}

func ExampleAdd() {
    var err error
	var newLink []byte = []byte(`"godoc.org/github.com/ecoshub"`)
	var json []byte = []byte(`{"user":"eco","links":["github.com/ecoshub"]}`)

	json, err = Add(json, newLink, "links")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println(string(json))
    // Output: {"user":"eco","links":["github.com/ecoshub","godoc.org/github.com/ecoshub"]}
}

func ExampleAddKeyValue() {
    var err error
	var newValue []byte = []byte(`"go"`)
	var json []byte = []byte(`{"user":"eco"}`)

	json, err = AddKeyValue(json, "language", newValue)
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println(string(json))
    // Output: {"user":"eco","language":"go"}
}

func ExampleDelete() {
    var err error
	var json []byte = []byte(`{"user":"eco","languages":["go","java","python","C", "Cpp"]}`)

	json, err = Delete(json, "languages", "1")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	// After first deletion.
    // {"user":"eco","languages":["go","python","C", "Cpp"]}

	json, err = Delete(json, "user")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println(string(json))
    // Output: {"languages":["go","python","C", "Cpp"]}
}

func ExampleFlatten() {
	var json []byte = []byte(`{
	"user": "eco",
	"languages": [
		"go",
		"java",
		"python",
		"C",
		"Cpp"
	],
	"following": {
		"social": "dev.to",
		"code": "github"
	}
}`)

	json = Flatten(json)
	fmt.Println(string(json))
    // Output: {"user":"eco","languages":["go","java","python","C","Cpp"],"following":{"social":"dev.to","code":"github"}}
}

func ExampleFormat() {
	var json []byte = []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"],"following":{"social":"dev.to","code":"github"}}`)

	json = Indent(json)
	fmt.Println(string(json))
    /* Output: {
	"user": "eco",
	"languages": [
		"go",
		"java",
		"python",
		"C",
		"Cpp"
	],
	"following": {
		"social": "dev.to",
		"code": "github"
	}
}*/
}

func ExampleGet() {
    var err error
    var value []byte
	var json []byte = []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"],"following":{"social":"dev.to","code":"github"}}`)

	value, err = Get(json, "following", "social")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println(string(value))
    // Output: dev.to
}

func ExampleInsert() {
    var err error
	var newValue []byte = []byte(`"visual basic`)
	var json []byte = []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"]}`)

	json, err = Insert(json, 2, newValue, "languages")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println(string(json))
    // Output: {"user":"eco","languages":["go","java","visual basic,"python","C","Cpp"]}
}

func ExampleIterateArray() {
    var err error
	var json []byte = []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"]}`)

	err = IterateArray(json, func(value []byte) bool {
		fmt.Println(string(value))
		// this return is some kind control mechanism for escape the iteration any time you want.
		// true means keep iterate. false means stop iteration.
		return true
	}, "languages")

	if err != nil {
		fmt.Println(err)
		return
	}
    /* Output: go
java
python
C*/
}

func ExampleIterateKeyValue() {
    var err error
	var json []byte = []byte(`{"index":42,"user":"eco","language":"go","uuid":"4a1531c25d5ef124295a","active":true}`)

	err = IterateKeyValue(json, func(key, value []byte) bool {
		fmt.Println("key  :", string(key))
		fmt.Println("value:", string(value))
		// this return is some kind control mechanism for escape the iteration any time you want.
		// true means keep iterate. false means stop iteration.
		return true
	})

	if err != nil {
		fmt.Println(err)
		return
	}
    /* Output: key  : index
value: 42
key  : user
value: eco
key  : language
value: go
key  : uuid
value: 4a1531c25d5ef124295a*/
}

func ExampleMakeArray() {
	var years []byte = MakeArray(2005,2009,2013,2019)
	// [2005,2009,2013,2019]
	var active []byte = MakeArray(false,true,true,true)
	// [false,true,true,true]
	var languages []byte = MakeArray("visual-basic","java","python","go")
	// ["visual-basic","java","python","go"]

	all := MakeArrayBytes(years, active, languages)
	fmt.Println(string(all))
	// Output: [[2005,2009,2013,2019],[false,true,true,true],["visual-basic","java","python","go"]]
}

func ExampleMakeJson() {
	var keys  []string = []string{"username","ip","mac","active"}
	var values []interface{} = []interface{}{"eco","192.168.1.108","bc:ae:c5:13:84:f9",true}

	var user []byte = MakeJson(keys, values)

	fmt.Println(string(user))
	// Output: {"username":"eco","ip":"192.168.1.108","mac":"bc:ae:c5:13:84:f9","active":true}
}

// func ExampleParser_Add() {
//     fmt.Println("hi")
//     // Output: hi
// }