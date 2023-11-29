// service/indices.go
package service

import (
	"github.com/billyagasi/maas/internal/elasticsearch"
	"github.com/billyagasi/maas/internal/model"
)

func GetIndicesInfo(client *elasticsearch.Client) ([]model.IndexInfo, error) {
	// Logic to fetch and process indices information.
	// Convert sizes, handle errors, etc.
}
