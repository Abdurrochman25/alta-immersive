package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string `gorm:"<-:false"`
}

type GormUserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *GormUserModel {
	return &GormUserModel{db: db}
}

type UserModel interface {
	GetAll() ([]User, error)
	Get(userId int) (User, error)
	Insert(User) (User, error)
	Edit(user User, userId int) (User, error)
	Delete(userId int) (User, error)
	// Login(email, password string) (User, error)
}

func (m *GormUserModel) GetAll() ([]User, error) {
	var user []User
	if err := m.db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (m *GormUserModel) Get(userId int) (User, error) {
	var user User
	if err := m.db.Find(&user, userId).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (m *GormUserModel) Insert(user User) (User, error) {
	if err := m.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (m *GormUserModel) Edit(newUser User, userId int) (User, error) {
	var user User
	if err := m.db.Find(&user, userId).Error; err != nil {
		return user, err
	}

	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = newUser.Password

	if err := m.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (m *GormUserModel) Delete(userId int) (User, error) {
	var user User
	if err := m.db.Find(&user, userId).Error; err != nil {
		return user, err
	}

	if err := m.db.Delete(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// func (m *GormUserModel) Login(email, password string) (User, error) {
// 	var user User
// 	var err error

// 	if err = m.db.Where("emial = ? AND password = ?", email, password).First(&user).Error; err != nil {
// 		return user, err
// 	}

// 	user.Token, err = middlewares.CreateToken(int(user.ID))

// 	if err != nil {
// 		return user, err
// 	}

// 	if err := m.db.Save(user).Error; err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }
