package handler

import (
	"github.com/Mohamadreza-shad/ucl-draw/service/draw"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DrawHandler struct {
	DrawService *draw.Service
	Validator   *validator.Validate
}

func (h *DrawHandler) Draw(c *gin.Context) {
	err := h.DrawService.Draw(c.Request.Context())
	if err != nil {
		MakeErrorResponseWithoutCode(c.Writer,err)
		return
	}
	MakeSuccessResponse(c.Writer,nil, "draw has been done successfully")
}

func NewDrawHandler(
	drawService *draw.Service,
	validator *validator.Validate,
) *DrawHandler {
	return &DrawHandler{
		DrawService: drawService,
		Validator:   validator,
	}
}
