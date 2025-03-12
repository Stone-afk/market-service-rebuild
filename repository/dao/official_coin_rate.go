package dao

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type officialCoinRateDAO struct {
	db *gorm.DB
}

func (dao *officialCoinRateDAO) ListOfficialCoinRatesByAsset(ctx context.Context, asset string) ([]OfficialCoinRate, error) {
	var rates []OfficialCoinRate
	err := dao.db.Model(&OfficialCoinRate{}).Find(&rates).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return rates, nil
}

func (dao *officialCoinRateDAO) CreateOfficialCoinRates(ctx context.Context, officialCoinRates []OfficialCoinRate) error {
	result := dao.db.Model(&OfficialCoinRate{}).
		CreateInBatches(&officialCoinRates, len(officialCoinRates))
	return result.Error
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
