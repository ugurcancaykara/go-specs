package main

import "fmt"

func main() {
	fmt.Println(Hello("Zion"))

}

func Hello(name string) string {
	return "Hello, " + name
}
