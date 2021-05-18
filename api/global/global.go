package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	GVA_DB    *gorm.DB
	GVA_REDIS *redis.Client
)
