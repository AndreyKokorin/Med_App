package router

import (
	"awesomeProject/internal/handlers/appointments"
	"awesomeProject/internal/handlers/logIn"
	"awesomeProject/internal/handlers/logUp"
	"awesomeProject/internal/handlers/users"
	"awesomeProject/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.POST("/register", logUp.LogUpUser)
	r.POST("/login", logIn.LogIn)

	userGroup := r.Group("/users")
	usersMiddleWare := middleware.AuthMiddleware("user")
	userGroup.Use(usersMiddleWare)
	userGroup.GET("/:id", users.GetUserID)

	adminGroup := r.Group("/admin")
	adminMiddleWare := middleware.AuthMiddleware("admin")
	adminGroup.Use(adminMiddleWare)
	adminGroup.GET("/getUsers", users.GetAllUsers)
	adminGroup.DELETE("/delete/:id", users.DeleteUser)

	AllAuthMiddleWare := middleware.AuthMiddleware("user", "admin", "doctor")
	allGroup := r.Group("/users")
	allGroup.Use(AllAuthMiddleWare)
	allGroup.PUT("/:id", users.UpdateUser)
	allGroup.POST("/newAppointment", appointments.AddAppointment)
	allGroup.GET("/newAppointment/:id", appointments.GetAppointment)
	allGroup.GET("/appointments", appointments.GetAllAppointment)
	allGroup.GET("/:id/appointments", appointments.GetUserAppointments)
}
