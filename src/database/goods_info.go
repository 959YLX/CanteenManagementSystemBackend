package database

import (
	"github.com/jinzhu/gorm"
)

// GoodsInfo 商品信息实体
type GoodsInfo struct {
	gorm.Model
	BelongTo string  `gorm:"not null"`
	Species  uint64  `gorm:"not_null"`
	Price    float64 `gorm:"not null;default:0.0"`
	Extra    []byte
}

// SaveGoods 保存商品信息
func SaveGoods(goods GoodsInfo) {
	if client.db.NewRecord(goods) {
		client.db.Create(&goods)
	}
}

// GetGoodsInfo 根据商品ID获取商品信息
func GetGoodsInfo(goodsID uint) (*GoodsInfo, error) {
	result := &GoodsInfo{}
	r := client.db.Where("id = ?", goodsID).First(result)
	if r.Error != nil {
		return nil, r.Error
	}
	return result, nil
}

// GetGoodsList 获取商品列表,belongTo传nil则搜索所有商品
func GetGoodsList(belongTo *string) ([]*GoodsInfo, error) {
	var goods []*GoodsInfo
	var r *gorm.DB
	if belongTo == nil {
		r = client.db.Find(&goods)
	} else {
		r = client.db.Where("belong_to = ?", belongTo).Find(&goods)
	}
	if r.Error != nil {
		return nil, r.Error
	}
	return goods, nil
}
