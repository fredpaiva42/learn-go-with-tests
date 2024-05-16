package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	// Domain
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	// Side-effect
	fmt.Println(Hello(""))
}
