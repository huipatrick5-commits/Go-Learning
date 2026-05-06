package main

import (
	"userapi/config"
	"userapi/handler"
	"userapi/repository"
	"userapi/service"
	"userapi/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"fmt"
)

func main() {
    // 加载配置
    cfg, err := config.Load()
    if err != nil {
        panic("配置文件加载失败: " + err.Error())
    }

    utils.InitJWT(cfg.JWT.Secret, cfg.JWT.Expire)

    // 用配置初始化数据库
    db, err := gorm.Open(sqlite.Open(cfg.Database.Path), &gorm.Config{})
    if err != nil {
        panic("数据库连接失败")
    }

    db.AutoMigrate(&repository.User{})

    repo := repository.NewUserRepository(db)
    svc := service.NewUserService(repo)
    h := handler.NewUserHandler(svc)

    r := gin.Default()

    r.POST("/login", h.Login)

    auth := r.Group("/api")
    auth.Use(handler.AuthMiddleware())
    {
        auth.GET("/user/:id", h.GetUser)
        auth.POST("/user", h.AddUser)
        auth.DELETE("/user/:id", h.DelUser)
        auth.PUT("/user", h.ModUser)
    }

    // 用配置的端口启动
    r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}