package elasticsearch

import (
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/yourusername/yourproject/internal/config"
)

func NewClient(cfg *config.Config) (*elasticsearch.Client, error) {
	esConfig := elasticsearch.Config{
		Addresses: []string{
			cfg.Elasticsearch.Scheme + "://" + cfg.Elasticsearch.Host + ":" + strconv.Itoa(cfg.Elasticsearch.Port),
		},
		Username: cfg.Elasticsearch.Username,
		Password: cfg.Elasticsearch.Password,
	}
	return elasticsearch.NewClient(esConfig)
}
