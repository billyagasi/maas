package elasticsearch

import (
	"log"
	"strconv"

	"github.com/billyagasi/maas/internal/config"
	"github.com/elastic/go-elasticsearch/v7"
)

// NewClient creates and returns a new Elasticsearch client
func NewClient(cfg *config.Config) (*elasticsearch.Client, error) {
	cfgElasticsearch := cfg.Elasticsearch

	esConfig := elasticsearch.Config{
		Addresses: []string{
			cfgElasticsearch.Scheme + "://" + cfgElasticsearch.Host + ":" + strconv.Itoa(cfgElasticsearch.Port),
		},
		// Uncomment the following lines if you need authentication
		// Username: cfgElasticsearch.Username,
		// Password: cfgElasticsearch.Password,
	}

	client, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return nil, err
	}

	// Optionally, you can check if the cluster is up and running
	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil, err
	}
	defer res.Body.Close()

	log.Println(res)

	return client, nil
}

// The above function assumes you have a configuration struct similar to:
//
// type Config struct {
//     Elasticsearch struct {
//         Host   string
//         Port   int
//         Scheme string
//         // Uncomment if using authentication
//         // Username string
//         // Password string
//     }
// }
//
// This struct should be defined in your `config` package and loaded before you call NewClient.
