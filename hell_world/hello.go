package main

import "fmt"

const (
	englishPrefixHello = "Hello, "
	spanishPrefixHello = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishPrefixHello + name
	}
	return englishPrefixHello + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
