package jin

// Set jin_object method for same function name with interpreter
func (j JO) Set(newValue []byte, path ...string) error {
	var err error
	j, err = Set(j, newValue, path...)
	return err
}

// SetString jin_object method for same function name with interpreter
func (j JO) SetString(newValue string, path ...string) error {
	var err error
	j, err = SetString(j, newValue, path...)
	return err
}

// SetInt jin_object method for same function name with interpreter
func (j JO) SetInt(newValue int, path ...string) error {
	var err error
	j, err = SetInt(j, newValue, path...)
	return err
}

// SetFloat jin_object method for same function name with interpreter
func (j JO) SetFloat(newValue float64, path ...string) error {
	var err error
	j, err = SetFloat(j, newValue, path...)
	return err
}

// SetBool jin_object method for same function name with interpreter
func (j JO) SetBool(newValue bool, path ...string) error {
	var err error
	j, err = SetBool(j, newValue, path...)
	return err
}

// SetKey jin_object method for same function name with interpreter
func (j JO) SetKey(newKey string, path ...string) error {
	var err error
	j, err = SetKey(j, newKey, path...)
	return err
}
