package services

import "github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"

type BoilerplateCreateRequest struct {
	Name string
}

type BoilerplateCreateResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type BoilerplateCreateService struct {
	Repository domain.BoilerplateRepository
}

func NewBoilerplateCreateService(repository domain.BoilerplateRepository) BoilerplateCreateService {
	return BoilerplateCreateService{
		Repository: repository,
	}
}

func (s BoilerplateCreateService) Do(request BoilerplateCreateRequest) (*BoilerplateCreateResponse, error) {
	var id string = domain.UUIDv4()
	boilerplate, err := domain.NewBoilerplate(id, request.Name)
	if err != nil {
		return nil, err
	}

	if err := s.Repository.Create(boilerplate); err != nil {
		return nil, err
	}

	return &BoilerplateCreateResponse{
		Id:   boilerplate.Id,
		Name: boilerplate.Name,
	}, nil
}
