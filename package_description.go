/*
Copyright (c) 2020 eco
license that can be found in the LICENSE file.


Package jin is a comprehensive JSON manipulation tools bundle.
It provides parse, interpret, build and format tools.
Third party packages only used for benchmark. No dependency need for core functions.

There is four type of tools in  Jin:
- Interpreter
- Parser
- Builders
- Formaters

PARSER AND INTERPRETER
Major difference between parsing and interpreting is,
parser has to read all data before answer your needs.
On the other hand interpreter reads up to find the path you need.

Once the parse is complete you can access any data with no time.
But there is a time cost, and this cost can increase as data content grows.

QUICK START

Interpreter:
Interpreter is core element of this package, no need for instanciate, just call which function you want!

	// All interperter functions need one JSON as byte slice format. 
	json := []byte(`{"git":"ecoshub","repo":{"id":233809925,"name":"ecoshub/jin","url":"https://api.github.com/repos/ecoshub/jin"}}`)
	
	// And most of them needs a path value as string slice
	path := []string{"repo", "url"}

	// path can be provided as variable
	value, err := jin.Get(json, path...)

	// or hard coded
	value, err := jin.Get(json, "repo", "url")

	// type of 'value' is byte slice

	// if needed any other type to return all variations implemented as different functions.
	// for example. if you need 'value' to return as string.
	value, err := jin.GetString(json, "repo", "url")

All interpreter functions (except function variations) has own example provided in godoc.
The Interpreter has 36 functions.
Importent functions;
- Get
- Add
- AddKeyValue
- Set
- SetKey
- Delete
- Insert
- IterateArray
- IterateKeyValue
*/
package jin
