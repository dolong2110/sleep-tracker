package service

import (
	"github.com/gin-gonic/gin"
	"sleep-tracker/internal/models"
	"sleep-tracker/pkg/zapx"
)

type SleepService struct {
	SleepRepository models.SleepRepositorier
	UserRepository  models.UserRepositorier
}

type SleepServiceConfig struct {
	SleepRepository models.SleepRepositorier
	UserRepository  models.UserRepositorier
}

func NewSleepService(config *SleepServiceConfig) models.SleepServicer {
	return &SleepService{
		SleepRepository: config.SleepRepository,
		UserRepository:  config.UserRepository,
	}
}

func (s *SleepService) Create(c *gin.Context, sleep *models.Sleep) error {
	var err error

	if err = s.UserRepository.FindByID(c, &models.User{ID: sleep.UserID}); err != nil {
		zapx.Error(c, "failed to find the user", err)
		return err
	}

	sleep.SleepDuration = sleep.WakeupTime.Sub(sleep.SleepTime)

	if err = s.SleepRepository.Create(c, sleep); err != nil {
		zapx.Error(c, "failed to find the sleep", err)
		return err
	}

	return nil
}

func (s *SleepService) Delete(c *gin.Context, sleep *models.Sleep) error {
	var err error

	if err = s.SleepRepository.Delete(c, sleep); err != nil {
		zapx.Error(c, "failed to delete sleep", err)
		return err
	}

	return nil
}

func (s *SleepService) Update(c *gin.Context, sleep *models.Sleep) error {
	var err error

	sleepFetched := &models.Sleep{
		ID:     sleep.ID,
		UserID: sleep.UserID,
		Date:   sleep.Date,
	}

	if !sleep.SleepTime.IsZero() {
		sleepFetched.SleepTime = sleep.SleepTime
	}

	if !sleep.WakeupTime.IsZero() {
		sleepFetched.WakeupTime = sleep.WakeupTime
	}

	sleepFetched.SleepDuration = sleepFetched.WakeupTime.Sub(sleep.SleepTime)
	*sleep = *sleepFetched

	if err = s.SleepRepository.Update(c, sleep); err != nil {
		zapx.Error(c, "failed to update sleep", err)
		return err
	}

	return nil
}
