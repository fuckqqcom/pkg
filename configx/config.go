package configx

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Config struct {
	Mysql struct {
		Read  sqlx.SqlConn
		Write sqlx.SqlConn
	}
	Cache cache.CacheConf //redis缓存
	Redis redis.Redis
	Opts  cache.Option
}
