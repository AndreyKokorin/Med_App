package router

import (
	"awesomeProject/internal/handlers/appointments"
	"awesomeProject/internal/handlers/logIn"
	"awesomeProject/internal/handlers/logUp"
	"awesomeProject/internal/handlers/med_records"
	"awesomeProject/internal/handlers/users"
	"awesomeProject/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// Аутентификация
	r.POST("/register", logUp.LogUpUser)
	r.POST("/login", logIn.LogIn)

	// Группа для пользователей (только "user")
	userGroup := r.Group("/users", middleware.AuthMiddleware("user"))
	{
		userGroup.GET("/:id", users.GetUserID)
	}

	// Группа для администраторов (только "admin")
	adminGroup := r.Group("/admin", middleware.AuthMiddleware("admin"))
	{
		adminGroup.GET("/users", users.GetAllUsers)
		adminGroup.DELETE("/users/:id", users.DeleteUser)
	}

	doctorGroup := r.Group("/doctor", middleware.AuthMiddleware("doctor", "user"))
	{
		doctorGroup.POST("/newRecord", med_records.NewRecord)
		doctorGroup.GET("/user/:id/records", med_records.GetUserRecords)
		doctorGroup.GET("/record/:id", med_records.GetRecordId)
		doctorGroup.DELETE("/record/:id", med_records.DeleteRecord)
	}

	// Группа с доступом для "user", "admin", "doctor"
	sharedGroup := r.Group("/", middleware.AuthMiddleware("user", "admin", "doctor"))
	{

		// Работа с записями на прием
		appointmentsGroup := sharedGroup.Group("/appointments")
		{
			appointmentsGroup.POST("/", appointments.AddAppointment)
			appointmentsGroup.GET("/:id", appointments.GetAppointment)
			appointmentsGroup.GET("/", appointments.GetAllAppointment)
			appointmentsGroup.GET("/user/:id", appointments.GetUserAppointments)
			appointmentsGroup.PUT("/:id/time", appointments.UpdateAppointDate)
			appointmentsGroup.PUT("/:id/delete", appointments.DeleteAppointment)
		}
	}

	// Группа с доступом для "user", "admin"
	adminUser := r.Group("/adminAndUser", middleware.AuthMiddleware("admin", "user"))
	{
		adminUser.PUT("/userUpdate/:id", users.UpdateUser)
	}

}
