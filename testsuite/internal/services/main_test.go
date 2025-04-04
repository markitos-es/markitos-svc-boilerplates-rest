package services_test

import (
	"os"
	"testing"

	"markitos-svc-boilerplates-rest/internal/domain"
	internal_test "markitos-svc-boilerplates-rest/testsuite/internal"
)

type MockSpyBoilerRepository struct {
	LastCreatedBoilerplateName     *string
	LastDeleteBoilerplateId        *string
	LastOneBoilerplateId           *string
	LastUpdatedBoilerplateId       *string
	LastUpdatedBoilerplateName     *string
	LastOneForUpdatedBoilerplateId *string
	LastAllHaveBeenCalled          bool
	LastUpdateHaveBeenCalled       bool
	LastSearchHaveBeenCalled       bool
}

func NewMockSpyBoilerRepository() *MockSpyBoilerRepository {
	return &MockSpyBoilerRepository{
		LastCreatedBoilerplateName:     nil,
		LastDeleteBoilerplateId:        nil,
		LastOneBoilerplateId:           nil,
		LastUpdatedBoilerplateId:       nil,
		LastUpdatedBoilerplateName:     nil,
		LastOneForUpdatedBoilerplateId: nil,
		LastAllHaveBeenCalled:          false,
		LastUpdateHaveBeenCalled:       false,
		LastSearchHaveBeenCalled:       false,
	}
}

func (m *MockSpyBoilerRepository) Create(boilerplate *domain.Boilerplate) error {
	m.LastCreatedBoilerplateName = &boilerplate.Name

	return nil
}

func (m *MockSpyBoilerRepository) CreateHaveBeenCalledWith(boilerplateName *string) bool {
	var result bool = m.LastCreatedBoilerplateName != nil && *m.LastCreatedBoilerplateName == *boilerplateName

	m.LastCreatedBoilerplateName = nil

	return result
}

func (m *MockSpyBoilerRepository) Delete(id *domain.BoilerplateId) error {
	value := id.Value()
	m.LastDeleteBoilerplateId = &value

	return nil
}

func (m *MockSpyBoilerRepository) DeleteHaveBeenCalledWith(boilerplateId *string) bool {
	var result bool = m.LastDeleteBoilerplateId != nil && *m.LastDeleteBoilerplateId == *boilerplateId

	m.LastDeleteBoilerplateId = nil

	return result
}

func (m *MockSpyBoilerRepository) Update(boilerplate *domain.Boilerplate) error {
	m.LastUpdateHaveBeenCalled = true
	m.LastUpdatedBoilerplateId = &boilerplate.Id
	m.LastUpdatedBoilerplateName = &boilerplate.Name
	m.LastOneForUpdatedBoilerplateId = &boilerplate.Id

	return nil
}

func (m *MockSpyBoilerRepository) UpdateHaveBeenCalledWith(id, name string) bool {
	var matchCalled bool = m.LastUpdateHaveBeenCalled
	var matchId bool = *m.LastUpdatedBoilerplateId == id
	var matchName bool = *m.LastUpdatedBoilerplateName == name

	m.LastUpdatedBoilerplateId = nil
	m.LastUpdatedBoilerplateName = nil
	m.LastUpdateHaveBeenCalled = false

	return matchCalled && matchId && matchName
}

func (m *MockSpyBoilerRepository) UpdateHaveBeenCalled() bool {
	var matchCalled bool = m.LastUpdateHaveBeenCalled

	m.LastUpdateHaveBeenCalled = false
	m.LastUpdatedBoilerplateId = nil
	m.LastUpdatedBoilerplateName = nil

	return matchCalled
}

func (m *MockSpyBoilerRepository) UpdateHaveBeenCalledOneWith(id string) bool {
	var matchId bool = *m.LastOneForUpdatedBoilerplateId == id

	m.LastOneForUpdatedBoilerplateId = nil

	return matchId
}

func (m *MockSpyBoilerRepository) One(id *domain.BoilerplateId) (*domain.Boilerplate, error) {
	value := id.Value()
	m.LastOneBoilerplateId = &value

	return internal_test.NewRandomBoilerplateWithCustomId(id), nil
}

func (m *MockSpyBoilerRepository) OneHaveBeenCalledWith(boilerplateId *string) bool {
	var result bool = m.LastOneBoilerplateId != nil && *m.LastOneBoilerplateId == *boilerplateId

	m.LastOneBoilerplateId = nil

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
