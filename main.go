package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	docs "application/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()
	app.SetTrustedProxies(nil)

	docs.SwaggerInfo.BasePath = "/"
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.GET("/health", health)

	app.POST("/hello", hello)
	app.Run("0.0.0.0:8000")
}

// @BasePath /health
// HealthTO godoc
// @Summary hello endpoint
// @Schemes
// @Description do health check
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /health [get]
func health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Message: "OK",
	})
}

// @BasePath /hello
// Hello godoc
// @Summary hello endpoint
// @Schemes
// @Description do greeting
// @Tags example
// @Accept json
// @Produce json
// @Param	Body body HelloRequest true "Request Body"
// @Success 200 {object} Response
// @Router /hello [post]
func hello(ctx *gin.Context) {
	var body HelloRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Message: "failed to parse request body",
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Message: fmt.Sprintf("Hello %s", body.Name),
	})
}
