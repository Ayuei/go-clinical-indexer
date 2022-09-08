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

func ProduceMarco(dataPath string, jobs chan string, wg *sync.WaitGroup, accurate bool) {
	defer wg.Done()

	var counter int64
	var readFile *os.File

	if accurate {
		var err error
		readFile, err = os.Open(dataPath)
		utils.CheckError(err, "Unable to read file")

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)
		counter, err = utils.LineCounter(readFile)

		utils.CheckError(err, "Unable to read file")
	} else {
		counter = 1e9
	}

	pbn := pb.Start64(counter)

	readFile, _ = os.Open(dataPath)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		row := fileScanner.Text()
		jobs <- row
		pbn.Add(1)
	}

	close(jobs)
}

func ProduceClinicalTrials(dataPath string, jobs chan string, wg *sync.WaitGroup, accurate bool) {
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

func ProduceTestClinicalTrials(dataPath string, jobs chan string, wg *sync.WaitGroup, accurate bool) {
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

func GenericProducer(dataPath string, jobs chan string, wg *sync.WaitGroup, accurate bool) {
	defer wg.Done()

	//glob, err := filepath.Glob(globString)
	glob, err := filepath.Glob(dataPath + "*")
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

func GenericLineReaderProducer(dataPath string, jobs chan string, wg *sync.WaitGroup, accurate bool) {
	defer wg.Done()
	fmt.Println("Started producer")

	glob, err := filepath.Glob(dataPath + "/*")
	fmt.Println("Found", len(glob), "files...")
	utils.CheckError(err, "Glob")
	length := int64(31050107)
	bar := pb.Start64(length)
	//skipto := 20000000

	if len(glob) == 0 {
		panic("Producer is empty!")
	}

	counter := 0

	for _, path := range glob {
		f, err := os.Open(path)
		utils.CheckError(err, "Reading File, Generic Line Producer")

		if accurate {
			length, err = utils.LineCounter(f)
			utils.CheckError(err, "Open file")
			_, err = f.Seek(0, 0)
			utils.CheckError(err, "File Seek")
		}

		r := bufio.NewReader(f)
		s, e := utils.Readln(r)

		for e == nil {
			//if counter >= skipto {
			//}
			jobs <- s
			s, e = utils.Readln(r)
			bar.Add(1)
			counter += 1
		}

		//fmt.Printf("Finished %d out of %d.", counter, len(glob))
	}
	bar.Finish()

	close(jobs)
}

func BioredditSubmissionCSVProducer(dataPath string, jobs chan csvs.BioRedditSubmissions, wg *sync.WaitGroup, accurate bool) {
	defer wg.Done()

	glob, err := filepath.Glob(dataPath + "/*.csv*")
	fmt.Println("Found", len(glob), "files...")
	utils.CheckError(err, "Glob")

	if len(glob) == 0 {
		panic("Producer is empty!")
	}

	for _, path := range glob {
		f, err := os.Open(path)
		utils.CheckError(err, "Open file")
		length, err := utils.LineCounter(f)
		utils.CheckError(err, "Open file")
		_, err = f.Seek(0, 0)
		utils.CheckError(err, "File Seek")

		pbn := pb.Start64(length)

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

			jobs <- r
			pbn.Add(1)
		}
		pbn.Finish()
	}

	close(jobs)
}

func BioredditCommentCSVProducer(dataPath string, jobs chan csvs.BioRedditComments, wg *sync.WaitGroup, accurate bool) {
	defer wg.Done()

	glob, err := filepath.Glob(dataPath + "/*.csv")
	fmt.Println("Found", len(glob), "files...")
	pbn := pb.StartNew(len(glob))
	utils.CheckError(err, "Glob")

	if len(glob) == 0 {
		panic("Producer is empty!")
	}

	for _, path := range glob {
		utils.CheckError(err, "Open file")
		f, err := os.Open(path)

		//fs, err := f.Stat()

		//if fs.Size() < 10 {
		//	pbn.Add(1)
		//	continue
		//}

		utils.CheckError(err, "File stats")

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
			jobs <- r
		}
		pbn.Add(1)
	}

	pbn.Finish()
	close(jobs)
}
