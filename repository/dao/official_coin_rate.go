package dao

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type officialCoinRateDAO struct {
	db *gorm.DB
}

func (o officialCoinRateDAO) ListOfficialCoinRatesByAsset(ctx context.Context, asset string) ([]OfficialCoinRate, error) {
	//TODO implement me
	panic("implement me")
}

func (o officialCoinRateDAO) CreateOfficialCoinRates(ctx context.Context, officialCoinRates []OfficialCoinRate) error {
	//TODO implement me
	panic("implement me")
}

func NewOfficialCoinRateDAO(db *gorm.DB) OfficialCoinRateDAO {
	return &officialCoinRateDAO{db: db}
}

type OfficialCoinRate struct {
	GUID      uuid.UUID `gorm:"primaryKey;type:varchar(255);column:guid"`
	AssetName string    `gorm:"type:varchar(255);column:asset_name;not null;index:idx_official_coin_rate_asset_name"`
	BaseAsset string    `gorm:"type:varchar(255);column:base_asset;not null"`
	Price     string    `gorm:"type:varchar(255);column:price;not null"`
	Timestamp int64     `gorm:"type:integer;column:timestamp;not null;check:timestamp > 0"`
	CreatedAt int64     `gorm:"type:integer;column:created_at;not null"`
	UpdatedAt int64     `gorm:"type:integer;column:updated_at;not null"`
}

func (OfficialCoinRate) TableName() string {
	return "official_coin_rate"
}
