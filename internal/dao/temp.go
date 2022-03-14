package dao

import (
	"shippo-server/internal/model"
)

type TempDao struct {
	*Dao
}

func NewTempDao(s *Dao) *TempDao {
	return &TempDao{s}
}

// 根据订单号查询订单
func (d *TempDao) Temp_trade_20220108_findByTradeId(id string) (p model.Temp_trade_20220108, err error) {
	err = d.db.Where("trade_id", id).Limit(1).Find(&p).Error
	return
}

// 根据用户QQ查询订单
func (d *TempDao) Temp_trade_20220108_findByUserQQ(qq string) (p []model.Temp_trade_20220108, err error) {
	err = d.db.Where("user_qq", qq).Find(&p).Error
	return
}

// 创建订单
func (d *TempDao) Temp_trade_20220108_save(p model.Temp_trade_20220108) (model.Temp_trade_20220108, error) {
	return p, d.db.Save(&p).Error
}

// 查询出订单金额 >= 233；订单状态为（0正常）的订单
func (d *TempDao) Temp_trade_20220108_findSuccess() (p []model.Temp_trade_20220108_FindSuccessResult, err error) {
	err = d.db.Model(&model.Temp_trade_20220108{}).Select("user_qq", "sum(trade_amount) as amount").
		Group("user_qq").Where("amount_status", 0).Having("amount>=233").Find(&p).Error
	return
}
