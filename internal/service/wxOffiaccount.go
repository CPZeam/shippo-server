package service

import (
	"errors"
	"gorm.io/gorm"
	"shippo-server/internal/model"
	"shippo-server/utils/ecode"
)

type WxOffiaccountService struct {
	*Service
}

func NewWxOffiaccountService(s *Service) *WxOffiaccountService {
	return &WxOffiaccountService{s}
}

//查询全部公众号
func (t *WxOffiaccountService) FindAll()(r []model.WxOffiaccount,err error){
	r,err=t.dao.WxOffiaccount.FindAll()
	return
}

//查询公众号根据username
func (t *WxOffiaccountService) Find(username string)(r *model.WxOffiaccount,err error){
	r,err=t.dao.WxOffiaccount.Find(username)
	if errors.Is(err,gorm.ErrRecordNotFound){
		err=ecode.ErrRecordNotFound
	}
	return
}