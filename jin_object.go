package jin

//InterJSONObject basically a byte array of JSON
type InterJSONObject []byte

// New InterJSONObject
func New(buffer []byte) InterJSONObject {
	return InterJSONObject(buffer)
}

// Get main get func
func (i InterJSONObject) Get(path ...string) ([]byte, error) {
	return Get(i, path...)
}

// GetString main get func
func (i InterJSONObject) GetString(path ...string) (string, error) {
	return GetString(i, path...)
}

// GetStringArray main get func
func (i InterJSONObject) GetStringArray(path ...string) ([]string, error) {
	return GetStringArray(i, path...)
}

// GetType main get func
func (i InterJSONObject) GetType(path ...string) (string, error) {
	return GetType(i, path...)
}

// GetBool main get func
func (i InterJSONObject) GetBool(path ...string) (bool, error) {
	return GetBool(i, path...)
}

// GetBoolArray main get func
func (i InterJSONObject) GetBoolArray(path ...string) ([]bool, error) {
	return GetBoolArray(i, path...)
}

// GetFloat main get func
func (i InterJSONObject) GetFloat(path ...string) (float64, error) {
	return GetFloat(i, path...)
}

// GetFloatArray main get func
func (i InterJSONObject) GetFloatArray(path ...string) ([]float64, error) {
	return GetFloatArray(i, path...)
}

// GetInt main get func
func (i InterJSONObject) GetInt(path ...string) (int, error) {
	return GetInt(i, path...)
}

// GetIntArray main get func
func (i InterJSONObject) GetIntArray(path ...string) ([]int, error) {
	return GetIntArray(i, path...)
}

// GetAll main get func
func (i InterJSONObject) GetAll(keys []string, path ...string) ([]string, error) {
	return GetAll(i, keys, path...)
}

// GetAllMap main get func
func (i InterJSONObject) GetAllMap(keys []string, path ...string) (map[string]string, error) {
	return GetAllMap(i, keys, path...)
}

// GetKeys main get func
func (i InterJSONObject) GetKeys(path ...string) ([]string, error) {
	return GetKeys(i, path...)
}

// GetValues main get func
func (i InterJSONObject) GetValues(path ...string) ([]string, error) {
	return GetValues(i, path...)
}

// GetKeysValues main get func
func (i InterJSONObject) GetKeysValues(path ...string) ([]string, []string, error) {
	return GetKeysValues(i, path...)
}

// GetMap main get func
func (i InterJSONObject) GetMap(path ...string) (map[string]string, error) {
	return GetMap(i, path...)
}
