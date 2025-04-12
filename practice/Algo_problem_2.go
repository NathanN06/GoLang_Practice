package main

import "fmt"

//quicksort function making it so that it needs to run with one slice(int) and return an int
func quicksort(numbers []int)[]int{

	if len(numbers) <= 1 {
		return numbers
	}
	//define left and right slice
	left := []int{}
	right := []int{}

	//create pivot and pivotValue
	pivot := len(numbers)/2
	pivotValue := numbers[pivot]

	//Create Recursive loop
	for i := 0; i < len(numbers); i++{
		if i == pivot{
			continue
		}

		//If numbers is less/more than pivot value append left/right
		if numbers[i] < pivotValue{
			left = append(left, numbers[i])
		}
		if numbers[i] >= pivotValue{
			right = append(right, numbers[i])
		}
	}

	//quicksort left and right slices
	sortedLeft := quicksort(left)
	sortedRight := quicksort(right)

	//Combine 2 slices and pivot
	temp := append(sortedLeft, pivotValue)
	sorted := append(temp, sortedRight...)

	//Return sorted
	return sorted
}

func main(){
	// define numbers
	numbers := []int{}

	//use made function to sort the numbers
	sort := quicksort(numbers)

	//print sorted numbers
	fmt.Println("Numbers in ascending order: ", sort)
}