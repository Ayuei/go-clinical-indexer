package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"indexer/consumers"
	"indexer/producers"
	"indexer/structs/csvs"
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
	NumWorkers := flag.Int("workers", 1, "Number of parallel consumers")
	Accurate := flag.Bool("accurate", false, "Accurate progress bar, count lines before producing")

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

	filter := utils.CreateFilterMap(*filterFp)

	var wg sync.WaitGroup
	var producer func(string, chan string, *sync.WaitGroup, bool)
	var consumer func(chan string, string, *elastic.BulkProcessor, *sync.WaitGroup, map[string]bool, bool)

	producer, consumer = ProducerConsumerFactory(*DocType)

	if producer == nil || consumer == nil {
		if *DocType == "bioreddit_submissions" {
			jobs := make(chan csvs.BioRedditSubmissions)
			wg.Add(1)
			go producers.BioredditSubmissionCSVProducer(*DataPath, jobs, &wg, *Accurate)
			for i := 0; i < *NumWorkers; i++ {
				fmt.Printf("Started Worker")
				go consumers.ParseBioRedditSubmission(jobs, *elasticIndex, p, &wg, filter, *exclude)

				wg.Add(1)
			}
			wg.Wait()
		} else if *DocType == "bioreddit_comments" {
			jobs := make(chan csvs.BioRedditComments)
			wg.Add(1)
			go producers.BioredditCommentCSVProducer(*DataPath, jobs, &wg, *Accurate)
			for i := 0; i < *NumWorkers; i++ {
				fmt.Printf("Started Worker")
				go consumers.ParseBioRedditComment(jobs, *elasticIndex, p, &wg, filter, *exclude)

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
				go consumer(jobs, *elasticIndex, p, &wg, filter, *exclude)

				wg.Add(1)
			}
		} else {
			//panic("No consumer created")
		}
	} else {
		//panic("No producer created")
	}

	wg.Wait()

	err = p.Flush()
	utils.CheckError(err, "flush")
	err = p.Close()
	utils.CheckError(err, "flush")
}
