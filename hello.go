package main

import "fmt"

const (
	portuguese = "Portuguese"
	spanish    = "Spanish"
	french     = "French"

	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	portugueseHelloPrefix = "Olá, "
	frenchHelloPrefix     = "Bonjour, "
)

func Hello(name string, language string) string {
	// Domain
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	// Side-effect
	fmt.Println(Hello("Fred", ""))
}
