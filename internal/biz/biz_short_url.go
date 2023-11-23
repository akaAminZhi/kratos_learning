package biz

import (
	"context"
	"errors"
	"strings"

	"fmt"
	"shortUrl/internal/data/model"
	"shortUrl/third_party/base62"
	"shortUrl/third_party/connect"
	"shortUrl/third_party/md5"
	"shortUrl/third_party/urltool"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// type ShortUrl struct {
// 	LongUrl string
// }

// GreeterRepo is a Greater repo.
type ShortRepo interface {
	Save(context.Context, *model.ShortUrlMap) error
	Update(context.Context, *model.ShortUrlMap) error
	Get(context.Context, *model.ShortUrlMap) error
	GetShortNum(context.Context) (uint64, error)
	CheckSurlExist(context.Context, string) bool
}

// GreeterUsecase is a Greeter usecase.
type ShortUrlUsecase struct {
	repo ShortRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewShortUsecase(repo ShortRepo, logger log.Logger) *ShortUrlUsecase {
	return &ShortUrlUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *ShortUrlUsecase) CreateShortUrl(ctx context.Context, s *model.ShortUrlMap) error {
	uc.log.WithContext(ctx).Infof("CreateShortUrl: %v", s.Lurl)

	// 1.verify the url can reach
	ok := connect.Get(*(s.Lurl))

	if !ok {
		uc.log.Infof("the url:%v can not reach", s.Lurl)
		return fmt.Errorf("the url:%v can not reach", s.Lurl)
	}

	// 1.1 check if the url is already a short url
	baseUrl, err := urltool.GetBasePath(*s.Lurl)
	if err != nil {
		uc.log.Infof("the url:%v is unormal", s.Lurl)
		return err
	}
	if uc.repo.CheckSurlExist(ctx, baseUrl) {
		return errors.New("this url is already surl")
	}

	// 2.generate md5
	lurl := strings.TrimRight((*s.Lurl), `/`)

	s.Lurl = &lurl
	md5 := md5.GetMd5([]byte(*(s.Lurl)))
	// 2.1 search md5 from db, check if exist
	temp := &model.ShortUrlMap{Md5: md5}
	err = uc.repo.Get(ctx, temp)

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("this url already convert")
	}
	s.Md5 = md5

	// get sequen number from mysql
	seq, _ := uc.repo.GetShortNum(ctx)
	// change to base62 number
	surl := base62.Base10ToBase62(seq)
	s.Surl = surl
	// fmt.Println("surl:", surl)
	return uc.repo.Save(ctx, s)
}

func (uc *ShortUrlUsecase) ShowUrl(ctx context.Context, s *model.ShortUrlMap) error {
	// www.aaa.com/aed/123
	// 1.get base url will return 123

	// 2.generate md5
	lurl := strings.TrimRight((*s.Lurl), `\`)
	s.Lurl = &lurl
	md5 := md5.GetMd5([]byte(*(s.Lurl)))
	// 2.1 search md5 from db, check if exist
	err := uc.repo.Get(ctx, &model.ShortUrlMap{Md5: md5})
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("this url already convert")
	}
	s.Md5 = md5

	// get url from mysql
	seq, _ := uc.repo.GetShortNum(ctx)
	surl := base62.Base10ToBase62(seq)
	s.Surl = surl
	// fmt.Println("surl:", surl)
	return uc.repo.Save(ctx, s)
}
