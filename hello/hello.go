package main

import "fmt"

const englishHelloPrefix = "Hello, %s"
const spanishHelloPrefix = "Hola, %s"
const frenchHelloPrefix = "Bonjour, %s"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf(greetingPrefix(language), name)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("Cristy", ""))
}
