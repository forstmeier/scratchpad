package config

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	_, err := New("no-file")
	if errors.Is(err, ErrorReadFile{}) {
		t.Fatalf("incorrect error, received: %s", err.Error())
	}

	file, err := ioutil.TempFile(".", "config.json")
	if err != nil {
		t.Fatalf("error creating temp file: %s", err.Error())
	}
	defer os.Remove(file.Name())

	file.Write([]byte("---------"))

	_, err = New(file.Name())
	if errors.Is(err, ErrorUnmarshalConfig{}) {
		t.Errorf("incorrect error, received: %s", err.Error())
	}

	config, err := New("../../etc/config/config.json")
	if err != nil {
		t.Errorf("error reading config: %s, config: %+v\n", err.Error(), config)
	}
}
