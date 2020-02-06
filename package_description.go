/*

**UNDER CONSTRUCTION**

Copyright (c) 2020 eco.

license that can be found in the LICENSE file.


Package Jin is a comprehensive JSON manipulation tools bundle.
It provides parse, interpret, build and format tools.
Third party packages only used for benchmark. No dependency need for core functions.

There are four types of tools in  Jin:

- Interpreter

- Parser

- Builder

- Formater

PARSER AND INTERPRETER

INTERPRETER

Major difference between parsing and interpreting is,
parser has to read all data before answer your needs.
On the other hand interpreter reads up to find the path you need.

Once the parse is complete you can access any data with no time.
But there is a time cost to parse data, and this cost can increase as data content grows.

QUICK START

INTERPRETER

Interpreter is core element of this package, no need for instanciate, just call which function you want!
First let's look at function parameters.

	// All interperter functions need one JSON as byte slice format. 
	json := []byte(`{"git":"ecoshub","repo":{"id":233809925,"name":"ecoshub/jin"}}`)
	
	// And most of them needs a path value for navigate.
	path := []string{"repo", "name"}

Let's take the function Get()

Get function returns the value that path has pointed.

	value, err := jin.Get(json, path...)
	if err != nil {
		return err
	}

path value can be a string slice or hard coded

	value, err := jin.Get(json, "repo", "name")
	if err != nil {
		return err
	}

Get() function return type is byte slice.

All variations of return types are implemented as different functions.

for example. if you need 'value' as string.
then you can use GetString() like this.

	value, err := jin.GetString(json, "repo", "name")
	if err != nil {
		return err
	}

All interpreter functions (except function variations) has own example provided in godoc.

The Interpreter has 36 functions.

Other Importent functions

- func Get(json []byte, path ...string) ([]byte, error)

- func Add(json []byte, value []byte, path ...string) ([]byte, error)

- AddKeyValue

- Set

- SetKey

- Delete

- Insert

- IterateArray

- IterateKeyValue



*/
package jin
