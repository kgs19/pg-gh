// create a hello world application in go
package main

import "fmt"

func greeting() string {
	return "Hello, World!"
}

// create an other greeting function to say hello to Dim
func greetingDim() string {
	return "Hello, Dim!"
}

// create an other greeting function to say hello to Jason
func greetingJason() string {
	return "Hello, Jason!"
}	


func main() {
	fmt.Println(greeting())
	fmt.Println(greetingDim())
}
