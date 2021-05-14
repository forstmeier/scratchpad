package entities

import "fmt"

// ErrorReadConfigFile wraps errors returned by ioutil.ReadFile in the New function.
type ErrorReadConfigFile struct {
	err error
}

func newErrorReadConfigFile(err error) ErrorReadConfigFile {
	return ErrorReadConfigFile{
		err: fmt.Errorf("entities: read config file: %w", err),
	}
}

func (e ErrorReadConfigFile) Error() string {
	return e.err.Error()
}

// ErrorUnmarshalConfigJSON wraps errors returned by json.Unmarshal in the New function.
type ErrorUnmarshalConfigJSON struct {
	err error
}

func newErrorUnmarshalConfigJSON(err error) ErrorUnmarshalConfigJSON {
	return ErrorUnmarshalConfigJSON{
		err: fmt.Errorf("entities: unmarshal config json: %w", err),
	}
}

func (e ErrorUnmarshalConfigJSON) Error() string {
	return e.err.Error()
}

// ErrorReadModelsDirectory wraps errors returned by ioutil.ReadDir in the New function.
type ErrorReadModelsDirectory struct {
	err error
}

func newErrorReadModelsDirectory(err error) ErrorReadModelsDirectory {
	return ErrorReadModelsDirectory{
		err: fmt.Errorf("entities: read models directory: %w", err),
	}
}

func (e ErrorReadModelsDirectory) Error() string {
	return e.err.Error()
}

// ErrorCreateDocument wraps errors returned by prose.NewDocument in the ClassifyEntities method.
type ErrorCreateDocument struct {
	err error
}

func newErrorCreateDocument(err error) ErrorCreateDocument {
	return ErrorCreateDocument{
		err: fmt.Errorf("entities: create document: %w", err),
	}
}

func (e ErrorCreateDocument) Error() string {
	return e.err.Error()
}
