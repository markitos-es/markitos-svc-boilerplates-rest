package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/services"
)

func (s Server) delete(ctx *gin.Context) {
	id, err := s.GetParam(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	request := services.BoilerplateDeleteRequest{Id: *id}
	var service services.BoilerplateDeleteService = services.NewBoilerplateDeleteService(s.repository)
	if err := service.Do(request); err != nil {
		ctx.JSON(s.GetHTTPCode(err), errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"deleted": request.Id})
}
