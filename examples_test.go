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

// func ExampleParser_Add() {
//     fmt.Println("hi")
//     // Output: hi
// }