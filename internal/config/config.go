package config

import (
	"errors"
	"os"
)

type Config struct {
	BybitTickersURL string
	BinanceRestBase string
}

var Cfg Config

// LoadEnv — загружает переменные окружения (без .env библиотек)
func LoadEnv() error {
	bybitURL := os.Getenv("BYBIT_TICKERS_URL")
	if bybitURL == "" {
		return errors.New("env BYBIT_TICKERS_URL is required")
	}

	binanceBase := os.Getenv("BINANCE_REST_BASE")
	if binanceBase == "" {
		return errors.New("env BINANCE_REST_BASE is required")
	}

	Cfg = Config{
		BybitTickersURL: bybitURL,
		BinanceRestBase: binanceBase,
	}

	return nil
}
