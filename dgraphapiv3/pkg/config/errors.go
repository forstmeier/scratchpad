package config

// ErrorReadFile wraps errors returned by ioutil.ReadFile in the New function.
type ErrorReadFile struct {
	err error
}

func (e ErrorReadFile) Error() string {
	return "config: read file: " + e.err.Error()
}

// ErrorUnmarshalConfig wraps errors returned by json.Unmarshal in the New function.
type ErrorUnmarshalConfig struct {
	err error
}

func (e ErrorUnmarshalConfig) Error() string {
	return "config: unmarshal: " + e.err.Error()
}
