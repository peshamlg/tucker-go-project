package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var answer int = rand.Intn(1000) + 1
	var guess int
	var guessCount int = 1
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Guess the number between 1 and 1000:")
		_, err := fmt.Scanf("%d\n", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a integer.")
			reader.ReadString('\n')
			continue
		} else if guess < 1 || guess > 1000 {
			fmt.Println("Invalid input. Please enter a number between 1 and 1000.")
			continue
		} else if guess < answer {
			fmt.Println("Your guess is lower than the answer.")
		} else if guess > answer {
			fmt.Println("Your guess is higher than the answer.")
		} else {
			fmt.Println("You got it!")
			break
		}
		guessCount++
	}
	fmt.Printf("You guessed the number in %d tries.", guessCount)
}
