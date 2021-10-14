package global

import (
	"gorm.io/gorm"
	"mxshop_srvs/goods_srv/config"
)

var (
	ServerConfig = &config.ServerConfig{}
	NacosConfig  = &config.NacosConfig{}
	DB           *gorm.DB
)
