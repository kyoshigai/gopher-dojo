package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func interactWithUser(fileName string, ch chan<- bool) {
	problems, _ := getProblems(fileName)
	rand.Seed(time.Now().Unix())

	sc := bufio.NewScanner(os.Stdin)
	for len(problems) > 0 {
		// fmt.Printf("problems: %v\n", problems)
		randNum := rand.Intn(len(problems))
		problem := problems[randNum]
		problems = append(problems[:randNum], problems[randNum+1:]...)
		fmt.Printf("Type %q: ", problem)
		sc.Scan()
		answer := sc.Text()
		ch <- answer == problem
	}
	close(ch)
}

func getProblems(path string) ([]string, error) {
	var problems []string

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	problemScanner := bufio.NewScanner(f)
	for problemScanner.Scan() {
		problems = append(problems, problemScanner.Text())
	}
	return problems, nil
}

func main() {
	answerdCh := make(chan bool)
	go interactWithUser("problems.txt", answerdCh)
	score := 0

L:
	for {
		select {
		case correct, ok := <-answerdCh:
			if !ok {
				break L
			}
			if correct {
				score++
			}
		case <-time.After(5 * time.Second):
			fmt.Println("\nTime Up")
			break L
		}
	}
	fmt.Println("---------------------------")
	fmt.Printf("Your score is %d\n", score)
}
