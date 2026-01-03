package bybit

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

func FetchTickers() (string, error) {
	url := os.Getenv("BYBIT_TICKERS_URL")
	if url == "" {
		return "", errors.New("BYBIT_TICKERS_URL is not set")
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	log.Println("Bybit tickers fetched")
	return string(body), nil
}
