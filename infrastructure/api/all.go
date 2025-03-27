package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/services"
)

func (s Server) all(ctx *gin.Context) {
	var service services.BoilerplateAllService = services.NewBoilerplateAllService(s.repository)
	response, err := service.Do()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Data)
}
