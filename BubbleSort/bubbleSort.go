package main

import (
	"fmt"
)

func Swap(sli []int, i int) {
	tmp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = tmp
}

func BubbleSort(sli []int) {
	for i := len(sli); i > 1; i-- {
		for j := 0; j < i-1; j++ {
			if sli[j+1] < sli[j] {
				Swap(sli, j)
			}
		}
	}
}

func main() {
	var sli []int
	var err error
	// Get user input
	fmt.Println("Enter up to 10 integer separated by space:")
	for i := 0; err == nil && i < 10; i++ {
		var d int
		_, err = fmt.Scan(&d)
		if err == nil {
			sli = append(sli, d)
		}
	}
	// Sort and print result
	BubbleSort(sli)
	fmt.Println(sli)
}
