package handlers

import "fmt"

const packageName = "handlers"

// ErrorCSVReadAll wraps errors returned by csv.ReadAll in the
// SpecimensFromCSV function.
type ErrorCSVReadAll struct {
	err error
}

func (e *ErrorCSVReadAll) Error() string {
	return fmt.Sprintf("%s: csv read all: %s", packageName, e.err.Error())
}

// ErrorNotEnoughRows indicates that at least one header row and one
// data row have not been provided.
type ErrorNotEnoughRows struct{}

func (e *ErrorNotEnoughRows) Error() string {
	return fmt.Sprintf("%s: at least one header and one row required in csv", packageName)
}

// ErrorJSONEncode wraps errors returned by json.Encode in the
// SpecimensFromCSV function.
type ErrorJSONEncode struct {
	err error
}

func (e *ErrorJSONEncode) Error() string {
	return fmt.Sprintf("%s: json encode: %s", packageName, e.err.Error())
}

// ErrorJSONDecode wraps errors returned by json.Encode in the
// SpecimensFromCSV function.
type ErrorJSONDecode struct {
	err error
}

func (e *ErrorJSONDecode) Error() string {
	return fmt.Sprintf("%s: json decode: %s", packageName, e.err.Error())
}
