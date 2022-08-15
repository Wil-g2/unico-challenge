package fair_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	"github.com/wil-g2/unico-challenge/domain/fair"
)

func (s *ServiceSuite) TestCreate() {
	s.repo.EXPECT().Create(gomock.Any()).Return(nil)
	s.validator.EXPECT().Validate(gomock.Any()).Times(1)
	err := s.service.Create(&fair.CreateDTO{
		-46548146,
		-23568390,
		"355030885000019",
		"3550308005040",
		87,
		"VILA FORMOSA",
		26,
		"ARICANDUVAS",
		"Leste",
		"Leste 1",
		"1",
		"RUA CODAJ-s",
		"45",
		"VILA FORMOSA",
		"PRAÇA DE TESTE",
	})
	s.NoError(err)
}

func (s *ServiceSuite) TestCreateWithError() {
	s.repo.EXPECT().Create(gomock.Any()).Return(errors.New("error on create"))
	s.validator.EXPECT().Validate(gomock.Any()).Times(1)
	err := s.service.Create(&fair.CreateDTO{
		-46548146,
		-23568390,
		"355030885000019",
		"3550308005040",
		87,
		"VILA FORMOSA",
		26,
		"ARICANDUVAS",
		"Leste",
		"Leste 1",
		"1",
		"RUA CODAJ-s",
		"45",
		"VILA FORMOSA",
		"PRAÇA DE TESTE",
	})
	s.Error(err)
}

func (s *ServiceSuite) TestList() {
	expected := []fair.Fair{*fair.NewFair(-46548146,
		-23568390,
		"355030885000019",
		"3550308005040",
		87,
		"VILA FORMOSA",
		26,
		"ARICANDUVAS",
		"Leste",
		"Leste 1",
		"1",
		"RUA CODAJ-s",
		"45",
		"VILA FORMOSA",
		"PRAÇA DE TESTE")}

	s.repo.EXPECT().List().Return(expected, nil)
	fairs, err := s.service.List()
	s.NoError(err)
	s.Equal(len(expected), 1)
	s.Equal(expected, fairs)
}

func (s *ServiceSuite) TestListWithError() {
	s.repo.EXPECT().List().Return(nil, errors.New("not found"))
	_, err := s.service.List()
	s.Error(err)
}

func (s *ServiceSuite) TestFind() {
	expected := *fair.NewFair(-46548146,
		-23568390,
		"355030885000019",
		"3550308005040",
		87,
		"VILA FORMOSA",
		26,
		"ARICANDUVAS",
		"Leste",
		"Leste 1",
		"1",
		"RUA CODAJ-s",
		"45",
		"VILA FORMOSA",
		"PRAÇA DE TESTE")

	s.repo.EXPECT().Find(gomock.Any()).Return(expected, nil)
	fairs, err := s.service.Find(1)
	s.NoError(err)
	s.Equal(expected, fairs)
}

func (s *ServiceSuite) TestFindWithError() {
	s.repo.EXPECT().Find(gomock.Any()).Return(fair.Fair{}, errors.New("not found"))
	_, err := s.service.Find(1)
	s.Error(err)
}

func (s *ServiceSuite) TestDelete() {
	s.repo.EXPECT().Delete(gomock.Any()).Return(nil)
	err := s.service.Delete(1)
	s.NoError(err)
}

func (s *ServiceSuite) TestDeleteWithError() {
	s.repo.EXPECT().Delete(gomock.Any()).Return(errors.New("not found"))
	err := s.service.Delete(1)
	s.Error(err)
}

func (s *ServiceSuite) TestUpdate() {
	expected := *fair.NewFair(-46548146,
		-23568390,
		"355030885000019",
		"3550308005040",
		87,
		"VILA FORMOSA",
		26,
		"ARICANDUVAS",
		"Leste",
		"Leste 1",
		"1",
		"RUA CODAJ-s",
		"45",
		"VILA FORMOSA",
		"PRAÇA DE TESTE")

	s.validator.EXPECT().Validate(gomock.Any()).Times(1)
	s.repo.EXPECT().Find(gomock.Any()).Return(expected, nil)
	s.repo.EXPECT().Update(gomock.Any()).Return(nil)
	_, err := s.service.Update(1, &fair.UpdateDTO{
		-46548146,
		-23568390,
		"355030885000019",
		"3550308005040",
		87,
		"VILA FORMOSA",
		26,
		"ARICANDUVAS",
		"Leste",
		"Leste 1",
		"1",
		"RUA CODAJ-s",
		"45",
		"VILA FORMOSA",
		"PRAÇA DE TESTE",
	})
	s.NoError(err)
}

func (s *ServiceSuite) TestUpdateWithError() {
	s.repo.EXPECT().Find(gomock.Any()).Return(fair.Fair{}, errors.New("not found"))
	s.validator.EXPECT().Validate(gomock.Any()).Times(1)
	_, err := s.service.Update(1, &fair.UpdateDTO{
		-46548146,
		-23568390,
		"355030885000019",
		"3550308005040",
		87,
		"VILA FORMOSA",
		26,
		"ARICANDUVAS",
		"Leste",
		"Leste 1",
		"1",
		"RUA CODAJ-s",
		"45",
		"VILA FORMOSA",
		"PRAÇA DE TESTE",
	})
	s.Error(err)
}

func (s *ServiceSuite) TestFindByName() {
	s.repo.EXPECT().FindByName(gomock.Any()).Return([]fair.Fair{}, nil)
	_, err := s.service.FindByName("test")
	s.NoError(err)
}

func (s *ServiceSuite) TestFindByNameWithError() {
	s.repo.EXPECT().FindByName(gomock.Any()).Return([]fair.Fair{}, errors.New("not found"))
	_, err := s.service.FindByName("test")
	s.Error(err)
}
