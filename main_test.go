package main

import (
	"fmt"
	"testing"
)

func Test_Foo(t *testing.T) {
	fmt.Println("Test Foo")
}

func Test_Bar(t *testing.T) {
	fmt.Println("Test Bar")
}

func Test_DB(t *testing.T) {
	fmt.Println("Test DB")
	dbInit()
}
