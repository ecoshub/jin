package example

import "fmt"
import "jin"

func Example() {

}

func ExampleAdd() {
    var err error
	var newLink []byte = []byte(`"godoc.org/github.com/ecoshub"`)
	var json []byte = []byte(`{"user":"eco","links":["github.com/ecoshub"]}`)

	json, err = jin.Add(json, newLink, "links")
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

	json, err = jin.AddKeyValue(json, "language", newValue)
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

	json, err = jin.Delete(json, "languages", "1")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	// After first deletion.
    // {"user":"eco","languages":["go","python","C", "Cpp"]}

	json, err = jin.Delete(json, "user")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println(string(json))
    // Output: {"languages":["go","python","C", "Cpp"]}
}

// func ExampleParser_Add() {
//     fmt.Println("hi")
//     // Output: hi
// }