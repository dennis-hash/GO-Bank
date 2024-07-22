package main

import "math/rand"

type Account struct {
	ID        int `json: "id"`
	firstName string `json: "firstName"`
	lastName  string `json: "lastName"`
	Number    int64  `json: "number"`
}

func newAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(1000),
		firstName: firstName,
		lastName:  lastName,
		Number:    int64(rand.Intn(100000)),
	}
}
