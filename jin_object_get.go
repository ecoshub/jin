package jin

//JO jin_object method for same function name with interpreter
type JO []byte

// New jin_object method for same function name with interpreter
func New(buffer []byte) JO {
	return JO(buffer)
}

// Get jin_object method for same function name with interpreter
func (j JO) Get(path ...string) ([]byte, error) {
	return Get(j, path...)
}

// GetString jin_object method for same function name with interpreter
func (j JO) GetString(path ...string) (string, error) {
	return GetString(j, path...)
}

// GetStringArray jin_object method for same function name with interpreter
func (j JO) GetStringArray(path ...string) ([]string, error) {
	return GetStringArray(j, path...)
}

// GetType jin_object method for same function name with interpreter
func (j JO) GetType(path ...string) (string, error) {
	return GetType(j, path...)
}

// GetBool jin_object method for same function name with interpreter
func (j JO) GetBool(path ...string) (bool, error) {
	return GetBool(j, path...)
}

// GetBoolArray jin_object method for same function name with interpreter
func (j JO) GetBoolArray(path ...string) ([]bool, error) {
	return GetBoolArray(j, path...)
}

// GetFloat jin_object method for same function name with interpreter
func (j JO) GetFloat(path ...string) (float64, error) {
	return GetFloat(j, path...)
}

// GetFloatArray jin_object method for same function name with interpreter
func (j JO) GetFloatArray(path ...string) ([]float64, error) {
	return GetFloatArray(j, path...)
}

// GetInt jin_object method for same function name with interpreter
func (j JO) GetInt(path ...string) (int, error) {
	return GetInt(j, path...)
}

// GetIntArray jin_object method for same function name with interpreter
func (j JO) GetIntArray(path ...string) ([]int, error) {
	return GetIntArray(j, path...)
}

// GetAll jin_object method for same function name with interpreter
func (j JO) GetAll(keys []string, path ...string) ([]string, error) {
	return GetAll(j, keys, path...)
}

// GetAllMap jin_object method for same function name with interpreter
func (j JO) GetAllMap(keys []string, path ...string) (map[string]string, error) {
	return GetAllMap(j, keys, path...)
}

// GetKeys jin_object method for same function name with interpreter
func (j JO) GetKeys(path ...string) ([]string, error) {
	return GetKeys(j, path...)
}

// GetValues jin_object method for same function name with interpreter
func (j JO) GetValues(path ...string) ([]string, error) {
	return GetValues(j, path...)
}

// GetKeysValues jin_object method for same function name with interpreter
func (j JO) GetKeysValues(path ...string) ([]string, []string, error) {
	return GetKeysValues(j, path...)
}

// GetMap jin_object method for same function name with interpreter
func (j JO) GetMap(path ...string) (map[string]string, error) {
	return GetMap(j, path...)
}
