package producers

import (
	"bufio"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"indexer/utils"
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
	glob, err := filepath.Glob(dataPath + "/*/*xml")
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
