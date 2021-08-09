package main

import "fmt"

func main() {

	var firstnumber, secondnumber float64
	var operator string

	fmt.Print("enter 1st number : ")
	fmt.Scanln(&firstnumber)

	fmt.Print("enter 2nd number : ")
	fmt.Scanln(&secondnumber)

	fmt.Print("enter operator (+ - * / ) : ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%5.2f %s %.2f = %.2f", firstnumber, operator, secondnumber, firstnumber+secondnumber)

	case "-":
		fmt.Printf("%.2f %s %.2f = %.2f", firstnumber, operator, secondnumber, firstnumber-secondnumber)

	case "*":
		fmt.Printf("%.2f %s %.2f = %.2f", firstnumber, operator, secondnumber, firstnumber*secondnumber)

	case "/":
		if secondnumber == 0.0 {
			fmt.Println("cant divide any thing by zero")
		} /* else {
				fmt.Printf("%.2f %s %.2f = %.2f", firstnumber, operator, secondnumber, firstnumber/secondnumber)
			}
		default:
			fmt.Println("not valid operator")*/
	}
}
