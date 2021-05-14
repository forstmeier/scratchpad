package blevedb

import "testing"

func TestType(t *testing.T) {
	bleveSpecimen := &BleveSpecimen{}
	received := bleveSpecimen.Type()
	expected := "bleve_specimen"
	if received != expected {
		t.Errorf("incorrect type, received: %s, expected: %s", received, expected)
	}
}
