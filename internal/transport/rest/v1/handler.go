package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/lapkomo2018/goServiceExample/pkg/model"
	"log"
)

type (
	StructService interface {
		ChangeMessage(id uint, message string) error
		Add(str *model.Struct) error
		FindByID(id uint) (*model.Struct, error)
	}

	Validator interface {
		Message(message string) error
	}

	Handler struct {
		structService StructService
		validator     Validator
	}
)

func New(structService StructService, validator Validator) *Handler {
	return &Handler{
		structService: structService,
		validator:     validator,
	}
}

func (h *Handler) Init(api *echo.Group) {
	log.Println("Initializing V1 api")
	v1 := api.Group("/v1")
	{
		h.initStructApi(v1)
	}
}
