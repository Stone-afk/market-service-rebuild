package service

import (
	"context"
	"fmt"
	"market-service/domain"
	"market-service/repository"
)

var _ MarketPricesService = (*marketPricesService)(nil)

type marketPricesService struct {
	repo repository.MarketPriceRepository
}

func (s *marketPricesService) ListMarketPricesByAsset(ctx context.Context, asset string) ([]domain.MarketPrice, error) {
	res, err := s.repo.ListMarketPricesByAsset(ctx, asset)
	if err != nil {
		return nil, fmt.Errorf("failed to list market prices by asset: %w", err)
	}
	return res, err
}

func (s *marketPricesService) CreateMarketPrices(ctx context.Context, marketPrices []domain.MarketPrice) error {
	err := s.repo.CreateMarketPrices(ctx, marketPrices)
	if err != nil {
		return fmt.Errorf("failed to create market prices: %w", err)
	}
	return nil
}

func NewMarketPricesService(repo repository.MarketPriceRepository) MarketPricesService {
	return &marketPricesService{repo: repo}
}
