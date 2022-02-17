package main

import (
	"context"
	"flag"
	"github.com/olivere/elastic/v7"
	"indexer/consumers"
	"indexer/producers"
	"indexer/utils"
	"strconv"
	"sync"
)

func main() {
	elasticUrl := flag.String("url", "http://127.0.0.1", "Elasticsearch URL endpoint")
	elasticPort := flag.Int("port", 9200, "Elasticsearch URL port")
	elasticIndex := flag.String("index", "", "Elasticsearch index name")
	DocType := flag.String("collection", "med-marco", "Data collection to index")
	filterFp := flag.String("filter", "", "List of doc ids (one per line) to include in index")
	exclude := flag.Bool("excludeFilter", false, "Negate the filter option to exclude documents rather than include")
	Delete := flag.Bool("delete", false, "Delete index if it exists")
	DataPath := flag.String("data_path", ".", "Data collection location")
	NumWorkers := flag.Int("workers", 8, "Number of parallel consumers")

	flag.Parse()

	client, err := elastic.NewClient(
		elastic.SetURL(*elasticUrl + ":" + strconv.Itoa(*elasticPort)))

	utils.CheckError(err, "client")

	p, err := client.BulkProcessor().
		Name("Indexer-1").
		Workers(*NumWorkers).
		Do(context.Background())

	utils.CheckError(err, "bulk")
	CreateIndexIfNotExists(client, *elasticIndex, *Delete)

	jobs := make(chan string, 10000)

	var wg sync.WaitGroup
	var producer func(string, chan string, *sync.WaitGroup)
	var consumer func(chan string, string, *elastic.BulkProcessor, *sync.WaitGroup, map[string]bool, bool)

	if *DocType == "marco" {
		producer = producers.ProduceMarco // ("collection.tsv", "med-msmarco-train.txt", jobs, &wg)
		consumer = consumers.ParseMarcoDocument
	} else if *DocType == "clinical-trials" {
		producer = producers.ProduceClinicalTrials //("/home/vin/Projects/CDS_2021/index/raw_data/*/*.xml", jobs, &wg)
		consumer = consumers.ParseClinicalDocument
	} else if *DocType == "test-trials" {
		producer = producers.ProduceTestClinicalTrials //("/home/vin/Projects/CDS_2021/index/raw_data/*/*.xml", jobs, &wg)
		consumer = consumers.ParseTestClinicalDocument
	} else if *DocType == "generic" {
		producer = producers.GenericProducer
		consumer = consumers.ParsePubmed
	} else {
		panic("Unable to find valid document type")
	}

	// Producer thread
	if producer != nil {
		wg.Add(1)
		go producer(*DataPath, jobs, &wg)
	}

	filter := utils.CreateFilterMap(*filterFp)

	// Consumers
	for i := 0; i < *NumWorkers; i++ {
		// fmt.Printf("Started Worker")
		if consumer != nil {
			go consumer(jobs, *elasticIndex, p, &wg, filter, *exclude)
		}

		wg.Add(1)
	}

	wg.Wait()

	err = p.Flush()
	utils.CheckError(err, "flush")
	err = p.Close()
	utils.CheckError(err, "flush")
}
