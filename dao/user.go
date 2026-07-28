package dao

import (
	"errors"
	"fmt"

	"github.com/Leon1235532/Bubble_Demo/models"
	"gorm.io/gorm"
)

// 通过username查count
func GetCountByUsername(username string) (count int64, err error) {
	err = DB.Model(&models.User{}).
		Where("username = ?", username).
		Count(&count).Error
	return
}

// 用户注册
func UserCreate(user *models.User) (err error) {
	err = DB.Create(user).Error
	return
}

// 用户登录
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("User information not found!")
		}
		return nil, err
	}
	return &user, nil
}

// userid查用户对象
func GetUserByID(userid uint) (*models.User, error) {
	var user models.User
	err := DB.Where("id = ?", userid).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("User information not found.")
		}
		return nil, err
	}
	return &user, nil
}

// 修改密码
func UpdatePwd(user *models.User, hashpwd string) (err error) {
	err = DB.Model(user).Update("pwd", hashpwd).Error
	return
}

// 注销用户
func DeleteUser(user *models.User) (err error) {
	err = DB.Delete(user).Error
	return
}
