package middlewares

import (
	"crawlab/model"
	"crawlab/routes"
	"crawlab/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TryGetAuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取token string
		tokenStr := c.GetHeader("Authorization")

		//// 校验token
		user, err := services.CheckToken(tokenStr)
		//
		//// 校验失败，返回错误响应
		if err != nil {
			c.Next()
			return
		}
		c.Set("currentUser", user)

		//// 如果为普通权限，校验请求地址是否符合要求
		//if user.Role == constants.RoleNormal {
		//	if strings.HasPrefix(strings.ToLower(c.Request.URL.Path), "/users") {
		//		c.AbortWithStatusJSON(http.StatusUnauthorized, routes.Response{
		//			Status:  "ok",
		//			Message: "unauthorized",
		//			Error:   "unauthorized",
		//		})
		//		return
		//	}
		//}

		// 校验成功
		c.Next()
	}
}
func AuthorizationMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 获取token string
		tokenStr := c.GetHeader("Authorization")

		//// 校验token
		user, err := services.CheckToken(tokenStr)
		//
		//// 校验失败，返回错误响应
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, routes.Response{
				Status:  "ok",
				Message: "unauthorized",
				Error:   "unauthorized",
			})
			return
		}
		c.Set("currentUser", user)

		//// 如果为普通权限，校验请求地址是否符合要求
		//if user.Role == constants.RoleNormal {
		//	if strings.HasPrefix(strings.ToLower(c.Request.URL.Path), "/users") {
		//		c.AbortWithStatusJSON(http.StatusUnauthorized, routes.Response{
		//			Status:  "ok",
		//			Message: "unauthorized",
		//			Error:   "unauthorized",
		//		})
		//		return
		//	}
		//}

		// 校验成功
		c.Next()
	}
}
