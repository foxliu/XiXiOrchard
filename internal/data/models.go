// Package data XiXiOrchard/internal/data/models.go
package data

import (
	"gorm.io/gorm"
	"time"
)

type MarketData struct {
	gorm.Model
	Symbol  string    `gorm:"size:10;not null"`
	Price   float64   `gorm:"not null"`
	Volume  float64   `gorm:"not null"`
	Time    time.Time `gorm:"not null"`
	OrderID uint      // 关联字段
}

type Order struct {
	gorm.Model
	OrderID    string       `gorm:"size:36;not null;unique"`
	OrderType  string       `gorm:"size:20;not null"`
	Quantity   float64      `gorm:"not null"`
	MarketData []MarketData `gorm:"foreignKey:OrderID"`
}
