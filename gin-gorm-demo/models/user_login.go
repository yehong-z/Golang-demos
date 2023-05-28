package models

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

type UserLogin struct {
	gorm.Model
	Username string
	Password string
}

type UserLoginDAO struct{}

var (
	userLoginDao  *UserLoginDAO
	userLoginOnce sync.Once
)

func NewUserLoginDao() *UserLoginDAO {
	userLoginOnce.Do(func() {
		userLoginDao = new(UserLoginDAO)
	})
	return userLoginDao
}

func (u *UserLoginDAO) QueryUserLogin(username, password string, login *UserLogin) error {
	if login == nil {
		return errors.New("UserLogin == nil")
	}
	DB.Where("username=? and password=?", username, password).First(login)
	if login.ID == 0 {
		return errors.New("用户名或密码错误")
	}
	return nil
}

func (u *UserLoginDAO) IsUserExistByUsername(username string) bool {
	var userLogin UserLogin
	DB.Where("username=?", username).First(&userLogin)
	if userLogin.ID == 0 {
		return false
	}
	return true
}

func (u *UserLoginDAO) AddUser(userinfo *UserLogin) error {
	if userinfo == nil {
		return errors.New("userinfo is null")
	}
	return DB.Create(userinfo).Error
}
