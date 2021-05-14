package user

import (
	"errors"
	"fmt"
)

// ErrorCreateUserFieldsNotSet wraps fields returned by checkCreateUserFieldsSet in
// the CreateUser method.
type ErrorCreateUserFieldsNotSet struct {
	err error
}

func newErrorCreateUserFieldsNotSet(fields string) ErrorCreateUserFieldsNotSet {
	return ErrorCreateUserFieldsNotSet{
		err: fmt.Errorf("users: create user fields not set: %s", fields),
	}
}

func (e ErrorCreateUserFieldsNotSet) Error() string {
	return e.err.Error()
}

// ErrorUpdateUserFieldsNotSet wraps fields returned by checkUpdateUserFieldsSet in
// the UpdateUser method.
type ErrorUpdateUserFieldsNotSet struct {
	err error
}

func newErrorUpdateUserFieldsNotSet(fields string) ErrorUpdateUserFieldsNotSet {
	return ErrorUpdateUserFieldsNotSet{
		err: fmt.Errorf("users: update user fields not set: %s", fields),
	}
}

func (e ErrorUpdateUserFieldsNotSet) Error() string {
	return e.err.Error()
}

// ErrorDeleteUserFieldNotSet handles a missing field in the DeleteUser method.
type ErrorDeleteUserFieldNotSet struct {
	err error
}

func newErrorDeleteUserFieldNotSet() ErrorDeleteUserFieldNotSet {
	return ErrorDeleteUserFieldNotSet{
		err: errors.New("users: delete user field not set: authZeroID"),
	}
}

func (e ErrorDeleteUserFieldNotSet) Error() string {
	return e.err.Error()
}

// ErrorMarshalRequest wraps errors returned by json.Marshal in the newRequest method.
type ErrorMarshalRequest struct {
	err error
}

func newErrorMarshalRequest(err error) ErrorMarshalRequest {
	return ErrorMarshalRequest{
		err: fmt.Errorf("users: marshal payload: %w", err),
	}
}

func (e ErrorMarshalRequest) Error() string {
	return e.err.Error()
}

// ErrorCreateRequest wraps errors returned by http.NewRequest in the newRequest method.
type ErrorCreateRequest struct {
	err error
}

func newErrorCreateRequest(err error) ErrorCreateRequest {
	return ErrorCreateRequest{
		err: fmt.Errorf("users: create request: %w", err),
	}
}

func (e ErrorCreateRequest) Error() string {
	return e.err.Error()
}

// ErrorClientDo wraps errors returned by http.Client.Do in the sendRequest method.
type ErrorClientDo struct {
	err error
}

func newErrorClientDo(err error) ErrorClientDo {
	return ErrorClientDo{
		err: fmt.Errorf("users: client do: %w", err),
	}
}

func (e ErrorClientDo) Error() string {
	return e.err.Error()
}

// ErrorCheckStatus wraps errors returned by checkStatus in the sendRequest method.
type ErrorCheckStatus struct {
	err error
}

func newErrorCheckStatus(statusCode int) ErrorCheckStatus {
	return ErrorCheckStatus{
		err: fmt.Errorf("users: check status: status code: %d", statusCode),
	}
}

func (e ErrorCheckStatus) Error() string {
	return e.err.Error()
}

// ErrorResponseReadBody wraps errors returned by ioutil.ReadAll in the sendRequest method.
type ErrorResponseReadBody struct {
	err error
}

func newErrorResponseReadBody(err error) ErrorResponseReadBody {
	return ErrorResponseReadBody{
		err: fmt.Errorf("users: response read body: %w", err),
	}
}

func (e ErrorResponseReadBody) Error() string {
	return e.err.Error()
}

// ErrorResponseUnmarshal wraps errors returned by json.Unmarshal in the CreateUser method.
type ErrorResponseUnmarshal struct {
	err error
}

func newErrorResponseUnmarshal(err error) ErrorResponseUnmarshal {
	return ErrorResponseUnmarshal{
		err: fmt.Errorf("users: response read body: %w", err),
	}
}

func (e ErrorResponseUnmarshal) Error() string {
	return e.err.Error()
}
