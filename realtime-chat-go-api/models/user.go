package models

// User ...
type User struct {
	ID           int    `json:"id"`
	UserName     string `json:"userName"`
	UserEmail    string `json:"userEmail"`
	Active       int    `json:"active"`
	OnlineStatus int    `json:"onlineStatus"`
}

// DbUser ...
type DbUser struct {
	ID           int    `json:"id"`
	UserName     string `json:"userName"`
	UserPassword string `json:"userPassword"`
	UserEmail    string `json:"userEmail"`
	Active       int    `json:"active"`
	OnlineStatus int    `json:"onlineStatus"`
}
