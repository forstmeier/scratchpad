package db

import (
	"fmt"
	"strings"
)

const packageName = "db"

// ErrorMissingID is returned if an empty ID value is provided to
// AddAccount.
type ErrorMissingID struct{}

func (e *ErrorMissingID) Error() string {
	return fmt.Sprintf("%s: add account: missing id", packageName)
}

// ErrorPutItem wraps errors returned by dynamodb.DynamoDB.PutItem
// in the AddAccount method.
type ErrorPutItem struct {
	err error
}

func (e *ErrorPutItem) Error() string {
	return fmt.Sprintf("%s: add account: %s", packageName, e.err.Error())
}

// ErrorMissingFields is returned if any Account fields are not populated.
type ErrorMissingFields struct {
	fields []string
}

func (e *ErrorMissingFields) Error() string {
	ids := strings.Join(e.fields, ", ")
	return fmt.Sprintf("%s: add account: missing ids: %s", packageName, ids)
}

// ErrorUpdateItem wraps errors returned by dynamodb.DynamoDB.UpdateItem
// in the AddSubscription and AddConfig methods.
type ErrorUpdateItem struct {
	item string
	err  error
}

func (e *ErrorUpdateItem) Error() string {
	return fmt.Sprintf("%s: add %s: %s", packageName, e.item, e.err.Error())
}

// ErrorNoConfigSelections is returned if a Config is provided with no
// Selections.
type ErrorNoConfigSelections struct{}

func (e *ErrorNoConfigSelections) Error() string {
	return fmt.Sprintf("%s: add config: missing config selections", packageName)
}

// ErrorMarshalConfig wraps errors returned by json.Marshal in the
// AddConfig method.
type ErrorMarshalConfig struct {
	err error
}

func (e *ErrorMarshalConfig) Error() string {
	return fmt.Sprintf("%s: add config: %s", packageName, e.err.Error())
}

// ErrorGetItem wraps errors returned by dynamodb.DynamoDB.GetItem
// in the GetAccount method.
type ErrorGetItem struct {
	err error
}

func (e *ErrorGetItem) Error() string {
	return fmt.Sprintf("%s: get account: %s", packageName, e.err.Error())
}

// ErrorNoAccount is returned if no Account value is found
// by GetAccount.
type ErrorNoAccount struct{}

func (e *ErrorNoAccount) Error() string {
	return fmt.Sprintf("%s: get account: no account found", packageName)
}

// ErrorNoSubscription is returned if no Subscription values were stored
// in the Account being retrieved by GetAccount.
type ErrorNoSubscription struct{}

func (e *ErrorNoSubscription) Error() string {
	return fmt.Sprintf("%s: get account: missing subscription values", packageName)
}

// ErrorUnmarshalConfig wraps errors returned by json.Unmarshal in the
// GetAccount method.
type ErrorUnmarshalConfig struct {
	err error
}

func (e *ErrorUnmarshalConfig) Error() string {
	return fmt.Sprintf("%s: get account: %s", packageName, e.err.Error())
}

// ErrorDeleteItem wraps errors returned by dynamodb.DynamoDB.DeleteItem
// in the RemoveAccount method.
type ErrorDeleteItem struct {
	err error
}

func (e *ErrorDeleteItem) Error() string {
	return fmt.Sprintf("%s: remove account: %s", packageName, e.err.Error())
}
