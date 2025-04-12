package main

import "fmt"

//Quicksort function accepting and returning a slice
func quicksort(numbers []int)[]int{
	//Create if to stop infinite loop
	if len(numbers) <= 1 {
		return numbers
	}

	//create left and right slice
	left := []int{}
	right := []int{}

	//define pivot
	pivot := len(numbers)/2
	pivotValue := numbers[pivot]

	//Create recursive loop
	for i := 0; i < len(numbers); i++{
		if i == pivot{
			continue
		}
		if numbers[i] >= pivotValue{
			left = append(left, numbers[i])
		}
		if numbers [i] < pivotValue{
			right = append(right, numbers[i])
		}
	}

	//sort out left and right slices then combine into one return sorted sequence
	sortedLeft := quicksort(left)
	sortedRight := quicksort(right)

	temp := append(sortedLeft, pivotValue)
	sorted := append(temp, sortedRight...)

	//Return sequence
	return sorted
}


//Main function 
func main(){
	
	//Define slice
	numbers := []int{42, 7, 19, 3, 25, 11, 31, 8, 2, 17}

	//Sort using quicksort
	sort := quicksort(numbers)

	//print sorted numbers
	fmt.Println("Numbers in descending order:", sort)
}