package services_test

import (
	"testing"

	"markitos-svc-boilerplates-rest/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestCanSearchResources(t *testing.T) {
	var service services.BoilerplateSearchService = services.NewBoilerplateSearchService(repository)
	boilerplate, err := service.Do(services.BoilerplateSearchRequest{})

	assert.Nil(t, err)
	assert.IsType(t, services.BoilerplateSearchResponse{}, *boilerplate)
	assert.True(t, repository.SearchHaveBeenCalled())
}
