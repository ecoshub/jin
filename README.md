[![Go Report Card](https://goreportcard.com/badge/github.com/ecoshub/jin)](https://goreportcard.com/report/github.com/ecoshub/jin) ![License](https://img.shields.io/dub/l/vibe-d.svg) [![GoDoc](https://godoc.org/github.com/ecoshub/jin?status.svg)](https://godoc.org/github.com/ecoshub/jin)

## Fast and Easy Way to Deal With JSON

__Jin__ is a comprehensive JSON manipulation tool bundle.
All functions tested with random data against __Node.js__.
All test-path and test-value cases created automatically with __Node.js__.

It provides parse, interpret, build and format tools for JSON.
Third-party packages only used for the benchmark. No dependency need for core functions.

---

Benchmark against.
```
    github.com/buger/jsonparser
    github.com/valyala/fastjson
    github.com/json-iterator/go
```
In conclusion, the result of the benchmark __Jin__ was the fastest and lightweight package.

For more information take a look at __BENCHMARK__ section.

---

### Documentation

There is a detailed doctumentation in __[GoDoc](https://godoc.org/github.com/ecoshub/jin)__ with lots of examples in it.

---

### QUICK START

### Parser vs Interpreter

Major difference between parsing and interpreting is,
parser has to read all data before answer your needs.
On the other hand interpreter reads up to find the path you need.

Once the parse is complete you can get access any data with no time.
But there is a time cost to parse data, and this cost can increase as data content grows.

If you need to access all keys of a JSON then we are simply recommend you to use Parser.
But if you need to access some keys of a JSON I strongly recommend you to use Interperter, it will be much faster than parser. 

#### Interpreter

Interpreter is core element of this package, no need for instantiate, just call which function you want!

First let's look at function parameters.
```go
	// All interpreter functions need one JSON as byte slice type. 
	json := []byte(`{"git":"ecoshub","repo":{"id":233809925,"name":"ecoshub/jin"}}`)

	// And most of them needs a path value for navigate.
	path := []string{"repo", "name"}
```
Let's take the function `Get()`

Get function returns the value that path has pointed.
```go
	value, err := jin.Get(json, path...)
	if err != nil {
		return err
	}
	//String Output: ecoshub/jin
```
path variable can be a string slice or hard coded
```go
	value, err := jin.Get(json, "repo", "name")
	if err != nil {
		return err
	}
	//String Output: ecoshub/jin
```
`Get()` function return type is byte slice.

All variations of return types are implemented as different functions.

And slice types. (`GetStringArray()`, `GetIntArray()`, etc.)

For example. If you need 'value' as string,

then you can use `GetString()` like this.
```go
	value, err := jin.GetString(json, "repo", "name")
	if err != nil {
		return err
	}
	//Output: ecoshub/jin
```
---
#### Parser

Parser is another alternative for JSON manipulation.

We recommend to use this structure when you need to access all or most of the keys in the JSON.

Parser constructor need only one parameter.
```go
	// Parser constructor function jin.Parse() need one JSON as byte slice format. 
	json := []byte(`{"git":"ecoshub","repo":{"id":233809925,"name":"ecoshub/jin"}}`)
```
Lets Parse it with Parse function.
```go
	prs, err := jin.Parse(json)
	if err != nil {
		return err
	}
```
Let's look at Parser.Get()
```go
	value, err := prs.Get("repo")
	if err != nil {
		return err
	}
	//String Output: {"id":233809925,"name":"ecoshub/jin"}
```
*About path value look above.* 

`Parser.Get()` functions return type is byte slice like `Get()` function of interpreter.

Like interpreter, parser has all variations of return types implemented as different functions to.

Even slice types. (`GetStringArray()`, `GetIntArray()`, etc.)

For example. If you need 'value' as string.

Then you can use `Parser.GetString()` like this.
```go
	value, err := prs.GetString("repo")
	if err != nil {
		return err
	}
	//String Output: {"id":233809925,"name":"ecoshub/jin"}
```
All interpreter/parser functions (except function variations line `GetString()`) has own example provided in __[GoDoc](https://godoc.org/github.com/ecoshub/jin)__.

Other important functions of interpreter/parser. 

-`Add()`, `AddKeyValue()`, `Set()`, `SetKey()` `Delete()`, `Insert()`, `IterateArray()`, `IterateKeyValue()`

---

### Iteration Tools

Iteration tools provide funcions for access each key-value pair or each values of an array

Let's look at `IterateArray()` function.
```go
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
	// Output: go
	// java
	// python
	// C
	// Cpp
```
Another usefull function is `IterateKeyValue()`. Description and examples in __[GoDoc](https://godoc.org/github.com/ecoshub/jin)__.

---

### Other Tools

#### Formatting

There are two formatting functions. Flatten() and Indent()

Both of them have their own examples in Documentation.

#### Control Functions

Control functions are simple and easy way to check value types of any path.

For example. `jin.IsArray()` controls the path, if path is points to an array.

It will return true
```go
	json := []byte(`{"repo":{"name":"ecoshub/jin"},"others":["jin","penman"]}`)

	fmt.Println(jin.IsArray(json, "repo"))
	// Output: false

	fmt.Println(jin.IsArray(json, "others"))
	// Output: true
```
Or get value type of the path
```go
	json := []byte(`{"git":"ecoshub","repo":["jin","wsftp","penman"]}`)

	fmt.Println(jin.GetType(json, "repo"))
	// Output: array
```
#### JSON Build Tools

There are lots of JSON build functions in this package and all of them has its own examples.

We just want to mention a couple of them.

`Scheme` is simple and powerfull tool for create JSON schemes.
```go
	// MakeScheme need keys for construct a JSON scheme.
	person := MakeScheme("name", "lastname", "age")

	// now we can instantiate a JSON with values.
	eco := person.MakeJson("eco", "hub", "28")
	// {"name":"eco","lastname":"hub","age":28}

	koko := person.MakeJson("koko", "Bloom", "42")
	//{"name":"koko","lastname":"Bloom","age":42}
```

`MakeJson()`, `MakeArray()` functions and other variations is easy to use functions.

Go and take a look.  __[GoDoc](https://godoc.org/github.com/ecoshub/jin)__.


---

### Testing

Almost all functions/methods tested with complicated randomly creted JSONs.

Like This.
```go
	{
		"g;}\\=LUG[5pwAizS!lfkdRULF=": true,
		"gL1GG'S+-U~#fUz^R^=#genWFVGA$O": {
			"Nmg}xK&V5Z": -1787764711,
			"=B7a(KoF%m5rqG#En}dl\"y`117)WC&w~": -572664066,
			"Dj_{6evoMr&< 4m+1u{W!'zf;cl": ":mqp<s6('&??yG#)qpMs=H?",
			",Qx_5V(ceN)%0d-h.\"\"0v}8fqG-zgEBz;!C{zHZ#9Hfg%no*": false,
			"l&d>": true
		},
		"jhww/SRq?,Y\"5O1'{": "]\"4s{WH]b9aR+[$-'PQm8WW:B",
		":e": "Lu9(>9IbrLyx60E;9R]NHml@A~} QHgAUR5$TUCm&z,]d\">",
		"e&Kk^`rz`T!EZopgIo\\5)GT'MkSCf]2<{dt+C_H": 599287421.0854483
	}
```
Some packages not event runs properly with this kind of JSONS.
We did not see such packages as competitors to ourselves.
And that's because we didn't even bother to benchmark against them.


Test files are in the /tests directory.
Main test function needs __Node.js__ for path and value creation.
Before make any test be sure that your machine has a valid version of __Node.js__.

This package developed with __Node.js__ v13.7.0.

If you want to test another JSON file that is not in the /tests folder just drag and drop it to the /tests folder, and ron `go test`.

---

### Benchmark

Benchmark results.

- *Benchmark prefix removed to make more room for results.*
- *Benchmark between 'buger/jsonparser' and 'ecoshub/jin' use the same payload (JSON test cases) that 'buger/jsonparser' package use for benchmark it self.*
```go
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
	IterateArrayGetJsonparser-8         12932 ns/op             0 B/op        0 allocs/op 
	IterateArrayGetJin-8                12787 ns/op             0 B/op        0 allocs/op 
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
```

---

### Limitations

__Jin__ can handle all kind of JSON. Except single content JSONs

Like those:
```go
	{"golang"}
	{42}
	{false}
```
That kind of JSONs are forbidden.

---

### Upcomming

We are currently working on, 

- `Marshall()` and `Unmarshall()` functions.

- http.Request parser/interperter

- Builder functions for http.ResponseWriter

---

### Contribute

If you want to contribute this work feel free to fork it.

We want to fill this section with contributors.
