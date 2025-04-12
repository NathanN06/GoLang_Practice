package main

import (
	"fmt"
)

//Interface
type Speaker interface {
	Speak() string
}

//Struct
type Singer struct {
	Name string
	Age float64
	Genre string
}

type Actor struct {
	Name   string
	Age    float64 
	Movies int
}


func (s Singer) Speak() string{
	return fmt.Sprintf("Hi, I'm %s, a %.0f-year-old %s singer!", s.Name, s.Age, s.Genre)
}

func (s Actor) Speak() string{
	return fmt.Sprintf("Hi, I'm %s, a %.0f-year-old actor who has acted in %d movies!", s.Name, s.Age, s.Movies)
}

func Introduce (s Speaker) {
	fmt.Println(s.Speak())
}



func main() {
	singer1 := Singer{Name: "Adele", Age: 35, Genre: "Pop"}
	actor1 := Actor{Name: "Leonardo DiCaprio", Age: 49, Movies: 40}
	// Call the function using the interface
	Introduce(singer1)
	Introduce(actor1)
}