package utils

import (
	"github.com/go-playground/validator/v10"
	"mindx/pkg/errs"
)

// GetInvalidArgs -extract invalid arguments from error (normally in request's error)
func GetInvalidArgs(err error) []errs.InvalidArgument {
	var invalidArgs []errs.InvalidArgument
	if valErr, ok := err.(validator.ValidationErrors); ok {
		var valFieldErr validator.FieldError
		for _, valFieldErr = range valErr {
			invalidArgs = append(invalidArgs, errs.InvalidArgument{
				Field: valFieldErr.Field(),
				Value: valFieldErr.Value().(string),
				Tag:   valFieldErr.Tag(),
				Param: valFieldErr.Param(),
			})
		}
	}

	return invalidArgs
}
