package jin

import (
	"errors"
	"fmt"
)

func nullPathError() error {
	return errors.New("error: path cannot be null error_code:00 ")
}
func nullNewValueError() error {
	return errors.New("error: new value cannot be null error_code:01 ")
}
func emptyArrayError() error {
	return errors.New("error: array is empty error_code:02 ")
}
func indexExpectedError() error {
	return errors.New("error: index expected, got key value error_code:03 ")
}
func keyExpectedError() error {
	return errors.New("error: key expected, got index error_code:04 ")
}
func objectExpectedError() error {
	return errors.New("error: last path must be pointed at an object error_code:05 ")
}
func arrayExpectedError() error {
	return errors.New("error: last path must be pointed at an array error_code:06 ")
}
func indexOutOfRangeError() error {
	return errors.New("error: index out of range error_code:07 ")
}
func keyNotFoundError(key string) error {
	return fmt.Errorf("error: key '%v' not found error_code:08 ", key)
}
func badJSONError(val int) error {
	return fmt.Errorf("error: bad json format. at:'%v' error_code:09 ", val)
}
func badKeyError(key string) error {
	return fmt.Errorf("error: key ('%v') value cannot contain quote symbol. error_code:10", key)
}
func keyAlreadyExistsError(key string) error {
	return fmt.Errorf("error: key '%v' already exist error_code:11", key)
}
func intParseError(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to int. error_code:12", val)
}
func floatParseError(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to float. error_code:13", val)
}
func boolParseError(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to bool. error_code:14", val)
}
func stringArrayParseError(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to []string. error_code:15", val)
}
func intArrayParseError(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to []int. error_code:16", val)
}
func floatArrayParseError(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to []float. error_code:17", val)
}
func boolArrayParseError(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to []bool. error_code:18", val)
}
func nullKeyError() error {
	return errors.New("error: new key cannot be null error_code:19 ")
}
func generalEmptyError() error {
	return errors.New("error: Object/Array is empty error_code:20 ")
}
