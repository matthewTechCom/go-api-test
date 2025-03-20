package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	GoogleID string `gorm:"uniqueIndex" json:"id"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Name     string `json:"name"`
	GroupIDs []int  `json:"group_ids,omitempty"`
}
