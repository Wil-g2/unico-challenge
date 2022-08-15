package fair

//go:generate go run github.com/golang/mock/mockgen --source=interface.go --destination=mocks/mock_fair.go
type Service interface {
	List() ([]Fair, error)
	Find(id int) (Fair, error)
	FindByName(name string) ([]Fair, error)
	Create(input *CreateDTO) error
	Update(id int, input *UpdateDTO) (Fair, error)
	Delete(id int) error
}

type FairRepository interface {
	List() ([]Fair, error)
	Find(id int) (Fair, error)
	FindByName(name string) ([]Fair, error)
	Create(fair *Fair) error
	Update(fair *Fair) error
	Delete(id int) error
}
