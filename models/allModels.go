package models

import (
	"errors"
	"sync"
	"time"
)

type User struct {
	mu           sync.Mutex
	UserID       int
	PersonalInfo Person
	Capital      float32
}

func (u *User) GetCapital() float32 {
	var transaction Transactions
	transaction.TID = time.Now() //Hashear
	transaction.UserID = u.UserID
	transaction.Amount = 0
	transaction.Type = 0
	transaction.Capital = u.Capital
	return u.Capital
}

func (u *User) Credit(money float32) {
	u.mu.Lock()
	u.Capital += money
	var transaction Transactions
	transaction.TID = time.Now() //Hashear
	transaction.UserID = u.UserID
	transaction.Amount = money
	transaction.Type = 1
	transaction.Capital = u.Capital
	u.mu.Unlock()
}

func (u *User) Debit(money float32) error {
	u.mu.Lock()
	if u.Capital-money < 0 {
		return errors.New("Not enough credit")
	}
	u.Capital -= money
	var transaction Transactions
	transaction.TID = time.Now() //Hashear
	transaction.UserID = u.UserID
	transaction.Amount = money
	transaction.Type = 2
	transaction.Capital = u.Capital
	u.mu.Unlock()
	return nil
}

type Transactions struct {
	TID     time.Time //hash
	UserID  int
	Type    int //o string(?
	Amount  float32
	Capital float32
}

type Person struct {
	FName          string
	LName          string
	Birth          string
	DocumentType   string
	DocumentNumber string
	Address        Address
}

type Address struct {
	Street  string
	Number  string
	City    string
	State   string
	Country string
}

type BodyCredit struct {
	UserID int `binding:"required"`
	Data   interface{}
	Amount float32 `binding:"required,gte=0"`
}

type BodyDebit struct {
	UserID int `binding:"required"`
	Data   interface{}
	Amount float32 `binding:"required,gte=0"`
}
