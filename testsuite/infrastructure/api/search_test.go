package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestCanSearchWithPattern(t *testing.T) {
	var pattern string = domain.RandomString(10)

	var ids []string
	for range 5 {
		id := domain.UUIDv4()
		ids = append(ids, id)
		name := pattern + domain.RandomPersonalName()
		boilerplate, _ := domain.NewBoilerplate(id, name)

		persistBoilerplate(boilerplate)
	}

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/boilerplates?search="+pattern+"&page=1&size=6", nil)
	request.Header.Set("Content-Type", "application/json")
	RESTRouter().ServeHTTP(recorder, request)

	var response services.BoilerplateSearchResponse
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, 5, len(response.Data))
	assert.Equal(t, http.StatusOK, recorder.Code)

	for _, id := range ids {
		assert.Contains(t, recorder.Body.String(), id)
	}

	for _, id := range ids {
		deletePersisteRandomBoilerplate(id)
	}
}

func TestCantSearchWithoutInvalidOptionalPage(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/v1/boilerplates?page=abc", nil)
	request.Header.Set("Content-Type", "application/json")
	RESTRouter().ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
