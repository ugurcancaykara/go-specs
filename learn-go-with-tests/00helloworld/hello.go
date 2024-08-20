package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Zion"))

}

func Hello(name string) string {
	if name != "" {

		return englishHelloPrefix + name
	}
	return "Hello, World"
}
