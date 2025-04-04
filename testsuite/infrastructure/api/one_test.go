package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"markitos-svc-boilerplates-rest/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestBoilerplateCanGetOne(t *testing.T) {
	var boilerplate *domain.Boilerplate = createPersistedRandomBoilerplate()

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/boilerplates/"+boilerplate.Id, nil)
	request.Header.Set("Content-Type", "application/json")
	RESTRouter().ServeHTTP(recorder, request)

	var response map[string]any
	json.NewDecoder(recorder.Body).Decode(&response)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, response["name"].(string), boilerplate.Name)
	assert.Equal(t, response["id"].(string), boilerplate.Id)

	deletePersisteRandomBoilerplate(response["id"].(string))
}

func TestBoilerplateCantGetInvalidId(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/boilerplates/an-invalid-id", nil)
	request.Header.Set("Content-Type", "application/json")
	RESTRouter().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestBoilerplateCantGetValidIdButNonExistingResource(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/boilerplates/"+domain.UUIDv4(), nil)
	request.Header.Set("Content-Type", "application/json")
	RESTRouter().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}
