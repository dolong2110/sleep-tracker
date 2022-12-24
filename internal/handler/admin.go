package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sleep-tracker/internal/consts"
	"sleep-tracker/internal/models"
	"sleep-tracker/pkg/errs"
	"sleep-tracker/pkg/httpx"
	"sleep-tracker/pkg/zapx"
)

type AdminHandler struct {
	AdminService models.AdminServicer
	UserService  models.UserServicer
	MaxBodyBytes int64
}

type AdminHandlerConfig struct {
	AdminService models.AdminServicer
	UserService  models.UserServicer
	MaxBodyBytes int64
}

func NewAdminHandler(config *AdminHandlerConfig) *AdminHandler {
	return &AdminHandler{
		AdminService: config.AdminService,
		UserService:  config.UserService,
		MaxBodyBytes: config.MaxBodyBytes,
	}
}

func (h *AdminHandler) Delete(c *gin.Context) {
	var err error
	req := new(models.UserRequest)
	if ok := bindData(c, req); !ok {
		zapx.Error(c, consts.InvalidRequestBody, nil)
		return
	}

	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err = h.UserService.Delete(c, user); err != nil {
		zapx.Error(c, "failed to delete user.", err)
		c.JSON(errs.Status(err), httpx.ApiJson{
			Error:   []error{err},
			Message: "failed to delete a user.",
		})
		return
	}

	c.JSON(http.StatusOK, httpx.ApiJson{
		Message: "successful delete user.",
	})
}

func (h *AdminHandler) Update(c *gin.Context) {
	var err error
	req := new(models.UserRequest)
	if ok := bindData(c, req); !ok {
		zapx.Error(c, consts.InvalidRequestBody, nil)
		return
	}

	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err = h.UserService.Update(c, user); err != nil {
		zapx.Error(c, "failed to update user.", err)
		c.JSON(errs.Status(err), httpx.ApiJson{
			Error:   []error{err},
			Message: "failed to update a user.",
		})
		return
	}

	c.JSON(http.StatusOK, httpx.ApiJson{
		Message: "successful update user.",
		Data:    []interface{}{user},
	})
}

func (h *AdminHandler) GetUsers(c *gin.Context) {
	var err error
	users := new([]models.User)
	if err = h.UserService.List(c, users); err != nil {
		zapx.Error(c, "failed get list of all users.", err)
		c.JSON(errs.Status(err), httpx.ApiJson{
			Error:   []error{err},
			Message: "failed get list of all users.",
		})
		return
	}

	c.JSON(http.StatusOK, httpx.ApiJson{
		Message: "successful get list of all users.",
		Data:    []interface{}{users},
	})
}
