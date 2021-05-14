package graphql

import "fmt"

// ErrorMarshalPayload wraps errors returned by json.Marshal in the Mutate method.
type ErrorMarshalPayload struct {
	err error
}

func newErrorMarshalPayload(err error) ErrorMarshalPayload {
	return ErrorMarshalPayload{
		err: fmt.Errorf("graphql: marshal payload: %w", err),
	}
}

func (e ErrorMarshalPayload) Error() string {
	return e.err.Error()
}

// ErrorNewRequest wraps errors returned by http.NewRequest in the Mutate method.
type ErrorNewRequest struct {
	err error
}

func newErrorNewRequest(err error) ErrorNewRequest {
	return ErrorNewRequest{
		err: fmt.Errorf("graphql: new request: %w", err),
	}
}

func (e ErrorNewRequest) Error() string {
	return e.err.Error()
}

// ErrorClientDo wraps errors returned by http.Client.Do in the Mutate method.
type ErrorClientDo struct {
	err error
}

func newErrorClientDo(err error) ErrorClientDo {
	return ErrorClientDo{
		err: fmt.Errorf("graphql: client do: %w", err),
	}
}

func (e ErrorClientDo) Error() string {
	return e.err.Error()
}

// ErrorJSONDecode wraps errors returned by json.Decoder.Decode in the Mutate method.
type ErrorJSONDecode struct {
	err error
}

func newErrorJSONDecode(err error) ErrorJSONDecode {
	return ErrorJSONDecode{
		err: fmt.Errorf("graphql: json decode: %w", err),
	}
}

func (e ErrorJSONDecode) Error() string {
	return e.err.Error()
}
