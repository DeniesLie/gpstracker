package response

import (
	"github.com/DeniesLie/gpstracker/internal/core/validation"
)

type Response[T any] struct {
	IsSuccess        bool               `json:"isSuccess"`
	Message          string             `json:"message"`
	Data             T                  `json:"data"`
	ValidationResult *validation.Result `json:"validationResult"`
}

func Success[T any](data T) Response[T] {
	return Response[T]{
		IsSuccess: true,
		Message:   SuccessMessage,
		Data:      data,
	}
}

func Error(message string) Response[any] {
	return Response[any]{
		IsSuccess: false,
		Message:   message,
	}
}

func InvalidData(res validation.Result) Response[any] {
	return Response[any]{
		IsSuccess:        false,
		Message:          InvalidRequestMessage,
		ValidationResult: &res,
	}
}
