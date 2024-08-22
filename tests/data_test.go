// XiXiOrchard/tests/data_test.go
package tests

import (
	"XiXiOrchard/internal/data"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestFetchHistoricalData(t *testing.T) {
	// 使用 SQLite 内存数据库进行测试
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect database: %v", err)
	}

	// 自动迁移数据库结构
	err = db.AutoMigrate(&data.MarketData{})
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	// 插入测试数据
	now := time.Now()
	db.Create(&data.MarketData{Symbol: "AAPL", Price: 150, Volume: 100, Time: now.Add(-48 * time.Hour)})
	db.Create(&data.MarketData{Symbol: "AAPL", Price: 155, Volume: 150, Time: now.Add(-24 * time.Hour)})
	db.Create(&data.MarketData{Symbol: "AAPL", Price: 160, Volume: 200, Time: now})

	// 测试 FetchHistoricalData 方法
	startTime := now.Add(-2 * 24 * time.Hour)
	endTime := now

	result, err := data.FetchHistoricalData(db, "AAPL", startTime, endTime)
	if err != nil {
		t.Fatalf("Failed to fetch historical data: %v", err)
	}

	if len(result) != 3 {
		t.Errorf("Expected 3 records, got %d", len(result))
	}
}
