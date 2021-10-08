package jin

import "strconv"

// AddKeyValue jin_object method for same function name with interpreter
func (j *JO) AddKeyValue(key string, value []byte, path ...string) error {
	var err error
	j.body, err = AddKeyValue(j.body, key, value, path...)
	return err
}

// Add jin_object method for same function name with interpreter
func (j *JO) Add(value []byte, path ...string) error {
	var err error
	j.body, err = Add(j.body, value, path...)
	return err
}

// Insert jin_object method for same function name with interpreter
func (j *JO) Insert(index int, value []byte, path ...string) error {
	var err error
	j.body, err = Insert(j.body, index, value, path...)
	return err
}

// AddKeyValueString jin_object method for same function name with interpreter
func (j *JO) AddKeyValueString(key string, value string, path ...string) error {
	var err error
	j.body, err = AddKeyValueString(j.body, key, value, path...)
	return err
}

// AddKeyValueInt jin_object method for same function name with interpreter
func (j *JO) AddKeyValueInt(key string, value int, path ...string) error {
	var err error
	j.body, err = AddKeyValueInt(j.body, key, value, path...)
	return err
}

// AddKeyValueFloat jin_object method for same function name with interpreter
func (j *JO) AddKeyValueFloat(key string, value float64, path ...string) error {
	var err error
	j.body, err = AddKeyValueFloat(j.body, key, value, path...)
	return err
}

// AddKeyValueBool jin_object method for same function name with interpreter
func (j *JO) AddKeyValueBool(key string, value bool, path ...string) error {
	var err error
	j.body, err = AddKeyValueBool(j.body, key, value, path...)
	return err
}

// AddString jin_object method for same function name with interpreter
func (j *JO) AddString(value string, path ...string) error {
	var err error
	j.body, err = AddString(j.body, value, path...)
	return err
}

// AddInt jin_object method for same function name with interpreter
func (j *JO) AddInt(value int, path ...string) error {
	var err error
	j.body, err = AddInt(j.body, value, path...)
	return err
}

// AddFloat jin_object method for same function name with interpreter
func (j *JO) AddFloat(value float64, path ...string) error {
	var err error
	j.body, err = AddFloat(j.body, value, path...)
	return err
}

// AddBool jin_object method for same function name with interpreter
func (j *JO) AddBool(value bool, path ...string) error {
	var err error
	j.body, err = AddBool(j.body, value, path...)
	return err
}

// InsertString jin_object method for same function name with interpreter
func (j *JO) InsertString(index int, value string, path ...string) error {
	var err error
	j.body, err = InsertString(j.body, index, value, path...)
	return err
}

// InsertInt jin_object method for same function name with interpreter
func (j *JO) InsertInt(index int, value int, path ...string) error {
	var err error
	j.body, err = InsertInt(j.body, index, value, path...)
	return err
}

// InsertFloat jin_object method for same function name with interpreter
func (j *JO) InsertFloat(index int, value float64, path ...string) error {
	var err error
	j.body, err = InsertFloat(j.body, index, value, path...)
	return err
}

// InsertBool jin_object method for same function name with interpreter
func (j *JO) InsertBool(index int, value bool, path ...string) error {
	var err error
	j.body, err = InsertBool(j.body, index, value, path...)
	return err
}

func (j *JO) Store(key string, value []byte, path ...string) error {
	_, start, end, err := core(j.body, false, append(path, key)...)
	if err != nil {
		if ErrEqual(err, ErrCodeKeyNotFound) {
			j.body, err = AddKeyValue(j.body, key, value, path...)
			return err
		}
		return err
	}
	if j.body[start-1] == 34 && j.body[end] == 34 {
		j.body = replace(j.body, value, start-1, end+1)
		return nil
	}
	j.body = replace(j.body, value, start, end)
	return nil
}

// StoreString jin_object method for same function name with interpreter
func (j *JO) StoreString(key string, value string, path ...string) error {
	var err error
	j.body, err = Store(j.body, key, []byte(value), path...)
	return err
}

// StoreInt jin_object method for same function name with interpreter
func (j *JO) StoreInt(key string, value int, path ...string) error {
	var err error
	j.body, err = Store(j.body, key, []byte(strconv.Itoa(value)), path...)
	return err
}

// StoreFloat jin_object method for same function name with interpreter
func (j *JO) StoreFloat(key string, value float64, path ...string) error {
	var err error
	j.body, err = Store(j.body, key, []byte(strconv.FormatFloat(value, 'f', -1, 64)), path...)
	return err
}

// StoreBool jin_object method for same function name with interpreter
func (j *JO) StoreBool(key string, value bool, path ...string) error {
	var err error
	if value {
		j.body, err = Store(j.body, key, []byte("true"), path...)
		return err
	}
	j.body, err = Store(j.body, key, []byte("false"), path...)
	return err
}
