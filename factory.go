package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"indexer/consumers"
	"indexer/producers"
	"sync"
)

func ProducerConsumerFactory(DocType string) (func(string, chan string, *sync.WaitGroup), func(chan string, string, *elastic.BulkProcessor, *sync.WaitGroup, map[string]bool, bool)) {
	fmt.Println(DocType)
	if DocType == "marco" {
		return producers.ProduceMarco, consumers.ParseMarcoDocument // ("collection.tsv", "med-msmarco-train.txt", jobs, &wg) )
	} else if DocType == "clinical-trials" {
		return producers.ProduceClinicalTrials, consumers.ParseClinicalDocument //("/home/vin/Projects/CDS_2021/index/raw_data/*/*.xml", jobs, &wg)
	} else if DocType == "test-trials" {
		return producers.ProduceTestClinicalTrials, consumers.ParseTestClinicalDocument //("/home/vin/Projects/CDS_2021/index/raw_data/*/*.xml", jobs, &wg)
	} else if DocType == "pubmed" {
		return producers.GenericProducer, consumers.ParsePubmed
	} else if DocType == "s2orc" {
		return producers.GenericLineReaderProducer, consumers.ParseS2ORC
	} else if DocType == "s2orc-meta" {
		return producers.GenericLineReaderProducer, consumers.ParseS2ORCMetadata
	} else {
		return nil, nil
	}
}
