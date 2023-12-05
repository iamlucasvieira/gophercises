package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// question is a struct that holds a question and an answer
type question struct {
	question string
	answer   string
}

// quiz is a struct that holds a slice of questions
type quiz struct {
	questions []question
}

// game is a struct that holds a quiz and a score
type game struct {
	quiz  quiz
	score int
}

// play starts the quiz game
func (g *game) play() {
	// Iterate over the quiz questions
	for _, q := range g.quiz.questions {
		// Print the question
		fmt.Printf("%s = ", q.question)

		// Read the answer from the user
		var answer string
		_, err := fmt.Scanf("%s\n", &answer)
		if err != nil {
			answer = ""
		}

		// Check if the answer is correct
		if answer == q.answer {
			g.score++
		}
	}

	// Print the score
	fmt.Printf("You scored %d out of %d.\n", g.score, len(g.quiz.questions))
}

// readQuiz reads the csv quiz file and returns a quiz struct and an error
func readQuiz(fileName string) (quiz, error) {
	// Open the file
	f, err := os.Open(fileName)
	if err != nil {
		return quiz{}, fmt.Errorf("error opening file: %v", err)
	}

	// Defer closing the file
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	// Read the csv file
	csvReader := csv.NewReader(f)

	q := quiz{}

	// Iterate over the csv file and create a question struct for each line
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// Add the question to the quiz
		q.questions = append(q.questions, question{rec[0], rec[1]})
	}

	return q, nil
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	fmt.Println("Opening file:", *csvFilename)

	q, err := readQuiz(*csvFilename)

	if err != nil {
		fmt.Println(err)
	}

	g := game{q, 0}
	g.play()

}
