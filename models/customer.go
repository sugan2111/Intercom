package models

type Customer struct {
	Latitude  string `json:"latitude"`
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	Longitude string `json:"longitude"`
}
