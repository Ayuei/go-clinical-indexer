package main

import (
	"context"
	"flag"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"indexer/consumers"
	"indexer/producers"
	"indexer/structs/csvs"
	"indexer/utils"
	"strconv"
	"sync"
)

func checkArgs(elasticIndex string, DataPath string, NumWorkers int) {

	if len(elasticIndex) == 0 {
		panic("Index not specified! Use -h flag to discover options.")
	}

	if len(DataPath) == 0 {
		panic("Data path not specified! Use -h flag to discover options.")
	}

	if NumWorkers <= 0 {
		panic("Invalid number of workers specified! Use -h flag to discover options.")
	}
}

func main() {
	elasticUrl := flag.String("url", "http://127.0.0.1", "Elasticsearch URL endpoint")
	elasticPort := flag.Int("port", 9200, "Elasticsearch URL port")
	elasticIndex := flag.String("index", "", "Elasticsearch index name")
	DocType := flag.String("collection", "med-marco", "Data collection to index")
	filterFp := flag.String("filter", "", "List of doc ids (one per line) to include in index")
	exclude := flag.Bool("excludeFilter", false, "Negate the filter option to exclude documents rather than include")
	Delete := flag.Bool("delete", false, "Delete index if it exists")
	DataPath := flag.String("data_path", ".", "Data collection location")
	NumWorkers := flag.Int("workers", 1, "Number of parallel consumers")
	Accurate := flag.Bool("accurate", false, "Accurate progress bar, count lines before producing")

	flag.Parse()

	checkArgs(*elasticIndex, *DataPath, *NumWorkers)

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{*elasticUrl + ":" + strconv.Itoa(*elasticPort)},
	})

	utils.CheckError(err, "client")

	p, err := esutil.NewBulkIndexer(
		esutil.BulkIndexerConfig{
			NumWorkers: *NumWorkers,
			FlushBytes: 65536,
			Client:     client,
			Index:      *elasticIndex,
		})

	utils.CheckError(err, "bulk")
	CreateIndexIfNotExists(client, *elasticIndex, *Delete)

	filter := utils.CreateFilterMap(*filterFp)

	var wg sync.WaitGroup
	var producer func(string, chan string, *sync.WaitGroup, bool)
	var consumer func(chan string, esutil.BulkIndexer, *sync.WaitGroup, map[string]bool, bool)

	producer, consumer = ProducerConsumerFactory(*DocType)

	if producer == nil || consumer == nil {
		if *DocType == "bioreddit_submissions" {
			jobs := make(chan csvs.BioRedditSubmissions)
			wg.Add(1)
			go producers.BioredditSubmissionCSVProducer(*DataPath, jobs, &wg, *Accurate)
			for i := 0; i < *NumWorkers; i++ {
				go consumers.ParseBioRedditSubmission(jobs, p, &wg, filter, *exclude)

				wg.Add(1)
			}
			wg.Wait()
		} else if *DocType == "bioreddit_comments" {
			jobs := make(chan csvs.BioRedditComments)
			wg.Add(1)
			go producers.BioredditCommentCSVProducer(*DataPath, jobs, &wg, *Accurate)
			for i := 0; i < *NumWorkers; i++ {
				//fmt.Printf("Started Worker")
				go consumers.ParseBioRedditComment(jobs, p, &wg, filter, *exclude)

				wg.Add(1)
			}
			wg.Wait()

		} else {
			panic("Unable to find valid document type")
		}
	}

	jobs := make(chan string, 1)
	// Producer thread
	if producer != nil {
		wg.Add(1)
		go producer(*DataPath, jobs, &wg, *Accurate)

		// Consumers
		if consumer != nil {
			for i := 0; i < *NumWorkers; i++ {
				// fmt.Printf("Started Worker")
				go consumer(jobs, p, &wg, filter, *exclude)

				wg.Add(1)
			}
		} else {
			//panic("No consumer created")
		}
	} else {
		//panic("No producer created")
	}

	wg.Wait()

	utils.CheckError(err, "flush")
	err = p.Close(context.Background())
	utils.CheckError(err, "flush")
}
