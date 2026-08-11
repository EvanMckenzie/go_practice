package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// set up cli flags
	fileName := flag.String("filename", "problems.csv", "Path to the input file")
	timeLimit := flag.Int("timelimit", 30, "time limit to complete the quiz in seconds")

	// get file name
	flag.Parse()

	// load file
	file, err := os.Open(*fileName) // getting the actual value from the pointer
	if err != nil {
		exit(fmt.Sprintf("Failed to open file [%s]", *fileName))
	}

	// get csv reader
	reader := csv.NewReader(file)

	// read all the data
	records, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file")
	}

	numCorrect := 0

	problems := parseLines(records)

	fmt.Print("Press Enter to start")
	fmt.Scanln()

	// set up timer channel
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// set up answer channel
	answerChannel := make(chan string)

	// ask the question and wait for a response
ProblemLoop:
	for i, p := range problems {

		fmt.Printf("Problem #%d: %s = ", i+1, p.question)

		go func() {
			buf := bufio.NewReader(os.Stdin)
			answer, _ := buf.ReadString('\n')
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime Up!")
			break ProblemLoop
		case answer := <-answerChannel:
			if strings.TrimSpace(answer) != p.answer {
				fmt.Println("Wrong Answer!")
			} else {
				numCorrect++
			}
		}
	}

	fmt.Printf("Performance: %d/%d\n", numCorrect, len(problems))
}

func parseLines(lines [][]string) []problem {
	res := make([]problem, len(lines)) // we do this because we know exactly how big []problem needs to be so we can just declare that upfront instead of letting append handle the resizing

	// res is a slice of problems, parsed from each line of the csv file (assuming that each line is a problem)
	for i, line := range lines {
		res[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return res
}

// by creating a problem struct, we can prevent needing to massage incoming data to be the correct type
// this is pretty slick
type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
