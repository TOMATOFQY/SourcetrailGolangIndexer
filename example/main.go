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

type I interface{ g() }

type S struct{ a int }

func (s S) g() {}

type T struct{ a int }

func (t T) g() {}
func (t T) h() {}
