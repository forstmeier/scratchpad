package handlers

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/specimenguru/api/pkg/data"
)

// specimensFromCSV reads CSV formatted data from the received input and parses
// it into a slice of Specimen.
func specimensFromCSV(input io.ReadCloser) ([]data.Specimen, error) {
	reader := csv.NewReader(input)

	content, err := reader.ReadAll()
	if err != nil {
		return nil, &ErrorCSVReadAll{err: err}
	}

	if len(content) < 2 {
		return nil, &ErrorNotEnoughRows{}
	}

	specimens := make([]data.Specimen, len(content)-1) // don't count header row

	headers := content[0]
	for i, row := range content[1:] {
		specimenBody := make(map[string]interface{})
		for j, header := range headers {
			specimenBody[header] = row[j]
		}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(specimenBody); err != nil {
			return nil, &ErrorJSONEncode{err: err}
		}

		specimens[i] = data.NewSpecimen(strings.TrimSuffix(buf.String(), "\n"))
	}

	return specimens, nil
}

type specimenDTO struct {
	Body string `json:"body"`
}

func (s *specimenDTO) UnmarshalJSON(input []byte) error {
	s.Body = string(input)
	return nil
}

// specimensFromJSON reads JSON formatted data from the received input and parses
// it into a slice of Specimen.
func specimensFromJSON(input io.ReadCloser) ([]data.Specimen, error) {
	specimenDTOs := []specimenDTO{}
	if err := json.NewDecoder(input).Decode(&specimenDTOs); err != nil {
		return nil, &ErrorJSONDecode{err: err}
	}

	specimens := make([]data.Specimen, len(specimenDTOs))
	for i, s := range specimenDTOs {
		specimens[i] = data.NewSpecimen(s.Body)
	}

	return specimens, nil
}

func errorMessage(message string) string {
	return fmt.Sprintf(`{"error":"%s"}`, message)
}
