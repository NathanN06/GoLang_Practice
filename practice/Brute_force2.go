package main

import "fmt"

func main() {
	for a := 'a'; a <= 'z'; a++ {
		for b := 'a'; b <= 'z'; b++ {
			for c := 'a'; c <= 'z'; c++ {
				fmt.Printf("%c%c%c\n", a, b, c)
			}
		}
	}
}
