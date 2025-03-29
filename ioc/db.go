package ioc

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"market-service/pkg/retry"
	"market-service/repository/dao"
)

func InitDB(sourceType string) *gorm.DB {
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
			panic(fmt.Errorf("failed to connect to database: %w", er))
		}
		return db, nil
	})
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
