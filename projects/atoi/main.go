package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Print("Ошибка!Нужно число.")
	}

	var s string
	var str = strconv.Itoa(x)
	var res string
	for i := 0; i < len(str); i++ {
		s = string(str[i])
		a, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		res = res + strconv.Itoa(a*a)
	}
	fmt.Print(res)
}
