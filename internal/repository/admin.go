package repository

import (
	"gorm.io/gorm"
	"mindx/internal/models"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) models.AdminRepositorier {
	return &AdminRepository{
		DB: db,
	}
}
