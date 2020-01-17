package jsoninterpreter

// import (
// 	"testing"
// )

// // func TestAddValueMixed(t *testing.T){
// // 	for i, p := range mixedDummyGetPaths {
// // 		value, done := Get(mixedDummyGetTest, p...)
// // 		if done != nil {
// // 			t.Errorf("Total Fail, path:%v\n", p)
// // 		}
// // 		if string(value) != mixedDummyGetTestResult[i] {
// // 			t.Errorf("Fail, not same answer path:%v, got:%v, expected:%v", p, string(value), mixedDummyGetTestResult[i])
// // 		}
// // 	}
// // }

// func TestAddValueArray(t *testing.T){
//   for i, p := range ArrayAddValueTestPaths {
//     value, done := AddValue(ArrayAddValueTest, []byte(`"test-string"`), p...)
//     if done != nil {
//       t.Errorf("Total Add Fail, path:%v\n", p)
//     }
//     value, done = Get(value, p...)
//     if done != nil {
//     	t.Errorf("Total Get Fail, path:%v\n", p)
//     }
//     if string(value) != ArrayAddValueTestResult[i] {
//     	t.Errorf("Fail, not same answer path:%v, got:%v, expected:%v", p, string(value), ArrayAddValueTestResult[i])
//     }
//   }
// }

// var ArrayAddValueTest []byte = []byte(`[[[[1,[31,62,69],5,[12,13,14,15]],"test-string",[40,41,42]]]]`)

// var ArrayAddValueTestPaths [][]string = [][]string{
//    []string{"0", "0"},
//    []string{"0", "0", "0"},
//    []string{"0", "0", "0", "1"},
//    []string{"0", "0", "0", "2"},
//    []string{"0", "0", "0", "3"},
//    []string{"0", "0", "1"},
//    []string{"0", "0", "2"}}

// var ArrayAddValueTestResult  []string = []string{
// `[[1,[31,62,69],5,[12,13,14,15]],"test-string",[40,41,42]]`,
// `[1,[31,62,69],5,[12,13,14,15],"test-string"]`,
// `[31,62,69,"test-string"]`,
// `[12,13,14,15,"test-string"]`,
// `[40,41,42,"test-string"]`}

// // var mixedDummyGetTest []byte = []byte(` 
// //    [ 
// //     {
// //   "event":"save",
// //   "mac":"bc:ae:c5:13:84:f9",
// //   "username":"eco",
// //   "content":"all",
// //   "bool":false,
// //   "testcase":{
// //     "hi":"guys",
// //     "activeYears":[2005,2020,["emq","john"]]
// //   },
// //   "2":"test-name",
// //   "main":{
// //     "event":"cmsg",
// //     "mac":"bc:ae:c5:13:84:f9",
// //     "msg":"\"hi\" everyone!",
// //     "colors":["red", "blue", "green"],
// //     "done":false,
// //     "id":{
// //       "username":["eco","ecomain"],
// //       "num":9129234,
// //       "id":{
// //         "username":["deadlock","test-name"]
// //       }
// //     },
// //     "eyes":["green","blue"]
// //   },
// //   "id":{"test-number":9129234,"numbers":[31,10,20,[1990,1991,1992],22,32], "active":false},
// //   "test-json":{"test-number":9129234,"test-array":[11,[10,11,12,13],32]},
// //   "UID":{"username":["deadlock"],"numbers":[{"int":"31"},{"float":"3.14"},{"double":"3.0"},{"bool":false}]},
// //   "eyes":["green","blue"],
// //   "bools":[true,false,true]
// // },
// // {
// //   "event":"first",
// //   "mac":"00:00:ba:ba:ba:ba",
// //   "name":"emq"
// // },
// // {
// //   "event":"second",
// //   "mac":"00:00:ba:ba:ba:ba",
// //   "bool":false
// // },
// // {
// //   "event":"third",
// //   "mac":"00:00:ba:ba:ba:ba",
// //   "username":{"name":"emq","surname":"test-name"}
// // },
// // [[1,[31,62,69],5,[12,13,14,15]],"test-string",[40,41,42]]
// // ]`)

