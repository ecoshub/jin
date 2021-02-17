[![Go Report Card](https://goreportcard.com/badge/github.com/ecoshub/jin)](https://goreportcard.com/report/github.com/ecoshub/jin) ![License](https://img.shields.io/dub/l/vibe-d.svg) [![GoDoc](https://godoc.org/github.com/ecoshub/jin?status.svg)](https://godoc.org/github.com/ecoshub/jin)

# Welcome To JIN

<p align="center">
  <img src="mascot.svg" width="640" height="640">
</p>

#### _"Your wish is my command"_

## Fast and Easy Way to Deal With JSON

**Jin** is a comprehensive JSON manipulation tool bundle.
All functions tested with random data with help of **Node.js**.
All test-path and test-value creation automated with **Node.js**.

**Jin** provides `parse`, `interpret`, `build` and `format` tools for JSON.
Third-party packages only used for the benchmark. No dependency need for core functions.

We make some benchmark with other packages like **Jin**.

```
    github.com/buger/jsonparser
    github.com/valyala/fastjson
    github.com/json-iterator/go
    github.com/tidwall/gjson
    github.com/tidwall/sjson
```

In Result, **Jin** is the fastest (op/ns) and more memory friendly then others (B/op).

For more information please take a look at **BENCHMARK** section below.

---

### What is New?

##### _04.01.2021_

**JO (JsonObject)** introduced!! Actually that is a fancy word for []byte type

You can use all interpreter functions with JO. just initialize and go.

`Get()` Example

```go
	// instead of repeating []byte for all functions.
	// initialize
	jsonObject := jin.New(json)
	// and go
	serial, err := jsonObject.GetString("info", "serial")
	if err != nil {
		return err
	}
	pooling, err := jsonObject.GetFloat("info", "polling_time")
	if err != nil {
		return err
	}

	// old declaration
	serial, err := jin.GetString(json, "info", "serial")
	if err != nil {
		return err
	}
	pooling, err := jin.GetFloat(json, "info", "polling_time")
	if err != nil {
		return err
	}

```

`Set()` Example

```go
	// instead of repeating []byte for all functions.
	// initialize
	jsonObject := jin.New(json)
	// and go
	err := jsonObject.SetString("at-28C02", "info", "serial")
	if err != nil {
		return err
	}

	// old declaration
	json, err := jin.SetString(json, "at-28C02", "info", "serial")
	if err != nil {
		return err
	}

```

##### _06.04.2020_

**7 new** functions **tested** and **added** to package. Examples in **[GoDoc](https://godoc.org/github.com/ecoshub/jin)**

-   `GetMap()` get objects as `map[string]string` structure with key values pairs
-   `GetAll()` get only specific keys values
-   `GetAllMap()` get only specific keys with `map[string]string`structure
-   `GetKeys()` get objects keys as string array
-   `GetValues()` get objects values as string array
-   `GetKeysValues()` get objects keys and values with separate string arrays
-   `Length()` get length of JSON array.

---

### Installation

```
	go get github.com/ecoshub/jin
```

And you are good to go. Import and start using.

---

### Documentation

There is a detailed documentation in **[GoDoc](https://godoc.org/github.com/ecoshub/jin)** with lots of examples.

---

### QUICK START

#### Parser vs Interpreter

Major difference between parsing and interpreting is
parser has to read all data before answer your needs.
On the other hand interpreter reads up to find the data you need.

With parser, once the parse is complete you can access any data with no time.
But there is a time cost to parse all data and this cost can increase as data content grows.

If you need to access all keys of a JSON then, we are simply recommend you to use `Parser`.
But if you need to access some keys of a JSON then we strongly recommend you to use `Interpreter`, it will be much faster and much more memory-friendly than parser.

#### Interpreter

`Interpreter` is core element of this package, no need to create an Interpreter type, just call which function you want.

First let's look at general function parameters.

```go

	// All interpreter functions need one JSON as byte slice.
	json := []byte(`{"git":"ecoshub","repo":{"id":233809925,"name":["eco","jin"]}}`)

	// And most of them needs a path value for navigate.
	// Path value determines which part to navigate.
	// In this example we want to access 'jin' value.
	// So path must be 'repo' (object) -> 'name' (array) -> '1' (second element)
	path := []string{"repo", "name", "1"}

```

We are gonna use `Get()` function to access the value of path has pointed. In this case 'jin'.

```go

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

```

Path value can consist hard coded values.

```go

	value, err := jin.Get(json, "repo", "name", "1")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(value))
	//String Output: jin

```

`Get()` function return type is `[]byte` but all other variations of return types are implemented with different functions.

For example. If you need "value" as string use `GetString()`.

```go

	value, err := jin.GetString(json, "repo", "name", "0")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(value))
	//Output: eco

```

For example. If you need "value" as string use `GetString()`.

```go

	value, err := jin.GetString(json, "repo", "name", "0")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(value))
	//Output: eco

```

---

#### Parser

`Parser` is another alternative for JSON manipulation.

We recommend to use this structure when you need to access all or most of the keys in the JSON.

Parser constructor need only one parameter.

```go

	// Parser constructor function jin.Parse() need one JSON as []byte.
	json := []byte(`
	{
		"title": "LICENSE",
		"repo": {
			"id": 233809925,
			"name": "ecoshub/jin",
			"url": "https://api.github.com/repos/ecoshub/jin"
			}
	}`)

```

We can parse it with `Parse()` function.

```go

	prs, err := jin.Parse(json)
	if err != nil {
		log.Println(err)
		return
	}

```

Let's look at `Parser.Get()`

```go

	value, err := prs.Get("repo", "url")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(value))
	//Output: https://api.github.com/repos/ecoshub/jin

```

_About path value look above._

There is all return type variations of `Parser.Get()` function like `Interpreter`.

For return string use `Parser.GetString()` like this,

```go

	value, err := prs.GetString("repo", "name")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(value)
	//String Output: ecoshub/jin

```

All functions has own example provided in **[GoDoc](https://godoc.org/github.com/ecoshub/jin)**.

**Other usefull functions of Jin.**

-`Add()`, `AddKeyValue()`, `Set()`, `SetKey()` `Delete()`, `Insert()`, `IterateArray()`, `IterateKeyValue()` `Tree()`.

---

### Iteration Tools

Iteration tools provide functions for access each key-value pair or each value of an array

Let's look at `IterateArray()` function.

```go
	// JSON that we want to access.
	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"]}`)

	// IterateArray() uses a callback function to return elements.
	err := jin.IterateArray(json, func(value []byte) (bool, error) {

		// printing current element as string.
		fmt.Println(string(value))

		// this return is some kind of control mechanism for escape from iteration any time.
		// true means keep iterate. false means stop the iteration.
		return true, nil
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

```

Another useful function is `IterateKeyValue()`check for example in **[GoDoc](https://godoc.org/github.com/ecoshub/jin)**.

---

### Other Tools

#### Formatting

There are two formatting functions. `Flatten()` and `Indent()`.

`Indent()` is adds indentation to JSON for nicer visualization and `Flatten()` removes this indentation.

Examples in **[GoDoc](https://godoc.org/github.com/ecoshub/jin)**.

#### Control Functions

Control functions are simple and easy way to check value types of any path.

For example. `IsArray()`.

```go

	json := []byte(`{"repo":{"name":"ecoshub/jin"},"others":["jin","penman"]}`)

	result, _ := jin.IsArray(json, "repo")
	fmt.Println(result)
	// Output: false

	result, _ = jin.IsArray(json, "others")
	fmt.Println(result)
	// Output: true


```

Or you can use `GetType()`.

```go

	json := []byte(`{"git":"ecoshub","repo":["jin","wsftp","penman"]}`)

	result, _ := jin.GetType(json, "repo")
	fmt.Println(result)
	// Output: array

```

#### JSON Build Tools

There are lots of JSON build functions in this package and all of them has its own examples.

We just want to mention a couple of them.

`Scheme` is simple and powerful tool for create JSON schemes.

```go

	// MakeScheme need keys for construct a JSON scheme.
	person := MakeScheme("name", "lastname", "age")

	// now we can instantiate a JSON with values.
	eco := person.MakeJson("eco", "hub", "28")
	// {"name":"eco","lastname":"hub","age":28}

	koko := person.MakeJson("koko", "Bloom", "42")
	//{"name":"koko","lastname":"Bloom","age":42}

```

`MakeJson()`, `MakeArray()` functions and other variations are easy to use functions. Go and take a look. **[GoDoc](https://godoc.org/github.com/ecoshub/jin)**.

---

### Testing

Testing is very important thing for this type of packages and it shows how reliable it is.

For that reasons we use **Node.js** for unit testing.

Lets look at folder arrangement and working principle.

-   **test/** folder:

    -   **test-json.json**, this is a temporary file for testing. all other test-cases copying here with this name so they can process by **test-case-creator.js**.

    -   **test-case-creator.js** is core path & value creation mechanism. When it executed with `executeNode()` function. It reads the **test-json.json** file and generates the paths and values from this files content. With command line arguments it can generate different paths and values. As a result, two files are created with this process. the first of these files is **test-json-paths.json** and the second is **test-json-values.json**

    -   **test-json-paths.json** has all the path values.

    -   **test-json-values.json** has all the values that corresponding to path values.

-   **tests/** folder

    -   All files in this folder is a test-case. But it doesn't mean that you can't change anything, on the contrary, all test-cases are creating automatically based on this folder content. You can add or remove any **.json** file that you want.

    -   All `GO` side test-case automation functions are in **core_test.go** file.

This package developed with **Node.js** v13.7.0. please make sure that your machine has a valid version of **Node.js** before testing.

All functions and methods are tested with complicated randomly genereted **.json** files.

Like this,

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

Most of JSON packages not even run properly with this kind of JSON streams.
We did't see such packages as competitors to ourselves.
And that's because we didn't even bother to benchmark against them.

---

### Benchmark

Benchmark results.

-   _Benchmark prefix removed from function names for make room to results._
-   Benchmark between 'buger/jsonparser' and 'ecoshub/jin' use the same payload (JSON test-cases) that 'buger/jsonparser' package use for benchmark it self.

    _github.com/ecoshub/jin -> Jin_

    _github.com/buger/jsonparser -> Jsonparser_

    _github.com/valyala/fastjson -> Fastjson_

    _github.com/json-iterator/go -> Jsoniterator_

    _github.com/tidwall/gjson -> gjson_

    _github.com/tidwall/sjson -> sjon_

```go

	goos: linux
	goarch: amd64
	pkg: jin/benchmark

	// Get function (interpert)
	JsoniteratorGetSmall-8             2862 ns/op          597 B/op          40 allocs/op
	GjsonGetSmall-8                     921 ns/op           64 B/op           3 allocs/op
	JsonparserGetSmall-8                787 ns/op            0 B/op           0 allocs/op
	JinGetSmall-8                       729 ns/op            0 B/op           0 allocs/op
	GjsonGetMedium-8                   7084 ns/op          152 B/op           4 allocs/op
	JsonparserGetMedium-8              7329 ns/op            0 B/op           0 allocs/op
	JinGetMedium-8                     5624 ns/op            0 B/op           0 allocs/op
	GjsonrGetLarge-8                 119925 ns/op        28672 B/op           2 allocs/op
	JsonparserGetLarge-8              65725 ns/op            0 B/op           0 allocs/op
	JinGetLarge-8                     61516 ns/op            0 B/op           0 allocs/op

	// Array iteration function (interpert)
	IterateArrayGetGjson-8            21966 ns/op         8192 B/op           1 allocs/op
	IterateArrayGetJsonparser-8       11814 ns/op            0 B/op           0 allocs/op
	IterateArrayGetJin-8              11639 ns/op            0 B/op           0 allocs/op
	IterateObjectGetGjson-8           11329 ns/op         2304 B/op           1 allocs/op
	IterateObjectGetJsonparser-8       6100 ns/op            0 B/op           0 allocs/op
	IterateObjectGetJin-8              4551 ns/op            0 B/op           0 allocs/op

	// Set function (interpert)
	SJonSetSmall-8                     2126 ns/op         1664 B/op           9 allocs/op
	JsonParserSetSmall-8               1261 ns/op          704 B/op           4 allocs/op
	JinSetSmall-8                      1244 ns/op          704 B/op           4 allocs/op
	SjsonSetMedium-8                  15412 ns/op        13008 B/op          11 allocs/op
	JsonParserSetMedium-8              6868 ns/op         6912 B/op           3 allocs/op
	JinSetMedium-8                     6169 ns/op         6912 B/op           3 allocs/op
	SjsonSetLarge-8                  257300 ns/op       136736 B/op          14 allocs/op
	JsonParserSetLarge-8             121874 ns/op       114688 B/op           4 allocs/op
	JinSetLarge-8                     86574 ns/op       114688 B/op           4 allocs/op

	// Delete function (interpert)
	JsonParserDeleteSmall-8            2015 ns/op          704 B/op           4 allocs/op
	JinDeleteSmall-8                   1198 ns/op          640 B/op           4 allocs/op
	JsonParserDeleteMedium-8          10321 ns/op         6912 B/op           3 allocs/op
	JinDeleteMedium-8                  5780 ns/op         6144 B/op           3 allocs/op
	JsonParserDeleteLarge-8          123737 ns/op       114688 B/op           4 allocs/op
	JinDeleteLarge-8                  87322 ns/op       114688 B/op           4 allocs/op

	// Get function (parse)
	FastjsonGetSmall-8                 2755 ns/op         3408 B/op          11 allocs/op
	JinParseGetSmall-8                 1981 ns/op         1252 B/op          28 allocs/op
	FastjsonGetMedium-8               14958 ns/op        17304 B/op          54 allocs/op
	JinParseGetMedium-8               14175 ns/op         8304 B/op         201 allocs/op
	FastjsonGetLarge-8               229188 ns/op       283200 B/op         540 allocs/op
	JinParseGetLarge-8               222246 ns/op       134704 B/op        2903 allocs/op

	// Set function (parse)
	FastjsonSetSmall-8                 3709 ns/op         3792 B/op          19 allocs/op
	JinParseSetSmall-8                 3265 ns/op         1968 B/op          36 allocs/op

```

---

### Upcoming

We are currently working on,

-   `Marshal()` and `Unmarshal()` functions.

-   http.Request parser/interpreter

-   Builder functions for http.ResponseWriter

---

### Contribute

If you want to contribute this work feel free to fork it.

We want to fill this section with contributors.
