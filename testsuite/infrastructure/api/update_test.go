package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-svc-boilerplates-rest/internal/domain"
	"markitos-svc-boilerplates-rest/internal/services"

	"github.com/stretchr/testify/assert"
)

func TestCanUpdateABoilerplate(t *testing.T) {
	var boilerplate *domain.Boilerplate = createPersistedRandomBoilerplate()

	name := boilerplate.Name + " UPDATED"
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.BoilerplateUpdateRequest{
		Name: name,
	})
	request, _ := http.NewRequest(http.MethodPatch, "/v1/boilerplates/"+boilerplate.Id, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	RESTRouter().ServeHTTP(recorder, request)

	var response map[string]any
	json.NewDecoder(recorder.Body).Decode(&response)
	assert.Equal(t, http.StatusOK, recorder.Code)

	deletePersisteRandomBoilerplate(boilerplate.Id)
}

func TestCantUpdateANonExistingBoilerplate(t *testing.T) {
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.BoilerplateUpdateRequest{
		Name: domain.RandomPersonalName(),
	})
	request, _ := http.NewRequest(http.MethodPatch, "/v1/boilerplates/"+domain.UUIDv4(), bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	RESTRouter().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestCantUpdateAnInvalidBoilerplateId(t *testing.T) {
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.BoilerplateUpdateRequest{
		Name: domain.RandomPersonalName(),
	})
	request, _ := http.NewRequest(http.MethodPatch, "/v1/boilerplates/an-invalid-id-format", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	RESTRouter().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
