package main

import (
    "userapi/handler"
    "userapi/repository"
    "userapi/service"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    // 初始化数据库
    db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
    if err != nil {
        panic("数据库连接失败")
    }

    // 自动建表
    db.AutoMigrate(&repository.User{})

    // 从下往上初始化
    repo := repository.NewUserRepository(db)
    svc := service.NewUserService(repo)
    h := handler.NewUserHandler(svc)

    r := gin.Default()
    
    // 不需要认证
    r.POST("/login", h.Login)

    // 需要认证的路由组
    auth := r.Group("/api")
    auth.Use(handler.AuthMiddleware())
    {
        auth.GET("/user/:id", h.GetUser)
        auth.POST("/user", h.AddUser)
        auth.DELETE("/user/:id", h.DelUser)
        auth.PUT("/user", h.ModUser)
    }
    
    r.Run(":8080")
}