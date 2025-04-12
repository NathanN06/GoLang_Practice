package main

import (
	"fmt"
	"time"
)

func Publisher(ch chan int){
	i := 0 //Sets counter to 0
	for { // creates loop
	i++ // int increases by one
	time.Sleep(10 * time.Second) //every 10 seconds
	ch <- i // channel is equal to current i value 
	}
	close(ch)
}

func Subscriber(ch chan int){
	for { // loop
		value := <- ch //Value is equal to ch
		fmt.Println("Recieved:", value) // print value
	}
	
}


func main () {
	c := make(chan int) // creates channel which is interger

	go Publisher(c) // go routine which allows the function publisher to run simultaeneously
	go Subscriber(c) // go routine 

	select{} // keeps programme running
}