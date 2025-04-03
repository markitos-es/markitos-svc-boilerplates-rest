package database

import (
	"fmt"

	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"
	"gorm.io/gorm"
)

type BoilerplatePostgresRepository struct {
	db *gorm.DB
}

func NewBoilerplatePostgresRepository(db *gorm.DB) BoilerplatePostgresRepository {
	return BoilerplatePostgresRepository{db: db}
}

func (r *BoilerplatePostgresRepository) Create(boilerplate *domain.Boilerplate) error {
	return r.db.Create(boilerplate).Error
}

func (r *BoilerplatePostgresRepository) Delete(id *domain.BoilerplateId) error {
	_, err := r.One(id)
	if err != nil {
		return domain.ErrBoilerplateNotFound
	}

	return r.db.Delete(&domain.Boilerplate{}, "id = ?", id.Value()).Error
}

func (r *BoilerplatePostgresRepository) Update(boilerplate *domain.Boilerplate) error {
	return r.db.Save(boilerplate).Error
}

func (r *BoilerplatePostgresRepository) One(id *domain.BoilerplateId) (*domain.Boilerplate, error) {
	var boilerplate domain.Boilerplate
	if err := r.db.First(&boilerplate, "id = ?", id.Value()).Error; err != nil {
		return nil, domain.ErrBoilerplateNotFound
	}
	return &boilerplate, nil
}

func (r *BoilerplatePostgresRepository) All() ([]*domain.Boilerplate, error) {
	var boilerplates []*domain.Boilerplate
	if err := r.db.Find(&boilerplates).Error; err != nil {
		return nil, domain.ErrBoilerplateBadRequest
	}

	return boilerplates, nil
}

func (r *BoilerplatePostgresRepository) SearchAndPaginate(
	searchTerm string,
	pageNumber int,
	pageSize int) ([]*domain.Boilerplate, error) {
	offset := (pageNumber - 1) * pageSize
	var boilerplates []*domain.Boilerplate
	if err := r.db.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", searchTerm)).
		Order("name").
		Limit(pageSize).
		Offset(offset).
		Find(&boilerplates).Error; err != nil {
		return nil, err
	}

	return boilerplates, nil
}
