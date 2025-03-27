package api_test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/infrastructure/api"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/infrastructure/database"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/testsuite/infrastructure/testdb"
	internal_test "github.com/markitos-es/markitos-svc-boilerplates-rest/testsuite/internal"
)

var boilerplatesApiServer *api.Server

func TestMain(m *testing.M) {
	setupRESTServer()
	os.Exit(m.Run())
}

func RESTServer() *api.Server {
	return boilerplatesApiServer
}

func setupRESTServer() {
	gin.SetMode(gin.TestMode)
	boilerplatesApiServer = api.NewServer(":8080", testdb.GetRepository())
}

func RESTRouter() *gin.Engine {
	return RESTServer().Router()
}

func createPersistedRandomBoilerplate() *domain.Boilerplate {
	var boiler *domain.Boilerplate = internal_test.NewRandomBoilerplate()
	testdb.GetRepository().Create(boiler)

	return boiler
}

func persistBoilerplate(boiler *domain.Boilerplate) {
	testdb.GetRepository().Create(boiler)
}

func deletePersisteRandomBoilerplate(boilerId string) {
	repository := database.NewBoilerPostgresRepository(testdb.GetDB())
	id, err := domain.NewBoilerplateId(boilerId)
	if err != nil {
		panic(err)
	}

	repository.Delete(id)
}
