package consumers

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"indexer/structs/csvs"
	json2 "indexer/structs/json"
	"indexer/structs/text"
	xml2 "indexer/structs/xml"
	"indexer/utils"
	"os"
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

func ParseClinicalDocument(jobs chan string, p esutil.BulkIndexer,
	wg *sync.WaitGroup, filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case path, hasMore := <-jobs:
			if !hasMore {
				return
			}

			var study xml2.ClinicalStudy
			f, err := os.ReadFile(path)
			utils.CheckError(err, "Open File: "+path)

			// Remove parse errors
			// f = []byte(strings.Map(clean, string(f)))

			err = xml.Unmarshal(f, &study)
			utils.CheckError(err, "Unmarshal")

			if checkFilter(study.IDInfo.NctID, filter, exclude) {
				documentReq, err := json.Marshal(study)
				utils.CheckError(err, "Unable to marshal")

				err = p.Add(
					context.Background(),
					esutil.BulkIndexerItem{
						// Action field configures the operation to perform (index, create, delete, update)
						Action: "index",

						// DocumentID is the optional document ID
						DocumentID: study.IDInfo.NctID,

						// Body is an `io.Reader` with the payload
						Body: bytes.NewReader(documentReq),
					},
				)
			}
		}
	}
}

func ParseTestClinicalDocument(jobs chan string, p esutil.BulkIndexer, wg *sync.WaitGroup,
	filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case path, hasMore := <-jobs:
			if !hasMore {
				return
			}

			var study xml2.ClinicalStudyGuido
			f, err := os.ReadFile(path)
			utils.CheckError(err, "Open File: "+path)

			err = xml.Unmarshal(f, &study)

			utils.CheckError(err, "Unmarshal")

			if checkFilter(study.IDInfo.NctID, filter, exclude) {
				documentReq, err := json.Marshal(study)
				utils.CheckError(err, "Unable to marshal")

				err = p.Add(
					context.Background(),
					esutil.BulkIndexerItem{
						// Action field configures the operation to perform (index, create, delete, update)
						Action: "index",

						// DocumentID is the optional document ID
						DocumentID: study.IDInfo.NctID,

						// Body is an `io.Reader` with the payload
						Body: bytes.NewReader(documentReq),
					},
				)
			}
		}
	}
}

func ParsePubmed(jobs chan string, p esutil.BulkIndexer, wg *sync.WaitGroup,
	filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case path, hasMore := <-jobs:
			if !hasMore {
				return
			}

			var study xml2.PubMedArticle
			f, err := os.ReadFile(path)
			utils.CheckError(err, "Open File: "+path)

			err = xml.Unmarshal(f, &study)

			utils.CheckError(err, "Unmarshal: "+path, false)

			documentReq, err := json.Marshal(study)
			utils.CheckError(err, "Unable to marshal")

			err = p.Add(
				context.Background(),
				esutil.BulkIndexerItem{
					// Action field configures the operation to perform (index, create, delete, update)
					Action: "index",

					// Body is an `io.Reader` with the payload
					Body: bytes.NewReader(documentReq),
				},
			)
		}
	}
}

func ParseMarcoDocument(jobs chan string, p esutil.BulkIndexer, wg *sync.WaitGroup,
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

				documentReq, err := json.Marshal(document)
				utils.CheckError(err, "Unable to marshal")

				err = p.Add(
					context.Background(),
					esutil.BulkIndexerItem{
						// Action field configures the operation to perform (index, create, delete, update)
						Action: "index",

						DocumentID: string(rune(document.Id)),

						// Body is an `io.Reader` with the payload
						Body: bytes.NewReader(documentReq),
					},
				)
			}
		}
	}
}

func ParseBioRedditSubmission(jobs chan csvs.BioRedditSubmissions, p esutil.BulkIndexer, wg *sync.WaitGroup,
	filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case document, hasMore := <-jobs:
			if !hasMore {
				return
			}

			documentReq, err := json.Marshal(document)
			utils.CheckError(err, "Unable to marshal")

			err = p.Add(
				context.Background(),
				esutil.BulkIndexerItem{
					// Action field configures the operation to perform (index, create, delete, update)
					Action: "index",

					// Body is an `io.Reader` with the payload
					Body: bytes.NewReader(documentReq),
				},
			)
		}
	}
}

func ParseBioRedditComment(jobs chan csvs.BioRedditComments, p esutil.BulkIndexer, wg *sync.WaitGroup,
	filter map[string]bool, exclude bool) {

	defer wg.Done()

	for {
		select {
		case document, hasMore := <-jobs:
			if !hasMore {
				return
			}

			documentReq, err := json.Marshal(document)
			utils.CheckError(err, "Unable to marshal")

			err = p.Add(
				context.Background(),
				esutil.BulkIndexerItem{
					// Action field configures the operation to perform (index, create, delete, update)
					Action: "index",

					// Body is an `io.Reader` with the payload
					Body: bytes.NewReader(documentReq),
				},
			)
		}
	}
}

func ParseS2ORC(jobs chan string, p esutil.BulkIndexer, wg *sync.WaitGroup,
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
			documentReq, err := json.Marshal(documentFlatten)

			err = p.Add(
				context.Background(),
				esutil.BulkIndexerItem{
					// Action field configures the operation to perform (index, create, delete, update)
					Action: "index",

					// DocumentID is the optional document ID
					DocumentID: documentFlatten.PaperID,

					// Body is an `io.Reader` with the payload
					Body: bytes.NewReader(documentReq),
				},
			)
		}
	}
}

func ParseS2ORCMetadata(jobs chan string, p esutil.BulkIndexer, wg *sync.WaitGroup,
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

			documentReq, err := json.Marshal(document)
			utils.CheckError(err, "Unable to marshal")

			err = p.Add(
				context.Background(),
				esutil.BulkIndexerItem{
					// Action field configures the operation to perform (index, create, delete, update)
					Action: "index",

					DocumentID: document.PaperID,

					// Body is an `io.Reader` with the payload
					Body: bytes.NewReader(documentReq),
				},
			)
		}
	}
}
