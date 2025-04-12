package main

import "fmt"

func fib(n int) int {
	//Recursive fibonacci function
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	//Define variable n as an int but empty
	var n int

	//Ask user for input
	fmt.Print("Enter any interger: ")
	//Save input as n
	fmt.Scanln(&n)

	//call fib function
	result := fib(n)

	//Print result
	fmt.Println("The Fibbonacci number at", n, "is", result)
}
