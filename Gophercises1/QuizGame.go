package main

import (
    "encoding/csv"
    "fmt"
	"io"
	"log"
	"os"
	"bufio"
	"strings"
)

func main() {
	fmt.Println("Welcome to Quiz Game!")
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Set the time limit for each question!")

	t := "0"
	
	if scanner.Scan() {
		t = scanner.Text()
	}

	fmt.Println(t)

	filePath := "QuestionAnswers.csv"
	file, _ := os.Open(filePath)

	r := csv.NewReader(bufio.NewReader(file))

	score := 0
	// timer := t
	state := 0

	for score >= 0{

		for {
			record, err := r.Read()
			
			if err == io.EOF{
				state = 1
				break
			}
	
			if err != nil {
				log.Fatal(err)
			}

			question := record[0]
			answer   := record[1]

			fmt.Println(question)

			userInput := ""
			
			if scanner.Scan(){
				userInput = scanner.Text()
			}

			fmt.Println(userInput)

			if strings.Compare(userInput, answer) != 0{
				fmt.Println("Wrong answer! -1")
				score = score-1
			} else{
				fmt.Println("Correct Answer! +1")
				score = score+1
			}

			fmt.Println("Score: ", score)

			if score < 0 {
				state = 2
				break
			}
		}

		if state == 1 {
			fmt.Println("End of Quiz!")
			break
		}else if state == 2 {
			fmt.Println("score became negative, game over")
			break
		}
	}

}