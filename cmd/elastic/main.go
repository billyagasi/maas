package main

import (
	"fmt"

	"github.com/billyagasi/maas/internal/config"
	"github.com/billyagasi/maas/internal/elasticsearch"
	"github.com/billyagasi/maas/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	indexInfo, err := service.GetIndicesInfo(client)
	if err != nil {
		panic(err)
	}

	fmt.Println(indexInfo)
}
