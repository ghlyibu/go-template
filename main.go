package main

import (
	"fmt"
	"go-template/controller"
	"go-template/dao/mysql"
	"go-template/dao/redis"
	"go-template/logger"
	"go-template/pkg/snowflake"
	"go-template/routes"
	"go-template/setting"
)

func main() {
	// 1. 加载配置
	if err := setting.Init("./config/dev.yaml"); err != nil {
		fmt.Printf("init setting failed, err:%v\n", err)
		return
	}
	//if len(os.Args) < 2 {
	//	fmt.Println("init config use ./config/dev.yaml")
	//
	//} else if err := setting.Init(os.Args[1]); err != nil {
	//	fmt.Printf("init setting failed, err:%v\n", err)
	//	return
	//}
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

	if err := snowflake.Init(setting.Conf.AppLConfig.StartTime, setting.Conf.AppLConfig.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// 初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}
	// 5. 注册路由
	r := routes.SetupRouter(setting.Conf.AppLConfig.Mode)
	// 6. 启动服务
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.AppLConfig.Port)); err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
