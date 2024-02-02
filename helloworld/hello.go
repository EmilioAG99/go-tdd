package main

import (
	"fmt"
	"strings"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}
	language = strings.ToLower(language)

	return greetingsPrefix(language) + name
}
func greetingsPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = spanishHelloPrefix
	case "french":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", "spanish"))
}
