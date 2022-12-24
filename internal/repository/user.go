package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mindx/internal/models"
	"mindx/pkg/errs"
	"mindx/pkg/zapx"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) models.UserRepositorier {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(c *gin.Context, userRequest *models.User) error {
	var (
		query *gorm.DB
	)
	if query = r.DB.WithContext(c).FirstOrCreate(userRequest, userRequest); query.Error != nil {
		zapx.Error(c, "failed to create a new userRequest", query.Error,
			zap.String("name", userRequest.Name), zap.String("email", userRequest.Email))
		if err, ok := query.Error.(*pgconn.PgError); ok && err.Code == pgerrcode.UniqueViolation {
			zapx.Error(c, fmt.Sprintf("userRequest with email %s already existed", userRequest.Email), err)
			return errs.NewConflict("userRequest", fmt.Sprintf("email - %s.", userRequest.Email))
		}

		return errs.NewInternal()
	}

	if query.RowsAffected == 0 {
		zapx.Error(c, "userRequest is already existed.", fmt.Errorf("userRequest is already existed"),
			zap.String("name", userRequest.Name), zap.String("email", userRequest.Email))
		return errs.NewConflict("userRequest", fmt.Sprintf("email - %s.", userRequest.Email))
	}

	return nil
}

func (r *UserRepository) FindByConditions(c *gin.Context, user *models.User) error {
	var query *gorm.DB
	if query = r.DB.WithContext(c).Find(user, user); query.Error != nil {
		zapx.Error(c, "could not find user", query.Error,
			zap.String("email", user.Email), zap.String("name", user.Name))
		return errs.NewInternal()
	}

	if query.RowsAffected == 0 {
		zapx.Error(c, "user is not existed", fmt.Errorf("user not found"),
			zap.String("email", user.Email), zap.String("name", user.Name))
		return errs.NewNotFound("user", fmt.Sprintf("email - %s", user.Email))
	}

	return nil
}

func (r *UserRepository) FindByID(c *gin.Context, user *models.User) error {
	var query *gorm.DB
	if query = r.DB.WithContext(c).Find(user); query.Error != nil {
		zapx.Error(c, "could not find user", query.Error,
			zap.String("email", user.Email), zap.String("name", user.Name))
		return errs.NewInternal()
	}

	if query.RowsAffected == 0 {
		zapx.Error(c, "user is not existed", fmt.Errorf("user not found"),
			zap.String("email", user.Email), zap.String("name", user.Name))
		return errs.NewNotFound("user", fmt.Sprintf("email - %s", user.Email))
	}

	return nil
}

func (r *UserRepository) Delete(c *gin.Context, user *models.User) error {
	var query *gorm.DB
	if query = r.DB.WithContext(c).Delete(user, user); query.Error != nil {
		zapx.Error(c, "failed to delete user", query.Error,
			zap.String("email", user.Email), zap.String("name", user.Name))
		return errs.NewInternal()
	}

	return nil
}

func (r *UserRepository) Update(c *gin.Context, user *models.User) error {
	if query := r.DB.WithContext(c).Save(user); query.Error != nil {
		zapx.Error(c, "failed to update user", query.Error,
			zap.String("email", user.Email), zap.String("name", user.Name))
		return errs.NewInternal()
	}

	return nil
}

func (r *UserRepository) List(c *gin.Context, users *[]models.User) error {
	if query := r.DB.WithContext(c).Find(users); query.Error != nil {
		zapx.Error(c, "failed to update user", query.Error)
		return errs.NewInternal()
	}

	return nil
}
