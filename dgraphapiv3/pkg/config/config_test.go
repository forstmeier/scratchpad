package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	esURL := "test_url"

	file, err := ioutil.TempFile(".", "config.json")
	if err != nil {
		t.Fatalf("error creating temp file: %s", err.Error())
	}
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(fmt.Sprintf(`{"API": {"ES_URL": "%s"}}`, esURL)))
	if err != nil {
		t.Fatalf("error writing temp file data: %s", err.Error())
	}

	cfg, err := New(file.Name())
	if err != nil {
		t.Fatalf("error creating test config: %s", err.Error())
	}

	if cfg.API.ESURL != esURL {
		t.Errorf("incorrect es url, received: %s, expected: %s", cfg.API.ESURL, esURL)
	}
}
