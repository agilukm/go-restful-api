package middleware

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func AuthApiKey() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		if c.GetHeader("X-API-KEY") != "RAHASIA" {
//
//			c.Header("Content-Type", "application/json")
//			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
//				"Code":   http.StatusUnauthorized,
//				"Status": "UNAUTHORIZED",
//			})
//		} else {
//			c.Next()
//		}
//	}
//
//}
