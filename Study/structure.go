package main

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Birthday    time.Time
}

func (u *User) Add() User {
	u.FirstName = "India"
	u.LastName = "Sabne"
	u.Age = 31
	u.PhoneNumber = "9987275806"
	return *u
}
func main() {
	user := User{
		FirstName:   "Bharat",
		LastName:    "Sabne",
		PhoneNumber: "9987275806",
		Age:         31,
		Birthday:    time.Date(1991, 8, 15, 00, 00, 00, 000, time.Local),
	}
	log.Println(user.FirstName, user.LastName, user.Birthday, user.Age)
	user = user.Add()
	log.Println(user.FirstName, user.LastName, user.Birthday, user.Age)

	myMap := make(map[string]string)
	myMap["Name"] = "Bharat"
	myMap["LastName"] = "Sabne"

	fmt.Println(myMap["Name"], myMap["LastName"])
}
