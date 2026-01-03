package core

import "strings"

// normalizeToCommon приводит название пары к нашему стандарту
// общий принцип — убрать разделители и привести к верхнему регистру
func normalizeToCommon(s string) string {
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "_", "")
	s = strings.ReplaceAll(s, "/", "")
	return s
}

// BybitNormalize -> к нашему формату
func BybitNormalize(s string) string {
	return normalizeToCommon(s)
}

// BinanceNormalize -> к нашему формату
func BinanceNormalize(s string) string {
	return normalizeToCommon(s)
}
