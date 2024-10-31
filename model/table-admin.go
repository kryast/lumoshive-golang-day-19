package model

type Admin struct {
	ID        int    `json:"id"`
	AdminName string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
