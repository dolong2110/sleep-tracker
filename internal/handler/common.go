package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"sleep-tracker/internal/utils"
	"sleep-tracker/pkg/errs"
	"sleep-tracker/pkg/httpx"
	"sleep-tracker/pkg/zapx"
)

// binData parse JSON body from requests
func bindData(c *gin.Context, req interface{}) bool {
	var (
		errResp *errs.Error
		err     error
	)

	if c.ContentType() != "application/json" {
		msg := fmt.Sprintf("%s only accepts Content - Template application/json.", c.FullPath())
		zapx.Error(context.TODO(), msg, errors.Errorf("not application/json content type."))
		errResp = errs.NewUnsupportedMediaType(msg)
		c.JSON(errResp.Code, httpx.ApiJson{
			Error:   []error{errResp},
			Message: "not application/json content type.",
		})
		return false
	}

	if err = c.ShouldBind(req); err != nil {
		zapx.Error(c, "error binding request data.", err)
		if _, ok := err.(validator.ValidationErrors); ok {
			invalidArgs := utils.GetInvalidArgs(err)

			zapx.Error(context.TODO(), "failed to parse the json data.", err)
			errResp = errs.NewBadRequest("invalid request parameters. See invalidArgs.")
			errResp.InvalidArgs = invalidArgs
			c.JSON(errResp.Code, httpx.ApiJson{
				Error:   []error{errResp},
				Message: "invalid body request data.",
			})
			return false
		}

		zapx.Error(context.TODO(), "failed to parse the json data.", err)
		errResp = errs.NewInternal()
		c.JSON(errResp.Code, httpx.ApiJson{
			Error:   []error{errResp},
			Message: "failed to parse body request data.",
		})
		return false
	}

	return true
}
