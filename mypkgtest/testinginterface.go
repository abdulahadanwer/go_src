package main

import (
	"mypkg"
)

func main() {
	var value mypkg.MyInterface
	value = mypkg.MyType(6)
	value.MethodWithoutParameters()
	value.MethodWithParamters(124.3)
	value.MethodWithReturnValue()
}
