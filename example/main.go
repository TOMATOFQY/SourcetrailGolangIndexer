package main

func main() {
	f()
	fib(42)
	FuncFromAnotherFile()
}

func f() {
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-1)
}
