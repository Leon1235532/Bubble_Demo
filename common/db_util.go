package common

import "gorm.io/gorm"

func RetCountErr(res *gorm.DB) (count int64, err error) {
	if res.Error != nil {
		return 0, res.Error
	}
	count = res.RowsAffected
	return
}
