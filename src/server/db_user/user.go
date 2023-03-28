package db_user

import "time"

type User struct {
	ID           uint `gorm:"primarykey"`
	Username     string
	Password     string
	DailyScore   uint
	WeeklyScore  uint
	MonthlyScore uint
	TotalScore   uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
