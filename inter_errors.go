package jsoninterpreter

import "errors"

func NULL_PATH_ERROR() 			error {return errors.New("ERROR_CODE:00 Error: Path can't be null. ")}
func NULL_NEW_VALUE_ERROR() 	error {return errors.New("ERROR_CODE:01 Error: New value can't be null. ")}
func INDEX_EXPECTED_ERROR() 	error {return errors.New("ERROR_CODE:02 Error: Index expected, got key value. ")}
func KEY_EXPECTED_ERROR() 		error {return errors.New("ERROR_CODE:03 Error: Key expected, got index. ")}
func OBJECT_EXPECTED_ERROR() 	error {return errors.New("ERROR_CODE:04 Error: Last path must be pointed at an object. ")}
func ARRAY_EXPECTED_ERROR() 	error {return errors.New("ERROR_CODE:05 Error: Last path must be pointed at an array. ")}
func INDEX_OUT_OF_RANGE_ERROR() error {return errors.New("ERROR_CODE:06 Error: Index out of range. ")}
func KEY_NOT_FOUND_ERROR() 		error {return errors.New("ERROR_CODE:07 Error: Key not found. ")}
func BAD_JSON_ERROR() 			error {return errors.New("ERROR_CODE:08 Error: Bad JSON format. ")}
func BAD_KEY_ERROR() 			error {return errors.New("ERROR_CODE:09 Error: Key value can't contain quote symbol. ")}
func KEY_ALREADY_EXIST_ERROR()	error {return errors.New("ERROR_CODE:10 Error: Key already exist. Use Set() function to change. ")}
func CAST_INT_ERROR() 			error {return errors.New("ERROR_CODE:11 Cast Error: Cast to int error. ")}
func CAST_FLOAT_ERROR() 		error {return errors.New("ERROR_CODE:12 Cast Error: Cast to float error. ")}
func CAST_BOOL_ERROR() 			error {return errors.New("ERROR_CODE:13 Cast Error: Cast to bool error. ")}
func CAST_STRING_ARRAY_ERROR() 	error {return errors.New("ERROR_CODE:14 Cast Error: Cast to []string error. ")}
func CAST_INT_ARRAY_ERROR() 	error {return errors.New("ERROR_CODE:15 Cast Error: Cast to []int error. ")}
func CAST_FLOAT_ARRAY_ERROR() 	error {return errors.New("ERROR_CODE:16 Cast Error: Cast to []float error. ")}
func CAST_BOOL_ARRAY_ERROR() 	error {return errors.New("ERROR_CODE:17 Cast Error: Cast to []bool error. ")}
func EMPTY_ARRAY_ERROR() 		error {return errors.New("ERROR_CODE:18 Error: Array is empty. ")}
