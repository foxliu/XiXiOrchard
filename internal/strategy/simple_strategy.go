// Package strategy XiXiOrchard/internal/strategy/simple_strategy.go
package strategy

import (
	"XiXiOrchard/internal/data"
	"errors"
)

type SimpleMovingAverageStrategy struct {
	ShortWindow int
	LongWindow  int
}

func NewSimpleMovingAverageStrategy(shortWindow, longWindow int) *SimpleMovingAverageStrategy {
	return &SimpleMovingAverageStrategy{
		ShortWindow: shortWindow,
		LongWindow:  longWindow,
	}
}

func (s *SimpleMovingAverageStrategy) Evaluate(marketData []data.MarketData) (signal string, err error) {
	if len(marketData) < s.LongWindow {
		return "", errors.New("not enough market data to evaluate strategy")
	}

	shortAvg := s.calculateMovingAverage(marketData[len(marketData)-s.ShortWindow:])
	longAvg := s.calculateMovingAverage(marketData[len(marketData)-s.LongWindow:])

	if shortAvg > longAvg {
		return "buy", nil
	} else if shortAvg < longAvg {
		return "sell", nil
	} else {
		return "hold", nil
	}
}

func (s *SimpleMovingAverageStrategy) calculateMovingAverage(data []data.MarketData) float64 {
	var sum float64
	for _, d := range data {
		sum += d.Price
	}
	return sum / float64(len(data))
}
