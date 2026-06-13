package main

import "fmt"

func main() {
	//* Varibales
	// var name string = "Infan"
	// name := "Infan"
	//! name string := "Mizan" Wrong Way
	// var name = "infan"
	//* Group Variable
	var (
		name string = "Infan"
		age  int    = 22
	)
	//* Multiple Varibales declaration
	// var x, y int
	// x = 35
	// y = 20
	const x, y string = "Infan", "Jioun"
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(x, y)
}
