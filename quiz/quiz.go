package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	fileName := flag.String("filename", "problems.csv", "string for name of CSV file to be used")
	timeLimit := flag.Int("timer", 30, "time limit (in seconds)")
	flag.Parse()

	csvReader, err := getCsvReader(*fileName)
	if err != nil {
		fmt.Println("Could not find specified filename")
		return
	}

	validAnswersCnt := 0
	totalQuestionsCnt := 0

	// Let user know time limit before starting
	fmt.Printf("Begin quiz by pressing Enter. You will have %v seconds to complete the quiz once you hit Enter. \n", *timeLimit)
	var userInput string
	fmt.Scanf("%s", &userInput)

	// Start goroutine to stop quiz after time limit has passed
	t := time.NewTimer(time.Second * time.Duration(*timeLimit))
	go func() {
		<-t.C
		fmt.Printf("\nTime's up!\n")
		printFinish(totalQuestionsCnt, validAnswersCnt)
		os.Exit(0)
	}()

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

	printFinish(totalQuestionsCnt, validAnswersCnt)
}

func getCsvReader(fileName string) (*csv.Reader, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	return csv.NewReader(f), err
}

func printFinish(totalQuestionsCnt int, validAnswersCnt int) {
	fmt.Printf("There were %v total questions asked. You answered %v correctly!\n", totalQuestionsCnt, validAnswersCnt)
}
