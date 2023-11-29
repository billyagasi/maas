package main

import (
	"log"

	"github.com/billyagasi/maas/internal/config"
	"github.com/billyagasi/maas/internal/elasticsearch"
)

func main() {
	cfg, err := config.LoadConfig("./configs")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}

	// Now you can use esClient with your services
	// ...
}
