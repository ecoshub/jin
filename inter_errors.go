package jsoninterpreter

import "errors"

func getError(err string) error {
	errs := make(map[string]error)
	errs["NULL_PATH_ERROR"] = errors.New("Error: Path can't be null.")
	errs["NULL_NEW_VALUE_ERROR"] = errors.New("Error: New value can't be null.")
	errs["INDEX_EXPECTED_ERROR"] = errors.New("Error: Index expected, got key value.")
	errs["KEY_EXPECTED_ERROR"] = errors.New("Error: Key value expected, got index.")
	errs["INDEX_OUT_OF_RANGE_ERROR"] = errors.New("Error: Index out of range.")
	errs["KEY_NOT_FOUND_ERROR"] = errors.New("Error: Key not found.")
	errs["BAD_JSON_ERROR"] = errors.New("Error: Bad JSON format.")
	errs["BAD_KEY_ERROR"] = errors.New("Error: Key value can't contain quote symbol.")
	errs["CAST_INT_ERROR"] = errors.New("Cast Error: Cast to int error.")
	errs["CAST_FLOAT_ERROR"] = errors.New("Cast Error: Cast to float error.")
	errs["CAST_BOOL_ERROR"] = errors.New("Cast Error: Cast to bool error.")
	errs["CAST_STRING_ARRAY_ERROR"] = errors.New("Cast Error: Cast to []string error.")
	errs["CAST_INT_ARRAY_ERROR"] = errors.New("Cast Error: Cast to []int error.")
	errs["CAST_FLOAT_ARRAY_ERROR"] = errors.New("Cast Error: Cast to []float error.")
	errs["CAST_BOOL_ARRAY_ERROR"] = errors.New("Cast Error: Cast to []bool error.")
	return errs[err]
}