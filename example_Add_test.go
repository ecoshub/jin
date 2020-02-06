package jin

import "fmt"

func Example() {

}

// Add() adds a value to an array
// it must point to an array
// otherwise it will provide an error message through 'err' variable
func ExampleAdd() {
    var err error
	var exampleJSON []byte = []byte(`{"user":"eco","year":2020,"links":["https://github.com/ecoshub"]}`)
	
	var newLink []byte = []byte(`"https://dev.to/eco9999"`)

	exampleJSON, err = jin.Add(exampleJSON, newLink, "links")
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	fmt.Println(string(exampleJSON))
    // Output: {"user":"eco","year":2020,"links":["https://github.com/ecoshub","https://dev.to/eco9999"]}
}

// func ExampleParser_Add() {
//     fmt.Println("hi")
//     // Output: hi
// }