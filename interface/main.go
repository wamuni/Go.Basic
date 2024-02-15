package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

// if we're not using receiver, we can ignore it
func (englishBot) getGreeting() string {
	return "Hello"
}

type spanishBot struct{}

func (bot spanishBot) getGreeting() string {
	return "Hola!"
}

type chineseBot struct{}

func (bot chineseBot) getGreeting() {
	fmt.Println("Hello")
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	cb := chineseBot{}
	printGreeting(eb)
	printGreeting(sb)
	// printGreeting(cb) can't do that since it's not following bot interface
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(bot englishBot) {
// 	fmt.Println(bot.getGreeting())
// }

// func printGreeting(bot spanishBot) {
// 	fmt.Println(bot.getGreeting())
// }
