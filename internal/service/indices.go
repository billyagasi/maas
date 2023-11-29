package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/billyagasi/maas/internal/elasticsearch"
	"github.com/billyagasi/maas/internal/model"
)

// GetIndicesInfo connects to Elasticsearch and retrieves index information
func GetIndicesInfo(client *elasticsearch.Client) ([]model.IndexInfo, error) {
	// Context for Elasticsearch operations
	ctx := context.Background()

	// Fetching all indices
	response, err := client.Cat.Indices(ctx, client.Cat.Indices.WithFormat("json"), client.Cat.Indices.WithH("index", "store.size", "status"))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Parse response to a slice of IndexInfo
	var indices []model.IndexInfo
	// ... code to unmarshal JSON from response into `indices` ...

	// Processing each index
	for i, index := range indices {
		index.SizeKB = convertSizeToKB(index.StoreSize)
		indices[i] = index

		// Fetching health status
		healthResponse, err := client.Cluster.Health(client.Cluster.Health.WithIndex(index.IndexName))
		if err != nil {
			indices[i].Health = "Unknown"
			continue
		}
		defer healthResponse.Body.Close()

		// Parse healthResponse to get the health status
		// ... code to unmarshal JSON from healthResponse and set `indices[i].Health` ...
	}

	return indices, nil
}

// convertSizeToKB converts size string to numeric KB
func convertSizeToKB(sizeStr string) float64 {
	// Splitting the string to isolate the numeric part
	numericPart := strings.Split(sizeStr, " ")[0]
	size, err := strconv.ParseFloat(numericPart, 64)
	if err != nil {
		return 0
	}
	return size
}
