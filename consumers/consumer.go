package consumers

import (
	"encoding/json"
	"encoding/xml"
	"github.com/olivere/elastic/v7"
	"indexer/structs/csvs"
	json2 "indexer/structs/json"
	"indexer/structs/text"
	xml2 "indexer/structs/xml"
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

			var study xml2.ClinicalStudy
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

			var study xml2.ClinicalStudyGuido
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

			var study xml2.PubMedArticle
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

			var document text.Marco
			parsedRow := strings.Split(row, "\t")
			Id, err := strconv.Atoi(parsedRow[0])

			utils.CheckError(err, "Parse to int")

			if checkFilter(parsedRow[0], filter, exclude) {
				document = text.Marco{
					Id:   Id,
					Text: parsedRow[1],
				}

				rq := elastic.NewBulkIndexRequest().Index(index).Doc(document)
				p.Add(rq)
			}
		}
	}
}

func ParseBioRedditSubmission(jobs chan csvs.BioRedditSubmissions, index string, p *elastic.BulkProcessor, wg *sync.WaitGroup,
	filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case document, hasMore := <-jobs:
			if !hasMore {
				return
			}

			rq := elastic.NewBulkIndexRequest().Index(index).Doc(document)
			p.Add(rq)
		}
	}
}

func ParseBioRedditComment(jobs chan csvs.BioRedditComments, index string, p *elastic.BulkProcessor, wg *sync.WaitGroup,
	filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case document, hasMore := <-jobs:
			if !hasMore {
				return
			}

			rq := elastic.NewBulkIndexRequest().Index(index).Doc(document)
			p.Add(rq)
		}
	}
}

func ParseS2ORC(jobs chan string, index string, p *elastic.BulkProcessor, wg *sync.WaitGroup,
	filter map[string]bool, exclude bool) {
	defer wg.Done()

	for {
		select {
		case line, hasMore := <-jobs:
			if !hasMore {
				return
			}

			var document json2.S2ORC
			err := json.Unmarshal([]byte(line), &document)

			if err != nil {
				panic(err)
			}

			utils.CheckError(err, "Unmarshal S2ORC")
			documentFlatten := document.Flatten()

			rq := elastic.NewBulkIndexRequest().Index(index).Id(documentFlatten.PaperID).Doc(documentFlatten)
			p.Add(rq)
		}
	}
}

func ParseS2ORCMetadata(jobs chan string, index string, p *elastic.BulkProcessor, wg *sync.WaitGroup,
	filter map[string]bool, exclude bool) {
	defer wg.Done()

	for {
		select {
		case line, hasMore := <-jobs:
			if !hasMore {
				return
			}

			var document json2.S2ORCMetadata
			err := json.Unmarshal([]byte(line), &document)

			if err != nil {
				panic(err)
			}

			utils.CheckError(err, "Unmarshal S2ORC")

			rq := elastic.NewBulkUpdateRequest().Index(index).Id(document.PaperID).Doc(document)
			p.Add(rq)
		}
	}
}
