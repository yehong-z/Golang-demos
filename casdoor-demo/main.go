package main

import (
	"casdoor_demo/api"
	"fmt"

	auth "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-gonic/gin"
)

// 使用demo网站上已经有的APP，省去注册应用程序的步骤

// 启动一个web服务用于验证是否成功获取到token
func main() {
	api.Init()
	r := gin.Default()
	r.GET("/callback", func(c *gin.Context) {
		code := c.Query("code")
		state := c.Query("state")
		token, _ := auth.GetOAuthToken(code, state)

		acToken := token.AccessToken
		fmt.Println(acToken)
		reToken := token.RefreshToken

		token, _ = auth.RefreshOAuthToken(reToken)

		fmt.Println(token.AccessToken)

		Jwt, _ := auth.ParseJwtToken(token.AccessToken)

		fmt.Println(Jwt.ExpiresAt)
		c.JSON(200, gin.H{
			"message":      "Hello, world!",
			"AccessToken":  token.AccessToken,
			"RefreshToken": token.RefreshToken,
			"TokenType":    token.TokenType,
			"Expiry":       token.Expiry,
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		// 从重定向 URL 的 GET 参数中获取代码和状态
		code := c.Query("code")
		state := c.Query("state")

		// 用代码和状态交换token
		token, err := auth.GetOAuthToken(code, state)
		if err != nil {
			panic(err)
		}

		// 验证访问令牌
		claims, err := auth.ParseJwtToken(token.AccessToken)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(200, gin.H{
			"msg": claims,
		})
	})
	// 启动 HTTP 服务
	err := r.Run(":9000")
	if err != nil {
		return
	}
}
