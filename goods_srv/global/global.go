package global

import (
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"mxshop_srvs/goods_srv/config"
)

var (
	ServerConfig = &config.ServerConfig{}
	NacosConfig  = &config.NacosConfig{}
	DB           *gorm.DB
	EsClient     *elastic.Client
)
