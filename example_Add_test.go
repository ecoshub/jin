package jin

import "fmt"

func Example() {

}

func ExampleAdd() {
    var err error
	var newLink []byte = []byte(`"https://dev.to/eco9999"`)
	var json []byte = []byte(`{"user":"eco","links":["https://github.com/ecoshub"]}`)

	json, err = jin.Add(json, newLink, "links")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println(string(json))
    // Output: {"user":"eco","links":["https://github.com/ecoshub","https://dev.to/eco9999"]}
}

// func ExampleParser_Add() {
//     fmt.Println("hi")
//     // Output: hi
// }