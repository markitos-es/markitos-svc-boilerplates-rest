package api

import (
	"net/http"

	"markitos-svc-boilerplates-rest/internal/services"

	"github.com/gin-gonic/gin"
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
