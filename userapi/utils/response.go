package utils

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Response struct {
    Code int    `json:"code"`
    Msg  string `json:"msg"`
    Data any    `json:"data"`
}

func Success(c *gin.Context, data any) {
    c.JSON(http.StatusOK, Response{
        Code: 0,
        Msg:  "success",
        Data: data,
    })
}

func Fail(c *gin.Context, code int, msg string) {
    c.JSON(http.StatusOK, Response{
        Code: code,
        Msg:  msg,
        Data: nil,
    })
}