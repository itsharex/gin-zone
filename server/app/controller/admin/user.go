package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (u *User) Login(c *gin.Context) {
	c.String(http.StatusOK, "请求输入用户名")
}
