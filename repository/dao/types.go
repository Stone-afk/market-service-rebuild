package dao

import "context"

type MarketPriceDAO interface {
	ListMarketPricesByAsset(ctx context.Context, asset string) ([]MarketPrice, error)
	CreateMarketPrices(ctx context.Context, marketPrices []MarketPrice) error
}

type OfficialCoinRateDAO interface {
	ListOfficialCoinRatesByAsset(ctx context.Context, asset string) ([]OfficialCoinRate, error)
	CreateOfficialCoinRates(ctx context.Context, officialCoinRates []OfficialCoinRate) error
}
