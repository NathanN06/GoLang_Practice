package main

import (
	"fmt"
)

//Global slice of numbers
var numbers = []int {9,8,7,6,5,4,3,2,1}

func main () {
	
	//Set min and max values
	min := numbers[0] //set first value to min
	max := numbers[0] //set first value to max
	
	//loop through all numbers in slice
	for i := 0; i < len(numbers); i ++{
		
		//If number is smaller than min replace minimum value
		if numbers[i] < min {
			min = numbers[i]
		}
		
		//If number is greater than max replace max
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	fmt.Println("Maximum Value:", max)//print max
	fmt.Println("Minimum Value:", min)//print min

}