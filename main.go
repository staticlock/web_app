package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/router"
	"web_app/settings"

	"go.uber.org/zap"
)

// go web开发通用脚手架
func main() {
	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("初始化配置失败：%v\n", err)
	} else {
		fmt.Printf("初始化配置成功!\n")
	}
	//2.初始化日志
	if err := logger.Init(settings.Config.LogConfig); err != nil {
		fmt.Printf("初始化日志失败：%v\n", err)
	} else {
		zap.L().Info("初始化日志成功!\n")
	}
	defer zap.L().Sync()
	//3.初始化mysql
	if err := mysql.Init(settings.Config.MysqlConfig); err != nil {
		zap.L().Error("初始化mysql失败:", zap.Error(err))
	} else {
		zap.L().Info("初始化mysql成功:\n")
	}
	defer mysql.DB.Close()
	//4.初始化redis
	if err := redis.Init(settings.Config.RedisConfig); err != nil {
		zap.L().Error("初始化redis失败:", zap.Error(err))
	} else {
		zap.L().Info("初始化redis成功:\n")
	}
	defer redis.Rdb.Close()

	//5.注册路由
	r := router.SetRouters()
	//6.启动服务（优雅关机）
	srv := &http.Server{
		Addr:    settings.Config.Port,
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen: %s\n", zap.Error(err))
		}
	}()
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}
	zap.L().Info("Server exiting")
}
