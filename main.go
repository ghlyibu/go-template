package main

import (
	"fmt"
	"go-template/dao/mysql"
	"go-template/dao/redis"
	"go-template/logger"
	"go-template/routes"
	"go-template/setting"
	"os"
)

func main() {
	// 1. 加载配置
	if len(os.Args) < 2 {
		fmt.Println("init config use ./config/dev.yaml")
		if err := setting.Init("./config/dev.yaml"); err != nil {
			fmt.Printf("init setting failed, err:%v\n", err)
			return
		}
		return
	} else if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("init setting failed, err:%v\n", err)
		return
	}
	// 2. 初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// 3. 初始化MySQL连接
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	// 4. 初始化Redis连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	// 5. 注册路由
	r := routes.SetupRouter(setting.Conf.AppLConfig.Mode)
	// 6. 启动服务
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.AppLConfig.Port)); err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
