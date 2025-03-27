package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/services"
	internal_test "github.com/markitos-es/markitos-svc-boilerplates-rest/testsuite/internal"
	"github.com/stretchr/testify/assert"
)

func TestBoilerplateCanCreate(t *testing.T) {
	recorder := httptest.NewRecorder()
	boiler := internal_test.NewRandomOnlyNameBoilerplate()
	requestBody, _ := json.Marshal(services.BoilerplateCreateRequest{
		Name: boiler.Name,
	})
	request, _ := http.NewRequest(http.MethodPost, "/v1/boilerplates", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	RESTRouter().ServeHTTP(recorder, request)

	var response map[string]any
	json.NewDecoder(recorder.Body).Decode(&response)
	responseId := response["id"].(string)
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, boiler.Name, response["name"])
	assert.NotEmpty(t, responseId)

	deletePersisteRandomBoilerplate(responseId)
}

func TestBoilerplateCantCreateWithoutName(t *testing.T) {
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.BoilerplateCreateRequest{})
	request, _ := http.NewRequest(http.MethodPost, "/v1/boilerplates", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	RESTRouter().ServeHTTP(recorder, request)

	var response map[string]any
	json.NewDecoder(recorder.Body).Decode(&response)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestBoilerplateCantCreateWithoutValidName(t *testing.T) {
	recorder := httptest.NewRecorder()
	requestBody, _ := json.Marshal(services.BoilerplateCreateRequest{
		Name: "!!!!!invalid!!!name!!!",
	})
	request, _ := http.NewRequest(http.MethodPost, "/v1/boilerplates", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")

	RESTRouter().ServeHTTP(recorder, request)

	var response map[string]any
	json.NewDecoder(recorder.Body).Decode(&response)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
