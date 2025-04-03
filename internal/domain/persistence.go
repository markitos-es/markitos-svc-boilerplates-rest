package domain

type BoilerplateRepository interface {
	Create(boilerplate *Boilerplate) error
	Delete(id *BoilerplateId) error
	One(id *BoilerplateId) (*Boilerplate, error)
	Update(boilerplate *Boilerplate) error
	All() ([]*Boilerplate, error)
	SearchAndPaginate(searchTerm string, pageNumber int, pageSize int) ([]*Boilerplate, error)
}
