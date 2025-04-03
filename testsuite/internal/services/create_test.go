package services_test

import (
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCanCreateAUser(t *testing.T) {
	var boilerplate domain.Boilerplate = domain.Boilerplate{
		Name: domain.RandomPersonalName(),
	}
	var request services.BoilerplateCreateRequest = services.BoilerplateCreateRequest{
		Name: boilerplate.Name,
	}

	var service services.BoilerplateCreateService = services.NewBoilerplateCreateService(repository)
	response, err := service.Do(request)

	assert.Nil(t, err)
	assert.True(t, repository.CreateHaveBeenCalledWith(&request.Name))
	assert.Equal(t, response.Name, request.Name)
	assert.NotEmpty(t, response.Id)
}

func TestCantCreateWithoutName(t *testing.T) {
	var request services.BoilerplateCreateRequest = services.BoilerplateCreateRequest{}

	var service services.BoilerplateCreateService = services.NewBoilerplateCreateService(repository)
	_, err := service.Do(request)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, domain.ErrBoilerplateBadRequest)
	assert.False(t, repository.CreateHaveBeenCalledWith(&request.Name))
}

func TestCantCreateWithoutValidName(t *testing.T) {
	var request services.BoilerplateCreateRequest = services.BoilerplateCreateRequest{
		Name: "",
	}

	var service services.BoilerplateCreateService = services.NewBoilerplateCreateService(repository)
	_, err := service.Do(request)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, domain.ErrBoilerplateBadRequest)
	assert.False(t, repository.CreateHaveBeenCalledWith(&request.Name))
}