// // var mixedDummyGetTestResult  []string = []string{`{
// //   "event":"save",
// //   "mac":"bc:ae:c5:13:84:f9",
// //   "username":"eco",
// //   "content":"all",
// //   "bool":false,
// //   "testcase":{
// //     "hi":"guys",
// //     "activeYears":[2005,2020,["emq","john"]]
// //   },
// //   "2":"test-name",
// //   "main":{
// //     "event":"cmsg",
// //     "mac":"bc:ae:c5:13:84:f9",
// //     "msg":"\"hi\" everyone!",
// //     "colors":["red", "blue", "green"],
// //     "done":false,
// //     "id":{
// //       "username":["eco","ecomain"],
// //       "num":9129234,
// //       "id":{
// //         "username":["deadlock","test-name"]
// //       }
// //     },
// //     "eyes":["green","blue"]
// //   },
// //   "id":{"test-number":9129234,"numbers":[31,10,20,[1990,1991,1992],22,32], "active":false},
// //   "test-json":{"test-number":9129234,"test-array":[11,[10,11,12,13],32]},
// //   "UID":{"username":["deadlock"],"numbers":[{"int":"31"},{"float":"3.14"},{"double":"3.0"},{"bool":false}]},
// //   "eyes":["green","blue"],
// //   "bools":[true,false,true]
// // }`,
// // `save`,
// // `bc:ae:c5:13:84:f9`,
// // `eco`,
// // `all`,
// // `false`,
// // `{
// //     "hi":"guys",
// //     "activeYears":[2005,2020,["emq","john"]]
// //   }`,
// // `guys`,
// // `[2005,2020,["emq","john"]]`,
// // `2005`,
// // `2020`,
// // `["emq","john"]`,
// // `emq`,
// // `john`,
// // `test-name`,
// // `{
// //     "event":"cmsg",
// //     "mac":"bc:ae:c5:13:84:f9",
// //     "msg":"\"hi\" everyone!",
// //     "colors":["red", "blue", "green"],
// //     "done":false,
// //     "id":{
// //       "username":["eco","ecomain"],
// //       "num":9129234,
// //       "id":{
// //         "username":["deadlock","test-name"]
// //       }
// //     },
// //     "eyes":["green","blue"]
// //   }`,
// // `cmsg`,
// // `bc:ae:c5:13:84:f9`,
// // `\"hi\" everyone!`,
// // `["red", "blue", "green"]`,
// // `red`,
// // `blue`,
// // `green`,
// // `{
// //       "username":["eco","ecomain"],
// //       "num":9129234,
// //       "id":{
// //         "username":["deadlock","test-name"]
// //       }
// //     }`,
// // `["eco","ecomain"]`,
// // `eco`,
// // `ecomain`,
// // `9129234`,
// // `{
// //         "username":["deadlock","test-name"]
// //       }`,
// // `["deadlock","test-name"]`,
// // `deadlock`,
// // `test-name`,
// // `{"test-number":9129234,"numbers":[31,10,20,[1990,1991,1992],22,32], "active":false}`,
// // `9129234`,
// // `[31,10,20,[1990,1991,1992],22,32]`,
// // `31`,
// // `10`,
// // `20`,
// // `[1990,1991,1992]`,
// // `1990`,
// // `1991`,
// // `1992`,
// // `22`,
// // `32`,
// // `false`,
// // `{"test-number":9129234,"test-array":[11,[10,11,12,13],32]}`,
// // `9129234`,
// // `[11,[10,11,12,13],32]`,
// // `11`,
// // `[10,11,12,13]`,
// // `10`,
// // `11`,
// // `12`,
// // `13`,
// // `32`,
// // `{"username":["deadlock"],"numbers":[{"int":"31"},{"float":"3.14"},{"double":"3.0"},{"bool":false}]}`,
// // `["deadlock"]`,
// // `deadlock`,
// // `[{"int":"31"},{"float":"3.14"},{"double":"3.0"},{"bool":false}]`,
// // `{"int":"31"}`,
// // `31`,
// // `{"float":"3.14"}`,
// // `3.14`,
// // `{"double":"3.0"}`,
// // `3.0`,
// // `{"bool":false}`,
// // `false`,
// // `["green","blue"]`,
// // `green`,
// // `blue`,
// // `[true,false,true]`,
// // `true`,
// // `false`,
// // `true`,
// // `{
// //   "event":"first",
// //   "mac":"00:00:ba:ba:ba:ba",
// //   "name":"emq"
// // }`,
// // `first`,
// // `00:00:ba:ba:ba:ba`,
// // `emq`,
// // `{
// //   "event":"second",
// //   "mac":"00:00:ba:ba:ba:ba",
// //   "bool":false
// // }`,
// // `second`,
// // `00:00:ba:ba:ba:ba`,
// // `false`,
// // `{
// //   "event":"third",
// //   "mac":"00:00:ba:ba:ba:ba",
// //   "username":{"name":"emq","surname":"test-name"}
// // }`,
// // `third`,
// // `00:00:ba:ba:ba:ba`,
// // `{"name":"emq","surname":"test-name"}`,
// // `emq`,
// // `test-name`,
// // `[[1,[31,62,69],5,[12,13,14,15]],"test-string",[40,41,42]]`,
// // `[1,[31,62,69],5,[12,13,14,15]]`,
// // `1`,
// // `[31,62,69]`,
// // `31`,
// // `62`,
// // `69`,
// // `5`,
// // `[12,13,14,15]`,
// // `12`,
// // `13`,
// // `14`,
// // `15`,
// // `test-string`}

