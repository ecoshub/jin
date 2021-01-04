package jin

// AddKeyValue jin_object method for same function name with interpreter
func (j JO) AddKeyValue(key string, value []byte, path ...string) error {
	var err error
	j, err = AddKeyValue(j, key, value, path...)
	return err
}

// Add jin_object method for same function name with interpreter
func (j JO) Add(value []byte, path ...string) error {
	var err error
	j, err = Add(j, value, path...)
	return err
}

// Insert jin_object method for same function name with interpreter
func (j JO) Insert(index int, value []byte, path ...string) error {
	var err error
	j, err = Insert(j, index, value, path...)
	return err
}

// AddKeyValueString jin_object method for same function name with interpreter
func (j JO) AddKeyValueString(key string, value string, path ...string) error {
	var err error
	j, err = AddKeyValueString(j, key, value, path...)
	return err
}

// AddKeyValueInt jin_object method for same function name with interpreter
func (j JO) AddKeyValueInt(key string, value int, path ...string) error {
	var err error
	j, err = AddKeyValueInt(j, key, value, path...)
	return err
}

// AddKeyValueFloat jin_object method for same function name with interpreter
func (j JO) AddKeyValueFloat(key string, value float64, path ...string) error {
	var err error
	j, err = AddKeyValueFloat(j, key, value, path...)
	return err
}

// AddKeyValueBool jin_object method for same function name with interpreter
func (j JO) AddKeyValueBool(key string, value bool, path ...string) error {
	var err error
	j, err = AddKeyValueBool(j, key, value, path...)
	return err
}

// AddString jin_object method for same function name with interpreter
func (j JO) AddString(value string, path ...string) error {
	var err error
	j, err = AddString(j, value, path...)
	return err
}

// AddInt jin_object method for same function name with interpreter
func (j JO) AddInt(value int, path ...string) error {
	var err error
	j, err = AddInt(j, value, path...)
	return err
}

// AddFloat jin_object method for same function name with interpreter
func (j JO) AddFloat(value float64, path ...string) error {
	var err error
	j, err = AddFloat(j, value, path...)
	return err
}

// AddBool jin_object method for same function name with interpreter
func (j JO) AddBool(value bool, path ...string) error {
	var err error
	j, err = AddBool(j, value, path...)
	return err
}

// InsertString jin_object method for same function name with interpreter
func (j JO) InsertString(index int, value string, path ...string) error {
	var err error
	j, err = InsertString(j, index, value, path...)
	return err
}

// InsertInt jin_object method for same function name with interpreter
func (j JO) InsertInt(index int, value int, path ...string) error {
	var err error
	j, err = InsertInt(j, index, value, path...)
	return err
}

// InsertFloat jin_object method for same function name with interpreter
func (j JO) InsertFloat(index int, value float64, path ...string) error {
	var err error
	j, err = InsertFloat(j, index, value, path...)
	return err
}

// InsertBool jin_object method for same function name with interpreter
func (j JO) InsertBool(index int, value bool, path ...string) error {
	var err error
	j, err = InsertBool(j, index, value, path...)
	return err
}
