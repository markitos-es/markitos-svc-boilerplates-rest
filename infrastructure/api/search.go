package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/services"
)

func (s *Server) search(ctx *gin.Context) {
	searchTerm := ctx.Query("search")
	securedSearchTerm, err := domain.NewBoilerplateSearchTerm(searchTerm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	pageNumberStr := ctx.DefaultQuery("page", "1")
	if pageNumberStr == "" {
		pageNumberStr = "1"
	}
	securedPageNumber, err := domain.NewBoilerplatePositiveNumber(pageNumberStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	pageSizeStr := ctx.DefaultQuery("size", "10")
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}
	securedPageSize, err := domain.NewBoilerplatePositiveNumber(pageSizeStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	var service services.BoilerplateSearchService = services.NewBoilerplateSearchService(s.repository)
	var request services.BoilerplateSearchRequest = services.BoilerplateSearchRequest{
		SearchTerm: securedSearchTerm.Value(),
		PageNumber: securedPageNumber.ValueToInt(),
		PageSize:   securedPageSize.ValueToInt(),
	}

	response, err := service.Do(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResonses(err))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
