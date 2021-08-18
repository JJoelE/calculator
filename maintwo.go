package main

import (
	"fmt"
)

func addy(x, y float64) float64 {
	return x + y
}
func minus(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}
func divide(x, y float64) float64 {
	return x / y
}

func main() {

	var firstnumber, secondnumber float64
	var operator string

	fmt.Print("enter 1st number : ")
	fmt.Scanln(&firstnumber)

	fmt.Print("enter 2nd number : ")
	fmt.Scanln(&secondnumber)

	fmt.Print("enter operator (+ - * / ) : ")
	fmt.Scanln(&operator)

	if operator == "+" {
		answer := addy(firstnumber, secondnumber)
		fmt.Println(answer)
	}
	if operator == "-" {
		answer := minus(firstnumber, secondnumber)
		fmt.Println(answer)
	}
	if operator == "*" {
		product := multiply(firstnumber, secondnumber)
		fmt.Println(product)
	}
	if operator == "/" {
		answer := divide(firstnumber, secondnumber)
		fmt.Println(answer)
	}
	if operator != "+" && operator != "-" && operator != "/" && operator != "*" {
		fmt.Println("sorry, invalid operator")
	}
}
