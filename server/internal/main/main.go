package main

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    handler := api.NewHandler("handler")
    handler.Register(r)
    r.Run()
}
