package internal_test

import (
	"time"

	"github.com/markitos-es/markitos-svc-boilerplates-rest/internal/domain"
)

func NewRandomBoilerplate() *domain.Boilerplate {
	return &domain.Boilerplate{
		Id:        domain.UUIDv4(),
		Name:      domain.RandomPersonalName(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewRandomOnlyNameBoilerplate() *domain.Boilerplate {
	return &domain.Boilerplate{
		Name: domain.RandomPersonalName(),
	}
}

func NewRandomBoilerplateWithCustomId(boilerplateId *domain.BoilerplateId) *domain.Boilerplate {
	return &domain.Boilerplate{
		Id:        boilerplateId.Value(),
		Name:      domain.RandomPersonalName(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
