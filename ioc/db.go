package ioc

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"market-service/pkg/retry"
)

func NewDB(sourceType string) (*gorm.DB, error) {
	type Config struct {
		DSN string `yaml:"dsn"`
	}
	c := Config{
		DSN: "root:root@tcp(localhost:13316)/market",
	}
	key := fmt.Sprintf("db.%s", sourceType)
	err := viper.UnmarshalKey(key, &c)
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	db, err := retry.Do[*gorm.DB](context.Background(), 10, retryStrategy, func() (*gorm.DB, error) {
		db, er := gorm.Open(postgres.Open(c.DSN))
		if er != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", er)
		}
		return db, nil
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
