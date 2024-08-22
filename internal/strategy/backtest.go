// XiXiOrchard/internal/strategy/backtest.go
package strategy

import (
	"XiXiOrchard/internal/data"
	"XiXiOrchard/internal/monitoring"
	"gorm.io/gorm"
	"log"
	"time"
)

type Backtest struct {
	Strategy Strategy
	DB       *gorm.DB
	Monitor  *monitoring.Monitor
}

func NewBacktest(strategy Strategy, db *gorm.DB, monitor *monitoring.Monitor) *Backtest {
	return &Backtest{
		Strategy: strategy,
		DB:       db,
		Monitor:  monitor,
	}
}

func (b *Backtest) Run(symbol string) {
	marketData, err := data.FetchHistoricalData(b.DB, symbol, time.Now().AddDate(0, -1, 0), time.Now())
	if err != nil {
		log.Fatalf("Error fetching historical data: %v", err)
	}

	signal, err := b.Strategy.Evaluate(marketData)
	if err != nil {
		log.Fatalf("Error evaluating strategy: %v", err)
	}

	log.Printf("Strategy signal for %s: %s", symbol, signal)

	prices := extractPrices(marketData)
	b.Monitor.MonitorResults(prices)
}

func extractPrices(marketData []data.MarketData) []float64 {
	prices := make([]float64, len(marketData))
	for i, datum := range marketData {
		prices[i] = datum.Price
	}
	return prices
}
