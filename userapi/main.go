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
    r.GET("/user/:id", h.GetUser)
    r.POST("/user", h.AddUser)
    r.DELETE("/user/:id", h.DelUser)
    r.PUT("/user", h.ModUser)

    r.Run(":8080")
}