package services

import (
	"markitos-svc-boilerplates-rest/internal/domain"
)

type BoilerplateSearchResponse struct {
	Data []*domain.Boilerplate `json:"data"`
}

type BoilerplateSearchRequest struct {
	SearchTerm string `json:"searchTerm"`
	PageNumber int    `json:"pageNumber" bindings:"min=1"`
	PageSize   int    `json:"pageSize" bindings:"min=10,max=100"`
}

type BoilerplateSearchService struct {
	Repository domain.BoilerplateRepository
}

func NewBoilerplateSearchService(repository domain.BoilerplateRepository) BoilerplateSearchService {
	return BoilerplateSearchService{
		Repository: repository,
	}
}

func (s BoilerplateSearchService) Do(request BoilerplateSearchRequest) (*BoilerplateSearchResponse, error) {
	response, err := s.Repository.SearchAndPaginate(request.SearchTerm, request.PageNumber, request.PageSize)
	if err != nil {
		return nil, err
	}

	return &BoilerplateSearchResponse{
		Data: response,
	}, nil
}
