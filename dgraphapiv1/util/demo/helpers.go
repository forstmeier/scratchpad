package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func randomID() string {
	return uuid.New().String()
}

func randomString(options []string) string {
	return options[rand.Intn(len(options))]
}

func randomDate() string {
	year := time.Now().Year() - (rand.Intn(50) + 20)
	month := rand.Intn(12) + 1
	day := rand.Intn(25) + 1

	dt := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).String()

	return dt
}
