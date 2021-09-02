package main

import (
	"bufio"
	"github.com/cheggaaa/pb/v3"
	"log"
	"os"
	"path/filepath"
	"sync"
)



func produceMarco(dataPath string, jobs chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	readFile, err := os.Open(dataPath+"/collection.tsv")

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
	checkError(err, "Unable to close file")
}

func produceClinicalTrials(dataPath string, jobs chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	glob, err := filepath.Glob(dataPath+"/*/*.xml")
	pbn := pb.StartNew(len(glob))
	checkError(err, "Glob")
	for _, path := range glob {
		jobs <- path
		pbn.Add(1)
	}

	close(jobs)
}

func produceTestClinicalTrials(dataPath string, jobs chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	glob, err := filepath.Glob(dataPath+"/*/*.xml")
	pbn := pb.StartNew(len(glob))
	checkError(err, "Glob")
	for _, path := range glob {
		jobs <- path
		pbn.Add(1)
	}

	close(jobs)
}