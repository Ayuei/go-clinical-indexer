package main

import (
	"github.com/elastic/go-elasticsearch/v8"
	"indexer/utils"
	"log"
	"strings"
)

func CreateIndexIfNotExists(es *elasticsearch.Client, index string, delete bool) {
	mappings := `{
		"settings":{
			"number_of_shards":1,
            "number_of_replicas":0,
            "index.mapping.ignore_malformed": true 
		}
	}`

	//indexExists, err := es.IndexExists(index).Do(context.Background())

	response, err := es.Indices.Exists([]string{index})

	if err != nil {
		log.Fatalf("Something went wrong when checking index exists %s", err)
		return
	}

	if response.StatusCode != 404 && delete {
		log.Print("Index exists, deleting")
		_, err := es.Indices.Delete([]string{index})
		// _, err := es.DeleteIndex(index).Do(context.Background())
		utils.CheckError(err, "Delete Index")
		return
	}

	response, err = es.Indices.Create(index)
	response, err = es.Indices.PutMapping([]string{index}, strings.NewReader(mappings))

	utils.CheckError(err, "Create index")
}
