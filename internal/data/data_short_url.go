package data

import (
	"context"
	"shortUrl/internal/biz"
	"shortUrl/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type shortRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewShortRepo(data *Data, logger log.Logger) biz.ShortRepo {
	return &shortRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *shortRepo) Save(ctx context.Context, model *model.ShortUrlMap) error {
	// s.data.mysqlDb.Exec("REPLACE INTO sequence (stub) VALUES ('a');")
	// fmt.Println(model)
	// s.data.mysqlDb.First(model.Sequence{})
	err := s.data.mysqlDb.Create(&model).Error
	if err != nil {
		s.log.Infof("create short url map fall to database\n")
		return err
	}
	return nil
}

func (s *shortRepo) Update(ctx context.Context, model *model.ShortUrlMap) error {
	return nil
}

func (s *shortRepo) Get(ctx context.Context, model *model.ShortUrlMap) error {

	return s.data.mysqlDb.First(model).Error
}

func (s *shortRepo) GetShortNum(context.Context) (uint64, error) {
	var id []uint64

	result := s.data.mysqlDb.Exec("REPLACE INTO sequence (stub) VALUES ('a');")
	err := result.Error
	if err != nil {
		s.log.Infof("fall to get sequence number\n")
		return 0, err
	}
	s.data.mysqlDb.Raw("SELECT LAST_INSERT_ID();").Pluck("id", &id)

	return id[0], nil
}
