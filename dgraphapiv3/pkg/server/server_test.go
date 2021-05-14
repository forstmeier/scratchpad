package server

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	server, err := New()
	if err != nil {
		t.Errorf("error ceating server: %s", err.Error())
	}

	if server == nil {
		t.Error("error creating server")
	}

	os.RemoveAll(databaseFile)
}
