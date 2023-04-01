package jdi

import "fmt"

type JDI struct {
}

func (jdi *JDI) Test() {
	fmt.Println("test")
}
func (jdi *JDI) Test2() {
	fmt.Println("test2")
}

func NewJDI() *JDI {
	j := &JDI{}

	return j
}
