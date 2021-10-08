package jin

// Set jin_object method for same function name with interpreter
func (j *JO) Set(newValue []byte, path ...string) error {
	var err error
	j.body, err = Set(j.body, newValue, path...)
	return err
}

// SetString jin_object method for same function name with interpreter
func (j *JO) SetString(newValue string, path ...string) error {
	var err error
	j.body, err = SetString(j.body, newValue, path...)
	return err
}

// SetInt jin_object method for same function name with interpreter
func (j *JO) SetInt(newValue int, path ...string) error {
	var err error
	j.body, err = SetInt(j.body, newValue, path...)
	return err
}

// SetFloat jin_object method for same function name with interpreter
func (j *JO) SetFloat(newValue float64, path ...string) error {
	var err error
	j.body, err = SetFloat(j.body, newValue, path...)
	return err
}

// SetBool jin_object method for same function name with interpreter
func (j *JO) SetBool(newValue bool, path ...string) error {
	var err error
	j.body, err = SetBool(j.body, newValue, path...)
	return err
}

// SetKey jin_object method for same function name with interpreter
func (j *JO) SetKey(newKey string, path ...string) error {
	var err error
	j.body, err = SetKey(j.body, newKey, path...)
	return err
}
