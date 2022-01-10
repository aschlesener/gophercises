package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := flag.String("filename", "problems.csv", "string for name of CSV file to be used")
	flag.Parse()

	csvReader, err := getCsvReader(*fileName)
	if err != nil {
		fmt.Println("Could not find specified filename")
		return
	}

	validAnswersCnt := 0
	totalQuestionsCnt := 0

	// Read questions from file until we reach the end
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if len(line) < 2 {
			fmt.Println("Error processing CSV line")
			continue
		}

		totalQuestionsCnt++

		// Display question
		fmt.Printf(line[0] + ": ")

		// Get answer
		var answer string
		fmt.Scanf("%s", &answer)

		if answer == line[1] {
			validAnswersCnt++
		}
	}

	fmt.Printf("There were %v total questions. You answered %v correctly!\n", totalQuestionsCnt, validAnswersCnt)
}

func getCsvReader(fileName string) (*csv.Reader, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return csv.NewReader(f), err
}
