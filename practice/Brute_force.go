package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Print("Enter any word: ")
	fmt.Scanln(&s)

	fmt.Print("Enter a letter to search for: ")
	letter, _, _ := reader.ReadRune() // reads 1 char from input

	// count + positions
	counter := 0
	var positions []int

	for i := 0; i < len(s); i++ {
		if letter == rune(s[i]) {
			counter++
			positions = append(positions, i+1)
		}
	}

	fmt.Println("Total occurrences:", counter)
	fmt.Println("Positions:", positions)
}