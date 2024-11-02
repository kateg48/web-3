package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var max byte
	for i := 0; i < len(s); i++ {
		if s[i] >= max {
			max = s[i]
		}
	}
	fmt.Printf("%c", max)
}
