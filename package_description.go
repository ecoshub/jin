/*
**UNDER CONSTRUCTION**

Copyright (c) 2020 eco.

License that can be found in the LICENSE file.


Package Jin is a comprehensive JSON manipulation tools bundle.
All functions tested with random data against NODEJS.
All test-path and test-value cases created automatically with NODEJS.
It provides parse, interpret, build and format tools.
Third-party packages only used for the benchmark. No dependency need for core functions.

Benchmarked against.

	"github.com/buger/jsonparser" (with interpreter)
	"github.com/valyala/fastjson" (with parser)
	"github.com/json-iterator/go" (with parser)

In conclusion the results of the benchmark 'Jin' was the fastest and lightweight package

For more information take a look at BENCHMARK section.


PARSER VS INTERPRETER

Major difference between parsing and interpreting is,
parser has to read all data before answer your needs.
On the other hand interpreter reads up to find the path you need.

Once the parse is complete you can access any data with no time.
But there is a time cost to parse data, and this cost can increase as data content grows.

If you need to access all keys of a JSON then I am simply recommend you to user Parser.
But if you need to access some keys of a JSON I strongly refoment you to use Interperter, it will be much faster than parser. 


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

For example. If you need 'value' as string,
then you can use GetString() like this.

	value, err := jin.GetString(json, "repo", "name")
	if err != nil {
	return err
	}

PARSER

Parser is another alternative for manipulation JSON.
We recommend to use this structure when you need to access all or most of the keys in the JSON.

Parser constructor need only one parameter.

	// Parser constructor function jin.Parse() need one JSON as byte slice format. 
	json := []byte(`{"git":"ecoshub","repo":{"id":233809925,"name":"ecoshub/jin"}}`)

Lets Parse it with Parse function.

	prs, err := jin.Parse(json)
	if err != nil {
	return err
	}

Let's take the function Parser.Get()

	value, err := prs.Get("repo")
	if err != nil {
	return err
	}

*About path value look above. 

Parser.Get() function return type is byte slice like Get() function of interpreter.

All variations of return types are implemented as different functions.

For example. If you need 'value' as string.
Then you can use Parser.GetString() like this.

	value, err := prs.GetString("repo")
	if err != nil {
	return err
	}

All interpreter/parser functions (except function variations line GetString()) has own example provided in godoc.

Other important functions of interpreter/parser. 

- func Add(), AddKeyValue(), Set(), SetKey() Delete(), Insert(), IterateArray(), IterateKeyValue()
*/
package jin
