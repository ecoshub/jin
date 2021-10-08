package jin

// Get jin_object method for same function name with interpreter
func (j *JO) Get(path ...string) ([]byte, error) {
	return Get(j.body, path...)
}

// GetString jin_object method for same function name with interpreter
func (j *JO) GetString(path ...string) (string, error) {
	return GetString(j.body, path...)
}

// GetStringArray jin_object method for same function name with interpreter
func (j *JO) GetStringArray(path ...string) ([]string, error) {
	return GetStringArray(j.body, path...)
}

// GetType jin_object method for same function name with interpreter
func (j *JO) GetType(path ...string) (string, error) {
	return GetType(j.body, path...)
}

// GetBool jin_object method for same function name with interpreter
func (j *JO) GetBool(path ...string) (bool, error) {
	return GetBool(j.body, path...)
}

// GetBoolArray jin_object method for same function name with interpreter
func (j *JO) GetBoolArray(path ...string) ([]bool, error) {
	return GetBoolArray(j.body, path...)
}

// GetFloat jin_object method for same function name with interpreter
func (j *JO) GetFloat(path ...string) (float64, error) {
	return GetFloat(j.body, path...)
}

// GetFloatArray jin_object method for same function name with interpreter
func (j *JO) GetFloatArray(path ...string) ([]float64, error) {
	return GetFloatArray(j.body, path...)
}

// GetInt jin_object method for same function name with interpreter
func (j *JO) GetInt(path ...string) (int, error) {
	return GetInt(j.body, path...)
}

// GetIntArray jin_object method for same function name with interpreter
func (j *JO) GetIntArray(path ...string) ([]int, error) {
	return GetIntArray(j.body, path...)
}

// GetAll jin_object method for same function name with interpreter
func (j *JO) GetAll(keys []string, path ...string) ([]string, error) {
	return GetAll(j.body, keys, path...)
}

// GetAllMap jin_object method for same function name with interpreter
func (j *JO) GetAllMap(keys []string, path ...string) (map[string]string, error) {
	return GetAllMap(j.body, keys, path...)
}

// GetKeys jin_object method for same function name with interpreter
func (j *JO) GetKeys(path ...string) ([]string, error) {
	return GetKeys(j.body, path...)
}

// GetValues jin_object method for same function name with interpreter
func (j *JO) GetValues(path ...string) ([]string, error) {
	return GetValues(j.body, path...)
}

// GetKeysValues jin_object method for same function name with interpreter
func (j *JO) GetKeysValues(path ...string) ([]string, []string, error) {
	return GetKeysValues(j.body, path...)
}

// GetMap jin_object method for same function name with interpreter
func (j *JO) GetMap(path ...string) (map[string]string, error) {
	return GetMap(j.body, path...)
}
