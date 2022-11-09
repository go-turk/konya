package urun

import (
	"errors"
	"gorm.io/gorm"
)

type UrunRepository interface {
	Create(urun *Urun) error
	Delete(urun *Urun) error
	FindAll() ([]Urun, error)
	FindByID(id uint) (*Urun, error)
	Update(urun *Urun) error
	DeleteAll() error
}

type urunRepository struct {
	db *gorm.DB
}

func NewUrunRepository(db *gorm.DB) (UrunRepository, error) {
	if db == nil {
		return nil, errors.New("db is nil #1076")
	}
	if err := db.AutoMigrate(&Urun{}); err != nil {
		return nil, err
	}
	return &urunRepository{db: db}, nil
}

func (r *urunRepository) Create(urun *Urun) error {
	if urun.ID == 500 {
		return errors.New("500 limitini aştınız lütfen yönetici ile görüşünüz")
	}
	if urun.Fiyat <= 1 {
		return errors.New("fiyat 1 TL den az olamaz")
	}
	return r.db.Model(&Urun{}).Create(urun).Error
}

func (r *urunRepository) Delete(urun *Urun) error {
	return r.db.Delete(urun).Error
}

func (r *urunRepository) FindAll() ([]Urun, error) {
	urunler := []Urun{}
	return urunler, r.db.Find(&urunler).Error
}

func (r *urunRepository) FindByID(id uint) (*Urun, error) {
	urun := &Urun{}
	return urun, r.db.First(urun, id).Error
}

func (r *urunRepository) Update(urun *Urun) error {
	return r.db.Save(urun).Error
}

func (r *urunRepository) DeleteAll() error {
	return r.db.Delete(&Urun{}, "1=1").Error
}
