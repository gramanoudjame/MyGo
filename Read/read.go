package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of FIELD_SIZE 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/
const FIELD_SIZE int = 20
const SEPARATOR string = " "
const EOL string = "\n"

type name struct {
	fname string
	lname string
}

func main() {
	// Get file name
	fmt.Println("Enter the name of the text file")
	var fileName string
	_, err := fmt.Scan(&fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read file
	f, _ := os.Open(fileName)
	r := bufio.NewReader(f)
	line, _, err := r.ReadLine()
	var names []name
	for err == nil {
		var n name
		tokens := strings.Split(string(line), SEPARATOR)
		if len(tokens[0]) > FIELD_SIZE {
			n.fname = tokens[0][:FIELD_SIZE]
		} else {
			n.fname = tokens[0]
		}
		if len(tokens[1]) > FIELD_SIZE {
			n.lname = tokens[1][:FIELD_SIZE]
		} else {
			n.lname = tokens[1]
		}
		names = append(names, n)
		line, _, err = r.ReadLine()
	}
	f.Close()

	// Print slice of names
	for _, name := range names {
		fmt.Printf("%v %v\n", name.fname, name.lname)
	}
}
