package service

import (
	"context"
	"fmt"
	"market-service/domain"
	"market-service/repository"
)

var _ OfficialCoinRateService = (*officialCoinRateService)(nil)

type officialCoinRateService struct {
	repo repository.OfficialCoinRateRepository
}

func (s *officialCoinRateService) ListOfficialCoinRatesByAsset(ctx context.Context, asset string) ([]domain.OfficialCoinRate, error) {
	res, err := s.repo.ListOfficialCoinRatesByAsset(ctx, asset)
	if err != nil {
		return nil, fmt.Errorf("failed to list official coin rates by asset: %w", err)
	}
	return res, err
}

func (s *officialCoinRateService) CreateOfficialCoinRates(ctx context.Context, officialCoinRates []domain.OfficialCoinRate) error {
	err := s.repo.CreateOfficialCoinRates(ctx, officialCoinRates)
	if err != nil {
		return fmt.Errorf("failed to create official coin rates: %w", err)
	}
	return nil
}

func NewOfficialCoinRateService(repo repository.OfficialCoinRateRepository) OfficialCoinRateService {
	return &officialCoinRateService{repo: repo}
}
