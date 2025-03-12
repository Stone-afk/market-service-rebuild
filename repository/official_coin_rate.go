package repository

import (
	"context"
	"github.com/ecodeclub/ekit/slice"
	"github.com/google/uuid"
	"market-service/domain"
	"market-service/repository/dao"
	"time"
)

var _ OfficialCoinRateRepository = (*officialCoinRateRepository)(nil)

type officialCoinRateRepository struct {
	dao dao.OfficialCoinRateDAO
}

func (repo *officialCoinRateRepository) ListOfficialCoinRatesByAsset(ctx context.Context, asset string) ([]domain.OfficialCoinRate, error) {
	data, err := repo.dao.ListOfficialCoinRatesByAsset(ctx, asset)
	if err != nil {
		return nil, err
	}
	return slice.Map(data, func(idx int, src dao.OfficialCoinRate) domain.OfficialCoinRate {
		return toOfficialCoinRateDomain(src)
	}), nil
}

func (repo *officialCoinRateRepository) CreateOfficialCoinRates(ctx context.Context, officialCoinRates []domain.OfficialCoinRate) error {
	now := time.Now().UnixMilli()
	return repo.dao.CreateOfficialCoinRates(ctx, slice.Map(officialCoinRates, func(idx int, src domain.OfficialCoinRate) dao.OfficialCoinRate {
		res := toOfficialCoinRateEntity(src)
		res.GUID = uuid.New()
		res.CreatedAt = now
		res.UpdatedAt = now
		return res
	}))

}

func NewOfficialCoinRateRepository(dao dao.OfficialCoinRateDAO) OfficialCoinRateRepository {
	return &officialCoinRateRepository{dao: dao}
}

func toOfficialCoinRateEntity(rate domain.OfficialCoinRate) dao.OfficialCoinRate {
	return dao.OfficialCoinRate{
		AssetName: rate.AssetName,
		Price:     rate.Price,
		BaseAsset: rate.BaseAsset,
		Timestamp: rate.Timestamp,
		CreatedAt: rate.CreatedAt,
		UpdatedAt: rate.UpdatedAt,
	}
}

func toOfficialCoinRateDomain(rate dao.OfficialCoinRate) domain.OfficialCoinRate {
	return domain.OfficialCoinRate{
		GUID:      rate.GUID.String(),
		AssetName: rate.AssetName,
		Price:     rate.Price,
		BaseAsset: rate.BaseAsset,
		Timestamp: rate.Timestamp,
		CreatedAt: rate.CreatedAt,
		UpdatedAt: rate.UpdatedAt,
	}
}
