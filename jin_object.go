package jin

import "io/ioutil"

type JO struct {
	body []byte
}

func NewJSON() *JO {
	return &JO{body: []byte("{}")}
}

// New jin_object method for same function name with interpreter
func New(buffer []byte) *JO {
	return &JO{body: buffer}
}

func (j *JO) JSON() []byte {
	return j.body
}

func (j *JO) String() string {
	return string(j.body)
}

// ReadJSONFile read json file and return a JSON OBject (JO)
func ReadJSONFile(path string) (*JO, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &JO{body: content}, nil
}

// ReadFile read json file and return a JSON OBject (JO)
func ReadFile(path string) (*JO, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &JO{body: content}, nil
}
