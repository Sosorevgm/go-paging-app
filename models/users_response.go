package models

type UsersResponse struct {
	Count int    `json:"count"`
	Users []User `json:"users"`
}
