package blevedb

import "fmt"

const packageName = "blevedb"

// ErrorNewBleveIndex wraps errors returned by bleve.New in the New function.
type ErrorNewBleveIndex struct {
	err error
}

func (e *ErrorNewBleveIndex) Error() string {
	return fmt.Sprintf("%s: new bleve index: %s", packageName, e.err.Error())
}

// ErrorOpenBleveIndex wraps errors returned by bleve.Open in the New function.
type ErrorOpenBleveIndex struct {
	err error
}

func (e *ErrorOpenBleveIndex) Error() string {
	return fmt.Sprintf("%s: open bleve index: %s", packageName, e.err.Error())
}

// ErrorNoAddData indicates no data has been provided to the Add method.
type ErrorNoAddData struct{}

func (e *ErrorNoAddData) Error() string {
	return fmt.Sprintf("%s: no data provided to add method", packageName)
}

// ErrorNoQuery indicates no full text query has been provided to the Query method.
type ErrorNoQuery struct{}

func (e *ErrorNoQuery) Error() string {
	return fmt.Sprintf("%s: no full text query provided to query method", packageName)
}

// ErrorRunBleveQuery wraps errors returned by bleve.Search in the Query method.
type ErrorRunBleveQuery struct {
	err error
}

func (e *ErrorRunBleveQuery) Error() string {
	return fmt.Sprintf("%s: run bleve query: %s", packageName, e.err.Error())
}
