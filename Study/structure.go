package main

import (
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

func main() {
	user := User{
		FirstName:   "Bharat",
		LastName:    "Sabne",
		PhoneNumber: "9987275806",
		Age:         31,
		Birthday:    time.Date(1991, 8, 15, 00, 00, 00, 000, time.Local),
	}
	log.Println(user.FirstName, user.LastName, user.Birthday)
}
