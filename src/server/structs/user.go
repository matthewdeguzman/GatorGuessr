package db_user

import "time"

type User struct {
	ID        uint `gorm:"primarykey"`
	Username  string
	Password  string
	Score     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
