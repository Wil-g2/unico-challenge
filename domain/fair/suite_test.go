package fair_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/wil-g2/unico-challenge/domain/fair"
	mock_fair "github.com/wil-g2/unico-challenge/domain/fair/mocks"
	mock_validator "github.com/wil-g2/unico-challenge/domain/shared/mocks"
)

type ServiceSuite struct {
	suite.Suite
	*require.Assertions
	ctrl      *gomock.Controller
	repo      *mock_fair.MockFairRepository
	validator *mock_validator.MockValidator
	service   fair.Service
}

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_fair.NewMockFairRepository(s.ctrl)
	s.validator = mock_validator.NewMockValidator(s.ctrl)
	s.service = fair.NewService(s.repo, s.validator)
}

func (s *ServiceSuite) TearDownTest() {
	s.ctrl.Finish()
}
