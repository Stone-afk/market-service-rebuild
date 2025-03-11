package dao

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type marketPriceDAO struct {
	db *gorm.DB
}

func (dao *marketPriceDAO) ListMarketPricesByAsset(ctx context.Context, asset string) ([]MarketPrice, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *marketPriceDAO) CreateMarketPrices(ctx context.Context, marketPrices []MarketPrice) error {
	//TODO implement me
	panic("implement me")
}

func NewMarketPriceDAO(db *gorm.DB) MarketPriceDAO {
	return &marketPriceDAO{db: db}
}

type MarketPrice struct {
	GUID      uuid.UUID `gorm:"primaryKey;type:varchar(255);column:guid"`
	AssetName string    `gorm:"type:varchar(255);column:asset_name;not null;index:idx_market_asset_name"`
	PriceUSDT string    `gorm:"type:varchar(255);column:price_usdt;not null"`
	Volume    string    `gorm:"type:varchar(255);column:volume;not null"`
	Rate      string    `gorm:"type:varchar(255);column:rate;not null"`
	Timestamp int64     `gorm:"type:integer;column:timestamp;not null;check:timestamp > 0"`
	CreatedAt int64     `gorm:"type:integer;column:created_at;not null"`
	UpdatedAt int64     `gorm:"type:integer;column:updated_at;not null"`
}

func (MarketPrice) TableName() string {
	return "market_price"
}
