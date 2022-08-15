package repositories

import (
	"github.com/wil-g2/unico-challenge/domain/fair"
	"gorm.io/gorm"
)

type fairRepository struct {
	database *gorm.DB
}

func NewFairRepository(database *gorm.DB) fair.FairRepository {
	return &fairRepository{database}
}

func (r *fairRepository) List() ([]fair.Fair, error) {
	fairs := []fair.Fair{}
	err := r.database.Find(&fairs).Error
	return fairs, err
}

func (r *fairRepository) Find(id int) (fair.Fair, error) {
	fair := fair.Fair{ID: id}
	err := r.database.First(&fair).Error
	return fair, err
}

func (r *fairRepository) Create(fair *fair.Fair) error {
	return r.database.Create(fair).Error
}

func (r *fairRepository) Update(fair *fair.Fair) error {

	return r.database.Save(fair).Error
}

func (r *fairRepository) Delete(id int) error {
	return r.database.Delete(&fair.Fair{}, id).Error
}

// FindByName implements fair.FairRepository
func (r *fairRepository) FindByName(name string) ([]fair.Fair, error) {
	fairs := []fair.Fair{}

	// fair := fair.Fair{FairName: "PC CHARLES MULLER"}
	// err := r.database.Debug().Where("cod_dist = ?", 26).Find(&fair.Fair{}).Error
	r.database.Debug().Where("fair_name LIKE ?", "%CHARLES%").Find(&fairs)
	// fairs = append(fairs, fair)
	return fairs, nil
}
