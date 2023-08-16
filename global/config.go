package global

import "time"

type Config struct {
	Server *ServerCfg
	DB     *DBCfg
	Rds    *RdsCfg
}

// 检查是否是DEBUG模式
func (o Config) IsDebug() bool {
	return o.Server.RunMode == "debug"
}

type ServerCfg struct {
	AppName        string
	RunMode        string
	HttpPort       string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	ContextTimeout time.Duration
	LogSavePath    string
	LogFileName    string
	LogFileExt     string
}

type DBCfg struct {
	TablePrefix  string
	DsnWrite     string
	DsnRead      []string
	MaxIdleConns int
	MaxOpenConns int
}

type RdsCfg struct {
	Addr         string
	Username     string
	Password     string
	DB           int
	MinIdleConns int
}
