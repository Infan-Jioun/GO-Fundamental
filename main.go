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
	// func(coffeeType string) {
	// 	fmt.Printf("Making hot %s....", coffeeType)
	// }("Latter")
	//* Variable Scope
	// sugar := 2
	// makeCoffee := func() {
	// 	coffee := "Cappaccino"
	// 	sugar := 3 // Modify
	// 	fmt.Printf("Making %s with %d spoon of sugar \n", coffee, sugar)
	// 	fmt.Println("Value of inner sugar", sugar)
	// }
	// makeCoffee()
	// fmt.Println("Value of outer sugar", sugar)
	//* if-else & switch/
	// score := 20
	// if score >= 80 {
	// 	fmt.Println("You Got Gold Medal...")
	// } else if score >= 70 {
	// 	fmt.Println("You Got Silver Medal...")
	// } else {
	// 	fmt.Println("You got participation certificate...")
	// }
	//* Scoped if else
	// if score := 60; score >= 80 {
	// 	prizeMoney := 1000
	// 	fmt.Println("You have won Gold Medal and prize money is", prizeMoney)
	// 	fmt.Println("You Got Gold Medal an your score is", score)
	// } else if score >= 70 {

	// 	fmt.Println("You Got Silver Medal and your score is", score)
	// } else {
	// 	fmt.Println("You got participation certificate and your score is", score)
	// }

	// if err := saveToDb(); err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// *Switch
	// day := "fri"
	// switch day {
	// case "sat":
	// 	println("Yes, work day!")
	// case "sun":
	// 	println("Yes, Second work day")
	// case "fri":
	// 	println("Yes, Jummah Day!")
	// default:
	// 	println("Another boring day")
	// }
	//* For Loop

	// for i := 0; i <= 5; i++ {
	// 	makeCoffee()
	// 	fmt.Println(i)
	// }
}
