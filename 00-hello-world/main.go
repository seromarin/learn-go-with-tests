package main

import "fmt"

const (
	english            = "English"
	englishHelloPrefix = "Hello, "
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "
	french             = "French"
	frenchHelloPrefix  = "Bonjour, "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

func main() {
	fmt.Println(Hello("World", "English"))
}
