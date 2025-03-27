package services

import "github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"

type BoilerplateUpdateRequest struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type BoilerplateUpdateService struct {
	Repository domain.BoilerplateRepository
}

func NewBoilerplateUpdateService(repository domain.BoilerplateRepository) BoilerplateUpdateService {
	return BoilerplateUpdateService{
		Repository: repository,
	}
}

func (s BoilerplateUpdateService) Do(request BoilerplateUpdateRequest) error {
	securedId, err := domain.NewBoilerplateId(request.Id)
	if err != nil {
		return err
	}
	securedName, err := domain.NewBoilerplateName(request.Name)
	if err != nil {
		return err
	}

	boiler, err := s.Repository.One(securedId)
	if err != nil {
		return err
	}

	boiler.Name = securedName.Value()

	return s.Repository.Update(boiler)
}
