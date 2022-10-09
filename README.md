## casbin RBAC website demo

### Build demo gin web server

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}
```

### Build basic web server

```go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
```

```go
package routers

import (
	"rbac/controllers"

	"github.com/gin-gonic/gin"
)

func APIRouter(r *gin.Engine) {
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/", controllers.Hello)
	}
}
```

```go
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
```

