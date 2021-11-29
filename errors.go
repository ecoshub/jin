package jin

import (
	"fmt"
	"strconv"
)

const (
	ErrCodeNullPath          int = 0
	ErrCodeEmptyArray        int = 2
	ErrCodeIndexExpected     int = 3
	ErrCodeObjectExpected    int = 5
	ErrCodeArrayExpected     int = 6
	ErrCodeIndexOutOfRange   int = 7
	ErrCodeKeyNotFound       int = 8
	ErrCodeBadJSON           int = 9
	ErrCodeKeyAlreadyExist   int = 11
	ErrCodeIntegerParse      int = 12
	ErrCodeFloatParse        int = 13
	ErrCodeBoolParse         int = 14
	ErrCodeStringArrayParse  int = 15
	ErrCodeIntegerArrayParse int = 16
	ErrCodeFloatArrayParse   int = 17
	ErrCodeBoolArrayParse    int = 18
	ErrCodeNullKey           int = 19
	ErrCodeEmpty             int = 20
)

func ErrEqual(err error, code int) bool {
	if err == nil {
		return false
	}
	str := err.Error()
	if len(str) < 2 {
		return false
	}
	codeStr := str[len(str)-2:]
	c, err := strconv.Atoi(codeStr)
	if err != nil {
		return false
	}
	return c == code
}

// errNullPath "error: path cannot be null error_code:00"
func errNullPath() error {
	return fmt.Errorf("error: path cannot be null error_code: %02d", ErrCodeNullPath)
}

// errEmptyArray "error: array is empty error_code:02"
func errEmptyArray() error {
	return fmt.Errorf("error: array is empty error_code: %02d", ErrCodeEmptyArray)
}

// errIndexExpected "error: index expected, got key value error_code:03"
func errIndexExpected() error {
	return fmt.Errorf("error: index expected, got key value error_code: %02d", ErrCodeIndexExpected)
}

// errObjectExpected "error: last path must be pointed at an object error_code:05"
func errObjectExpected() error {
	return fmt.Errorf("error: last path must be pointed at an object error_code: %02d", ErrCodeObjectExpected)
}

// errArrayExpected "error: last path must be pointed at an array error_code:06"
func errArrayExpected() error {
	return fmt.Errorf("error: last path must be pointed at an array error_code: %02d", ErrCodeArrayExpected)
}

// errIndexOutOfRange "error: index out of range error_code:07"
func errIndexOutOfRange() error {
	return fmt.Errorf("error: index out of range error_code: %02d", ErrCodeIndexOutOfRange)
}

// errKeyNotFound "error: key '%v' not found error_code:08"
func errKeyNotFound(key string) error {
	return fmt.Errorf("error: key '%v' not found error_code: %02d", key, ErrCodeKeyNotFound)
}

// errBadJSON "error: bad json format. at:'%v' error_code:09"
func errBadJSON(val int) error {
	return fmt.Errorf("error: bad json format. at:'%v' error_code: %02d", val, ErrCodeBadJSON)
}

// errKeyAlreadyExist "error: key '%v' already exist error_code:11"
func errKeyAlreadyExist(key string) error {
	return fmt.Errorf("error: key '%v' already exist error_code: %02d", key, ErrCodeKeyAlreadyExist)
}

// errIntegerParse "parse error: '%v' cannot be converted to int. error_code:12"
func errIntegerParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to int. error_code: %02d", val, ErrCodeIntegerParse)
}

// errFloatParse "parse error: '%v' cannot be converted to float. error_code:13"
func errFloatParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to float. error_code: %02d", val, ErrCodeFloatParse)
}

// errBoolParse "parse error: '%v' cannot be converted to bool. error_code:14"
func errBoolParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to bool. error_code: %02d", val, ErrCodeBoolParse)
}

// errStringArrayParse "parse error: '%v' cannot be converted to []string. error_code:15"
func errStringArrayParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to []string. error_code: %02d", val, ErrCodeStringArrayParse)
}

// errIntegerArrayParse "parse error: '%v' cannot be converted to []int. error_code:16"
func errIntegerArrayParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to []int. error_code: %02d", val, ErrCodeIntegerArrayParse)
}

// errFloatArrayParse "parse error: '%v' cannot be converted to []float. error_code:17"
func errFloatArrayParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to []float. error_code: %02d", val, ErrCodeFloatArrayParse)
}

// errBoolArrayParse "parse error: '%v' cannot be converted to []bool. error_code:18"
func errBoolArrayParse(val string) error {
	return fmt.Errorf("parse error: '%v' cannot be converted to []bool. error_code: %02d", val, ErrCodeBoolArrayParse)
}

// errNullKey "error: new key cannot be null error_code:19"
func errNullKey() error {
	return fmt.Errorf("error: new key cannot be null error_code: %02d", ErrCodeNullKey)
}

// errEmpty "error: Object/Array is empty error_code:20"
func errEmpty() error {
	return fmt.Errorf("error: Object/Array is empty error_code: %02d", ErrCodeEmpty)
}
