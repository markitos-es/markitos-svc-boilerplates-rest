package database

import (
	"fmt"

	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"
	"gorm.io/gorm"
)

type BoilerPostgresRepository struct {
	db *gorm.DB
}

func NewBoilerPostgresRepository(db *gorm.DB) BoilerPostgresRepository {
	return BoilerPostgresRepository{db: db}
}

func (r *BoilerPostgresRepository) Create(boilerplate *domain.Boilerplate) error {
	return r.db.Create(boilerplate).Error
}

func (r *BoilerPostgresRepository) Delete(id *domain.BoilerplateId) error {
	_, err := r.One(id)
	if err != nil {
		return domain.ErrBoilerplateNotFound
	}

	return r.db.Delete(&domain.Boilerplate{}, "id = ?", id.Value()).Error
}

func (r *BoilerPostgresRepository) Update(boilerplate *domain.Boilerplate) error {
	return r.db.Save(boilerplate).Error
}

func (r *BoilerPostgresRepository) One(id *domain.BoilerplateId) (*domain.Boilerplate, error) {
	var boiler domain.Boilerplate
	if err := r.db.First(&boiler, "id = ?", id.Value()).Error; err != nil {
		return nil, domain.ErrBoilerplateNotFound
	}
	return &boiler, nil
}

func (r *BoilerPostgresRepository) All() ([]*domain.Boilerplate, error) {
	var boilers []*domain.Boilerplate
	if err := r.db.Find(&boilers).Error; err != nil {
		return nil, domain.ErrBoilerplateBadRequest
	}

	return boilers, nil
}

func (r *BoilerPostgresRepository) SearchAndPaginate(
	searchTerm string,
	pageNumber int,
	pageSize int) ([]*domain.Boilerplate, error) {
	offset := (pageNumber - 1) * pageSize
	var boilers []*domain.Boilerplate
	if err := r.db.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", searchTerm)).
		Order("name").
		Limit(pageSize).
		Offset(offset).
		Find(&boilers).Error; err != nil {
		return nil, err
	}

	return boilers, nil
}
