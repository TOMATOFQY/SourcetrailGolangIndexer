package main

func main() {
	f()
	g(1)
	FuncFromAnotherFile()
}

func g(a int) int {
	a = a + 1
	return a
}

func f() {
	b := 1
	go g(b)
	go func() {}()
}
