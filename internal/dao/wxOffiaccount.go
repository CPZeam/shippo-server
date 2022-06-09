package dao

import (
	"shippo-server/internal/model"
)

type WxOffiaccountDao struct {
	*Dao
}

func NewWxOffiaccountDao(s *Dao) *WxOffiaccountDao {
	return &WxOffiaccountDao{s}
}

//FindAll 查询所有公众号
func (t *WxOffiaccountDao) FindAll() (r []model.WxOffiaccount, err error) {
	err = t.db.Model(&model.WxOffiaccount{}).Find(&r).Error
	return
}

//Find 根据username查询公众号
func (t *WxOffiaccountDao) Find(username string) (u *model.WxOffiaccount, err error) {
	err = t.db.Model(&model.WxOffiaccount{}).Where("username", username).Limit(1).Find(&u).Error
	return
}
