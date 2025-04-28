package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(ctx context.Context, ticker string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetch(ticker)
}

var PriceMocks = map[string]float64{
	"BTC": 5.44,
	"MTC": 3.333,
}

func MockPriceFetch(ticker string) (float64, error) {
	price, ok := PriceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("error %s", ticker)
	}
	return price, nil
}
