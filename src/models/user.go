package models

import "time"

// User Struct
//
// Represents the default user model
type User struct {
	ID         uint64    `json:"id,omitempty"`
	Username   string    `json:"username,omitempty"`
	Name       string    `json:"name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Password   string    `json:"password,omitempty"`
	JoinedDate time.Time `json:"joineddate,omitempty"`
}
