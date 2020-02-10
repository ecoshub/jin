/*
Package jin

Copyright (c) 2020 eco.

License that can be found in the LICENSE file.

WELCOME TO JIN

"Your wish is my command"


FAST AND EASY WAY TO DEAL WITH JSON

Jin is a comprehensive JSON manipulation tool bundle.
All functions tested with random data with help of Node.js.
All test-path and test-value created automatically with Node.js.

Jin provides parse, interpret, build and format tools for JSON.
Third-party packages only used for the benchmark. No dependency need for core functions.

We make some benchmark with other packages like us.

    github.com/buger/jsonparser
    github.com/valyala/fastjson
    github.com/json-iterator/go

In Result, Jin is the fastest (op/ns) and more memory friendly then others (B/op).

For more information please take a look at BENCHMARK section below.


INSTALLATION

	go get github.com/ecoshub/jin

And you are good to go. Import and start using.

QUICK START

PARSER VS INTERPRETER

Major difference between parsing and interpreting is
parser has to read all data before answer to your commands.
On the other hand interpreter reads up to find the data you need.

With parser, once the parse is complete you can get access any data with no time.
But there is a time cost to parse a data and this cost can increase as data content grows.

If you need to access all keys of a JSON then we are simply recommend you to use Parser.
But if you need to access some keys of a JSON we strongly recommend you to use Interpreter, it will be much faster than parser.


INTERPRETER

Interpreter is core element of this package, no need to create an Interpreter type, just call which function you want!

First let's look at function parameters.


	// All interpreter functions need one JSON as byte slice.
	json := []byte({"git":"ecoshub","repo":{"id":233809925,"name":["eco","jin"]}})

	// And most of them needs a path value for navigate.
	// Path value determines which part to navigate.
	// In this example we want to access 'jin' value.
	// So path must be 'repo' object -> 'name' array -> '1'
	// second element with index of one.
	path := []string{"repo", "name", "1"}


We are gonna use Get() function to return the value of path has pointed. In this case 'jin'.


	value, err := jin.Get(json, path...)
	if err != nil {
		log.Println(err)
		return
	}
	// the Get() functions return type is []byte
	// To understand its value,
	// first we have to convert it to string.
	fmt.Println(string(value))
	// Output: jin


Path value can consist hard coded values.


	value, err := jin.Get(json, "repo", "name", "1")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(value))
	//String Output: jin


Get() function return type is []byte but all other variations of return types are implemented with different functions.

For example. If you need "value" as string,

There is a function called GetString().


	value, err := jin.GetString(json, "repo", "name", "0")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(value))
	//Output: eco


PARSER

Parser is another alternative for JSON manipulation.

We recommend to use this structure when you need to access all or most of the keys in the JSON.

Parser constructor need only one parameter.


	// Parser constructor function jin.Parse() need one JSON as []byte.
	json := []byte(
	{
		"title": "LICENSE",
		"repo": {
			"id": 233809925,
			"name": "ecoshub/jin",
			"url": "https://api.github.com/repos/ecoshub/jin"
			}
	})


Parse it with Parse function.


	prs, err := jin.Parse(json)
	if err != nil {
		log.Println(err)
		return
	}


Let's look at Parser.Get()


	value, err := prs.Get("repo", "url")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(value))
	//Output: https://api.github.com/repos/ecoshub/jin


About path value look above.

There is all kind of return type methods for Parser like Interpreter.

You can use Parser.GetString() like this.


	value, err := prs.GetString("repo", "name")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(value))
	//String Output: ecoshub/jin


All interpreter/parser functions (except function variations like GetString()) has own example provided in here.

Other useful functions of interpreter/parser.

-Add(), AddKeyValue(), Set(), SetKey() Delete(), Insert(), IterateArray(), IterateKeyValue() Tree().


ITERATION TOOLS

Iteration tools provide functions for access each key-value pair or each values of an array

Let's look at IterateArray() function.

	// JSON that we want to access.
	json := []byte({"user":"eco","languages":["go","java","python","C","Cpp"]})

	// IterateArray() uses a callback function to return elements.
	err := jin.IterateArray(json, func(value []byte) bool {

		// printing current element as string.
		fmt.Println(string(value))

		// this return is some kind of control mechanism for escape from iteration any time.
		// true means keep iterate. false means stop the iteration.
		return true
	// last parameter is path. Its currently pointing at "language" array.
	}, "languages")

	// standard error definition
	if err != nil {
		log.Println(err)
		return
	}

	// Output: go
	// java
	// python
	// C
	// Cpp


Another useful function is IterateKeyValue(). Description and examples in here..

OTHER TOOLS

FORMATTING

There are two formatting functions. Flatten() and Indent()

Both of them have their own examples in here...

 Control Functions

Control functions are simple and easy way to check value types of any path.

For example. IsArray().


	json := []byte({"repo":{"name":"ecoshub/jin"},"others":["jin","penman"]})

	result, _ := jin.IsArray(json, "repo")
	fmt.Println(result)
	// Output: false

	result, _ = jin.IsArray(json, "others")
	fmt.Println(result)
	// Output: true



Or you can use GetType().


	json := []byte({"git":"ecoshub","repo":["jin","wsftp","penman"]})

	result, _ := jin.GetType(json, "repo")
	fmt.Println(result)
	// Output: array


JSON BUILD TOOLS

There are lots of JSON build functions in this package and all of them has its own examples.

We just want to mention a couple of them.

Scheme is simple and powerful tool for create JSON schemes.


	// MakeScheme need keys for construct a JSON scheme.
	person := MakeScheme("name", "lastname", "age")

	// now we can instantiate a JSON with values.
	eco := person.MakeJson("eco", "hub", "28")
	// {"name":"eco","lastname":"hub","age":28}

	koko := person.MakeJson("koko", "Bloom", "42")
	//{"name":"koko","lastname":"Bloom","age":42}



MakeJson(), MakeArray() functions and other variations is easy to use functions. Go and take a look.  in here..


TESTING

Testing is very important for this type of packages and it shows how reliable it is.

For that reasons we use Node.js for unit testing.

Lets look at folder arrangement and working principle.

- test/ folder:

	- test-json.json, this is a temporary file for testing. all other test-cases copying here with this name so they can process by test-case-creator.js.

	- test-case-creator.js is core path & value creation mechanism.	When it executed with executeNode() function. It reads the test-json.json file and generates the paths and values from this files content. With command line arguments it can generate different paths and values. As a result, two files are created. the first of these files is test-json-paths.json and the second is test-json-values.json

	- test-json-paths.json keeps all the path values.

	- test-json-values.json keeps all the values that corresponding to path values.

- tests/ folder

	- All files in this folder is a test-case. But it doesn't mean that you can't change anything, on the contrary, all test-cases are creating automatically based on this folder content. You can add or remove any .json file that you want.

	- All GO side test-case automation functions are in core_test.go file.

This package developed with Node.js v13.7.0. please make sure that your machine has a valid version of Node.js befoure testing.

All functions and methods are tested with complicated randomly created .json files.

Like this.

	{
		"g;}\\=LUG[5pwAizS!lfkdRULF=": true,
		"gL1GG'S+-U~fUz^R^=genWFVGA$O": {
			"Nmg}xK&V5Z": -1787764711,
			"=B7a(KoF%m5rqGEn}dl\"y117)WC&w~": -572664066,
			"Dj_{6evoMr&< 4m+1u{W!'zf;cl": ":mqp<s6('&??yG)qpMs=H?",
			",Qx_5V(ceN)%0d-h.\"\"0v}8fqG-zgEBz;!C{zHZ9Hfg%no": false,
			"l&d>": true
		},
		"jhww/SRq?,Y\"5O1'{": "]\"4s{WH]b9aR+[$-'PQm8WW:B",
		":e": "Lu9(>9IbrLyx60E;9R]NHml@A~} QHgAUR5$TUCm&z,]d\">",
		"e&Kk^rzT!EZopgIo\\5)GT'MkSCf]2<{dt+C_H": 599287421.0854483
	}


Most of JSON packages not even run properly with this kind of JSON streams.
We did't see such packages as competitors to ourselves.
And that's because we didn't even bother to benchmark against them.


BENCHMARK

Benchmark results.


- Benchmark prefix removed from function names for make room to results.
- Benchmark between 'buger/jsonparser' and 'ecoshub/jin' use the same payload (JSON test-cases) that 'buger/jsonparser' package use for benchmark it self.

	github.com/ecoshub/jin		-> Jin
	github.com/buger/jsonparser	-> Jsonparser
	github.com/valyala/fastjson	-> Fastjson
	github.com/json-iterator/go	-> Jsoniterator


	goos: linux
	goarch: amd64
	pkg: jin/benchmark

	// Get Function.
	JsonparserGetSmall-8                  826 ns/op             0 B/op        0 allocs/op
	JinGetSmall-8                         792 ns/op             0 B/op        0 allocs/op
	JsonparserGetMedium-8                7734 ns/op             0 B/op        0 allocs/op
	JinGetMedium-8                       5793 ns/op             0 B/op        0 allocs/op
	JsonparserGetLarge-8                62319 ns/op             0 B/op        0 allocs/op
	JinGetLarge-8                       56575 ns/op             0 B/op        0 allocs/op

	// Set Function.
	JsonParserSetSmall-8                 1268 ns/op           704 B/op        4 allocs/op
	JinSetSmall-8                        1213 ns/op           704 B/op        4 allocs/op
	JsonParserSetMedium-8                7014 ns/op          6912 B/op        3 allocs/op
	JinSetMedium-8                       5767 ns/op          6912 B/op        3 allocs/op
	JsonParserSetLarge-8               126726 ns/op        114688 B/op        4 allocs/op
	JinSetLarge-8                       87239 ns/op        114688 B/op        4 allocs/op

	// Delete Function.
	JsonParserDeleteSmall-8              2092 ns/op           704 B/op        4 allocs/op
	JinDeleteSmall-8                     1211 ns/op           640 B/op        4 allocs/op
	JsonParserDeleteMedium-8            11096 ns/op          6912 B/op        3 allocs/op
	JinDeleteMedium-8                    5429 ns/op          6144 B/op        3 allocs/op
	JsonParserDeleteLarge-8            130838 ns/op        114688 B/op        4 allocs/op
	JinDeleteLarge-8                    85999 ns/op        114688 B/op        4 allocs/op

	// Iterators Function.
	IterateArrayGetJsonparser-8         12296 ns/op             0 B/op        0 allocs/op
	IterateArrayGetJin-8                11441 ns/op             0 B/op        0 allocs/op
	IterateObjectGetJsonparser-8         6381 ns/op             0 B/op        0 allocs/op
	IterateObjectGetJin-8                4638 ns/op             0 B/op        0 allocs/op

	// Parser Get Small Function.
	JsoniteratorGetSmall-8               4006 ns/op           874 B/op        1 allocs/op
	FastjsonGetSmall-8                   2773 ns/op          3408 B/op        1 allocs/op
	JinParseGetSmall-8                   2040 ns/op          1252 B/op        8 allocs/op

	// Parser Get Medium Function.
	JsoniteratorGetMedium-8             29936 ns/op          9730 B/op        5 allocs/op
	FastjsonGetMedium-8                 16190 ns/op         17304 B/op        4 allocs/op
	JinParseGetMedium-8                 14016 ns/op          8304 B/op        1 allocs/op

	// Parser Get Large Function.
	JsoniteratorGetLarge-8             634964 ns/op        219307 B/op        3 allocs/op
	FastjsonGetLarge-8                 221918 ns/op        283200 B/op        0 allocs/op
	JinParseGetLarge-8                 218904 ns/op        134704 B/op        3 allocs/op

	// Parser Set Function.
	FastjsonSetSmall-8                   3662 ns/op          3792 B/op        9 allocs/op
	JinParseSetSmall-8                   3382 ns/op          1968 B/op        6 allocs/op



LIMITATIONS

Jin can handle all kind of JSON. Except single content JSONs

Like those:

	{"golang"}
	{42}
	{false}


That kind of JSONs are forbidden.



UPCOMING

We are currently working on,

- Marshall() and Unmarshall() functions.

- http.Request parser/interpreter

- Builder functions for http.ResponseWriter



CONTRIBUTE

If you want to contribute this work feel free to fork it.

We want to fill this section with contributors.
*/
package jin
