package main

import (
	"log"

	"github.com/billyagasi/maas/internal/config"
	"github.com/billyagasi/maas/internal/elasticsearch"
	"github.com/billyagasi/maas/internal/service"
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

	// Example of using esClient
	indicesInfo, err := service.GetIndicesInfo(esClient)
	if err != nil {
		log.Fatalf("Error getting indices information: %v", err)
	}

	// Do something with the indicesInfo, like printing it
	for _, info := range indicesInfo {
		log.Printf("Index: %s, Size (KB): %f, Status: %s, Health: %s\n", info.IndexName, info.SizeKB, info.Status, info.Health)
	}
}
