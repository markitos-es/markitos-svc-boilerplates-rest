package domain

import "time"

type Boilerplate struct {
	Id        string    `json:"id" binding:"required,uuid"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required,datetime" default:"now"`
	UpdatedAt time.Time `json:"updated_at" binding:"required,datetime" default:"now"`
}

func NewBoilerplate(id, name string) (*Boilerplate, error) {
	secureId, err := NewBoilerplateId(id)
	if err != nil {
		return nil, err
	}

	secureName, err := NewBoilerplateName(name)
	if err != nil {
		return nil, err
	}

	return &Boilerplate{
		Id:        secureId.value,
		Name:      secureName.value,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (b *Boilerplate) GetId() *BoilerplateId {
	result, _ := NewBoilerplateId(b.Id)

	return result
}
