package data

import (
	"context"
	"fmt"
	"shortUrl/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewShortRepo, NewSequenceDataBase, NewRedis)

// Data .
type Data struct {
	// TODO wrapped database client
	mysqlDb *gorm.DB
	rdb     *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb *redis.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	ctx := context.Background()
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	return &Data{mysqlDb: db, rdb: rdb}, cleanup, nil
}

func NewSequenceDataBase(c *conf.Data, logger log.Logger) *gorm.DB {
	dsn := c.Database.Source
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.NewHelper(logger).Infof("init mysql database err:%v\n", err)
		panic("init database fail")

	}
	// err = db.AutoMigrate(&model.ShortUrlMap{})
	// if err != nil {
	// 	log.NewHelper(logger).Infof("auto migrate mysql database err:%v\n", err)
	// 	panic("auto migrate fail")

	// }
	return db

}

func NewRedis(c *conf.Data, logger log.Logger) *redis.Client {
	addr := c.Redis.Addr

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// err = db.AutoMigrate(&model.ShortUrlMap{})
	// if err != nil {
	// 	log.NewHelper(logger).Infof("auto migrate mysql database err:%v\n", err)
	// 	panic("auto migrate fail")

	// }

	return client

}
