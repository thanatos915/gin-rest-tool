package cmd

import (
	"context"
	"fmt"
	"gin-rest-tool/global"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

//定义自己的Writer
type MyWriter struct {
	mlog *logrus.Logger
}

//实现gorm/logger.Writer接口
func (m *MyWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	//利用loggus记录日志
	m.mlog.Info(logstr)
}

func NewMyWriter() *MyWriter {
	log := logrus.New()
	//配置logrus
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return &MyWriter{mlog: log}
}

// 初始化应用
func InitApp() {
	// 解析配置文件
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Load App Config Error: %v", err)
	}

	c := new(global.Config)
	err = cfg.MapTo(c)
	if err != nil {
		log.Fatalf("Load Database Config Error: %v", err)
	}
	c.Server.ReadTimeout *= time.Second
	c.Server.WriteTimeout *= time.Second
	c.Server.ContextTimeout *= time.Second
	global.C = c

	// 初始化日志功能
	initLogger()
	initDB()
	initRds()
}

// 系统日志
func initLogger() {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	fileName := global.C.Server.LogSavePath + "/" + global.C.Server.LogFileName + global.C.Server.LogFileExt
	l.SetOutput(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    50,
		MaxAge:     10,
		MaxBackups: 30,
		LocalTime:  true,
	})

	if global.C.IsDebug() {
		l.SetLevel(logrus.DebugLevel)
	} else {
		l.SetLevel(logrus.ErrorLevel)
	}

	global.Logger = l
}

// 初始化数据库
func initDB() {
	var err error

	// init database
	cof := logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	}

	var dbLog logger.Interface

	if global.C.IsDebug() {
		cof.LogLevel = logger.Info
		dbLog = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), cof)
	} else {
		cof.LogLevel = logger.Error
		dbLog = logger.New(NewMyWriter(), cof)
	}

	db, err := gorm.Open(mysql.Open(global.C.DB.DsnWrite), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 dbLog,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   global.C.DB.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("Init Database Error: %v", err)
	}
	global.DB = db
}

// 初始化Redis
func initRds() {
	global.RDS = redis.NewClient(&redis.Options{
		Addr:         global.C.Rds.Addr,
		Username:     global.C.Rds.Username,
		Password:     global.C.Rds.Password,
		DB:           global.C.Rds.DB,
		MinIdleConns: global.C.Rds.MinIdleConns,
		MaxRetries:   3,
	})
	_, err := global.RDS.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Init RDS Error: redis not config")
		return
	}
}
