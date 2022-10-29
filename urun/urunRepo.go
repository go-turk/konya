package urun

import (
	"errors"
	"gorm.io/gorm"
)

// TODO: Onumuzdeki derste interface ile tanımalam yapmanın faydalarını konuşalım
type UrunRepository struct {
	db *gorm.DB
}

func NewUrunRepository(db *gorm.DB) (*UrunRepository, error) {
	if db == nil {
		return nil, errors.New("db is nil")
	}
	if err := db.AutoMigrate(&Urun{}); err != nil {
		return nil, err
	}
	return &UrunRepository{db: db}, nil
}

func (r *UrunRepository) Create(urun *Urun) error {
	return r.db.Model(&Urun{}).Create(urun).Error
}

func (r *UrunRepository) Delete(urun *Urun) error {
	return r.db.Delete(urun).Error
}

func (r *UrunRepository) FindAll() ([]Urun, error) {
	urunler := []Urun{}
	return urunler, r.db.Find(&urunler).Error
}

func (r *UrunRepository) FindByID(id uint) (*Urun, error) {
	urun := &Urun{}
	return urun, r.db.First(urun, id).Error
}

func (r *UrunRepository) Update(urun *Urun) error {
	return r.db.Save(urun).Error
}
