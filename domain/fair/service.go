package fair

import (
	"github.com/wil-g2/unico-challenge/infra/repositories"
	"github.com/wil-g2/unico-challenge/infra/validator"
)

type Service interface {
	List() ([]Fair, error)
	Find(id int) (Fair, error)
	Create(input *CreateDTO) error
	Update(id int, input *UpdateDTO) (Fair, error)
	Delete(id int) error
}

type service struct {
	repository repositories.FairRepository
	validator  validator.Validator
}

func NewService(repository repositories.FairRepository, validator validator.Validator) Service {
	return &service{repository: repository, validator: validator}
}

// Create implements Service
func (s *service) Create(input *CreateDTO) error {
	fair := domain.NewFair(
		input.Long,
		input.Lat,
		input.Setcens,
		input.Areap,
		input.CodDist,
		input.Distrito,
		input.CodSubPref,
		input.Regiao5,
		input.Regiao8,
		input.FairName,
		input.Record,
		input.Street,
		input.Number,
		input.Neighborhood,
		input.Reference)
	// Long:         input.Long,
	// Lat:          input.Lat,
	// Setcens:      input.Setcens,
	// Areap:        input.Areap,
	// CodDist:      input.CodDist,
	// Distrito:     input.Distrito,
	// CodSubPref:   input.CodSubPref,
	// Regiao5:      input.Regiao5,
	// Regiao8:      input.Regiao8,
	// FairName:     input.FairName,
	// Record:       input.Record,
	// Street:       input.Street,
	// Number:       input.Number,
	// Neighborhood: input.Neighborhood,
	// Reference:    input.Reference,
	// }
	return s.repository.Create(fair)
}

// Delete implements Service
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

// Find implements Service
func (s *service) Find(id int) (Fair, error) {
	return s.repository.Find(id)
}

// List implements Service
func (s *service) List() ([]Fair, error) {
	return s.repository.List()
}

// Update implements Service
func (s *service) Update(id int, input *UpdateDTO) (Fair, error) {
	fair, err := s.repository.Find(id)
	if err != nil {
		return Fair{}, err
	}

	fair.Areap = input.Areap
	fair.CodDist = input.CodDist
	fair.CodSubPref = input.CodSubPref
	fair.Distrito = input.Distrito
	fair.FairName = input.FairName
	fair.Lat = input.Lat
	fair.Long = input.Long
	fair.Neighborhood = input.Neighborhood
	fair.Number = input.Number
	fair.Record = input.Record
	fair.Reference = input.Reference
	fair.Regiao5 = input.Regiao5
	fair.Regiao8 = input.Regiao8
	fair.Setcens = input.Setcens
	fair.Street = input.Street

	if err := s.repository.Update(&fair); err != nil {
		return Fair{}, err
	}
	return fair, err
}
