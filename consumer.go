package main

import (
	"encoding/xml"
	"github.com/olivere/elastic/v7"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

func checkError(err error, where string) {
	if err != nil {
		println(where)
		panic(err)
	}
}

func clean (r rune) rune {
	if unicode.IsPrint(r) {
		return r
	}
	return -1
}

func checkFilter(docid string, filter map[string]bool, exclude bool) bool {
	// Return true if it passes the filter
	if filter != nil {
		_, exists := filter[docid]

		if exclude {
			return !exists // It exists, so we exclude it
		}

		return exists // it exists, so we include it
	}
	return true // If filter doesn't exist, don't worry about excluding
}

func parseClinicalDocument(jobs chan string, index string, p *elastic.BulkProcessor,
	wg *sync.WaitGroup, filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case path, hasMore := <- jobs:
			if !hasMore {
				return
			}

			var study ClinicalStudy
			f, err := ioutil.ReadFile(path)
			checkError(err, "Open File: "+path)

			// Remove parse errors
			// f = []byte(strings.Map(clean, string(f)))

			err = xml.Unmarshal(f, &study)
			checkError(err, "Unmarshal")

			if checkFilter(study.IDInfo.NctID, filter, exclude) {
				rq := elastic.NewBulkIndexRequest().Index(index).Doc(study)

				p.Add(rq)
			}
		}
	}
}

func parseTestClinicalDocument(jobs chan string, index string, p *elastic.BulkProcessor, wg *sync.WaitGroup,
	filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case path, hasMore := <-jobs:
			if !hasMore {
				return
			}

			var study TestClinicalStudy
			f, err := ioutil.ReadFile(path)
			checkError(err, "Open File: "+path)

			err = xml.Unmarshal(f, &study)

			checkError(err, "Unmarshal")

			if checkFilter(study.IDInfo.NctID, filter, exclude) {
				rq := elastic.NewBulkIndexRequest().Index(index).Doc(study)

				p.Add(rq)
			}
		}
	}
}

func parseMarcoDocument(jobs chan string, index string, p *elastic.BulkProcessor, wg *sync.WaitGroup,
	filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case row, hasMore := <- jobs:
			if !hasMore {
				return
			}

			var document Marco
			parsedRow := strings.Split(row, "\t")
			Id, err := strconv.Atoi(parsedRow[0])

			checkError(err, "Parse to int")

			if checkFilter(parsedRow[0], filter, exclude) {
				document = Marco{
					Id: Id,
					Text: parsedRow[1],
				}

				rq := elastic.NewBulkIndexRequest().Index(index).Doc(document)
				p.Add(rq)
			}
		}
	}
}
