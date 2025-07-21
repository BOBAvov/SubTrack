package handler

import (
	"github.com/go-playground/validator/v10"
	"time"
)

func isDateValid(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	_, err := time.Parse("01.2006", dateStr)
	return err == nil
}
