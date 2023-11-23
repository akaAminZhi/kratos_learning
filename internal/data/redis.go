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

func (s *shortRepo) getRedisValue(ctx context.Context, key string) (string, error) {
	value, err := s.data.rdb.Get(ctx, key).Result()
	if err != nil {
		s.log.Infof("get short url map fall to redis database\n")
		return "", err
	}
	return value, nil
}
