// Package strategy XiXiOrchard/internal/strategy/ema_strategy.go
package strategy

import (
	"XiXiOrchard/internal/data"
	"errors"
)

type ExponentialMovingAverageStrategy struct {
	ShortWindow int
	LongWindow  int
	Alpha       float64 // 平滑因子
}

func NewExponentialMovingAverageStrategy(shortWindow, longWindow int, alpha float64) *ExponentialMovingAverageStrategy {
	return &ExponentialMovingAverageStrategy{
		ShortWindow: shortWindow,
		LongWindow:  longWindow,
		Alpha:       alpha,
	}
}

func (s *ExponentialMovingAverageStrategy) Evaluate(marketData []data.MarketData) (signal string, err error) {
	if len(marketData) < s.LongWindow {
		return "", errors.New("not enough market data to evaluate strategy")
	}

	shortEMA := s.calculateEMA(marketData[len(marketData)-s.ShortWindow:])
	longEMA := s.calculateEMA(marketData[len(marketData)-s.LongWindow:])

	if shortEMA > longEMA {
		return "buy", nil
	} else if shortEMA < longEMA {
		return "sell", nil
	} else {
		return "hold", nil
	}
}

func (s *ExponentialMovingAverageStrategy) calculateEMA(data []data.MarketData) float64 {
	ema := data[0].Price
	for i := 1; i < len(data); i++ {
		ema = s.Alpha*data[i].Price + (1-s.Alpha)*ema
	}
	return ema
}
