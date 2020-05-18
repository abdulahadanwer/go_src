package mypkg

import "fmt"

// MyInterface deffine
type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParamters(float64)
	MethodWithReturnValue() string
}

// MyType value
type MyType int

// MethodWithoutParameters value
func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

// MethodWithParamters value
func (m MyType) MethodWithParamters(f float64) {
	fmt.Println("MethodWithParamters called with", f, m)
}

// MethodWithReturnValue value
func (m MyType) MethodWithReturnValue() string {
	return "hi from MethodWithReturnValue called"
}
