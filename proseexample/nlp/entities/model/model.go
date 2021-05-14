package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"

	"github.com/jdkato/prose/v2"
)

type docEntity struct {
	Sentence string `json:"sentence"`
	Word     string `json:"word"`
	Start    int    `json:"start"`
	End      int    `json:"end"`
	Label    string `json:"label"`
}

func main() {
	inputFile := flag.String("input", "", "target json file in data/ containing training data")
	trainPercent := flag.Float64("train", 0.8, "data split for model training")
	targetWord := flag.String("word", "", "target word for trained model")

	flag.Parse()

	message := "%s arg must be provided\n"
	if inputFile == nil {
		log.Fatalf(message, "input")
	}
	if trainPercent == nil {
		log.Fatalf(message, "trainPercent")
	}
	if targetWord == nil {
		log.Fatalf(message, "targetWord")
	}

	trainingData, err := readData(filepath.Join("data", *inputFile))
	if err != nil {
		log.Fatalf("error reading training data: %s", err.Error())
	}

	train, test := splitData(trainingData, *trainPercent)

	model := prose.ModelFromData(*targetWord, prose.UsingEntities(train))

	correct := 0.0
	for _, entry := range test {
		doc, err := prose.NewDocument(
			entry.Sentence,
			prose.WithSegmentation(false),
			prose.UsingModel(model),
		)
		if err != nil {
			log.Fatalf("error creating test document: %s", err.Error())
		}

		docEntities := doc.Entities()
		if len(docEntities) == 0 {
			correct++
		} else {
			expected := entry.Sentence[entry.Start:entry.End]
			if reflect.DeepEqual(expected, docEntities) {
				correct++
			}
		}
	}

	log.Printf("correct (%%): %f\n", correct/float64(len(test)))
	model.Write(filepath.Join("models", *targetWord))
}

func readData(path string) ([]docEntity, error) {
	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(bytes.NewReader(jsonData))
	entities := []docEntity{}
	for {
		entity := docEntity{}
		err := decoder.Decode(&entity)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("error reading json data:", err.Error())
		}

		entities = append(entities, entity)
	}
	return entities, nil
}

func splitData(data []docEntity, splitPercent float64) ([]prose.EntityContext, []docEntity) {
	cutoff := int(float64(len(data)) * splitPercent)
	train, test := []prose.EntityContext{}, []docEntity{}
	for i, entity := range data {
		if i < cutoff {
			train = append(train, prose.EntityContext{
				Text: entity.Sentence,
				Spans: []prose.LabeledEntity{
					prose.LabeledEntity{
						Start: entity.Start,
						End:   entity.End,
						Label: entity.Label,
					},
				},
			})
		} else {
			test = append(test, entity)
		}
	}

	return train, test
}
