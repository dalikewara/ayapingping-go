package entity

import "time"

type Role struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
