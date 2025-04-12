package main

import "fmt"

// Create function that takes in an interger and outputs an int
func fib(n int) int {

	// Define empty slice which in n+1 in the fib sequence
	fibn := make([]int, n+1)

	//Define the first and second term in the sequence
	fibn[0] = 0
	fibn[1] = 1

	//create loop starting at 2
	for i := 2; i <= n; i++ {
		fibn[i] = fibn[i-1] + fibn[i-2]
	}
	return fibn[n]
}

// Create main function
func main() {

	//Create variable n
	var n int

	//ask user for any interger
	fmt.Print("Please enter any interger: ")
	fmt.Scan(&n)

	sorted := fib(n)

	//print result
	fmt.Println("The", n, "term in the fibonacci sequence is: ", sorted)
}
