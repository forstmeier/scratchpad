package data

import "testing"

func TestNewSpecimen(t *testing.T) {
	body := `{"text":"data"}`

	specimen := NewSpecimen(body)
	if specimen.ID == "" || specimen.Timestamp.IsZero() || specimen.Body != body {
		t.Errorf("error creating specimen, received: %+v\n", specimen)
	}
}
