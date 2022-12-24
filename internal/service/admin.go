package service

import "sleep-tracker/internal/models"

type AdminService struct {
	AdminRepository models.AdminRepositorier
	UserRepository  models.UserRepositorier
}

type AdminServiceConfig struct {
	AdminRepository models.AdminRepositorier
	UserRepository  models.UserRepositorier
}

func NewAdminService(config *AdminServiceConfig) models.AdminServicer {
	return &AdminService{
		AdminRepository: config.AdminRepository,
		UserRepository:  config.UserRepository,
	}
}
