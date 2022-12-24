package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"gorm.io/gorm"
	"sleep-tracker/internal/models"
	"sleep-tracker/pkg/errs"
	"sleep-tracker/pkg/zapx"
)

type SleepRepository struct {
	DB *gorm.DB
}

func NewSleepRepository(db *gorm.DB) models.SleepRepositorier {
	return &SleepRepository{
		DB: db,
	}
}

func (r *SleepRepository) Create(c *gin.Context, sleep *models.Sleep) error {
	var (
		query *gorm.DB
	)
	if query = r.DB.WithContext(c).FirstOrCreate(sleep, sleep); query.Error != nil {
		zapx.Error(c, "failed to create a new userRequest", query.Error)
		if err, ok := query.Error.(*pgconn.PgError); ok && err.Code == pgerrcode.UniqueViolation {
			zapx.Error(c, fmt.Sprintf("sleep with id %s already existed", sleep.ID), err)
			return errs.NewConflict("sleep", fmt.Sprintf("id - %s.", sleep.ID))
		}

		return errs.NewInternal()
	}

	if query.RowsAffected == 0 {
		zapx.Error(c, "userRequest is already existed.", fmt.Errorf("sleep is already existed"))
		return errs.NewConflict("sleep", fmt.Sprintf("id - %s.", sleep.ID))
	}

	return nil
}

func (r *SleepRepository) Delete(c *gin.Context, sleep *models.Sleep) error {
	var query *gorm.DB
	if query = r.DB.WithContext(c).Delete(sleep, sleep); query.Error != nil {
		zapx.Error(c, "failed to delete sleep", query.Error)
		return errs.NewInternal()
	}

	if query.RowsAffected == 0 {
		zapx.Error(c, "sleep is not existed", fmt.Errorf("sleep not found"))
		return errs.NewNotFound("sleep", fmt.Sprintf("id - %s", sleep.ID))
	}

	return nil
}

func (r *SleepRepository) Update(c *gin.Context, sleep *models.Sleep) error {
	var query *gorm.DB
	if query = r.DB.WithContext(c).Save(sleep); query.Error != nil {
		zapx.Error(c, "failed to update sleep", query.Error)
		return errs.NewInternal()
	}

	if query.RowsAffected == 0 {
		zapx.Error(c, "sleep is not existed", fmt.Errorf("sleep not found"))
		return errs.NewNotFound("sleep", fmt.Sprintf("id - %s", sleep.ID))
	}

	return nil
}
