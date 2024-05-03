package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func quizMaster(fileName string, timeLimit int) {
	var correctAnswer int
	var incorrectAnswer int
	var totalQuestions int

	fmt.Println("The time limit is", timeLimit)
	quizFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("An error with opening the file ::", err)
	}
	defer quizFile.Close()

	reader := csv.NewReader(quizFile)
	allData, _ := reader.ReadAll()

	for _, record := range allData {
		question := record[0]
		answer := record[1]
		var input string

		fmt.Print("Question: ", question, ": ")
		fmt.Scan(&input)
		if input == answer {
			fmt.Println("The answer is correct")
			correctAnswer++
			totalQuestions++
		} else {
			fmt.Println("The answer is wrong")
			incorrectAnswer++
			totalQuestions++
		}
	}

	quizResult := float64(correctAnswer) / float64(totalQuestions) * 100
	fmt.Println("You got", correctAnswer, "correct answers.")
	fmt.Println("You got", incorrectAnswer, "incorrect answers.")
	fmt.Printf("You got a %.2f%% on the quiz.\n", quizResult)
}

func main() {
	csv := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer' (default 'problems.csv')")
	limit := flag.Int("limit", 30, "The time limit for the quiz in seconds (default 30)")
	flag.Parse()
	quizMaster(*csv, *limit)
}
