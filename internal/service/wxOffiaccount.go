package service

import (
	"errors"
	"gorm.io/gorm"
	"shippo-server/internal/model"
	"shippo-server/utils/ecode"
)

type WxOffiaccount struct {
	*Service
}

func (t *WxOffiaccount) FindAll()(r []model.WxOffiaccount,err error){
	r,err=t.dao.WxOffiaccount.FindAll()
	return
}

func (t *WxOffiaccount) Find(username string)(r model.WxOffiaccount,err error){
	r,err=t.dao.WxOffiaccount.Find(username)
	if errors.Is(err,gorm.ErrRecordNotFound){
		err=ecode.ErrRecordNotFound
	}
	return
}