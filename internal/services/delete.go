package services

import "markitos-svc-boilerplates-rest/internal/domain"

type BoilerplateDeleteRequest struct {
	Id string `json:"id"`
}

type BoilerplateDeleteService struct {
	Repository domain.BoilerplateRepository
}

func NewBoilerplateDeleteService(repository domain.BoilerplateRepository) BoilerplateDeleteService {
	return BoilerplateDeleteService{
		Repository: repository,
	}
}

func (s BoilerplateDeleteService) Do(request BoilerplateDeleteRequest) error {
	boilerplateId, err := domain.NewBoilerplateId(request.Id)
	if err != nil {
		return err
	}

	if err := s.Repository.Delete(boilerplateId); err != nil {
		return err
	}

	return nil
}
