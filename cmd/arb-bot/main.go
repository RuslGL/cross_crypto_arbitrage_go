package main

import (
	"log"
	"time"

	"cross_crypto_arbitrage_go/internal/config"
	"cross_crypto_arbitrage_go/internal/core"
	"cross_crypto_arbitrage_go/internal/exchange/binance"
	"cross_crypto_arbitrage_go/internal/exchange/bybit"
)

func main() {
	log.Println("arb-bot started")

	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("=== Fetch round ===")

		// --- BYBIT ---
		bybitJson, err := bybit.FetchTickers(cfg.BybitTickersURL)
		if err != nil {
			log.Println("Bybit error:", err)
		} else {
			log.Println("Bybit OK, normalizing…")

			symbols := bybit.ExtractSymbols(bybitJson)

			var normalized []string
			for _, s := range symbols {
				normalized = append(normalized, core.BybitNormalize(s))
			}

			log.Println("BybitPairs:", normalized)
		}

		// --- BINANCE ---
		binanceJson, err := binance.FetchTickers(cfg.BinanceRestBase)
		if err != nil {
			log.Println("Binance error:", err)
		} else {
			log.Println("Binance OK, normalizing…")

			symbols := binance.ExtractSymbols(binanceJson)

			var normalized []string
			for _, s := range symbols {
				normalized = append(normalized, core.BinanceNormalize(s))
			}

			log.Println("BinancePairs:", normalized)
		}

		time.Sleep(30 * time.Second)
	}
}
