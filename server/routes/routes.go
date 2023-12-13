package routes

import (
	"gin/server/controllers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleRequests() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {

		requestID := uuid.New()

		c.Writer.Header().Set("X-Request-ID", requestID.String())

		c.Set("RequestID", requestID)

		c.Next()

	})

	//teste

	r.GET("api/v1/ping", controllers.Ping)

	// usuarios

	r.POST("api/v1/users", controllers.CreateUser)

	r.GET("api/v1/users", controllers.GetUser)

	r.PUT("api/v1/users", controllers.UpdateUser)

	r.DELETE("api/v1/users", controllers.DeleteUser)

	r.Run()

}
