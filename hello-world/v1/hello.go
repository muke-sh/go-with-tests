package main

import "fmt"

const englishPrefix string = "Hello, "
const spanishPrefix string = "Hola, "
const frenchPrefix string = "Bonjour, "

const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
