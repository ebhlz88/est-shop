package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggerService struct {
	next PriceFetcher
}

func NewLoggerService(next PriceFetcher) *LoggerService {
	return &LoggerService{
		next: next,
	}
}

func (s *LoggerService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestId": ctx.Value("requestId"),
			"timeTaken": time.Since(begin),
			"price":     price,
		}).Info("fetched Successfully")
	}(time.Now())
	return s.next.FetchPrice(ctx, ticker)
}
