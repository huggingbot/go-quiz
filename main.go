package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type QuestionAnswer struct {
	question string
	answer string
}

func main() {
	filename, timeLimit, shuffle := readArgs()
	questionAnswers := readCsvFile(filename)
	if shuffle {
		questionAnswers = shuffleQuestions(questionAnswers)
	}
	score := askQuestions(questionAnswers, timeLimit)
	fmt.Println("Score:", score, "/", len(questionAnswers))
}

func readArgs() (string, int, bool) {
	filePtr := flag.String("filename", "problems.csv", "CSV file containing quiz questions")
	timeLimitPtr := flag.Int("timeLimit", 30, "Time limit for each question")
	shufflePtr := flag.Bool("shuffle", false, "Should questions be shuffled?")
	flag.Parse()

	return *filePtr, *timeLimitPtr, *shufflePtr
}

func readCsvFile(filename string) []QuestionAnswer {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file for file " + filename, "\n", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Error parsing file as CSV for file " + filename, "\n", err)
	}

	questionAnswers := []QuestionAnswer{}
	for _, record := range records {
		question, answer := record[0], record[1]
		questionAnswer := QuestionAnswer{question: question, answer: answer}
		questionAnswers = append(questionAnswers, questionAnswer)
	}
	return questionAnswers
}

func shuffleQuestions(questionAnswers []QuestionAnswer) []QuestionAnswer {
	source := rand.NewSource(time.Now().Unix())
	randGen := rand.New(source)
	length := len(questionAnswers)

	for i := 0; i < length; i++ {
		randInt := randGen.Intn(length - 1)
		questionAnswers[i], questionAnswers[randInt] = questionAnswers[randInt], questionAnswers[i]
	}
	return questionAnswers
}

func getInput(reader *bufio.Reader, answerCh chan<- string) {
	result, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input string", "\n", err)
	}
	answerCh<-result
}

func askQuestions(questionAnswers []QuestionAnswer, timeLimit int) int {
	score := 0
	answerCh := make(chan string)
	reader := bufio.NewReader(os.Stdin)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for _, questionAnswer := range questionAnswers {
		fmt.Println(questionAnswer.question)

		go getInput(reader, answerCh)

		select {
		case <-timer.C:
			return score
		case answer := <-answerCh:
			if strings.Compare(strings.Trim(strings.ToLower(answer), "\n"), questionAnswer.answer) == 0 {
				score++
			}
		}
	}
	return score
}
