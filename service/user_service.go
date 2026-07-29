package service

import (
	"errors"
	"fmt"

	"github.com/Leon1235532/GoTask/auth"
	"github.com/Leon1235532/GoTask/dao"
	"github.com/Leon1235532/GoTask/models"
	"github.com/Leon1235532/GoTask/schemas"
	"gorm.io/gorm"
)

func UserRegister(userinfo *schemas.RegisterRequest) (err error) {
	var count int64
	if count, err = dao.GetCountByUsername(userinfo.Username); err != nil {
		return
	}
	if count > 0 {
		return fmt.Errorf("Username is already occupied.")
	} else {
		var hashPwd []byte
		if hashPwd, err = auth.HashPassword(userinfo.Password); err != nil {
			return fmt.Errorf("Password hashing failed: %w", err)
		}
		user := models.User{
			Username: userinfo.Username,
			Pwd:      string(hashPwd),
		}
		if err = dao.UserCreate(&user); err != nil {
			return
		}
		return nil
	}
}

func UserLogin(userinfo *schemas.LoginRequest) (tokenString string, err error) {
	var user *models.User
	user, err = dao.GetUserByUsername(userinfo.Username)
	if err != nil {
		return "", err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) || !auth.VerifyPwd([]byte(user.Pwd), userinfo.Password) {
		return "", fmt.Errorf("Username don't exists or password is incorrect.")
	}

	tokenString, err = auth.CreateAccToken(user)
	if err != nil {
		return
	}
	return
}

func ModifyPwd(userid uint, pwd schemas.PwdChange) error {
	user, err := dao.GetUserByID(userid)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("user don't exist!")
	}

	if !auth.VerifyPwd([]byte(user.Pwd), pwd.OldPwd) {
		return fmt.Errorf("Incorrect original password entered!")
	}

	if pwd.NewPwd == pwd.OldPwd {
		return fmt.Errorf("The new password cannot be the same as the old password!")
	}
	hash, err := auth.HashPassword(pwd.NewPwd)

	if err != nil {
		return err
	}
	return dao.UpdatePwd(user, string(hash))
}

func LogOutUser(userid uint, inputpwd schemas.LogOutRequest) (user *models.User, err error) {
	if user, err = dao.GetUserByID(userid); err != nil {
		return
	}
	if !auth.VerifyPwd([]byte(user.Pwd), inputpwd.Pwd) {
		return nil, fmt.Errorf("Incorrect password entered!")
	}
	if err = dao.DeleteUser(user); err != nil {
		return
	}
	return
}
