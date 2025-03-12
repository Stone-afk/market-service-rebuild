package repository

import (
	"context"
	"github.com/ecodeclub/ekit/slice"
	"github.com/google/uuid"
	"market-service/domain"
	"market-service/repository/dao"
	"time"
)

var _ MarketPriceRepository = (*marketPriceRepository)(nil)

type marketPriceRepository struct {
	dao dao.MarketPriceDAO
}

func (repo *marketPriceRepository) ListMarketPricesByAsset(ctx context.Context, asset string) ([]domain.MarketPrice, error) {
	data, err := repo.dao.ListMarketPricesByAsset(ctx, asset)
	if err != nil {
		return nil, err
	}
	return slice.Map(data, func(idx int, src dao.MarketPrice) domain.MarketPrice {
		return toMarketPriceDomain(src)
	}), nil
}

func (repo *marketPriceRepository) CreateMarketPrices(ctx context.Context, marketPrices []domain.MarketPrice) error {
	now := time.Now().UnixMilli()
	return repo.dao.CreateMarketPrices(ctx, slice.Map(marketPrices, func(idx int, src domain.MarketPrice) dao.MarketPrice {
		res := toMarketPriceEntity(src)
		res.GUID = uuid.New()
		res.CreatedAt = now
		res.UpdatedAt = now
		return res
	}))
}

func NewMarketPriceRepository(dao dao.MarketPriceDAO) MarketPriceRepository {
	return &marketPriceRepository{dao: dao}
}

func toMarketPriceEntity(marketPrice domain.MarketPrice) dao.MarketPrice {
	return dao.MarketPrice{
		AssetName: marketPrice.AssetName,
		PriceUSDT: marketPrice.AssetPrice,
		Volume:    marketPrice.AssetVolume,
		Rate:      marketPrice.AssetRate,
		Timestamp: marketPrice.Timestamp,
		CreatedAt: marketPrice.CreatedAt,
		UpdatedAt: marketPrice.UpdatedAt,
	}
}

func toMarketPriceDomain(marketPrice dao.MarketPrice) domain.MarketPrice {
	return domain.MarketPrice{
		GUID:        marketPrice.GUID.String(),
		AssetName:   marketPrice.AssetName,
		AssetPrice:  marketPrice.PriceUSDT,
		AssetVolume: marketPrice.Volume,
		AssetRate:   marketPrice.Rate,
		Timestamp:   marketPrice.Timestamp,
		CreatedAt:   marketPrice.CreatedAt,
		UpdatedAt:   marketPrice.UpdatedAt,
	}
}
