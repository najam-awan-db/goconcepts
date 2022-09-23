package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(10))
	for i := 1; i <= 10; i = i + 1 {
		fmt.Println(fib(i))
	}
}
