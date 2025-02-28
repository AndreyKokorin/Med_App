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
		adminGroup.DELETE("/delete/user/:id", users.DeleteUser)
	}

	// Группа для докторов ("doctor") и пользователей ("user")
	doctorGroup := r.Group("/doctor", middleware.AuthMiddleware("doctor", "admin"))
	{
		doctorGroup.POST("/newRecord", med_records.NewRecord)
		doctorGroup.GET("/record/:id", med_records.GetRecordId)
		doctorGroup.DELETE("/record/:id", med_records.DeleteRecord)
		doctorGroup.POST("/addSchedules", schedules.AddSchedules)
		doctorGroup.GET("/users/filter", users.GetFilterUsers)
	}

	// Группа с доступом для "user", "admin", "doctor"
	sharedGroup := r.Group("/allRoles", middleware.AuthMiddleware("user", "admin", "doctor"))
	{
		sharedGroup.POST("appointments/add", appointments.AddAppointment)
		sharedGroup.GET("appointments/:id", appointments.GetAppointment)
		sharedGroup.GET("appointments/all", appointments.GetAllAppointment)
		sharedGroup.GET("appointments/user/:id", appointments.GetUserAppointments)
		sharedGroup.PUT("appointments/:id/time", appointments.UpdateAppointDate)
		sharedGroup.DELETE("appointments/:id", appointments.DeleteAppointment)
		sharedGroup.GET("/doctors", users.GetAllDoctors)
		sharedGroup.GET("/user/:id/info", users.GetUserID)
		sharedGroup.GET("/filter", appointments.GetFilterAppointments)
		sharedGroup.PUT("/userUpdate/:id", users.UpdateUser)
		sharedGroup.GET("/user/:id/records", med_records.GetUserRecords)
		sharedGroup.GET("/profile", users.GetProfile)
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
