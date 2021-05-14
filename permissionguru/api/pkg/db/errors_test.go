package db

import (
	"errors"
	"testing"
)

func TestErrorMissingID(t *testing.T) {
	err := &ErrorMissingID{}

	recieved := err.Error()
	expected := "db: add account: missing id"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorPutItem(t *testing.T) {
	err := &ErrorPutItem{err: errors.New("mock put item error")}

	recieved := err.Error()
	expected := "db: add account: mock put item error"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorMissingFields(t *testing.T) {
	err := &ErrorMissingFields{
		fields: []string{
			"accountID",
			"subscriptionID",
		},
	}

	recieved := err.Error()
	expected := "db: add account: missing ids: accountID, subscriptionID"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorUpdateItem(t *testing.T) {
	err := &ErrorUpdateItem{
		item: "subscription",
		err:  errors.New("mock update item error"),
	}

	recieved := err.Error()
	expected := "db: add subscription: mock update item error"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorNoConfigSelections(t *testing.T) {
	err := &ErrorNoConfigSelections{}

	recieved := err.Error()
	expected := "db: add config: missing config selections"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorMarshalConfig(t *testing.T) {
	err := &ErrorMarshalConfig{err: errors.New("mock marshal config error")}

	recieved := err.Error()
	expected := "db: add config: mock marshal config error"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorGetItem(t *testing.T) {
	err := &ErrorGetItem{err: errors.New("mock get item error")}

	recieved := err.Error()
	expected := "db: get account: mock get item error"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorNoSubscription(t *testing.T) {
	err := &ErrorNoSubscription{}

	recieved := err.Error()
	expected := "db: get account: missing subscription values"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorUnmarshalConfig(t *testing.T) {
	err := &ErrorUnmarshalConfig{err: errors.New("mock unmarshal config error")}

	recieved := err.Error()
	expected := "db: get account: mock unmarshal config error"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorNoAccount(t *testing.T) {
	err := &ErrorNoAccount{}

	recieved := err.Error()
	expected := "db: get account: no account found"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorDeleteItem(t *testing.T) {
	err := &ErrorDeleteItem{err: errors.New("mock delete item error")}

	recieved := err.Error()
	expected := "db: remove account: mock delete item error"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}
