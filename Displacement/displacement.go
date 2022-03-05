package main

import "fmt"

func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var a, vo, so, t float64
	fmt.Println("Enter values for acceleration, initial velocity and initiale displacment:")
	fmt.Scan(&a)
	fmt.Scan(&vo)
	fmt.Scan(&so)
	fmt.Println("Enter values for time:")
	fmt.Scan(&t)
	fn := GenDisplaceFn(a, vo, so)
	fmt.Println(fn(t))
}
