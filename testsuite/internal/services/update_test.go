package services_test

import (
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCanUpdateABoiler(t *testing.T) {
	var request services.BoilerplateUpdateRequest = services.BoilerplateUpdateRequest{
		Id:   domain.UUIDv4(),
		Name: domain.RandomPersonalName(),
	}

	var service services.BoilerplateUpdateService = services.NewBoilerplateUpdateService(repository)
	err := service.Do(request)

	assert.Nil(t, err)
	assert.True(t, repository.UpdateHaveBeenCalledWith(request.Id, request.Name))
	assert.True(t, repository.UpdateHaveBeenCalledOneWith(request.Id))
}

func TestCantUpdateWithoutName(t *testing.T) {
	var request services.BoilerplateUpdateRequest = services.BoilerplateUpdateRequest{
		Id: domain.UUIDv4(),
	}

	var service services.BoilerplateUpdateService = services.NewBoilerplateUpdateService(repository)
	err := service.Do(request)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, domain.ErrBoilerplateBadRequest)
	assert.False(t, repository.UpdateHaveBeenCalled())
}

func TestCantUpdateWithoutId(t *testing.T) {
	var request services.BoilerplateUpdateRequest = services.BoilerplateUpdateRequest{
		Name: domain.RandomPersonalName(),
	}

	var service services.BoilerplateUpdateService = services.NewBoilerplateUpdateService(repository)
	err := service.Do(request)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, domain.ErrBoilerplateBadRequest)
	assert.False(t, repository.UpdateHaveBeenCalled())
}

func TestCantUpdateWithInvalidId(t *testing.T) {
	var request services.BoilerplateUpdateRequest = services.BoilerplateUpdateRequest{
		Id:   "invalid-id",
		Name: domain.RandomPersonalName(),
	}

	var service services.BoilerplateUpdateService = services.NewBoilerplateUpdateService(repository)
	err := service.Do(request)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, domain.ErrBoilerplateBadRequest)
	assert.False(t, repository.UpdateHaveBeenCalled())
}
