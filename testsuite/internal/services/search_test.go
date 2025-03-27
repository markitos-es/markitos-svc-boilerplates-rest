package services_test

import (
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCanSearchResources(t *testing.T) {
	var service services.BoilerplateSearchService = services.NewBoilerplateSearchService(repository)
	boiler, err := service.Do(services.BoilerplateSearchRequest{})

	assert.Nil(t, err)
	assert.IsType(t, services.BoilerplateSearchResponse{}, *boiler)
	assert.True(t, repository.SearchHaveBeenCalled())
}
