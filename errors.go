package jin

import "errors"
import "fmt"

func nullPathError() error {
	return errors.New("ERROR_CODE:00 Error: Path cannot be null")
}
func nullNewValueError() error {
	return errors.New("ERROR_CODE:01 Error: New value cannot be null")
}
func emptyArrayError() error {
	return errors.New("ERROR_CODE:02 Error: Array is empty")
}
func indexExpectedError() error {
	return errors.New("ERROR_CODE:03 Error: Index expected, got key value")
}
func keyExpectedError() error {
	return errors.New("ERROR_CODE:04 Error: Key expected, got index")
}
func objectExpectedError() error {
	return errors.New("ERROR_CODE:05 Error: Last path must be pointed at an object")
}
func arrayExpectedError() error {
	return errors.New("ERROR_CODE:06 Error: Last path must be pointed at an array")
}
func indexOutOfRangeError() error {
	return errors.New("ERROR_CODE:07 Error: Index out of range")
}
func keyNotFoundError() error {
	return errors.New("ERROR_CODE:08 Error: Key not found")
}
func badJSONError(index int) error {
	return fmt.Errorf("ERROR_CODE:09 Error: Bad JSON format. at:%v", index)
}
func badKeyError() error {
	return errors.New("ERROR_CODE:10 Error: Key value cannot contain quote symbol")
}
func keyAlreadyExistsError() error {
	return errors.New("ERROR_CODE:11 Error: Key already exist")
}
func intParseError(val string) error {
	return fmt.Errorf("ERROR_CODE:12 Parse Error: '%v' Cannot be converted to int", val)
}
func floatParseError(val string) error {
	return fmt.Errorf("ERROR_CODE:13 Parse Error: '%v' Cannot be converted to float", val)
}
func boolParseError(val string) error {
	return fmt.Errorf("ERROR_CODE:14 Parse Error: '%v' Cannot be converted to bool", val)
}
func stringArrayParseError(val string) error {
	return fmt.Errorf("ERROR_CODE:15 Parse Error: '%v' Cannot be converted to []string", val)
}
func intArrayParseError(val string) error {
	return fmt.Errorf("ERROR_CODE:16 Parse Error: '%v' Cannot be converted to []int", val)
}
func floatArrayParseError(val string) error {
	return fmt.Errorf("ERROR_CODE:17 Parse Error: '%v' Cannot be converted to []float", val)
}
func boolArrayParseError(val string) error {
	return fmt.Errorf("ERROR_CODE:18 Parse Error: '%v' Cannot be converted to []bool", val)
}
func nullKeyError() error {
	return errors.New("ERROR_CODE:01 Error: New key cannot be null")
}
