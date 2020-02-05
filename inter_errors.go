package jin

import "errors"
import "fmt"

func error_null_path() error {
	return errors.New("ERROR_CODE:00 Error: Path cannot be null.")
}
func error_null_new_value() error {
	return errors.New("ERROR_CODE:01 Error: New value cannot be null.")
}
func error_empty_array() error {
	return errors.New("ERROR_CODE:02 Error: Array is empty.")
}
func error_index_expected() error {
	return errors.New("ERROR_CODE:03 Error: Index expected, got key value.")
}
func error_key_expected() error {
	return errors.New("ERROR_CODE:04 Error: Key expected, got index.")
}
func error_object_expected() error {
	return errors.New("ERROR_CODE:05 Error: Last path must be pointed at an object.")
}
func error_array_expected() error {
	return errors.New("ERROR_CODE:06 Error: Last path must be pointed at an array.")
}
func error_index_out_of_range() error {
	return errors.New("ERROR_CODE:07 Error: Index out of range.")
}
func error_key_not_found() error {
	return errors.New("ERROR_CODE:08 Error: Key not found.")
}
func error_bad_json(index int) error {
	return fmt.Errorf("ERROR_CODE:09 Error: Bad JSON format. at:%v", index)
}
func error_bad_key() error {
	return errors.New("ERROR_CODE:10 Error: Key value cannot contain quote symbol.")
}
func error_key_already_exists() error {
	return errors.New("ERROR_CODE:11 Error: Key already exist.")
}
func error_parse_int(val string) error {
	return fmt.Errorf("ERROR_CODE:12 Parse Error: '%v' Cannot be converted to int.", val)
}
func error_parse_float(val string) error {
	return fmt.Errorf("ERROR_CODE:13 Parse Error: '%v' Cannot be converted to float.", val)
}
func error_parse_bool(val string) error {
	return fmt.Errorf("ERROR_CODE:14 Parse Error: '%v' Cannot be converted to bool.", val)
}
func error_parse_string_array(val string) error {
	return fmt.Errorf("ERROR_CODE:15 Parse Error: '%v' Cannot be converted to []string.", val)
}
func error_parse_int_array(val string) error {
	return fmt.Errorf("ERROR_CODE:16 Parse Error: '%v' Cannot be converted to []int.", val)
}
func error_parse_float_array(val string) error {
	return fmt.Errorf("ERROR_CODE:17 Parse Error: '%v' Cannot be converted to []float.", val)
}
func error_parse_bool_array(val string) error {
	return fmt.Errorf("ERROR_CODE:18 Parse Error: '%v' Cannot be converted to []bool.", val)
}
func error_null_key() error {
	return errors.New("ERROR_CODE:01 Error: New key cannot be null.")
}
