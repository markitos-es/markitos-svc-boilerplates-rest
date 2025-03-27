package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/services"
)

func (s *Server) update(ctx *gin.Context) {
	request, shouldExitByError := createRequestOrExitWithError(ctx)
	if shouldExitByError {
		return
	}

	var service services.BoilerplateUpdateService = services.NewBoilerplateUpdateService(s.repository)
	if err := service.Do(request); err != nil {
		ctx.JSON(s.GetHTTPCode(err), errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"updated": request.Id})
}

type BoilerplateUpdateRequestUri struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type BoilerplateUpdateRequestBody struct {
	Name string `json:"name" binding:"required"`
}

func createRequestOrExitWithError(ctx *gin.Context) (services.BoilerplateUpdateRequest, bool) {
	var requestUri BoilerplateUpdateRequestUri
	if err := ctx.ShouldBindUri(&requestUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return services.BoilerplateUpdateRequest{}, true
	}
	var requestBody BoilerplateUpdateRequestBody
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return services.BoilerplateUpdateRequest{}, true
	}

	var request services.BoilerplateUpdateRequest = services.BoilerplateUpdateRequest{
		Id:   requestUri.Id,
		Name: requestBody.Name,
	}

	return request, false
}
