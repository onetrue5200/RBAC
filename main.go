package main

import (
	"rbac/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.APIRouter(r)

	r.Run(":3000")
}
