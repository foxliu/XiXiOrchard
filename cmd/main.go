// XiXiOrchard/cmd/main.go
package main

import (
	"XiXiOrchard/config"
	"XiXiOrchard/internal/data"
	"XiXiOrchard/internal/logger"
	"XiXiOrchard/internal/monitoring"
	"XiXiOrchard/internal/strategy"
)

func main() {
	// 初始化日志
	logger.InitLogger()

	// 加载配置文件
	err := config.LoadConfig("config/config.yaml")
	if err != nil {
		logger.ErrorLogger.Fatalf("Could not load config: %v", err)
	}
	logger.InfoLogger.Println("Config loaded successfully")

	// 连接 PostgreSQL 数据库
	db, err := data.ConnectDB()
	if err != nil {
		logger.ErrorLogger.Fatalf("Could not connect to database: %v", err)
	}
	logger.InfoLogger.Println("Connected to database successfully")

	// 初始化策略和回测模块
	smaStrategy := strategy.NewSimpleMovingAverageStrategy(5, 20)
	monitor := monitoring.NewMonitor(false, "", "", "", "", 996)
	backtest := strategy.NewBacktest(smaStrategy, db, monitor)

	// 运行策略回测
	logger.InfoLogger.Println("Starting backtest for AAPL")
	backtest.Run("AAPL")
}
