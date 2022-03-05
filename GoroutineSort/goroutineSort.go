package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func swap(i, j *int) {
	tmp := *i
	*i = *j
	*j = tmp
}

func bubbleSort(sli []int) {
	for i := len(sli); i > 1; i-- {
		for j := 0; j < i-1; j++ {
			if sli[j+1] < sli[j] {
				swap(&sli[j], &sli[j+1])
			}
		}
	}
}

func scanSliceOfInts(sli *[]int) error {
	var err error
	var s string
	in := bufio.NewReader(os.Stdin)
	s, err = in.ReadString('\n')
	if err != nil {
		return err
	}
	s = strings.TrimSuffix(s, "\n")
	for _, e := range strings.Split(s, ",") {
		var integer int
		integer, err = strconv.Atoi(e)
		if err != nil {
			return err
		}
		*sli = append(*sli, integer)
	}
	return nil
}

func split(sli []int, n int) [][]int {
	result := make([][]int, n)
	if len(sli) == 0 {
		return result
	}
	q := len(sli) / n
	r := len(sli) % n
	for i, j := 0, 0; i < n; i++ {
		if j < r {
			result[i] = sli[j*(q+1) : (j+1)*(q+1)]
			j++
		} else {
			result[i] = sli[i*q+j : (i+1)*q+j]
		}
	}
	return result
}

func main() {
	// Get user input
	var input []int
	fmt.Println("Enter a serie of integers separated by ',':")
	for err := scanSliceOfInts(&input); err != nil; err = scanSliceOfInts(&input) {
		input = input[:0]
		fmt.Printf("Try again : %v\n", err)
	}
	// Get 4 sub slices
	subSli := split(input, 4)
	// Launch goroutines for sorting each sub slice
	var wg sync.WaitGroup
	for _, s := range subSli {
		wg.Add(1)
		go func(wg *sync.WaitGroup, s []int) {
			defer wg.Done()
			bubbleSort(s)
			fmt.Println(s)
		}(&wg, s)
	}
	wg.Wait()
	// Slice is already merged as sub slices were references inside the initial slice
	// So just sort it
	bubbleSort(input)
	fmt.Println(input)
}
