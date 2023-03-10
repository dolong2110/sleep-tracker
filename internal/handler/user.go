package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sleep-tracker/internal/consts"
	"sleep-tracker/internal/models"
	"sleep-tracker/internal/utils"
	"sleep-tracker/pkg/errs"
	"sleep-tracker/pkg/httpx"
	"sleep-tracker/pkg/zapx"
	"time"
)

type UserHandler struct {
	UserService  models.UserServicer
	SleepService models.SleepServicer
	MaxBodyBytes int64
}

type UserHandlerConfig struct {
	UserService  models.UserServicer
	SleepService models.SleepServicer
	MaxBodyBytes int64
}

func NewUserHandler(config *UserHandlerConfig) *UserHandler {
	return &UserHandler{
		UserService:  config.UserService,
		SleepService: config.SleepService,
		MaxBodyBytes: config.MaxBodyBytes,
	}
}

func (h *UserHandler) Signup(c *gin.Context) {
	var err error
	req := new(models.UserRequest)
	if ok := bindData(c, req); !ok {
		zapx.Error(c, consts.InvalidRequestBody, nil)
		return
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err = h.UserService.Signup(c, user); err != nil {
		zapx.Error(c, "failed to signup user.", err)
		c.JSON(errs.Status(err), httpx.ApiJson{
			Error:   []error{err},
			Message: "failed to create new user.",
		})
		return
	}

	c.JSON(http.StatusCreated, httpx.ApiJson{
		Message: "successful to create new user.",
		Data:    []interface{}{user},
	})
}

func (h *UserHandler) Signin(c *gin.Context) {
	var err error
	req := new(models.UserRequest)
	if ok := bindData(c, req); !ok {
		zapx.Error(c, consts.InvalidRequestBody, nil)
		return
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err = h.UserService.Signin(c, user); err != nil {
		zapx.Error(c, "failed to signin user.", err)
		c.JSON(errs.Status(err), httpx.ApiJson{
			Error:   []error{err},
			Message: "failed to signin user.",
		})
		return
	}

	c.JSON(http.StatusOK, httpx.ApiJson{
		Message: "successful to signin.",
		Data:    []interface{}{user},
	})
}

func (h *UserHandler) CreateSleep(c *gin.Context) {
	var err error

	sleepRequest := new(models.SleepRequest)
	if ok := bindData(c, sleepRequest); !ok {
		zapx.Error(c, consts.InvalidRequestBody, nil)
		return
	}

	sleep, err := utils.GetSleepFromRequest(sleepRequest)
	if err != nil {
		zapx.Error(c, consts.InvalidRequestBody, nil)
		err = errs.NewBadRequest("invalid request body")
		c.JSON(errs.Status(err), httpx.ApiJson{
			Error:   []error{err},
			Message: "invalid request body.",
		})
		return
	}

	if err = h.SleepService.Create(c, sleep); err != nil {
		zapx.Error(c, "failed to create new sleep entry.", err)
		c.JSON(errs.Status(err), httpx.ApiJson{
			Error:   []error{err},
			Message: "failed to create new sleep entry.",
		})
		return
	}

	sleepRequest.ID = sleep.ID
	sleepRequest.SleepDuration = int64(sleep.SleepDuration / time.Millisecond)

	c.JSON(http.StatusCreated, httpx.ApiJson{
		Message: "successful to create new sleep entry.",
		Data:    []interface{}{sleepRequest},
	})
}

func (h *UserHandler) DeleteSleep(c *gin.Context) {
	var err error
	sleep := new(models.Sleep)
	if ok := bindData(c, sleep); !ok {
		zapx.Error(c, consts.InvalidRequestBody, nil)
		return
	}

	if err = h.SleepService.Delete(c, sleep); err != nil {
		zapx.Error(c, "failed to delete sleep.", err)
		c.JSON(errs.Status(err), httpx.ApiJson{
			Error:   []error{err},
			Message: "failed to delete a user.",
		})
		return
	}

	c.JSON(http.StatusOK, httpx.ApiJson{
		Message: "successful delete sleep.",
	})
}

func (h *UserHandler) UpdateSleep(c *gin.Context) {
	var err error
	sleepRequest := new(models.SleepRequest)
	if ok := bindData(c, sleepRequest); !ok {
		zapx.Error(c, consts.InvalidRequestBody, nil)
		return
	}

	sleep, err := utils.GetSleepFromRequest(sleepRequest)
	if err != nil {
		zapx.Error(c, consts.InvalidRequestBody, nil)
		err = errs.NewBadRequest("invalid request body")
		c.JSON(errs.Status(err), httpx.ApiJson{
			Error:   []error{err},
			Message: "invalid request body.",
		})
		return
	}

	if err = h.SleepService.Update(c, sleep); err != nil {
		zapx.Error(c, "failed to update sleep.", err)
		c.JSON(errs.Status(err), httpx.ApiJson{
			Error:   []error{err},
			Message: "failed to update a sleep.",
		})
		return
	}

	c.JSON(http.StatusOK, httpx.ApiJson{
		Message: "successful update user.",
		Data:    []interface{}{sleep},
	})
}
