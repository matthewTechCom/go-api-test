package model

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	GroupIDs []int  `json:"group_ids,omitempty"`
}
