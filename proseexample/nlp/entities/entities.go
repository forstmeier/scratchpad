package entities

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/jdkato/prose/v2"
)

// Classifier interface defines the contract for NLP features.
type Classifier interface {
	ClassifyEntities(blob string, doc string) (map[string][]string, error)
}

// Processor holds trained entity models exposes the entity classification method.
//
// Struct fields:
// - entitiesConfig -> key: doc name, value: target words
// - entitiesModels -> key: doc name, value: target words models
type Processor struct {
	entitiesConfig map[string]config
	entitiesModels map[string][]prose.DocOpt
}

type config struct {
	DocType          string   `json:"docType"`
	TrainedEntities  []string `json:"trainedEntities"`
	ProvidedEntities []string `json:"providedEntities"`
}

// New generates an instance of the Classifier interface.
func New(configPath, modelsDir string) (Classifier, error) {
	// read in document type configuration file
	configJSON := []config{}
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, newErrorReadConfigFile(err)
	}

	if err := json.Unmarshal(configData, &configJSON); err != nil {
		return nil, newErrorUnmarshalConfigJSON(err)
	}

	// read in model files
	modelsFiles, err := ioutil.ReadDir(modelsDir)
	if err != nil {
		return nil, newErrorReadModelsDirectory(err)
	}

	models := map[string]prose.DocOpt{}
	for _, modelFile := range modelsFiles {
		modelName := modelFile.Name()
		model := prose.UsingModel(prose.ModelFromDisk(filepath.Join(modelsDir, modelName)))
		models[modelName] = model
	}

	// create config/model references by document type
	entitiesConfig := map[string]config{}
	entitiesModels := map[string][]prose.DocOpt{}
	for _, docConfig := range configJSON {
		wordModels := []prose.DocOpt{}
		for _, word := range docConfig.TrainedEntities {
			wordModels = append(wordModels, models[word])
		}

		entitiesConfig[strings.ToLower(docConfig.DocType)] = docConfig
		entitiesModels[strings.ToLower(docConfig.DocType)] = wordModels
	}

	return &Processor{
		entitiesConfig: entitiesConfig,
		entitiesModels: entitiesModels,
	}, nil

}

// ClassifyEntities identifies and collects any target words for a given text blob and doc type.
func (p *Processor) ClassifyEntities(blob string, docType string) (map[string][]string, error) {
	document, err := prose.NewDocument(blob, p.entitiesModels[strings.ToLower(docType)]...)
	if err != nil {
		return nil, newErrorCreateDocument(err)
	}

	results := make(map[string][]string)

	// NOTE: this produces duplicates between the TrainedEntities and ProvidedEntities if
	// there is a match on the latter - this logic should eventually be cleaned up
	for _, entity := range document.Tokens() {
		for _, trainedWord := range p.entitiesConfig[strings.ToLower(docType)].TrainedEntities {
			if strings.Contains(entity.Label, trainedWord) {
				if _, ok := results[trainedWord]; ok {
					results[trainedWord] = append(results[trainedWord], entity.Text)
				} else {
					results[trainedWord] = []string{entity.Text}
				}
			}
		}

		for _, providedWord := range p.entitiesConfig[strings.ToLower(docType)].ProvidedEntities {
			if providedWord == entity.Tag {
				if _, ok := results[providedWord]; ok {
					results[providedWord] = append(results[providedWord], entity.Text)
				} else {
					results[providedWord] = []string{entity.Text}
				}
			}
		}
	}

	return results, nil
}
