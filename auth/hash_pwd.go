package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) ([]byte, error) {
	password := []byte(pwd)
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

}

func VerifyPwd(hashPwd []byte, pwd string) bool {
	password := []byte(pwd)
	return bcrypt.CompareHashAndPassword(hashPwd, password) == nil
}
