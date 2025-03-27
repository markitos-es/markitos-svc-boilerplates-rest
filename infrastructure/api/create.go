package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/services"
)

func (s Server) create(ctx *gin.Context) {
	var request services.BoilerplateCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	var service services.BoilerplateCreateService = services.NewBoilerplateCreateService(s.repository)
	response, err := service.Do(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
