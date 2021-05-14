package config

import (
	"errors"
	"testing"
)

func TestErrorReadFileError(t *testing.T) {
	input := errors.New("input error")

	err := ErrorReadFile{err: input}

	received := err.Error()
	expected := "config: read file: input error"

	if received != expected {
		t.Errorf("error received: %s, expected: %s", received, expected)
	}
}

func TestErrorUnmarshalConfigError(t *testing.T) {
	input := errors.New("input error")

	err := ErrorUnmarshalConfig{err: input}

	received := err.Error()
	expected := "config: unmarshal: input error"

	if received != expected {
		t.Errorf("error received: %s, expected: %s", received, expected)
	}
}
