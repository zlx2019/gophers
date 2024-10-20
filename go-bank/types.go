package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// 账户
type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"number"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		// ID:        rand.Intn(10_0000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Int63n(1000000000),
		CreatedAt: time.Now().In(time.FixedZone("UTC+8", 60*60*8)),
	}
}
