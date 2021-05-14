package graphql

import "strings"

// ErrorRequestPayloadMarshal wraps errors returned by json.Marshal in the newRequest
// function.
type ErrorRequestPayloadMarshal struct {
	err error
}

func (e ErrorRequestPayloadMarshal) Error() string {
	return "graphql: request payload marshal: " + e.err.Error()
}

// ErrorHTTPNewRequest wraps errors returned by http.NewRequest in the newRequest
// function.
type ErrorHTTPNewRequest struct {
	err error
}

func (e ErrorHTTPNewRequest) Error() string {
	return "graphql: http new request: " + e.err.Error()
}

// ErrorMutateDo wraps errors returned by http.Client.Do in the Mutate method.
type ErrorMutateDo struct {
	err error
}

func (e ErrorMutateDo) Error() string {
	return "graphql: http do: " + e.err.Error()
}

// ErrorBodyReadAll wraps errors returned by ioutil.ReadAll in the Mutate method.
type ErrorBodyReadAll struct {
	err error
}

func (e ErrorBodyReadAll) Error() string {
	return "graphql: response body read all: " + e.err.Error()
}

// ErrorBodyUnmarshal wraps errors returned json.Unmarshal in the Mutate method.
type ErrorBodyUnmarshal struct {
	err error
}

func (e ErrorBodyUnmarshal) Error() string {
	return "graphql: response body unmarshal: " + e.err.Error()
}

// ErrorGraphQLResponse wraps errors returned by Dgraph.
type ErrorGraphQLResponse struct {
	errs []graphQLError
}

func (e ErrorGraphQLResponse) Error() string {
	messages := make([]string, len(e.errs))

	for i, err := range e.errs {
		messages[i] = err.Message
	}

	message := "graphql: response errors: message: " + strings.Join(messages, ", message: ")
	return message
}