// // var mixedDummyGetPaths [][]string = [][]string{
// //    []string{ "0"},
// //    []string{"0", "event"},
// //    []string{"0", "mac"},
// //    []string{"0", "username"},
// //    []string{"0", "content"},
// //    []string{"0", "bool"},
// //    []string{"0", "testcase"},
// //    []string{"0", "testcase", "hi"},
// //    []string{"0", "testcase", "activeYears"},
// //    []string{"0", "testcase", "activeYears", "0"},
// //    []string{"0", "testcase", "activeYears", "1"},
// //    []string{"0", "testcase", "activeYears", "2"},
// //    []string{"0", "testcase", "activeYears", "2", "0"},
// //    []string{"0", "testcase", "activeYears", "2", "1"},
// //    []string{"0", "2"},
// //    []string{"0", "main"},
// //    []string{"0", "main", "event"},
// //    []string{"0", "main", "mac"},
// //    []string{"0", "main", "msg"},
// //    []string{"0", "main", "colors"},
// //    []string{"0", "main", "colors", "0"},
// //    []string{"0", "main", "colors", "1"},
// //    []string{"0", "main", "colors", "2"},
// //    []string{"0", "main", "id"},
// //    []string{"0", "main", "id", "username"},
// //    []string{"0", "main", "id", "username", "0"},
// //    []string{"0", "main", "id", "username", "1"},
// //    []string{"0", "main", "id", "num"},
// //    []string{"0", "main", "id", "id"},
// //    []string{"0", "main", "id", "id", "username"},
// //    []string{"0", "main", "id", "id", "username", "0"},
// //    []string{"0", "main", "id", "id", "username", "1"},
// //    []string{"0", "id"},
// //    []string{"0", "id", "test-number"},
// //    []string{"0", "id", "numbers"},
// //    []string{"0", "id", "numbers", "0"},
// //    []string{"0", "id", "numbers", "1"},
// //    []string{"0", "id", "numbers", "2"},
// //    []string{"0", "id", "numbers", "3"},
// //    []string{"0", "id", "numbers", "3", "0"},
// //    []string{"0", "id", "numbers", "3", "1"},
// //    []string{"0", "id", "numbers", "3", "2"},
// //    []string{"0", "id", "numbers", "4"},
// //    []string{"0", "id", "numbers", "5"},
// //    []string{"0", "id", "active"},
// //    []string{"0", "test-json"},
// //    []string{"0", "test-json", "test-number"},
// //    []string{"0", "test-json", "test-array"},
// //    []string{"0", "test-json", "test-array", "0"},
// //    []string{"0", "test-json", "test-array", "1"},
// //    []string{"0", "test-json", "test-array", "1", "0"},
// //    []string{"0", "test-json", "test-array", "1", "1"},
// //    []string{"0", "test-json", "test-array", "1", "2"},
// //    []string{"0", "test-json", "test-array", "1", "3"},
// //    []string{"0", "test-json", "test-array", "2"},
// //    []string{"0", "UID"},
// //    []string{"0", "UID", "username"},
// //    []string{"0", "UID", "username", "0"},
// //    []string{"0", "UID", "numbers"},
// //    []string{"0", "UID", "numbers", "0"},
// //    []string{"0", "UID", "numbers", "0", "int"},
// //    []string{"0", "UID", "numbers", "1"},
// //    []string{"0", "UID", "numbers", "1", "float"},
// //    []string{"0", "UID", "numbers", "2"},
// //    []string{"0", "UID", "numbers", "2", "double"},
// //    []string{"0", "UID", "numbers", "3"},
// //    []string{"0", "UID", "numbers", "3", "bool"},
// //    []string{"0", "main", "eyes"},
// //    []string{"0", "main", "eyes", "0"},
// //    []string{"0", "main", "eyes", "1"},
// //    []string{"0", "bools"},
// //    []string{"0", "bools", "0"},
// //    []string{"0", "bools", "1"},
// //    []string{"0", "bools", "2"},
// //    []string{"1"},
// //    []string{"1", "event"},
// //    []string{"1", "mac"},
// //    []string{"1", "name"},
// //    []string{"2"},
// //    []string{"2", "event"},
// //    []string{"2", "mac"},
// //    []string{"2", "bool"},
// //    []string{"3"},
// //    []string{"3", "event"},
// //    []string{"3", "mac"},
// //    []string{"3", "username"},
// //    []string{"3", "username", "name"},
// //    []string{"3", "username", "surname"},
// //    []string{"4"},
// //    []string{"4", "0"},
// //    []string{"4", "0", "0"},
// //    []string{"4", "0", "1"},
// //    []string{"4", "0", "1", "0"},
// //    []string{"4", "0", "1", "1"},
// //    []string{"4", "0", "1", "2"},
// //    []string{"4", "0", "2"},
// //    []string{"4", "0", "3"},
// //    []string{"4", "0", "3", "0"},
// //    []string{"4", "0", "3", "1"},
// //    []string{"4", "0", "3", "2"},
// //    []string{"4", "0", "3", "3"},
// //    []string{"4", "1"}}