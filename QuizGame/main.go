package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func quizMaster(fileName string, timeLimit int) {
	var correctAnswer int
	var incorrectAnswer int
	var totalQuestions int

	fmt.Println("The time limit is", timeLimit, "seconds")
	var input string
	fmt.Print("Are you ready?(y or n): ")
	fmt.Scanln(&input)
	if input == "n" {
		os.Exit(1)
	} else {

	}

	quizFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("An error with opening the file:", err)
		return
	}
	defer quizFile.Close()

	reader := csv.NewReader(quizFile)
	allData, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse the provided CSV file:", err)
		return
	}

	answers := make(chan bool)

	go func() {
		for _, record := range allData {
			question := record[0]
			answer := record[1]

			fmt.Print("Question: ", question, ": ")
			var input string
			fmt.Scanln(&input)
			if input == answer {
				answers <- true
			} else {
				answers <- false
			}
		}
		close(answers)
	}()

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for {
		select {
		case <-timer.C:
			fmt.Println("\nTime's up! Try Again!")
			return
		case answer, ok := <-answers:
			if !ok {
				goto result
			}
			totalQuestions++
			if answer {
				correctAnswer++
			} else {
				incorrectAnswer++
			}
		}
	}

result:
	quizResult := float64(correctAnswer) / float64(totalQuestions) * 100
	fmt.Println("You got", correctAnswer, "correct answers.")
	fmt.Println("You got", incorrectAnswer, "incorrect answers.")
	fmt.Printf("You got a %.2f%% on the quiz.\n", quizResult)
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer' (default 'problems.csv')")
	limit := flag.Int("limit", 30, "The time limit for the quiz in seconds (default 30)")
	flag.Parse()

	quizMaster(*csvFile, *limit)
}
