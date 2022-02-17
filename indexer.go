package main

import (
	"context"
	"github.com/olivere/elastic/v7"
	"indexer/utils"
	"log"
)

func CreateIndexIfNotExists(es *elastic.Client, index string, delete bool) {
	mappings := `{
		"settings":{
			"number_of_shards":1,
            "number_of_replicas":0,
            "index.mapping.ignore_malformed": true 
		}
	}`

	indexExists, err := es.IndexExists(index).Do(context.Background())

	if err != nil {
		log.Fatalf("Something went wrong when checking index exists %s", err)
		return
	}

	if indexExists && delete {
		log.Print("Index exists, deleting")
		_, err := es.DeleteIndex(index).Do(context.Background())
		utils.CheckError(err, "Delete Index")
		return
	}

	_, err = es.CreateIndex(index).BodyString(mappings).Do(context.Background())
	utils.CheckError(err, "Create index")
}
