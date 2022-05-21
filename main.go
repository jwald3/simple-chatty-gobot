package main

import (
	"fmt"
	"strconv"
)

// Generate basic greeting using provided bot name and year created as string inputs
func greet(name, year string) {
	fmt.Println("Hello! My name is " + name + ".")
	fmt.Println("I was created in " + year + ".")
}

// Ask and acknowledge user name using standard input
func showName() {
	var name string
	fmt.Println("Please, remind me your name.")
	fmt.Scan(&name) // "&" ensures that user input is stored in the memory address designated as "name"
	fmt.Println("What a great name you have, " + name + "!")
}

// Guesses user age based on algorithm
func guessAge() {
	var rem3, rem5, rem7, age int

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter the remainders of dividing your age by 3, 5, and 7 (separated by spaces please)")
	fmt.Scan(&rem3, &rem5, &rem7)

	// algorithm guesses age based on remainders and constants
	age = (rem3*70 + rem5*21 + rem7*15) % 105
	fmt.Println("Your age  is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}

// Count from 0 to user-provided number
func count() {
	var n int

	fmt.Println("Now I will prove to you I can count to any number you want.")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ { // basic for-loop using java-style syntax of "init"; "condition"; "iteration"
		fmt.Printf("%d!\n", i)
	}
}

// Ask user a question and keep asking until the correct answer is provided
func startQuiz() {
	fmt.Println("Let's test your programming knowledge.")
	fmt.Println("Why do we use methods?")
	fmt.Println("1. To repeat a statement multiple times.")
	fmt.Println("2. To decompose a program into several small subroutines.")
	fmt.Println("3. To determine the execution time of a program.")
	fmt.Println("4. To interrupt the execution of a program.")

	var answeredCorrectly bool = false
	var userAnswer string
	var correctAnswer string = "2"

	fmt.Scan(&userAnswer)

	// if user answers correctly the first time, it will register as "true" and the quiz will be done
	answeredCorrectly = userAnswer == correctAnswer

	// if initial answer is not correct, the user will be asked to provide another answer until they get it right
	// Go does not have while loop, so this acts as an "until condition is met" loop
	for !answeredCorrectly {
		fmt.Println("Please, try again.")
		fmt.Scan(&userAnswer)
		// if right, loop is ended by "true" otherwise loop continues
		answeredCorrectly = userAnswer == correctAnswer
	}

	fmt.Println("Correct answer, excellent.")
}

// End program with a friendly message
func sayGoodbye() {
	fmt.Println("Congratulations, have a nice day!")
}

func main() {
	greet("Aid", "2022")
	showName()
	guessAge()
	count()
	startQuiz()
	sayGoodbye()
}
