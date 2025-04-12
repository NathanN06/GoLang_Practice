package main

import "fmt"

var numbers = []int{14, 3, 29, 8, 21, 1, 17, 6, 13, 5}

func main () {
	for i := 0; i < len(numbers); i++{
		for j := 0; j < len(numbers)-1; j++{
			if numbers[j] < numbers[j+1] {
				temp := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = temp
			}
		}
	}
	fmt.Println("Numbers in Descending order:", numbers)
}