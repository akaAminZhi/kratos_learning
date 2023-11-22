package service

import (
	"context"
	"fmt"

	pb "shortUrl/api/shortUrl/v1"
	"shortUrl/internal/biz"
	"shortUrl/internal/data/model"
)

type ShortUrlService struct {
	pb.UnimplementedShortUrlServer
	uc *biz.ShortUrlUsecase
}

func NewShortUrlService(uc *biz.ShortUrlUsecase) *ShortUrlService {
	return &ShortUrlService{uc: uc}
}

func (s *ShortUrlService) CreateShortUrl(ctx context.Context, req *pb.CreateShortUrlRequest) (*pb.CreateShortUrlReply, error) {
	shortUrlMap := model.ShortUrlMap{
		Lurl: &req.LongUrl,
	}

	err := s.uc.CreateShortUrl(ctx, &shortUrlMap)
	if err != nil {
		return nil, err
	}
	surl := "akaQZM.com/" + shortUrlMap.Surl
	return &pb.CreateShortUrlReply{ShortUrl: surl}, nil
}
func (s *ShortUrlService) UpdateShortUrl(ctx context.Context, req *pb.UpdateShortUrlRequest) (*pb.UpdateShortUrlReply, error) {

	return &pb.UpdateShortUrlReply{}, nil
}
func (s *ShortUrlService) DeleteShortUrl(ctx context.Context, req *pb.DeleteShortUrlRequest) (*pb.DeleteShortUrlReply, error) {
	return &pb.DeleteShortUrlReply{}, nil
}
func (s *ShortUrlService) GetUrl(ctx context.Context, req *pb.GetUrlRequest) (*pb.GetUrlReply, error) {
	fmt.Println(req.ShortUrl)
	return &pb.GetUrlReply{}, nil
}
