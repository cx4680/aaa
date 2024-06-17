package main

import (
	"github.com/joho/godotenv"
	"github.com/opentrx/seata-golang/v2/pkg/util/log"
	"guangfa-fund/internal/api"
	"guangfa-fund/internal/web"
)

func main() {
	//加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Errorf("Error loading .env file: %v", err)
	}
	//启动项目
	web.InitGinEngine(api.Router)
}
