package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	check_err_number(n, err)
	for i := 0; i < n; i++ {
		var p int
		_, err := fmt.Scanf("%d", &p)
		check_err_number(p, err)
		if p > 3 {
			p = 1
		}
		fmt.Println(strings.Repeat("*", p))
	}
}

func check_err_number(n int, err error) {
	if err != nil {
		panic(err)
	}
	if n <= 0 {
		panic(fmt.Errorf("Invalid number given %d", n))
	}
}
