package config

import (
	"log"
	"strings"

	"github.com/qdrant/go-client/qdrant"
)

func NewQDrantConfig(uri, apiKey string, port int) *qdrant.Client {
	// Strip scheme if present
	host := strings.TrimPrefix(uri, "https://")
	host = strings.TrimPrefix(host, "http://")

	client, err := qdrant.NewClient(&qdrant.Config{
		Host:   host,
		Port:   port,
		APIKey: apiKey,
		UseTLS: true,
	})

	if err != nil {
		panic("Failed to connect to Qdrant: " + err.Error())
	}

	log.Println("Qdrant Connected")

	return client

}
