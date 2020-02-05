package jin

import "errors"
import "fmt"

func ERROR_NULL_PATH() error {
	return errors.New("ERROR_CODE:00 Error: Path cannot be null.")
}
func ERROR_NULL_NEW_VALUE() error {
	return errors.New("ERROR_CODE:01 Error: New value cannot be null.")
}
func ERROR_EMPTY_ARRAY() error {
	return errors.New("ERROR_CODE:02 Error: Array is empty.")
}
func ERROR_INDEX_EXPECTED() error {
	return errors.New("ERROR_CODE:03 Error: Index expected, got key value.")
}
func ERROR_KEY_EXPECTED() error {
	return errors.New("ERROR_CODE:04 Error: Key expected, got index.")
}
func ERROR_OBJECT_EXPECTED() error {
	return errors.New("ERROR_CODE:05 Error: Last path must be pointed at an object.")
}
func ERROR_ARRAY_EXPECTED() error {
	return errors.New("ERROR_CODE:06 Error: Last path must be pointed at an array.")
}
func ERROR_INDEX_OUT_OF_RANGE() error {
	return errors.New("ERROR_CODE:07 Error: Index out of range.")
}
func ERROR_KEY_NOT_FOUND() error {
	return errors.New("ERROR_CODE:08 Error: Key not found.")
}
func ERROR_BAD_JSON(index int) error {
	return fmt.Errorf("ERROR_CODE:09 Error: Bad JSON format. at:%v", index)
}
func ERROR_BAD_KEY() error {
	return errors.New("ERROR_CODE:10 Error: Key value cannot contain quote symbol.")
}
func ERROR_KEY_ALREADY_EXISTS() error {
	return errors.New("ERROR_CODE:11 Error: Key already exist.")
}
func ERROR_PARSE_INT(val string) error {
	return fmt.Errorf("ERROR_CODE:12 Parse Error: '%v' Cannot be converted to int.", val)
}
func ERROR_PARSE_FLOAT(val string) error {
	return fmt.Errorf("ERROR_CODE:13 Parse Error: '%v' Cannot be converted to float.", val)
}
func ERROR_PARSE_BOOL(val string) error {
	return fmt.Errorf("ERROR_CODE:14 Parse Error: '%v' Cannot be converted to bool.", val)
}
func ERROR_PARSE_STRING_ARRAY(val string) error {
	return fmt.Errorf("ERROR_CODE:15 Parse Error: '%v' Cannot be converted to []string.", val)
}
func ERROR_PARSE_INT_ARRAY(val string) error {
	return fmt.Errorf("ERROR_CODE:16 Parse Error: '%v' Cannot be converted to []int.", val)
}
func ERROR_PARSE_FLOAT_ARRAY(val string) error {
	return fmt.Errorf("ERROR_CODE:17 Parse Error: '%v' Cannot be converted to []float.", val)
}
func ERROR_PARSE_BOOL_ARRAY(val string) error {
	return fmt.Errorf("ERROR_CODE:18 Parse Error: '%v' Cannot be converted to []bool.", val)
}
func ERROR_NULL_KEY() error {
	return errors.New("ERROR_CODE:01 Error: New key cannot be null.")
}
