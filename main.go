package main

import (
	"rbac/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	routers.APIRouter(r)

	r.Run(":3000")
}
