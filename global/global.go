package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	C      *Config
	DB     *gorm.DB
	RDS    *redis.Client
	Logger *logrus.Logger
)
