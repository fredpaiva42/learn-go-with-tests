package main

import "fmt"

func Hello(name string) string {
	// Domain
	return "Hello, " + name
}

func main() {
	// Side-effect
	fmt.Println(Hello("world"))
}
