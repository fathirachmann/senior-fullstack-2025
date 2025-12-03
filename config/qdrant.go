package config

import (
	"log"

	"github.com/qdrant/go-client/qdrant"
)

func NewQDrantConfig(uri, apiKey string) *qdrant.Client {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host:   uri,
		Port:   6334,
		APIKey: apiKey,
		UseTLS: true,
	})

	if err != nil {
		panic("Failed to connect to Qdrant: " + err.Error())
	}

	log.Println("Qdrant Connected")

	return client

}
