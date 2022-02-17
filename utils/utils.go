package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func CreateFilterMap(filterFp string) map[string]bool {
	if len(strings.TrimSpace(filterFp)) == 0 {
		return nil
	}

	var filter map[string]bool
	filter = make(map[string]bool)

	//readFile, err := os.Open(dataPath+"/med-msmarco-train.txt")
	readFile, err := os.Open(filterFp)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		filter[strings.TrimSpace(fileScanner.Text())] = true
		CheckError(err, "Cannot convert to integer")
	}

	return filter
}

func CheckError(err error, where string, raise ...bool) {
	if err != nil && raise[0] == true {
		println(where)
		panic(err)
	}
}
