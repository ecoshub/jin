
# Change Log
All notable changes to this project will be documented in this file.
 
## 08.10.2021
### float notation change:

-	float notation changed to 'f' from 'e'. '2e10-2' -> '0.02'

example:
```go
    if jin.ErrEqual(err, jin.ErrCodeKeyNotFound) {
        // err is a KeyNotFound error
    }
```

### jin object (JO) type->struct correction

### ReadJSONFile -> ReadFile function name allias added

## [NEW]:

**Store** New! store function. store function can set or override a value like a real JSON object.

implemented for **parser**, **interpreter** and **JO**

```go
    json := []byte(`{"name":"eco","lastname":"hub","arr":[0,1,2,3,4,5]}`)
	var err error
    // adds age key to value 30
	json, err = jin.StoreInt(json, "age", 30)
	if err != nil {
        fmt.Println(err)
		return
	}
    // sets age value to 40
	json, err = jin.StoreInt(j, "age", 40)
	if err != nil {
		fmt.Println(err)
		return
	}
```

### 21.11.2021:
-   Benchmark files separated from master.
