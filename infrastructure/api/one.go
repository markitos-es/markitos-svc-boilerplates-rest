package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/services"
)

func (s Server) one(ctx *gin.Context) {
	id, err := s.GetParam(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	request := services.BoilerplateOneRequest{Id: *id}
	var service services.BoilerplateOneService = services.NewBoilerplateOneService(s.repository)
	response, err := service.Do(request)
	if err != nil {
		ctx.JSON(s.GetHTTPCode(err), errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Data)
}
