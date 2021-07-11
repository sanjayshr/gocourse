package main

import "fmt"

func User(name string) string {
	return name
}

func main() {
	fmt.Printf(User("test"))
}
