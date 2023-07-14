package models

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Team string `json:"team"`
}
