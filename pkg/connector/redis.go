package connector

import (
	"movie-service/pkg/logger"

	"github.com/redis/go-redis/v9"
)

type RdsOpts struct {
	Uri string
}

func InitRedis(opts RdsOpts) *redis.Client {
	options, err := redis.ParseURL(opts.Uri)
	if err != nil {
		logger.Panic(err)
	}
	rdb := redis.NewClient(options)

	return rdb
}
