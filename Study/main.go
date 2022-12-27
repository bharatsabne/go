package main

import "fmt"

func main() {
	fmt.Println("This is test!")
	var someVariable string
	var i int
	someVariable = "String from Variable"
	fmt.Println(someVariable)
	i = 9987275806
	fmt.Println("I Value :", i)
	returnFromString, otherThingFromFunction := SomeFunction()
	fmt.Println("Function returned: ", returnFromString, otherThingFromFunction)
}

func SomeFunction() (string, string) {
	return "Something!", "LOL"
}
