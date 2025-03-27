package services_test

import (
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCanGetAResource(t *testing.T) {
	var request services.BoilerplateOneRequest = services.BoilerplateOneRequest{
		Id: domain.UUIDv4(),
	}

	var service services.BoilerplateOneService = services.NewBoilerplateOneService(repository)
	boiler, err := service.Do(request)

	assert.Nil(t, err)
	assert.IsType(t, services.BoilerplateOneResponse{}, *boiler)
	assert.True(t, repository.OneHaveBeenCalledWith(&request.Id))
}

func TestCantGetWithoutId(t *testing.T) {
	var request services.BoilerplateOneRequest = services.BoilerplateOneRequest{}

	var service services.BoilerplateOneService = services.NewBoilerplateOneService(repository)
	_, err := service.Do(request)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, domain.ErrBoilerplateBadRequest)
	assert.False(t, repository.OneHaveBeenCalledWith(&request.Id))
}

func TestCantGetWithoutInvalidId(t *testing.T) {
	var request services.BoilerplateOneRequest = services.BoilerplateOneRequest{
		Id: "invalid-id",
	}

	var service services.BoilerplateOneService = services.NewBoilerplateOneService(repository)
	_, err := service.Do(request)

	assert.ErrorIs(t, err, domain.ErrBoilerplateBadRequest)
	assert.False(t, repository.OneHaveBeenCalledWith(&request.Id))
}
