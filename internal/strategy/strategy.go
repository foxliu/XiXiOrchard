// Package strategy XiXiOrchard/internal/strategy/strategy.go
package strategy

import "XiXiOrchard/internal/data"

// Strategy 接口定义
type Strategy interface {
	Evaluate(marketData []data.MarketData) (signal string, err error)
}
