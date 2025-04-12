package main

import "fmt"

//Define Holders struct
type Holders struct{
	id string
	balance int
}

//create a quicksort function to take a slice and produce a slice
func quicksort(holders []Holders)[]Holders{
	// create if statement to stop recursive loop
	if len(holders) <= 1 {
		return holders
	}

	//Define pivot

	pivot := len(holders)/2
	pivotValue := holders[pivot].balance

	//define left and right slice
	left := []Holders{}
	right := []Holders{}

	//Create recusrive loop
	for i := 0 ; i < len(holders);i ++{
		if holders[i].balance == pivotValue{
			continue
		}
		if holders[i].balance >= pivotValue{
			left = append(left, holders[i])
		}
		if holders[i].balance < pivotValue{
			right = append(right, holders[i])
		}
	}
	sortedLeft := quicksort(left)
	sortedRight := quicksort(right)
	
	pivotHolder := holders[pivot]

	//Combine pivot left and right and return right
	temp := append(sortedLeft, pivotHolder)
	sorted := append(temp, sortedRight...)

	return sorted

}

func main() {
	holders := []Holders{
		{"alice", 120},
		{"bob", 300},
		{"charlie", 250},
		{"diana", 400},
		{"eve", 150},
		{"frank", 90},
		{"grace", 220},
		{"heidi", 180},
		{"ivan", 330},
		{"judy", 275},
	}

	sorted := quicksort(holders)
	for i, h := range sorted[:3] {
	fmt.Printf("Top %d: %s with $%d\n", i+1, h.id, h.balance)
}
}