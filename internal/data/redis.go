package data

import "context"

func (s *shortRepo) setRedisKey(ctx context.Context, key string, value string) error {
	err := s.data.rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		s.log.Infof("create short url map fall to redis database\n")
		return err
	}
	return nil
}
