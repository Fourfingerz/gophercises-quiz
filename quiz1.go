package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

var fileDir string

type problem struct {
	question string
	answer   string
}

func main() {
	flag.StringVar(&fileDir, "f", "problems.csv", "quiz dir")
	flag.Parse()
	quiz := convertCsvToQuiz(fileDir)
	questionsCorrect, quizLength := giveQuiz(quiz)
	fmt.Println(`You got`, questionsCorrect, `questions out of`, quizLength, `correct.`)
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

func giveQuiz(quiz []problem) (correctAnswers, quizLength int) {
	quizLength = len(quiz)

	// timer := time.NewTimer(3 * time.Second)
		for k, v := range quiz {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Question", k, ":")
			fmt.Println(v.question)
			fmt.Println("Type your answer below: ")
			text, _ := reader.ReadString('\n')
			fmt.Print("You just answered: ")
			fmt.Println(text)
			fmt.Print("The actual answer is: ")
			fmt.Println(v.answer)
			fmt.Println("")
			if (strings.TrimRight(text, "\n") == v.answer) {
				correctAnswers++
			}
		}
	return
}