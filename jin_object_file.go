package jin

import (
	"io/ioutil"
)

// ReadJSONFile read json file and return a JSON OBject (JO)
func ReadJSONFile(path string) (JO, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return JO(content), nil
}
