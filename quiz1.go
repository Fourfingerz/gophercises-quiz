package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

var fileDir string

type problem struct {
	question string
	answer   string
}

func main() {
	timeLimit := flag.Int("limit", 30, "time limit for the quiz in seconds")
	flag.StringVar(&fileDir, "f", "problems.csv", "quiz dir")
	flag.Parse()

	quiz := convertCsvToQuiz(fileDir)


	giveQuiz(quiz, *timeLimit)
}

func convertCsvToQuiz(filename string) []problem {
	var quiz []problem

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		data := problem{
			question: line[0],
			answer:   line[1],
		}
		quiz = append(quiz, data)
	}

	return quiz
}

func giveQuiz(quiz []problem, timeLimit int) () {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	quizLength := len(quiz)
	correctAnswers := 0

	problemloop:
	for i, p := range quiz {
		fmt.Printf("Question #%d: %s = ", i+1, p.question)
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerChan:
			if answer == p.answer {
				correctAnswers++
			}
		}
	}
	fmt.Println(`You got`, correctAnswers, `questions out of`, quizLength, `correct.`)
}