package models

import (
	"github.com/google/uuid"
	"time"
)

type Sleep struct {
	ID            uuid.UUID     `gorm:"column:id; primary_key" json:"id"`
	UserID        uuid.UUID     `gorm:"column:user_id" json:"user_id"`
	Date          time.Time     `gorm:"column:date; type:date" json:"date"`
	SleepTime     time.Time     `gorm:"column:sleep_time; type:time" json:"sleep_time"`
	WakeupTime    time.Time     `gorm:"column:wakeup_time; type:time" json:"wakeup_time"`
	SleepDuration time.Duration `gorm:"column:sleep_duration" json:"sleep_duration"`
}

type SleepRequest struct {
	ID            uuid.UUID `json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	Date          string    `json:"date"`
	SleepTime     string    `json:"sleep_time,omitempty"`
	WakeupTime    string    `json:"wakeup_time,omitempty"`
	SleepDuration int64     `json:"sleep_duration,omitempty"`
}
