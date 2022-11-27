# Exercise: Quiz Game

A program that will read in a quiz provided via a CSV file and will then give the quiz to a user keeping track of how many questions they get right out of the total questions. Regardless of whether the answer is correct or wrong the next question will be asked immediately afterwards.

To run the program:

```
go run main.go
```

The CSV file defaults to `problems.csv`, but the filename can be customized via a flag.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

Quizzes are relatively short (< 100 questions) and have single word/number answers.

The default time limit is 30 seconds, and can be customized via a flag.

The quiz will stop as soon as the time limit is exceeded and print out the score.

The timer starts when the program starts, and the questions will be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question will be asked.

At the end of the quiz the program will output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.

Input answers will be trimmed, lowercased and cleared of whitespaces.

There is also a flag to shuffle the quiz order each time it is run.
