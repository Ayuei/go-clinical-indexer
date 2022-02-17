package consumers

import (
	"encoding/xml"
	"github.com/olivere/elastic/v7"
	"indexer/structs"
	"indexer/utils"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

func clean(r rune) rune {
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

func ParseClinicalDocument(jobs chan string, index string, p *elastic.BulkProcessor,
	wg *sync.WaitGroup, filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case path, hasMore := <-jobs:
			if !hasMore {
				return
			}

			var study structs.ClinicalStudy
			f, err := ioutil.ReadFile(path)
			utils.CheckError(err, "Open File: "+path)

			// Remove parse errors
			// f = []byte(strings.Map(clean, string(f)))

			err = xml.Unmarshal(f, &study)
			utils.CheckError(err, "Unmarshal")

			if checkFilter(study.IDInfo.NctID, filter, exclude) {
				rq := elastic.NewBulkIndexRequest().Index(index).Doc(study)

				p.Add(rq)
			}
		}
	}
}

func ParseTestClinicalDocument(jobs chan string, index string, p *elastic.BulkProcessor, wg *sync.WaitGroup,
	filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case path, hasMore := <-jobs:
			if !hasMore {
				return
			}

			var study structs.TestClinicalStudy
			f, err := ioutil.ReadFile(path)
			utils.CheckError(err, "Open File: "+path)

			err = xml.Unmarshal(f, &study)

			utils.CheckError(err, "Unmarshal")

			if checkFilter(study.IDInfo.NctID, filter, exclude) {
				rq := elastic.NewBulkIndexRequest().Index(index).Doc(study)

				p.Add(rq)
			}
		}
	}
}

func ParsePubmed(jobs chan string, index string, p *elastic.BulkProcessor, wg *sync.WaitGroup,
	filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case path, hasMore := <-jobs:
			if !hasMore {
				return
			}

			var study structs.PubMedArticle
			f, err := ioutil.ReadFile(path)
			utils.CheckError(err, "Open File: "+path)

			err = xml.Unmarshal(f, &study)

			utils.CheckError(err, "Unmarshal: "+path, false)

			rq := elastic.NewBulkIndexRequest().Index(index).Doc(study)
			p.Add(rq)
		}
	}
}

func ParseMarcoDocument(jobs chan string, index string, p *elastic.BulkProcessor, wg *sync.WaitGroup,
	filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case row, hasMore := <-jobs:
			if !hasMore {
				return
			}

			var document structs.Marco
			parsedRow := strings.Split(row, "\t")
			Id, err := strconv.Atoi(parsedRow[0])

			utils.CheckError(err, "Parse to int")

			if checkFilter(parsedRow[0], filter, exclude) {
				document = structs.Marco{
					Id:   Id,
					Text: parsedRow[1],
				}

				rq := elastic.NewBulkIndexRequest().Index(index).Doc(document)
				p.Add(rq)
			}
		}
	}
}
