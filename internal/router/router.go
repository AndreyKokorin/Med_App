package router

import (
	"awesomeProject/internal/handlers/appointments"
	"awesomeProject/internal/handlers/logIn"
	"awesomeProject/internal/handlers/logUp"
	"awesomeProject/internal/handlers/med_records"
	"awesomeProject/internal/handlers/schedules"
	"awesomeProject/internal/handlers/timeSlots"
	"awesomeProject/internal/handlers/users"
	"awesomeProject/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func SetupRouter(r *gin.Engine) {
	r.Use(Cors)

	r.POST("/logup", logUp.LogUpUser)
	r.POST("/login", logIn.LogIn)
	r.POST("/sendEmailCode", users.ChangePasswordSendEmail)
	r.POST("/changePassword", users.ChangePassword)
	r.POST("refresh", logIn.Refresh)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		sharedGroup.GET("/doctors", users.GetAllDoctors)
		sharedGroup.GET("/user/:id/info", users.GetUserID)
		sharedGroup.PUT("/userUpdate/:id", users.UpdateUser)
		sharedGroup.GET("/user/:id/records", med_records.GetUserRecords)
		sharedGroup.GET("/profile", users.GetProfile)
		sharedGroup.GET("/doctor/:id/actualSlots", timeSlots.GetActualTimeSlotsForDoctor)
		sharedGroup.PUT("appointments/:id/cancel", appointments.CancelAppointment)
		sharedGroup.GET("/appointments", appointments.GetAppointmentDetails)
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

/*docker exec -it f7dd40fb36c6 psql -U postgres -d fitness_api
 */
