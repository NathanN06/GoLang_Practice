package main

import "fmt"

//Define struct for balance and id
type Holder struct{
	id string
	balance int
}

//Define slice with data
var holders = []Holder{
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

//Bubble sort main with print
func main (){
	for i := 0; i < len(holders); i++{
		for j := 0; j < len(holders)-1-i; j++{
			if holders[j].balance < holders[j+1].balance{
				temp := holders[j]
				holders[j] = holders[j+1]
				holders[j+1] = temp
			}
		}
	}
	//Define top 3
	top3 := holders[0:3]

	//print top 3
	fmt.Println("Top 3 Balances", top3)
}