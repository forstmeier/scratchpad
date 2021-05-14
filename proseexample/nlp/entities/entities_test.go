package entities

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/jdkato/prose/v2"
)

func TestNew(t *testing.T) {

	configFile, err := ioutil.TempFile("", "config.json")
	if err != nil {
		t.Fatalf("error creating temp config file: %s", err.Error())
	}
	defer os.Remove(configFile.Name())

	configContent := []byte(`[
		{
			"docType": "consent",
			"trainedEntities": ["JEDI"],
			"providedEntities": ["NNP"]
		}	
	]`)

	if _, err := configFile.Write(configContent); err != nil {
		t.Fatalf("error writing data to temp config file: %s", err.Error())
	}

	modelsDir, err := ioutil.TempDir("", "models")
	if err != nil {
		t.Fatalf("error create temporary directory: %s", err.Error())
	}
	defer os.RemoveAll(modelsDir)

	m := prose.ModelFromData("consent", prose.UsingEntities([]prose.EntityContext{
		prose.EntityContext{
			Accept: true,
			Spans: []prose.LabeledEntity{
				prose.LabeledEntity{
					Start: 0,
					End:   3,
					Label: "NN",
				},
			},
			Text: "text",
		},
	}))

	if err := m.Write(modelsDir + "/consent.model"); err != nil {
		t.Fatalf("error creating model file: %s", err.Error())
	}

	processor, err := New(configFile.Name(), modelsDir)
	if processor == nil || err != nil {
		t.Error("error creating nlp:", err.Error())
	}
}

func TestClassifyEntities(t *testing.T) {
	tests := []struct {
		description   string
		configContent []byte
		modelEntities []prose.EntityContext
		inputBlob     string
		inputDoc      string
		error         string
		results       map[string][]string
	}{
		{
			description: "successful classification invocation with trained model",
			configContent: []byte(`[
				{
					"docType": "consent",
					"trainedEntities": ["JEDI"],
					"providedEntities": ["NNP"]
				}	
			]`),
			modelEntities: []prose.EntityContext{
				prose.EntityContext{
					Spans: []prose.LabeledEntity{
						prose.LabeledEntity{
							Start: 10,
							End:   13,
							Label: "JEDI",
						},
					},
					Text:   "John is a Jedi.",
					Accept: true,
				},
				prose.EntityContext{
					Spans: []prose.LabeledEntity{
						prose.LabeledEntity{
							Start: 2,
							End:   5,
							Label: "JEDI",
						},
					},
					Text:   "A Jedi, John is.",
					Accept: true,
				},
			},
			inputBlob: "John is a Jedi.",
			inputDoc:  "consent",
			error:     "",
			results: map[string][]string{
				"NNP":  []string{"John", "Jedi"},
				"JEDI": []string{"Jedi"},
			},
		},
	}

	// NOTE: refactor to generate different models in tests struct slice
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			configFile, err := ioutil.TempFile("", "config.json")
			if err != nil {
				t.Fatalf("error creating temp config file: %s", err.Error())
			}
			defer os.Remove(configFile.Name())

			if _, err := configFile.Write(test.configContent); err != nil {
				t.Fatalf("error writing data to temp config file: %s", err.Error())
			}

			model := prose.ModelFromData("JEDI", prose.UsingEntities(test.modelEntities))
			modelsDir, err := ioutil.TempDir("", "models")
			if err != nil {
				t.Fatalf("error creating temp model directory: %s", err.Error())
			}
			defer os.RemoveAll(modelsDir)

			if err := model.Write(filepath.Join(modelsDir, "JEDI")); err != nil {
				t.Fatalf("error writing model file to temp model file: %s", err.Error())
			}

			classifier, err := New(configFile.Name(), modelsDir)
			if err != nil {
				t.Fatalf("error creating new nlp: %s", err.Error())
			}

			results, err := classifier.ClassifyEntities(test.inputBlob, test.inputDoc)
			if err != nil && err.Error() != test.error {
				t.Error("error classifying words:", err.Error())
			}

			if !reflect.DeepEqual(results, test.results) {
				t.Errorf("incorrect target words received: %+v, expected: %+v\n", results, test.results)
			}
		})
	}
}
