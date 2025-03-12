package repository

import (
	"context"
	"market-service/domain"
)

type MarketPriceRepository interface {
	ListMarketPricesByAsset(ctx context.Context, asset string) ([]domain.MarketPrice, error)
	CreateMarketPrices(ctx context.Context, marketPrices []domain.MarketPrice) error
}

type OfficialCoinRateRepository interface {
	ListOfficialCoinRatesByAsset(ctx context.Context, asset string) ([]domain.OfficialCoinRate, error)
	CreateOfficialCoinRates(ctx context.Context, officialCoinRates []domain.OfficialCoinRate) error
}
