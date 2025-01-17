package main

import "fmt"

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func foo1() {
	fmt.Println("foo")
}

func bar1() {
	fmt.Println("bar")
}


func init() {
	foo()
	bar()
	foo1()
	bar1()
}
