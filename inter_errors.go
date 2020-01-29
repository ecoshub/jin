package jint

import "errors"
import "fmt"

func NULL_PATH_ERROR() error {
	return errors.New("ERROR_CODE:00 Error: Path cannot be null.")
}
func NULL_NEW_VALUE_ERROR() error {
	return errors.New("ERROR_CODE:01 Error: New value cannot be null.")
}
func EMPTY_ARRAY_ERROR() error {
	return errors.New("ERROR_CODE:02 Error: Array is empty.")
}
func INDEX_EXPECTED_ERROR() error {
	return errors.New("ERROR_CODE:03 Error: Index expected, got key value.")
}
func KEY_EXPECTED_ERROR() error {
	return errors.New("ERROR_CODE:04 Error: Key expected, got index.")
}
func OBJECT_EXPECTED_ERROR() error {
	return errors.New("ERROR_CODE:05 Error: Last path must be pointed at an object.")
}
func ARRAY_EXPECTED_ERROR() error {
	return errors.New("ERROR_CODE:06 Error: Last path must be pointed at an array.")
}
func INDEX_OUT_OF_RANGE_ERROR() error {
	return errors.New("ERROR_CODE:07 Error: Index out of range.")
}
func KEY_NOT_FOUND_ERROR() error {
	return errors.New("ERROR_CODE:08 Error: Key not found.")
}
func BAD_JSON_ERROR(index int) error {
	return errors.New(fmt.Sprintf("ERROR_CODE:09 Error: Bad JSON format. at:%v", index))
}
func BAD_KEY_ERROR() error {
	return errors.New("ERROR_CODE:10 Error: Key value cannot contain quote symbol.")
}
func KEY_ALREADY_EXIST_ERROR() error {
	return errors.New("ERROR_CODE:11 Error: Key already exist.")
}
func PARSE_INT_ERROR(val string) error {
	return errors.New(fmt.Sprintf("ERROR_CODE:12 Parse Error: '%v' Cannot be converted to int.", val))
}
func PARSE_FLOAT_ERROR(val string) error {
	return errors.New(fmt.Sprintf("ERROR_CODE:13 Parse Error: '%v' Cannot be converted to float.", val))
}
func PARSE_BOOL_ERROR(val string) error {
	return errors.New(fmt.Sprintf("ERROR_CODE:14 Parse Error: '%v' Cannot be converted to bool.", val))
}
func PARSE_STRING_ARRAY_ERROR(val string) error {
	return errors.New(fmt.Sprintf("ERROR_CODE:15 Parse Error: '%v' Cannot be converted to []string.", val))
}
func PARSE_INT_ARRAY_ERROR(val string) error {
	return errors.New(fmt.Sprintf("ERROR_CODE:16 Parse Error: '%v' Cannot be converted to []int.", val))
}
func PARSE_FLOAT_ARRAY_ERROR(val string) error {
	return errors.New(fmt.Sprintf("ERROR_CODE:17 Parse Error: '%v' Cannot be converted to []float.", val))
}
func PARSE_BOOL_ARRAY_ERROR(val string) error {
	return errors.New(fmt.Sprintf("ERROR_CODE:18 Parse Error: '%v' Cannot be converted to []bool.", val))
}
func END_OF_ITERATION() error {
	return errors.New("ERROR_CODE:19 Iteration Ended. If you want to restart the iteration use Reset() function.")
}
