package jin

// Delete jin_object method for same function name with interpreter
func (j *JO) Delete(path ...string) error {
	var err error
	j.body, err = Delete(j.body, path...)
	return err
}

// IterateArray jin_object method for same function name with interpreter
func (j *JO) IterateArray(callback func([]byte) (bool, error), path ...string) error {
	return IterateArray(j.body, callback, path...)
}

// IterateKeyValue jin_object method for same function name with interpreter
func (j *JO) IterateKeyValue(callback func([]byte, []byte) (bool, error), path ...string) error {
	return IterateKeyValue(j.body, callback, path...)
}

// IsObject jin_object method for same function name with interpreter
func (j *JO) IsObject(path ...string) (bool, error) {
	return IsObject(j.body, path...)
}

// IsArray jin_object method for same function name with interpreter
func (j *JO) IsArray(path ...string) (bool, error) {
	return IsArray(j.body, path...)
}

// IsValue jin_object method for same function name with interpreter
func (j *JO) IsValue(path ...string) (bool, error) {
	return IsValue(j.body, path...)
}

// IsEmpty jin_object method for same function name with interpreter
func (j *JO) IsEmpty(path ...string) (bool, error) {
	return IsEmpty(j.body, path...)
}
