package yf

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Config *TickerConfig
	Price  float32
}

type TickerConfig struct {
	Ticker   string
	Period   string
	Interval string
}

type TickerRes struct {
	Chart struct {
		Result *[]struct {
			Meta struct {
				Price float32 `json:"regularMarketPrice"`
			} `json:"meta"`
		} `json:"result"`
		Error *struct {
			Description string `json:"description"`
		} `json:"error"`
	} `json:"chart"`
}

type fetchError struct {
	config  *TickerConfig
	message string
}

func (e *fetchError) Error() string {
	return fmt.Sprintf("error fetching %s: %s", *e.config, e.message)
}

func Fetch(config *TickerConfig) (*Data, error) {
	resp, err := http.Get(fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?period=%s&interval=%s", config.Ticker, config.Period, config.Interval))
	if err != nil {
		return nil, &fetchError{
			config:  config,
			message: err.Error(),
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, &fetchError{
			config:  config,
			message: fmt.Sprintf("error fetching ticker data, response code: %d", resp.StatusCode),
		}
	}

	buf := new(TickerRes)
	if err := json.NewDecoder(resp.Body).Decode(buf); err != nil {
		return nil, &fetchError{
			config:  config,
			message: err.Error(),
		}
	}

	if buf.Chart.Error != nil {
		return nil, &fetchError{
			config:  config,
			message: buf.Chart.Error.Description,
		}
	}

	return &Data{Config: config, Price: (*buf.Chart.Result)[0].Meta.Price}, nil
}
