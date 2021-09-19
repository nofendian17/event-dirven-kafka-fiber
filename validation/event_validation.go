package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"kafka-fibre/exception"
	"kafka-fibre/model"
)

func ValidateCreateEvent(request model.EventRequest)  {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Key, validation.Required),
		validation.Field(&request.Value, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
