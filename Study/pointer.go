package main

import "log"

func main() {
	var myStirng string
	myStirng = "Green"
	log.Println("myString is Set: ", myStirng)
	ChangeUsingPointer(&myStirng)
	log.Println("myString is Set using pointer: ", myStirng)
}
func ChangeUsingPointer(s *string) {
	log.Println("Address of string s", s)
	newValue := "Red"
	*s = newValue
}
