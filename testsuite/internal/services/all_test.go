package services_test

import (
	"testing"

	"markitos-svc-boilerplates-rest/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestCanGetAllResources(t *testing.T) {
	var service services.BoilerplateAllService = services.NewBoilerplateAllService(repository)
	boilerplate, err := service.Do()

	assert.Nil(t, err)
	assert.IsType(t, services.BoilerplateAllResponse{}, *boilerplate)
	assert.True(t, repository.AllHaveBeenCalled())
}
