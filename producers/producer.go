package producers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/jszwec/csvutil"
	"indexer/structs/csvs"
	"indexer/utils"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func ProduceMarco(dataPath string, jobs chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	readFile, err := os.Open(dataPath + "/collection.tsv")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	pbn := pb.StartNew(8841823)

	for fileScanner.Scan() {
		row := fileScanner.Text()
		jobs <- row
		pbn.Add(1)
	}

	err = readFile.Close()
	close(jobs)
	utils.CheckError(err, "Unable to close file")
}

func ProduceClinicalTrials(dataPath string, jobs chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	glob, err := filepath.Glob(dataPath + "/*/*.xml")
	pbn := pb.StartNew(len(glob))
	utils.CheckError(err, "Glob")
	for _, path := range glob {
		jobs <- path
		pbn.Add(1)
	}

	close(jobs)
}

func ProduceTestClinicalTrials(dataPath string, jobs chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	glob, err := filepath.Glob(dataPath + "/*/*.xml")
	pbn := pb.StartNew(len(glob))
	utils.CheckError(err, "Glob")
	for _, path := range glob {
		jobs <- path
		pbn.Add(1)
	}

	close(jobs)
}

func GenericProducer(dataPath string, jobs chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	//glob, err := filepath.Glob(globString)
	glob, err := filepath.Glob(dataPath + "/*/*")
	fmt.Println("Found", len(glob), "files...")
	pbn := pb.StartNew(len(glob))
	utils.CheckError(err, "Glob")

	if len(glob) == 0 {
		panic("Producer is empty!")
	}

	for _, path := range glob {
		jobs <- path
		pbn.Add(1)
	}

	close(jobs)
}

func BioredditSubmissionCSVProducer(dataPath string, jobs chan csvs.BioRedditSubmissions, wg *sync.WaitGroup) {
	defer wg.Done()

	glob, err := filepath.Glob(dataPath+"/*.csv")
	fmt.Println("Found", len(glob), "files...")
	pbn := pb.StartNew(len(glob))
	utils.CheckError(err, "Glob")

	if len(glob) == 0 {
		panic("Producer is empty!")
	}

	for _, path := range glob {
		utils.CheckError(err, "Open file")
		f, err := os.Open(path)

		csvReader := csv.NewReader(f)
		dec, err := csvutil.NewDecoder(csvReader)
		utils.CheckError(err, "Decoding file")

		for {
			r := csvs.BioRedditSubmissions{}

			if err := dec.Decode(&r); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			utils.CheckError(err, "Unmarshal")

			fmt.Println(r)

			jobs <- r
			pbn.Add(1)

		}
	}

	close(jobs)
}

func BioredditCommentCSVProducer(dataPath string, jobs chan csvs.BioRedditComments, wg *sync.WaitGroup) {
	defer wg.Done()

	glob, err := filepath.Glob(dataPath+"/*.csv")
	fmt.Println("Found", len(glob), "files...")
	pbn := pb.StartNew(len(glob))
	utils.CheckError(err, "Glob")

	if len(glob) == 0 {
		panic("Producer is empty!")
	}

	for _, path := range glob {
		utils.CheckError(err, "Open file")
		f, err := os.Open(path)

		csvReader := csv.NewReader(f)
		dec, err := csvutil.NewDecoder(csvReader)
		utils.CheckError(err, "Decoding file")

		for {
			r := csvs.BioRedditComments{}

			if err := dec.Decode(&r); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			utils.CheckError(err, "Unmarshal")

			fmt.Println(r)

			jobs <- r
			pbn.Add(1)
		}
	}

	close(jobs)
}