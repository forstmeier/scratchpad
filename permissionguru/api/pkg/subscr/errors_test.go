package subscr

import (
	"errors"
	"testing"
)

func TestErrorMissingFields(t *testing.T) {
	err := &ErrorMissingFields{
		fields: []string{
			"email",
			"zip",
		},
	}

	recieved := err.Error()
	expected := "subscr: create subscription: missing ids: email, zip"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorNewPaymentMethod(t *testing.T) {
	err := &ErrorNewPaymentMethod{err: errors.New("mock new payment method error")}

	recieved := err.Error()
	expected := "subscr: create subscription: mock new payment method error"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorNewCustomer(t *testing.T) {
	err := &ErrorNewCustomer{err: errors.New("mock new customer error")}

	recieved := err.Error()
	expected := "subscr: create subscription: mock new customer error"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorNewSubscription(t *testing.T) {
	err := &ErrorNewSubscription{err: errors.New("mock new subscription error")}

	recieved := err.Error()
	expected := "subscr: create subscription: mock new subscription error"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorDeleteCustomer(t *testing.T) {
	err := &ErrorDeleteCustomer{err: errors.New("mock delete customer error")}

	recieved := err.Error()
	expected := "subscr: remove subscription: mock delete customer error"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}

func TestErrorCreateUsageRecord(t *testing.T) {
	err := &ErrorCreateUsageRecord{err: errors.New("mock create usage record error")}

	recieved := err.Error()
	expected := "subscr: add usage: mock create usage record error"

	if recieved != expected {
		t.Errorf("incorrect error message, received: %s, expected: %s", recieved, expected)
	}
}
