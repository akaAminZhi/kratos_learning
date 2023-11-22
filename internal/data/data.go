package data

import (
	"fmt"
	"shortUrl/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewShortRepo, NewSequenceDataBase)

// Data .
type Data struct {
	// TODO wrapped database client
	mysqlDb *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{mysqlDb: db}, cleanup, nil
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
