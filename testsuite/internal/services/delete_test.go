package services_test

import (
	"testing"

	"markitos-svc-boilerplates-rest/internal/domain"
	"markitos-svc-boilerplates-rest/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestCanDeleteAUser(t *testing.T) {
	var request services.BoilerplateDeleteRequest = services.BoilerplateDeleteRequest{
		Id: domain.UUIDv4(),
	}

	var service services.BoilerplateDeleteService = services.NewBoilerplateDeleteService(repository)
	err := service.Do(request)
	assert.Nil(t, err)
	assert.True(t, repository.DeleteHaveBeenCalledWith(&request.Id))
}

func TestCantDeleteWithoutId(t *testing.T) {
	var request services.BoilerplateDeleteRequest = services.BoilerplateDeleteRequest{}

	var service services.BoilerplateDeleteService = services.NewBoilerplateDeleteService(repository)
	err := service.Do(request)

	assert.ErrorIs(t, err, domain.ErrBoilerplateBadRequest)
	assert.NotNil(t, err)
	assert.False(t, repository.DeleteHaveBeenCalledWith(&request.Id))
}

func TestCantDeleteWithInvalidId(t *testing.T) {
	var request services.BoilerplateDeleteRequest = services.BoilerplateDeleteRequest{
		Id: "invalid-id",
	}

	var service services.BoilerplateDeleteService = services.NewBoilerplateDeleteService(repository)
	err := service.Do(request)

	assert.ErrorIs(t, err, domain.ErrBoilerplateBadRequest)
	assert.False(t, repository.DeleteHaveBeenCalledWith(&request.Id))
}
