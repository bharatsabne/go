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
	returnFromString, otherThingFromFunction, yetAnotherString := SomeFunction()
	fmt.Println("Function returned: ", returnFromString, otherThingFromFunction, yetAnotherString)
}

func SomeFunction() (string, string,string) {
	return "Something!", "LOL", "Yet another string"
}
