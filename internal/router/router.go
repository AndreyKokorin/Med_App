package router

import (
	"awesomeProject/internal/handlers/appointments"
	"awesomeProject/internal/handlers/logIn"
	"awesomeProject/internal/handlers/logUp"
	"awesomeProject/internal/handlers/med_records"
	"awesomeProject/internal/handlers/schedules"
	"awesomeProject/internal/handlers/users"
	"awesomeProject/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(r *gin.Engine) {
	r.Use(Cors)

	r.POST("/logup", logUp.LogUpUser)
	r.POST("/login", logIn.LogIn)
	r.POST("/sendEmailCode", users.ChangePasswordSendEmail)
	r.POST("/changePassword", users.ChangePassword)
	r.POST("refresh", logIn.Refresh)

	// Группа для администраторов (только "admin")
	adminGroup := r.Group("/admin", middleware.AuthMiddleware("admin"))
	{
		adminGroup.GET("/users", users.GetAllUsers)
		adminGroup.DELETE("/users/:id", users.DeleteUser)
		adminGroup.GET("/users/filter", users.GetFilterUsers) // Переименовано для согласованности
	}

	// Группа для докторов ("doctor") и пользователей ("user")
	doctorGroup := r.Group("/doctor", middleware.AuthMiddleware("doctor", "user"))
	{
		doctorGroup.POST("/newRecord", med_records.NewRecord)
		doctorGroup.GET("/user/:id/records", med_records.GetUserRecords)
		doctorGroup.GET("/record/:id", med_records.GetRecordId)
		doctorGroup.DELETE("/record/:id", med_records.DeleteRecord)
		doctorGroup.POST("/addSchedules", schedules.AddSchedules)
	}

	// Группа с доступом для "user", "admin", "doctor"
	sharedGroup := r.Group("/appointments", middleware.AuthMiddleware("user", "admin", "doctor"))
	{
		sharedGroup.POST("/add", appointments.AddAppointment)
		sharedGroup.GET("/:id", appointments.GetAppointment)
		sharedGroup.GET("/all", appointments.GetAllAppointment) // Переименовано
		sharedGroup.GET("/user/:id", appointments.GetUserAppointments)
		sharedGroup.PUT("/:id/time", appointments.UpdateAppointDate)
		sharedGroup.DELETE("/:id", appointments.DeleteAppointment) // Замена PUT на DELETE
		sharedGroup.GET("/doctors", users.GetAllDoctors)
		sharedGroup.GET("/user/:id/info", users.GetUserID) // Разрешение конфликта
		sharedGroup.GET("/filter", appointments.GetFilterAppointments)
	}

	// Группа с доступом для "user", "admin"
	adminUser := r.Group("/profile", middleware.AuthMiddleware("admin", "user")) // Переименована
	{
		adminUser.PUT("/userUpdate/:id", users.UpdateUser)
	}
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}
