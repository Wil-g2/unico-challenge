package fair

import (
	"github.com/wil-g2/unico-challenge/infra/log"
	"github.com/wil-g2/unico-challenge/infra/validator"
)

type service struct {
	repository FairRepository
	validator  validator.Validator
}

func NewService(repository FairRepository, validator validator.Validator) Service {
	return &service{repository: repository, validator: validator}
}

// Create implements Service
func (s *service) Create(input *CreateDTO) error {
	log.Info("create method run on service file", nil)
	err := s.validator.Validate(input)
	if err != nil {
		return err
	}
	fair := NewFair(
		input.Long,
		input.Lat,
		input.Setcens,
		input.Areap,
		input.CodDist,
		input.District,
		input.CodSubPref,
		input.Region5,
		input.Region8,
		input.FairName,
		input.Record,
		input.Street,
		input.Number,
		input.Neighborhood,
		input.Reference)
	return s.repository.Create(fair)
}

// Delete implements Service
func (s *service) Delete(id int) error {
	log.Info("delete method run on service file", nil)
	return s.repository.Delete(id)
}

// Find implements Service
func (s *service) Find(id int) (Fair, error) {
	log.Info("find method run on service file", nil)
	return s.repository.Find(id)
}

// List implements Service
func (s *service) List() ([]Fair, error) {
	log.Info("list method run on service file", nil)
	return s.repository.List()
}

// Update implements Service
func (s *service) Update(id int, input *UpdateDTO) (Fair, error) {
	log.Info("update method run on service file", nil)
	err := s.validator.Validate(input)
	if err != nil {
		return Fair{}, err
	}

	fair, err := s.repository.Find(id)
	if err != nil {
		return Fair{}, err
	}

	fair.Areap = input.Areap
	fair.CodDist = input.CodDist
	fair.CodSubPref = input.CodSubPref
	fair.District = input.District
	fair.FairName = input.FairName
	fair.Lat = input.Lat
	fair.Long = input.Long
	fair.Neighborhood = input.Neighborhood
	fair.Number = input.Number
	fair.Record = input.Record
	fair.Reference = input.Reference
	fair.Region5 = input.Region5
	fair.Region8 = input.Region8
	fair.Setcens = input.Setcens
	fair.Street = input.Street

	if err := s.repository.Update(&fair); err != nil {
		return Fair{}, err
	}
	return fair, err
}

// Find Fair by Name
func (s *service) FindByName(name string) ([]Fair, error) {
	log.Info("find method run on service file", nil)
	return s.repository.FindByName(name)
}
