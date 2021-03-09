package main

type Bar interface {

	bar()
}

type Foo interface {


}

type S struct {}

func (s *S) bar(){}

func g(a int){
	a = a + 1
}

func f(){
	b := 1
	go g(b)
	go g(b)
	go g(b)
	go func(){}()
}

func main(){
	f()

}
