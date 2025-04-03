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
	var boilerplate *domain.Boilerplate = internal_test.NewRandomBoilerplate()
	err := testdb.GetRepository().Create(boilerplate)
	require.NoError(t, err)

	var result domain.Boilerplate
	err = testdb.GetDB().First(&result, "id = ?", boilerplate.Id).Error
	require.NoError(t, err)
	require.Equal(t, boilerplate.Id, result.Id)
	require.Equal(t, boilerplate.Name, result.Name)
	require.WithinDuration(t, boilerplate.CreatedAt, result.CreatedAt, time.Second)
	require.WithinDuration(t, boilerplate.UpdatedAt, result.UpdatedAt, time.Second)

	testdb.GetDB().Delete(&result)
}

func TestBoilerplateDelete(t *testing.T) {
	var boilerplate *domain.Boilerplate = internal_test.NewRandomBoilerplate()
	testdb.GetRepository().Create(boilerplate)

	repository := database.NewBoilerplatePostgresRepository(testdb.GetDB())

	id, _ := domain.NewBoilerplateId(boilerplate.Id)
	err := repository.Delete(id)
	require.NoError(t, err)
}

func TestBoilerplateOne(t *testing.T) {
	var boilerplate *domain.Boilerplate = internal_test.NewRandomBoilerplate()
	testdb.GetRepository().Create(boilerplate)

	repository := database.NewBoilerplatePostgresRepository(testdb.GetDB())

	result, err := repository.One(boilerplate.GetId())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, boilerplate.Id, result.Id)
	require.Equal(t, boilerplate.Name, result.Name)

	id, _ := domain.NewBoilerplateId(boilerplate.Id)
	err = repository.Delete(id)
	require.NoError(t, err)
}

func TestBoilerUpdate(t *testing.T) {
	var boilerplate *domain.Boilerplate = internal_test.NewRandomBoilerplate()
	testdb.GetRepository().Create(boilerplate)

	repository := database.NewBoilerplatePostgresRepository(testdb.GetDB())

	boilerplate.Name = boilerplate.Name + "Updated"
	err := repository.Update(boilerplate)
	require.NoError(t, err)

	var result domain.Boilerplate
	err = testdb.GetDB().First(&result, "id = ?", boilerplate.Id).Error
	require.NoError(t, err)
	require.Equal(t, boilerplate.Name, result.Name)

	id, _ := domain.NewBoilerplateId(boilerplate.Id)
	err = repository.Delete(id)
	require.NoError(t, err)
}
