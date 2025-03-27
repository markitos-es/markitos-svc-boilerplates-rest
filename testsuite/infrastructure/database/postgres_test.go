package database_test

import (
	"testing"
	"time"

	"github.com/markitos-es/markitos-svc-boilerplates-rest/infrastructure/database"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"
	"github.com/markitos-es/markitos-svc-boilerplates-rest/testsuite/infrastructure/testdb"
	internal_test "github.com/markitos-es/markitos-svc-boilerplates-rest/testsuite/internal"
	"github.com/stretchr/testify/require"
)

func TestBoilerplateCreate(t *testing.T) {
	var boiler *domain.Boilerplate = internal_test.NewRandomBoilerplate()
	err := testdb.GetRepository().Create(boiler)
	require.NoError(t, err)

	var result domain.Boilerplate
	err = testdb.GetDB().First(&result, "id = ?", boiler.Id).Error
	require.NoError(t, err)
	require.Equal(t, boiler.Id, result.Id)
	require.Equal(t, boiler.Name, result.Name)
	require.WithinDuration(t, boiler.CreatedAt, result.CreatedAt, time.Second)
	require.WithinDuration(t, boiler.UpdatedAt, result.UpdatedAt, time.Second)

	testdb.GetDB().Delete(&result)
}

func TestBoilerplateDelete(t *testing.T) {
	var boiler *domain.Boilerplate = internal_test.NewRandomBoilerplate()
	testdb.GetRepository().Create(boiler)

	repository := database.NewBoilerPostgresRepository(testdb.GetDB())

	id, _ := domain.NewBoilerplateId(boiler.Id)
	err := repository.Delete(id)
	require.NoError(t, err)
}

func TestBoilerplateOne(t *testing.T) {
	var boiler *domain.Boilerplate = internal_test.NewRandomBoilerplate()
	testdb.GetRepository().Create(boiler)

	repository := database.NewBoilerPostgresRepository(testdb.GetDB())

	result, err := repository.One(boiler.GetId())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, boiler.Id, result.Id)
	require.Equal(t, boiler.Name, result.Name)

	id, _ := domain.NewBoilerplateId(boiler.Id)
	err = repository.Delete(id)
	require.NoError(t, err)
}

func TestBoilerUpdate(t *testing.T) {
	var boiler *domain.Boilerplate = internal_test.NewRandomBoilerplate()
	testdb.GetRepository().Create(boiler)

	repository := database.NewBoilerPostgresRepository(testdb.GetDB())

	boiler.Name = boiler.Name + "Updated"
	err := repository.Update(boiler)
	require.NoError(t, err)

	var result domain.Boilerplate
	err = testdb.GetDB().First(&result, "id = ?", boiler.Id).Error
	require.NoError(t, err)
	require.Equal(t, boiler.Name, result.Name)

	id, _ := domain.NewBoilerplateId(boiler.Id)
	err = repository.Delete(id)
	require.NoError(t, err)
}
