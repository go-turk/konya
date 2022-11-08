package user

import (
	"errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *User) error
	Delete(user *User) error
	FinAll() ([]User, error)
	FindByID(id uint) (*User, error)
	Update(user *User) error
	DeleteAll() error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (UserRepository, error) {
	if db == nil {
		return nil, errors.New("DB is nil!")
	}
	//Users tablosu yoksa bu aşamada oluşturulur.
	if err := db.AutoMigrate(&User{}); err != nil {
		return nil, err
	}
	return &userRepository{db: db}, nil
}

func (u *userRepository) Create(user *User) error {
	return u.db.Model(&User{}).Create(user).Error
}

func (u *userRepository) Delete(user *User) error {
	return u.db.Delete(user).Error
}

func (u *userRepository) FinAll() ([]User, error) {
	users := []User{}
	return users, u.db.Find(&users).Error
}

func (u *userRepository) FindByID(id uint) (*User, error) {
	user := &User{}
	return user, u.db.First(user, id).Error
}

func (u *userRepository) Update(user *User) error {
	return u.db.Save(user).Error
}

func (u *userRepository) DeleteAll() error {
	return u.db.Delete(&User{}, "1=1").Error
}
