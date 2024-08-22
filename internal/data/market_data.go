// Package data XiXiOrchard/internal/data/market_data.go
package data

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"XiXiOrchard/config"
)

func FetchRealTimeData(symbol string) (*MarketData, error) {
	// 构建 API 请求 URL
	url := fmt.Sprintf("%s/realtime?symbol=%s&api_key=%s", config.Cfg.MarketAPI, symbol, config.Cfg.APIKey)

	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch real-time data: %s", resp.Status)
	}

	// 解析 JSON 响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data MarketData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	data.Time = time.Now() // 将时间设置为当前时间
	return &data, nil
}
