package main

import(
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Guess the number!")
	fmt.Println("Please enter a number: ...")

	// Create a reader that will be reading the STDIN
	// where the centinel is the endline '\n'
	reader := bufio.NewReader(os.Stdin)
	guess, _ := reader.ReadString('\n')

	// Create a rand number with a non-fixed seed, so we
	// will be sure that the rand number changes everytime 
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	//Our secret number will be modulo 100.
	secret := rnd.Int31n(100)
	fmt.Printf("The secret number is: %v\n", secret)
	fmt.Printf("You guessed: %v\n", guess)

	//For converting our string guess number to int, we must
	//replace the endline to a non-character, so the Atoi
	//method could run withour errors
	guess = strings.Replace(guess, "\n", "", -1)
	guessNumber, err := strconv.Atoi(guess)
	if err != nil {
		log.Printf("Error converting...\nError: %v\n", err)
	}
	//Switch condition true, its quite useful for repeative
	//if/else condition. In Go! there's no differences between
	//if-else and switch, switch its not more efficient.
	switch {
	case int32(guessNumber) < secret:
		fmt.Println("Too small!")
	case int32(guessNumber) > secret:
		fmt.Println("Too big!")
	default:
		fmt.Println("You win!")
	}
}
