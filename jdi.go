package jdi

import "fmt"

type JDI struct {
}

func (jdi *JDI) Test() {
	fmt.Println("test")
}

func NewJDI() *JDI {
	j := &JDI{}

	return j
}
