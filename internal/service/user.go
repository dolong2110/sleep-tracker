package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mindx/internal/models"
	"mindx/internal/utils"
	"mindx/pkg/errs"
	"mindx/pkg/zapx"
)

type UserService struct {
	UserRepository models.UserRepositorier
}

type UserServiceConfig struct {
	UserRepository models.UserRepositorier
}

func NewUserService(config *UserServiceConfig) models.UserServicer {
	return &UserService{
		UserRepository: config.UserRepository,
	}
}

func (s *UserService) Signup(c *gin.Context, user *models.User) error {
	var (
		err error
	)

	if user.Password == "" {
		zapx.Error(c, "password not provided", fmt.Errorf("password not provided"))
		return errs.NewBadRequest("password should not be empty")
	}

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		zapx.Error(c, "failed to hash password", err)
		return errs.NewInternal()
	}

	user.Password = password
	if err = s.UserRepository.Create(c, user); err != nil {
		zapx.Error(c, "failed to create new userRequest", err)
		return err
	}

	return nil
}

func (s *UserService) Signin(c *gin.Context, user *models.User) error {
	var err error
	userFetch := &models.User{
		Email: user.Email,
	}
	if err = s.UserRepository.FindByConditions(c, userFetch); err != nil {
		zapx.Error(c, "failed to find user", err)
		return err
	}

	match, err := utils.ComparePasswords(userFetch.Password, user.Password)
	if err != nil {
		zapx.Error(c, "failed to compare password", err)
		return errs.NewInternal()
	}

	if !match {
		zapx.Error(c, "password not matched", fmt.Errorf("password not matched"))
		return errs.NewAuthorization("Invalid email or password")
	}

	*user = *userFetch

	return nil
}

func (s *UserService) Delete(c *gin.Context, user *models.User) error {
	var err error

	if err = s.UserRepository.FindByConditions(c, user); err != nil {
		zapx.Error(c, "failed to get user", err)
		return err
	}

	if err = s.UserRepository.Delete(c, user); err != nil {
		zapx.Error(c, "failed to delete user", err)
		return err
	}

	return nil
}

func (s *UserService) Update(c *gin.Context, user *models.User) error {
	var err error

	userFetched := &models.User{
		Email: user.Email,
	}

	if err = s.UserRepository.FindByConditions(c, userFetched); err != nil {
		zapx.Error(c, "failed to get user", err)
		return err
	}

	userFetched.Name = user.Name
	*user = *userFetched

	if err = s.UserRepository.Update(c, user); err != nil {
		zapx.Error(c, "failed to update user", err)
		return err
	}

	return nil
}

func (s *UserService) List(c *gin.Context, users *[]models.User) error {
	var err error

	if err = s.UserRepository.List(c, users); err != nil {
		zapx.Error(c, "failed to get the list of all users.", err)
		return err
	}

	return nil
}
