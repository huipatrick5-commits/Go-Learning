package handler

import (
	"net/http"
	"strconv"
	"userapi/repository"
	"userapi/service"
    "userapi/utils"
    "strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
    svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
    return &UserHandler{svc: svc}
}

func (h *UserHandler) GetUser(c *gin.Context) {
    // 从 URL 取参数，转成 int
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        utils.Fail(c, http.StatusBadRequest, "id 必须是数字")
        return
    }

    user, err := h.svc.GetUser(id)
    if err != nil {
        utils.Fail(c, http.StatusNotFound, err.Error())
        return
    }

    utils.Success(c, gin.H{"id": user.ID, "name": user.Name})
}

func (h *UserHandler) AddUser(c *gin.Context) {
    var user repository.User
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.Fail(c, http.StatusBadRequest, err.Error())
        return
    }
    err := h.svc.AddUser(user)
    if err != nil {
        utils.Fail(c, http.StatusNotFound, err.Error())
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "加入成功",
        "id":   user.ID,
        "name": user.Name,
        "age":  user.Age,
    })
}

func (h *UserHandler) DelUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        utils.Fail(c, http.StatusBadRequest, "id 必须是数字")
        return
    }

    err = h.svc.DelUser(id)
    if err != nil {
        utils.Fail(c, http.StatusNotFound, err.Error())
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "删除成功",
        "id": id,
    })
}

func (h *UserHandler) ModUser(c *gin.Context) {
    var user repository.User
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.Fail(c, http.StatusBadRequest, err.Error())
        return
    }
    err := h.svc.ModUser(user)
    if err != nil {
        utils.Fail(c, http.StatusNotFound, err.Error())
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "加入成功",
        "id":   user.ID,
        "name": user.Name,
        "age":  user.Age,
    })
}

func (h *UserHandler) Login(c *gin.Context) {
    var req struct {
        ID int `json:"id"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.Fail(c, http.StatusBadRequest, err.Error())
        return
    }

    // 验证用户存在
    user, err := h.svc.GetUser(req.ID)
    if err != nil {
        utils.Fail(c, http.StatusNotFound, "用户不存在")
        return
    }

    // 生成 token
    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        utils.Fail(c, http.StatusInternalServerError, "生成token失败")
        return
    }

    utils.Success(c, gin.H{
        "token": token,
        "user":  user.Name,
    })
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenStr := c.GetHeader("Authorization")
        if tokenStr == "" {
            utils.Fail(c, http.StatusUnauthorized, "未授权")
            c.Abort()
            return
        }

        // 去掉 "Bearer " 前缀
        tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

        claims, err := utils.ParseToken(tokenStr)
        if err != nil {
            utils.Fail(c, http.StatusUnauthorized, "token无效")
            c.Abort()
            return
        }

        // 把用户ID存到context
        c.Set("userID", claims.UserID)
        c.Next()
    }
}