package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/
func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a name")
	name, _ := in.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	fmt.Println("Enter an adress")
	adress, _ := in.ReadString('\n')
	adress = strings.TrimSuffix(adress, "\n")

	m := map[string]string{"name": name, "adress": adress}
	j, _ := json.Marshal(m)
	fmt.Printf("%s\n", j)
}
