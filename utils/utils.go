package utils

import (
	"bufio"
	"bytes"
	"io"
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

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

// LineCounter efficently calculate the number of lines in a file
func LineCounter(r io.Reader) (int64, error) {
	buf := make([]byte, 32*1024)
	count := int64(0)
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += int64(bytes.Count(buf[:c], lineSep))

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
