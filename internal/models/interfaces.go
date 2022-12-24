package models

import (
	"github.com/gin-gonic/gin"
)

type UserServicer interface {
	Signup(c *gin.Context, user *User) error
	Signin(c *gin.Context, user *User) error
	Delete(c *gin.Context, user *User) error
	Update(c *gin.Context, user *User) error
	List(c *gin.Context, users *[]User) error
}

type UserRepositorier interface {
	Create(c *gin.Context, user *User) error
	FindByConditions(c *gin.Context, user *User) error
	FindByID(c *gin.Context, user *User) error
	Delete(c *gin.Context, user *User) error
	Update(c *gin.Context, user *User) error
	List(c *gin.Context, users *[]User) error
}

type AdminServicer interface {
}

type AdminRepositorier interface {
}

type SleepServicer interface {
	Create(c *gin.Context, sleep *Sleep) error
	Delete(c *gin.Context, sleep *Sleep) error
	Update(c *gin.Context, user *Sleep) error
}

type SleepRepositorier interface {
	Create(c *gin.Context, sleep *Sleep) error
	FindByConditions(c *gin.Context, user *Sleep) error
	FindByID(c *gin.Context, sleep *Sleep) error
	Delete(c *gin.Context, sleep *Sleep) error
	Update(c *gin.Context, sleep *Sleep) error
}
