package handler

import (
	"net/http"
	"strconv"
	"userapi/repository"
	"userapi/service"

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
        c.JSON(http.StatusBadRequest, gin.H{"error": "id 必须是数字"})
        return
    }

    user, err := h.svc.GetUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "id":   user.ID,
        "name": user.Name,
        "age":  user.Age,
    })
}

func (h *UserHandler) AddUser(c *gin.Context) {
    var user repository.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := h.svc.AddUser(user)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
        c.JSON(http.StatusBadRequest, gin.H{"error": "id 必须是数字"})
        return
    }

    err = h.svc.DelUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := h.svc.ModUser(user)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "加入成功",
        "id":   user.ID,
        "name": user.Name,
        "age":  user.Age,
    })
}