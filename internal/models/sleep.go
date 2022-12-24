package models

import (
	"github.com/google/uuid"
	"time"
)

type Sleep struct {
	ID            uuid.UUID     `gorm:"column:id; primary_key" json:"id"`
	UserID        uuid.UUID     `gorm:"column:user_id" json:"user_id"`
	Date          time.Time     `gorm:"column:date" json:"date" time_format:"2006-01-02"`
	SleepTime     time.Time     `gorm:"column:sleep_time" json:"sleep_time" time_format:"2006-01-02"`
	WakeupTime    time.Time     `gorm:"column:wakeup_time" json:"wakeup_time" time_format:"2006-01-024"`
	SleepDuration time.Duration `gorm:"column:sleep_duration" json:"sleep_duration"`
}
