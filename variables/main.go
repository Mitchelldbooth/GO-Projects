package main

import "fmt"

func main() {
	var name = "Mitchell"
	var age int32 = 19
	var iscool = false
	iscool = true

	fmt.Println(name, age, iscool)

	fmt.Printf("%T\n", iscool)

	hometown := "Vancouver"

	fmt.Println(hometown)
	fmt.Printf("%T\n", hometown)

}
