package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestBoilerplateCanListAllResources(t *testing.T) {
	var boilerplate1 *domain.Boilerplate = createPersistedRandomBoilerplate()
	var boilerplate2 *domain.Boilerplate = createPersistedRandomBoilerplate()

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/boilerplates/all", nil)
	request.Header.Set("Content-Type", "application/json")
	RESTRouter().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), boilerplate1.Id)
	assert.Contains(t, recorder.Body.String(), boilerplate2.Id)

	deletePersisteRandomBoilerplate(boilerplate1.Id)
	deletePersisteRandomBoilerplate(boilerplate2.Id)
}
