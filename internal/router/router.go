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

	r.POST("/register", logUp.LogUpUser)
	r.POST("/login", logIn.LogIn)
	r.GET("/appointsfilter", appointments.GetFilterAppointments)

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
		adminGroup.GET("/getUserFilter", users.GetFilterUsers)
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
		sharedGroup.POST("/", appointments.AddAppointment)
		sharedGroup.GET("/:id", appointments.GetAppointment)
		sharedGroup.GET("/", appointments.GetAllAppointment)
		sharedGroup.GET("/user/:id", appointments.GetUserAppointments)
		sharedGroup.PUT("/:id/time", appointments.UpdateAppointDate)
		sharedGroup.PUT("/:id/delete", appointments.DeleteAppointment)
		sharedGroup.GET("doctors", users.GetAllDoctors)
	}

	// Группа с доступом для "user", "admin"
	adminUser := r.Group("/adminAndUser", middleware.AuthMiddleware("admin", "user"))
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
