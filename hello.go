package main

import "fmt"

var languagePrefixes = map[string]string{
	"English": "Hello, ",
	"Spanish": "Hola, ",
	"French":  "Bonjour, ",
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "" {
		language = "English"
	}
	return languagePrefixes[language] + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
