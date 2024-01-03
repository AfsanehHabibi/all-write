package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) Register(v1 *gin.Engine) {
	v1.POST("/login", gin.BasicAuth(gin.Accounts{
        "admin": "secret",
    }),login)

	v1.GET("/resource", func(c *gin.Context) {
        bearerToken := c.Request.Header.Get("Authorization")
        reqToken := strings.Split(bearerToken, " ")[1]
        claims := &Claims{}
        tkn, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })
        if err != nil {
            if err == jwt.ErrSignatureInvalid {
                c.JSON(http.StatusUnauthorized, gin.H{
                    "message": "unauthorized",
                })
                return
            }
            c.JSON(http.StatusBadRequest, gin.H{
                "message": "bad request",
            })
            return
        }
        if !tkn.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{
                "message": "unauthorized",
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "data": "resource data",
        })
    })
	//lists := v1.Group("/lists")
}
