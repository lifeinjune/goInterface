package main

import "fmt"

type bot interface { //declare of interface bot
	getGreeting() string // any types that implement the function getGreeting() and returns string is also member of bot type

}
type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string { // if the receiver that passed to the function is not been used, the name of the variable name on receiver can be deleted
	//e.g (e englishBot) -> (englishBot)
	// custom logic for englishBot
	return "hi there"
}

func (spanishBot) getGreeting() string {
	//custom logic for spanishBot
	return "hola"
}

func printGreeting(b bot) { // any member of type bot can call the printGreeting function becuase it accept the argument bot type variable
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting)
// }

// func printGreeting(sb spanishGreeting) {
// 	fmt.Println(sb.getGreeting)
// }

//interface used to make the code more reusable with different types of argument on same logic to process it
