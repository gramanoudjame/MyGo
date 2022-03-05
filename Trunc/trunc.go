package main

import (
	"fmt"
)

/*
Write a program which prompts the user to enter a floating point number
and prints the integer which is a truncated version of the floating point number that was entered.
Truncation is the process of removing the digits to the right of the decimal place.
*/
func main() {
	var f64 float64
	fmt.Println("Enter a floating point number :")
	fmt.Scan(&f64)
	i64 := int64(f64)
	fmt.Printf("The integer part of %f is %d\n", f64, i64)
}
