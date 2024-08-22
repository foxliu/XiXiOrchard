// Package data XiXiOrchard/internal/data/operations.go
package data

import (
	"gorm.io/gorm"
	"time"
)

// InsertOrderWithMarketData 插入订单及关联的市场数据
func InsertOrderWithMarketData(db *gorm.DB, order Order, marketData []MarketData) error {
	order.MarketData = marketData
	result := db.Create(&order)
	return result.Error
}

// GetOrderWithMarketData 查询某个订单的市场数据
func GetOrderWithMarketData(db *gorm.DB, orderID string) (Order, error) {
	var order Order
	result := db.Preload("MarketData").First(&order, "order_id = ?", orderID)
	return order, result.Error
}

// FetchHistoricalData 使用 GORM 实现 FetchHistoricalData 方法
func FetchHistoricalData(db *gorm.DB, symbol string, startTime, endTime time.Time) ([]MarketData, error) {
	var historicalData []MarketData
	result := db.Where("symbol = ? AND time BETWEEN ? AND ?", symbol, startTime, endTime).
		Order("time ASC").
		Find(&historicalData)
	return historicalData, result.Error
}
