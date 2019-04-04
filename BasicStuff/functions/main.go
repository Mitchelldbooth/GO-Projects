package main

import "fmt"

func greeting(name string) string {
	return "hello" + name
}

func main() {
	fmt.Println(greeting(" Mitchell"))
}
