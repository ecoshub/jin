package jsoninterpreter

import "errors"

func NULL_PATH_ERROR() error {return errors.New("Error: Path can't be null.")}
func NULL_NEW_VALUE_ERROR() error {return errors.New("Error: New value can't be null.")}
func INDEX_EXPECTED_ERROR() error {return errors.New("Error: Index expected, got key value.")}
func KEY_EXPECTED_ERROR() error {return errors.New("Error: Key value expected, got index.")}
func INDEX_OUT_OF_RANGE_ERROR() error {return errors.New("Error: Index out of range.")}
func KEY_NOT_FOUND_ERROR() error {return errors.New("Error: Key not found.")}
func BAD_JSON_ERROR() error {return errors.New("Error: Bad JSON format.")}
func BAD_KEY_ERROR() error {return errors.New("Error: Key value can't contain quote symbol.")}
func CAST_INT_ERROR() error {return errors.New("Cast Error: Cast to int error.")}
func CAST_FLOAT_ERROR() error {return errors.New("Cast Error: Cast to float error.")}
func CAST_BOOL_ERROR() error {return errors.New("Cast Error: Cast to bool error.")}
func CAST_STRING_ARRAY_ERROR() error {return errors.New("Cast Error: Cast to []string error.")}
func CAST_INT_ARRAY_ERROR() error {return errors.New("Cast Error: Cast to []int error.")}
func CAST_FLOAT_ARRAY_ERROR() error {return errors.New("Cast Error: Cast to []float error.")}
func CAST_BOOL_ARRAY_ERROR() error {return errors.New("Cast Error: Cast to []bool error.")}