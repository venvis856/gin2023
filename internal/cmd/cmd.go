package cmd

import (
	"context"
	"fmt"
	"gin/internal/cmd/system/cfgstruct"
	"gin/internal/cmd/system/process"
	"gin/internal/config"
	"gin/internal/library/helper"
	"gin/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

var (
	configFile string
	MainCmd    = &cobra.Command{
		Use:   "Main",
		Short: "Main",
	}
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "运行程序system_api",
		RunE:  Run,
	}
	appCmd = &cobra.Command{
		Use:   "app",
		Short: "运行程序app",
		RunE:  App,
	}
	setupCmd = &cobra.Command{
		Use:         "setup",
		Short:       "初始话配置文件",
		RunE:        cmdSetup,
		Annotations: map[string]string{"type": "setup"},
	}
	runConfig   config.Config
	setupConfig config.Config
)

func InitSystem() {
	defaultConfig := helper.ApplicationAbsFileDir(process.DefaultCfgFilename)
	cfgstruct.SetupFlag(MainCmd, &configFile, "consts", defaultConfig, "配置文件")
	//根据环境读取默认配置
	defaults := cfgstruct.EnvsFlag(MainCmd)
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//程序目录
	rootDir := cfgstruct.ConfigVar("ROOT", currentDir)
	//设置系统的HOME变量
	envHome := cfgstruct.ConfigVar("HOME", os.Getenv("HOME"))
	MainCmd.AddCommand(runCmd)
	MainCmd.AddCommand(appCmd)
	MainCmd.AddCommand(setupCmd)
	process.Bind(runCmd, &runConfig, defaults, cfgstruct.ConfigFile(configFile), envHome, rootDir)
	process.Bind(appCmd, &runConfig, defaults, cfgstruct.ConfigFile(configFile), envHome, rootDir)
	process.Bind(setupCmd, &setupConfig, defaults, cfgstruct.ConfigFile(configFile), envHome, cfgstruct.SetupMode(), rootDir)

	process.Exec(MainCmd)

}

func Run(cmd *cobra.Command, args []string) (err error) { //禁用控制台颜色
	fmt.Println(config.Version)
	gin.DisableConsoleColor()
	//设置模式
	//gin.DebugMode、gin.ReleaseMode、gin.TestMode  debug release test
	if cfgstruct.DefaultsType() == cfgstruct.DefaultsRelease {
		gin.SetMode(gin.ReleaseMode)
	}
	// 初始化全局配置, 出错会抛出panic
	InitServer(&runConfig)
	//gin实例
	ginDefault := gin.Default()
	//设置静态资源
	ginDefault.Static("/public", "public")
	ginDefault.Static("/resource", "resource")
	//路由
	router.InitRoutes(ginDefault)
	//优雅重启
	srv := &http.Server{
		Addr:           runConfig.Api.Server,
		Handler:        ginDefault,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("运行的服务端口为:", runConfig.Api.Server)
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
		return err
	}
	log.Println("Server exiting")
	return nil
}

func App(cmd *cobra.Command, args []string) (err error) { //禁用控制台颜色
	fmt.Println(config.Version)
	gin.DisableConsoleColor()
	//设置模式
	//gin.DebugMode、gin.ReleaseMode、gin.TestMode  debug release test
	if cfgstruct.DefaultsType() == cfgstruct.DefaultsRelease {
		gin.SetMode(gin.ReleaseMode)
	}
	// 初始化全局配置, 出错会抛出panic
	InitServer(&runConfig)
	//gin实例
	ginDefault := gin.Default()
	//设置静态资源
	ginDefault.Static("/public", "public")
	ginDefault.Static("/resource", "resource")
	//路由
	router.InitRoutes(ginDefault)
	//优雅重启
	srv := &http.Server{
		Addr:           runConfig.Api.Server,
		Handler:        ginDefault,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("运行的服务端口为:", runConfig.Api.Server)
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
		return err
	}
	log.Println("Server exiting")
	return nil
}

func cmdSetup(cmd *cobra.Command, args []string) error {
	return process.SaveConfig(cmd, configFile)
}
