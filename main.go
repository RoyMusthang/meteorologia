package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i string
	fmt.Scan(&i)
	n, _ := strconv.Atoi(i)
	fmt.Println(fizzbuzz(n))
}

func fizzbuzz(i int) string {
	var result string

	if i%3 == 0 && i%5 == 0 {
		result = "FizzBuzz"
	} else if i%3 == 0 {
		result = "Fizz"
	} else if i%5 == 0 {
		result = "Buzz"
	} else {
		return ""
	}
	return result
}
