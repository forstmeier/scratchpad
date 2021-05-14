// Package nlp provides the interface for interacting with various
// natural language processing features
package nlp

/*
nlp                      -> parent package for all language processing packages
├── doc.go
└── entities             -> package for parsing words per doc type
    ├── config.json      -> maps word entity models to doc types
    ├── entities.go      -> in-memory entity models and classifier method
    ├── entities_test.go
    ├── errors.go
    └── model            -> executable for generating entity models
        ├── data         -> folder for entity training json files
        ├── model.go     -> model generator script
        └── models       -> folder generated model files
*/
