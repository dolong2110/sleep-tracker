package utils

import (
	"sleep-tracker/internal/models"
	"sleep-tracker/pkg/timex"
)

func GetSleepFromRequest(sleepRequest *models.SleepRequest) (*models.Sleep, error) {
	date, err := timex.GetDateFromString(sleepRequest.Date)
	if err != nil {
		return nil, err
	}

	sleepTime, err := timex.GetDayTimeFromString(sleepRequest.SleepTime)
	if err != nil {
		return nil, err
	}

	wakeupTime, err := timex.GetDayTimeFromString(sleepRequest.WakeupTime)
	if err != nil {
		return nil, err
	}

	return &models.Sleep{
		ID:         sleepRequest.ID,
		UserID:     sleepRequest.UserID,
		Date:       date,
		SleepTime:  sleepTime,
		WakeupTime: wakeupTime,
	}, nil
}
