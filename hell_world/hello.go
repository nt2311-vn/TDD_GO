package main

import "fmt"

const (
	englishPrefixHello = "Hello, "
	spanishPrefixHello = "Hola, "
	frenchPrefixHello  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case "Spanish":
		return spanishPrefixHello + name

	case "French":
		return frenchPrefixHello + name

	default:
		return englishPrefixHello + name

	}
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
