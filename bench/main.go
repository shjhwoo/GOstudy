package main

func main() {}

func fibonacchi(n int) int {
	if n < 0 {
		return 1
	}

	if n < 2 {
		return n
	}

	return fibonacchi(n-2) + fibonacchi(n-1)
}
