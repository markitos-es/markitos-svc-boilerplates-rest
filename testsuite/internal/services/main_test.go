package services_test

import (
	"os"
	"testing"

	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"
	internal_test "github.com/markitos-es/markitos-svc-boilerplates-rest/testsuite/internal"
)

type MockSpyBoilerRepository struct {
	LastCreatedBoilerName     *string
	LastDeleteBoilerId        *string
	LastOneBoilerId           *string
	LastUpdatedBoilerId       *string
	LastUpdatedBoilerName     *string
	LastOneForUpdatedBoilerId *string
	LastAllHaveBeenCalled     bool
	LastUpdateHaveBeenCalled  bool
	LastSearchHaveBeenCalled  bool
}

func NewMockSpyBoilerRepository() *MockSpyBoilerRepository {
	return &MockSpyBoilerRepository{
		LastCreatedBoilerName:     nil,
		LastDeleteBoilerId:        nil,
		LastOneBoilerId:           nil,
		LastUpdatedBoilerId:       nil,
		LastUpdatedBoilerName:     nil,
		LastOneForUpdatedBoilerId: nil,
		LastAllHaveBeenCalled:     false,
		LastUpdateHaveBeenCalled:  false,
		LastSearchHaveBeenCalled:  false,
	}
}

func (m *MockSpyBoilerRepository) Create(boiler *domain.Boilerplate) error {
	m.LastCreatedBoilerName = &boiler.Name

	return nil
}

func (m *MockSpyBoilerRepository) CreateHaveBeenCalledWith(boilerName *string) bool {
	var result bool = m.LastCreatedBoilerName != nil && *m.LastCreatedBoilerName == *boilerName

	m.LastCreatedBoilerName = nil

	return result
}

func (m *MockSpyBoilerRepository) Delete(id *domain.BoilerplateId) error {
	value := id.Value()
	m.LastDeleteBoilerId = &value

	return nil
}

func (m *MockSpyBoilerRepository) DeleteHaveBeenCalledWith(boilerId *string) bool {
	var result bool = m.LastDeleteBoilerId != nil && *m.LastDeleteBoilerId == *boilerId

	m.LastDeleteBoilerId = nil

	return result
}

func (m *MockSpyBoilerRepository) Update(boiler *domain.Boilerplate) error {
	m.LastUpdateHaveBeenCalled = true
	m.LastUpdatedBoilerId = &boiler.Id
	m.LastUpdatedBoilerName = &boiler.Name
	m.LastOneForUpdatedBoilerId = &boiler.Id

	return nil
}

func (m *MockSpyBoilerRepository) UpdateHaveBeenCalledWith(id, name string) bool {
	var matchCalled bool = m.LastUpdateHaveBeenCalled
	var matchId bool = *m.LastUpdatedBoilerId == id
	var matchName bool = *m.LastUpdatedBoilerName == name

	m.LastUpdatedBoilerId = nil
	m.LastUpdatedBoilerName = nil
	m.LastUpdateHaveBeenCalled = false

	return matchCalled && matchId && matchName
}

func (m *MockSpyBoilerRepository) UpdateHaveBeenCalled() bool {
	var matchCalled bool = m.LastUpdateHaveBeenCalled

	m.LastUpdateHaveBeenCalled = false
	m.LastUpdatedBoilerId = nil
	m.LastUpdatedBoilerName = nil

	return matchCalled
}

func (m *MockSpyBoilerRepository) UpdateHaveBeenCalledOneWith(id string) bool {
	var matchId bool = *m.LastOneForUpdatedBoilerId == id

	m.LastOneForUpdatedBoilerId = nil

	return matchId
}

func (m *MockSpyBoilerRepository) One(id *domain.BoilerplateId) (*domain.Boilerplate, error) {
	value := id.Value()
	m.LastOneBoilerId = &value

	return internal_test.NewRandomBoilerplateWithCustomId(id), nil
}

func (m *MockSpyBoilerRepository) OneHaveBeenCalledWith(boilerId *string) bool {
	var result bool = m.LastOneBoilerId != nil && *m.LastOneBoilerId == *boilerId

	m.LastOneBoilerId = nil

	return result
}

func (m *MockSpyBoilerRepository) All() ([]*domain.Boilerplate, error) {
	m.LastAllHaveBeenCalled = true

	return nil, nil
}

func (m *MockSpyBoilerRepository) AllHaveBeenCalled() bool {
	result := m.LastAllHaveBeenCalled
	m.LastAllHaveBeenCalled = false

	return result
}

func (m *MockSpyBoilerRepository) SearchAndPaginate(
	searchTerm string,
	pageNumber int,
	pageSize int) ([]*domain.Boilerplate, error) {
	m.LastSearchHaveBeenCalled = true

	return nil, nil
}

func (m *MockSpyBoilerRepository) SearchHaveBeenCalled() bool {
	result := m.LastSearchHaveBeenCalled

	m.LastSearchHaveBeenCalled = false

	return result
}

var repository *MockSpyBoilerRepository

func TestMain(m *testing.M) {
	repository = NewMockSpyBoilerRepository()

	os.Exit(m.Run())
}
