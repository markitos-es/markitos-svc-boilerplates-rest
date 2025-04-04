package services

import "markitos-svc-boilerplates-rest/internal/domain"

type BoilerplateAllResponse struct {
	Data []*domain.Boilerplate `json:"data"`
}

type BoilerplateAllService struct {
	Repository domain.BoilerplateRepository
}

func NewBoilerplateAllService(repository domain.BoilerplateRepository) BoilerplateAllService {
	return BoilerplateAllService{
		Repository: repository,
	}
}

func (s BoilerplateAllService) Do() (*BoilerplateAllResponse, error) {
	boilerplates, err := s.Repository.All()
	if err != nil {
		return nil, err
	}

	return &BoilerplateAllResponse{
		Data: boilerplates,
	}, nil
}
