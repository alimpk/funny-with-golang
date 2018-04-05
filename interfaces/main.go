package main

import "fmt"

type bot interface {
	getGreeting() string
}
type engilshBot struct{}
type fingilshBot struct{}

func main() {
	eb := engilshBot{}
	fb := fingilshBot{}

	printGreeting(eb)
	printGreeting(fb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/**
func printGreeting(eb engilshBot) {
	fmt.Println(eb.getGreeting)
}

func printGreeting(fb engilshBot) {
	fmt.Println(fb.getGreeting)
}
**/

func (engilshBot) getGreeting() string {
	return "Hello!"
}

func (fingilshBot) getGreeting() string {
	return "Dorood!"
}
