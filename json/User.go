package main

type User struct {
	Id          int    `json:"id,omitzero"`
	Name        string `json:"name,omitempty"`
	Age         uint   `json:"age,omitzero"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}
