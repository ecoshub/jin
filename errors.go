package jin

import (
	"errors"
	"fmt"
)

// ErrNullPath "error: path cannot be null error_code:00 "
func ErrNullPath() error {
	return errors.New("error: path cannot be null error_code:00 ")
}

// ErrNullNewValue "error: new value cannot be null error_code:01 "
func ErrNullNewValue() error {
	return errors.New("error: new value cannot be null error_code:01 ")
}

// ErrEmptyArray "error: array is empty error_code:02 "
func ErrEmptyArray() error {
	return errors.New("error: array is empty error_code:02 ")
}

// ErrIndexExpected "error: index expected, got key value error_code:03 "
func ErrIndexExpected() error {
	return errors.New("error: index expected, got key value error_code:03 ")
}

// ErrKeyExpected "error: key expected, got index error_code:04 "
func ErrKeyExpected() error {
	return errors.New("error: key expected, got index error_code:04 ")
}

// ErrObjectExpected "error: last path must be pointed at an object error_code:05 "
func ErrObjectExpected() error {
	return errors.New("error: last path must be pointed at an object error_code:05 ")
}

// ErrArrayExpected "error: last path must be pointed at an array error_code:06 "
func ErrArrayExpected() error {
	return errors.New("error: last path must be pointed at an array error_code:06 ")
}

// ErrIndexOutOfRange "error: index out of range error_code:07 "
func ErrIndexOutOfRange() error {
	return errors.New("error: index out of range error_code:07 ")
}

// ErrKeyNotFound "error: key '%v' not found error_code:08 "
func ErrKeyNotFound(key string) error {
	return fmt.Errorf("error: key '%v' not found error_code:08 ", key)
}

// ErrBadJSON "error: bad json format. at:'%v' error_code:09 "
func ErrBadJSON(val int) error {
	return fmt.Errorf("error: bad json format. at:'%v' error_code:09 ", val)
}

// ErrBadKey "error: key ('%v') value cannot contain quote symbol. error_code:10"
func ErrBadKey(key string) error {
	return fmt.Errorf("error: key ('%v') value cannot contain quote symbol. error_code:10", key)
}

// ErrKeyAlreadyExist "error: key '%v' already exist error_code:11"
func ErrKeyAlreadyExist(key string) error {
	return fmt.Errorf("error: key '%v' already exist error_code:11", key)
}

// ErrIntegerParse "parse error: '%v' cannot be converted to int. error_code:12"
func ErrIntegerParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to int. error_code:12", val)
}

// ErrFloatParse "parse error: '%v' cannot be converted to float. error_code:13"
func ErrFloatParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to float. error_code:13", val)
}

// ErrBoolParse "parse error: '%v' cannot be converted to bool. error_code:14"
func ErrBoolParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to bool. error_code:14", val)
}

// ErrStringArrayParse "parse error: '%v' cannot be converted to []string. error_code:15"
func ErrStringArrayParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to []string. error_code:15", val)
}

// ErrIntegerArrayPars "parse error: '%v' cannot be converted to []int. error_code:16"
func ErrIntegerArrayParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to []int. error_code:16", val)
}

// ErrFloatArrayParse "parse error: '%v' cannot be converted to []float. error_code:17"
func ErrFloatArrayParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to []float. error_code:17", val)
}

// ErrBoolArrayParse "parse error: '%v' cannot be converted to []bool. error_code:18"
func ErrBoolArrayParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to []bool. error_code:18", val)
}

// ErrNullKey "error: new key cannot be null error_code:19"
func ErrNullKey() error {
	return errors.New("error: new key cannot be null error_code:19 ")
}

// ErrEmpty "error: Object/Array is empty error_code:20"
func ErrEmpty() error {
	return errors.New("error: Object/Array is empty error_code:20 ")
}
