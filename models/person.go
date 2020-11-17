package models

type User struct {
	UserID       int64
	PersonalInfo Person
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
