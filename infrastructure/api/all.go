package api

import (
	"net/http"

	"markitos-svc-boilerplates-rest/internal/services"

	"github.com/gin-gonic/gin"
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
