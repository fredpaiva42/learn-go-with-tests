package main

import "fmt"

func Hello() string {
	// Domain
	return "Hello, world"
}

func main() {
	// Side-effect
	fmt.Println(Hello())
}
