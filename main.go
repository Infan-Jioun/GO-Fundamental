package main

import "fmt"

//* Function
func makeCoffee() {
	fmt.Println("Making Coffee .....")
}
func main() {
	//* Varibales
	// var name string = "Infan"
	// name := "Infan"
	//! name string := "Mizan" Wrong Way
	// var name = "infan"
	//* Group Variable
	// var (
	// 	name string = "Infan"
	// 	age  int    = 22
	// )
	//* Multiple Varibales declaration
	// var x, y int
	// x = 35
	// y = 20
	// const x, y string = "Infan", "Jioun"
	// fmt.Println(name)
	// fmt.Println(age)
	// fmt.Println(x, y)
	// * Zero Variable
	// var age int //* 0
	// fmt.Println((age))
	// var name string //* ""
	// fmt.Println(name)
	// var isAdmin bool
	// fmt.Println(isAdmin) //* false
	// var score float64    //* 0
	// fmt.Println(score)
	// *  Function Call
	// makeCoffee()
	// makeCoffee()
	// makeCoffee()
	//* Anonymus Functions
	// makeCoffee := func() {
	// 	fmt.Printf("Making Coffee")
	// }
	// makeCoffee()
	//* IIFE Functions
	func(coffeeType string) {
		fmt.Printf("Making hot %s....", coffeeType)
	}("Latter")
}
