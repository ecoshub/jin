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

	https://github.com/buger/jsonparser (with interpreter)
	https://github.com/valyala/fastjson (with parser)
	https://github.com/json-iterator/go (with parser)

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

ITERATION TOOLS

Iteration tools provide funcions for access each key-value pair or each values of an array

For example let's look at IterateKey() function.

	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"]}`)

	err := jin.IterateArray(json, func(value []byte) bool {
	    fmt.Println(string(value))
	    // this return is some kind control mechanism for escape the iteration any time you want.
	    // true means keep iterate. false means stop iteration.
	    return true
	}, "languages")

	if err != nil {
	    fmt.Println(err)
	    return
	}

	//Output: 
	//go
	//java
	//python
	//C

IterateKeyValue() function description and examples are in below.

OTHER TOOLS

FORMATTING

There are two formatting functions. Flatten() and Indent()
Both of them have their own examples on below.

CONTROL FUNCTIONS

Control functions are simple and easy way to check value types of any path.
For example. jin.IsArray() controls the path, if path is points to an array.
It will return true

	json := []byte(`{"repo":{"name":"ecoshub/jin"},"others":["jin","penman"]}`)

	fmt.Println(jin.IsArray(json, "repo"))
	// Output: false

	fmt.Println(jin.IsArray(json, "others"))
	// Output: true

Or get value type of the path

	json := []byte(`{"git":"ecoshub","repo":["jin","wsftp","penman"]}`)

	fmt.Println(jin.GetType(json, "repo"))
	// Output: array



TESTING

Test files are in the /test directory.
Main test function needs NODEJS for path and value creation.
Befour make any test be sure that your machine has a valid version of NODEJS.
This package developed with NODEJS v13.7.0.

If you want to test another JSON file that is not in the tests folder.
Just drag and drop it to the tests folder all process is automated.

BENCHMARK

Benchmark results.

	goos: linux
	goarch: amd64
	pkg: jin/benchmark

	// Delete Function.
	BenchmarkJsonParserDeleteSmall-8    	  543441	    2092 ns/op	 	   704 B/op	4 allocs/op
	BenchmarkJinDeleteSmall-8           	  998816	    1211 ns/op	 	   640 B/op	4 allocs/op
	BenchmarkJsonParserDeleteMedium-8   	   99340	   11096 ns/op	 	  6912 B/op	3 allocs/op
	BenchmarkJinDeleteMedium-8          	  211236	    5429 ns/op	 	  6144 B/op	3 allocs/op
	BenchmarkJsonParserDeleteLarge-8    	    7863	  130838 ns/op	 	114688 B/op	4 allocs/op
	BenchmarkJinDeleteLarge-8           	   14088	   85999 ns/op	 	114688 B/op	4 allocs/op

	// Get Function.
	BenchmarkJsonparserGetSmall-8       	 1457156	     826 ns/op	 	     0 B/op	0 allocs/op
	BenchmarkJinGetSmall-8              	 1515829	     792 ns/op	 	     0 B/op	0 allocs/op
	BenchmarkJsonparserGetMedium-8      	  148725	    7734 ns/op	 	     0 B/op	0 allocs/op
	BenchmarkJinGetMedium-8             	  201876	    5793 ns/op	 	     0 B/op	0 allocs/op
	BenchmarkJsonparserGetLarge-8       	   18853	   62319 ns/op	 	     0 B/op	0 allocs/op
	BenchmarkJinGetLarge-8                	   20449	   56575 ns/op	 	     0 B/op	0 allocs/op

	// Iterators Function.
	BenchmarkIterateArrayGetJsonparser-8       90210	   12932 ns/op	 	     0 B/op	0 allocs/op
	BenchmarkIterateArrayGetJin-8         	   92814	   12787 ns/op	 	     0 B/op	0 allocs/op
	BenchmarkIterateObjectGetJsonparser-8     173838	    6381 ns/op	 	     0 B/op	0 allocs/op
	BenchmarkIterateObjectGetJin-8      	  248893	    4638 ns/op	 	     0 B/op	0 allocs/op

	// Set Function.
	BenchmarkJsonParserSetSmall-8       	  942067	    1268 ns/op	 	   704 B/op	4 allocs/op
	BenchmarkJinSetSmall-8              	  850646	    1213 ns/op	 	   704 B/op	4 allocs/op
	BenchmarkJsonParserSetMedium-8      	  162196	    7014 ns/op	 	  6912 B/op	3 allocs/op
	BenchmarkJinSetMedium-8             	  207309	    5767 ns/op	 	  6912 B/op	3 allocs/op
	BenchmarkJsonParserSetLarge-8       	    9375	  126726 ns/op	 	114688 B/op	4 allocs/op
	BenchmarkJinSetLarge-8              	   14323	   87239 ns/op	 	114688 B/op	4 allocs/op

	// Parser Get Small Function.
	BenchmarkJsoniteratorGetSmall-8     	  284790	    4006 ns/op	 	   874 B/op	1 allocs/op
	BenchmarkFastjsonGetSmall-8         	  454034	    2773 ns/op	 	  3408 B/op	1 allocs/op
	BenchmarkJinParseGetSmall-8         	  613586	    2040 ns/op	 	  1252 B/op	8 allocs/op

	// Parser Get Medium Function.
	BenchmarkJsoniteratorGetMedium-8    	   38133	   29936 ns/op	 	  9730 B/op	5 allocs/op
	BenchmarkFastjsonGetMedium-8        	   75524	   16190 ns/op	 	 17304 B/op	4 allocs/op
	BenchmarkJinParseGetMedium-8        	   81415	   14016 ns/op	 	  8304 B/op	1 allocs/op

	// Parser Get Large Function.
	BenchmarkJsoniteratorGetLarge-8     	    1912	  634964 ns/op	 	219307 B/op	3 allocs/op
	BenchmarkFastjsonGetLarge-8         	    6583	  221918 ns/op	 	283200 B/op	0 allocs/op
	BenchmarkJinParseGetLarge-8         	    5158	  218904 ns/op	 	134704 B/op	3 allocs/op

	// Parser Set Function.
	BenchmarkFastjsonSetSmall-8         	  345870	    3662 ns/op	 	  3792 B/op	9 allocs/op
	BenchmarkJinParseSetSmall-8         	  388720	    3382 ns/op	 	  1968 B/op	6 allocs/op



*/
package jin
