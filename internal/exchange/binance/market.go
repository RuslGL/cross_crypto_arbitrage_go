package binance

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

func FetchTickers() (string, error) {
	base := os.Getenv("BINANCE_REST_BASE")
	if base == "" {
		return "", errors.New("BINANCE_REST_BASE is not set")
	}

	url := base + "/api/v3/ticker/bookTicker"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Println("Binance tickers fetched")
	return string(body), nil
}
