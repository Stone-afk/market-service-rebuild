package dao

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	err := db.AutoMigrate(&MarketPrice{}, &OfficialCoinRate{})
	return err
}
