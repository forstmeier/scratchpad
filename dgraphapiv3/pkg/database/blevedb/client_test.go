package blevedb

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	filename := "test_new.bleve"

	client, err := New(filename)
	if err != nil {
		t.Errorf("error creating client: %s", err.Error())
	}

	if client == nil {
		t.Errorf("nil client received")
	}

	os.RemoveAll(filename)
}
