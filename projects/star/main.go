package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		if i != len(s)-1 {
			fmt.Printf("%c", s[i])
			fmt.Print("*")
		} else {
			fmt.Printf("%c", s[i])
		}
	}
}
