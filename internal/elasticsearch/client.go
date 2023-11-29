// elasticsearch/client.go
package elasticsearch

import (
	"strconv"

	"github.com/billyagasi/maas/internal/config"
	"github.com/elastic/go-elasticsearch/v8"
)

func NewClient(cfg *config.Config) (*elasticsearch.Client, error) {
	esCfg := elasticsearch.Config{
		Addresses: []string{
			cfg.Elasticsearch.Scheme + "://" + cfg.Elasticsearch.Host + ":" + strconv.Itoa(cfg.Elasticsearch.Port),
		},
	}
	return elasticsearch.NewClient(esCfg)
}
