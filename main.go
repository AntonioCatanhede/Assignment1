package main

import "fmt"

func main() {

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range s {

		if s[number]%2 != 0 {
			fmt.Println(s[number], " is odd \n")

		} else {

			fmt.Println(s[number], " is even \n")
		}
	}

}
