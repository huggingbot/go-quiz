package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestReadCsvFile(t *testing.T) {
	str := "3+3,6\n5+7,12\n13+14,27\n"
	questionAnswers := readCsvFile(strings.NewReader(str))

	var qA [3]QuestionAnswer
	qA[0].question = "3+3"
	qA[0].answer = "6"
	qA[1].question = "5+7"
	qA[1].answer = "12"
	qA[2].question = "13+14"
	qA[2].answer = "27"

	assert.Equal(t, qA[0], questionAnswers[0])
	assert.Equal(t, qA[1], questionAnswers[1])
	assert.Equal(t, qA[2], questionAnswers[2])
}

func TestAskSingleQuestion(t *testing.T) {
	answerCh := make(chan string)
	reader := bufio.NewReader(os.Stdin)
	timer := time.NewTimer(time.Duration(3) * time.Second)

	qA := QuestionAnswer{question: "4+4", answer: "8"}

	getInput = func(reader *bufio.Reader, answerCh chan<- string) {
		answerCh<-"8"
	}

	score := askSingleQuestion(qA, answerCh, reader, timer)

	assert.Equal(t, score, 1)
}
