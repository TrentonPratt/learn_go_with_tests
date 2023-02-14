package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("this value would be for the actual program, does not affect the test"))
}
