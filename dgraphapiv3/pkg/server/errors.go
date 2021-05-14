package server

import "fmt"

const packageName = "server"

// ErrorNewDatabase wraps errors returned by blevedb.New in the
// New function.
type ErrorNewDatabase struct {
	err error
}

func (e *ErrorNewDatabase) Error() string {
	return fmt.Sprintf("%s: new database: %s", packageName, e.err.Error())
}

// ErrorNewHandlersClient wraps errors returned by handlers.NewClient in the
// New function.
type ErrorNewHandlersClient struct {
	err error
}

func (e *ErrorNewHandlersClient) Error() string {
	return fmt.Sprintf("%s: new handlers client: %s", packageName, e.err.Error())
}
