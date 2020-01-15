package jsoninterpreter

import "errors"

func NULL_PATH_ERROR() 			error {return errors.New("Error: Path can't be null. ERROR_CODE:00")}
func NULL_NEW_VALUE_ERROR() 	error {return errors.New("Error: New value can't be null. ERROR_CODE:01")}

func INDEX_EXPECTED_ERROR() 	error {return errors.New("Error: Index expected, got key value. ERROR_CODE:02")}
func KEY_EXPECTED_ERROR() 		error {return errors.New("Error: Key value expected, got index. ERROR_CODE:03")}
func OBJECT_EXPECTED_ERROR() 	error {return errors.New("Error: Last path must be pointed at an object. ERROR_CODE:04")}
func ARRAY_EXPECTED_ERROR() 	error {return errors.New("Error: Last path must be pointed at an array. ERROR_CODE:05")}

func INDEX_OUT_OF_RANGE_ERROR() error {return errors.New("Error: Index out of range. ERROR_CODE:06")}
func KEY_NOT_FOUND_ERROR() 		error {return errors.New("Error: Key not found. ERROR_CODE:07")}
func BAD_JSON_ERROR() 			error {return errors.New("Error: Bad JSON format. ERROR_CODE:08")}
func BAD_KEY_ERROR() 			error {return errors.New("Error: Key value can't contain quote symbol. ERROR_CODE:09")}
func KEY_ALREADY_EXIST() 		error {return errors.New("Error: Key already exist. Use Set() function to change. ERROR_CODE:10")}

func CAST_INT_ERROR() 			error {return errors.New("Cast Error: Cast to int error. ERROR_CODE:11")}
func CAST_FLOAT_ERROR() 		error {return errors.New("Cast Error: Cast to float error. ERROR_CODE:12")}
func CAST_BOOL_ERROR() 			error {return errors.New("Cast Error: Cast to bool error. ERROR_CODE:13")}
func CAST_STRING_ARRAY_ERROR() 	error {return errors.New("Cast Error: Cast to []string error. ERROR_CODE:14")}
func CAST_INT_ARRAY_ERROR() 	error {return errors.New("Cast Error: Cast to []int error. ERROR_CODE:15")}
func CAST_FLOAT_ARRAY_ERROR() 	error {return errors.New("Cast Error: Cast to []float error. ERROR_CODE:16")}
func CAST_BOOL_ARRAY_ERROR() 	error {return errors.New("Cast Error: Cast to []bool error. ERROR_CODE:17")}

func EMPTY_ARRAY_ERROR() 				error {return errors.New("Error: Array is empty. ERROR_CODE:18")}