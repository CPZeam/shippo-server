package dao

import (
	"shippo-server/internal/model"
)

type WxOffiaccount struct {
	*Dao
}

//FindAll 查询所有公众号
func (t *WxOffiaccount) FindAll() (r []model.WxOffiaccount, err error) {
	err = t.db.Find(&r).Error
	return
}

//Find 根据username查询公众号
func (t *WxOffiaccount) Find(username string) (u model.WxOffiaccount, err error) {
	err = t.db.Model(&model.WxOffiaccount{}).First(&u).Error
	return
}
