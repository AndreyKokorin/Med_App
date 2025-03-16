package router

import (
	"awesomeProject/internal/handlers/appointments"
	"awesomeProject/internal/handlers/logIn"
	"awesomeProject/internal/handlers/logUp"
	"awesomeProject/internal/handlers/med_records"
	"awesomeProject/internal/handlers/schedules"
	"awesomeProject/internal/handlers/timeSlots"
	"awesomeProject/internal/handlers/users"
	"awesomeProject/internal/handlers/users/doctors"
	"awesomeProject/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func SetupRouter(r *gin.Engine) {
	r.Use(Cors)

	// Базовые маршруты (без авторизации)
	api := r.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", logUp.LogUpUser) // Используем "register" вместо "logup"
		auth.POST("/login", logIn.LogIn)
		auth.POST("/password/reset", users.ChangePasswordSendEmail) // Используем "password/reset" для отправки кода
		auth.POST("/password/change", users.ChangePassword)         // Используем "password/change" для изменения пароля
		auth.POST("/token/refresh", logIn.Refresh)                  // Используем "token/refresh" для обновления токена
		auth.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Группа для администраторов (только "admin")
	adminGroup := api.Group("/admin", middleware.AuthMiddleware("admin"))
	{
		adminGroup.DELETE("/users/:id", users.DeleteUser) // Удаление пользователя
	}

	// Группа для докторов ("doctor") и администраторов ("admin")
	doctorGroup := api.Group("/doctors", middleware.AuthMiddleware("doctor", "admin"))
	{
		doctorGroup.POST("/records", med_records.NewRecord)          // Создание новой медицинской записи
		doctorGroup.GET("/records/:id", med_records.GetRecordId)     // Получение медицинской записи по ID
		doctorGroup.DELETE("/records/:id", med_records.DeleteRecord) // Удаление медицинской записи
		doctorGroup.POST("/schedules", schedules.AddSchedules)       // Добавление расписания
		doctorGroup.GET("/users/filter", users.GetFilterUsers)       // Получение отфильтрованных пользователей
		doctorGroup.PUT("/update", doctors.UpdateDoctorsData)        // Обновление данных доктора
	}

	// Группа с доступом для всех ролей ("user", "admin", "doctor")
	sharedGroup := api.Group("/shared", middleware.AuthMiddleware("user", "admin", "doctor"))
	{
		// Маршруты для пользователей
		sharedGroup.GET("/users", users.GetAllUsers)                      // Получение всех пользователей
		sharedGroup.GET("/users/:id", users.GetUserID)                    // Получение информации о пользователе по ID
		sharedGroup.PUT("/users/update", users.UpdateUser)                // Обновление информации о пользователе
		sharedGroup.GET("/users/:id/records", med_records.GetUserRecords) // Получение медицинских записей пользователя
		sharedGroup.GET("/profile", users.GetProfile)                     // Получение профиля текущего пользователя
		sharedGroup.POST("/users/details", users.AddDetailsData)          // Добавление дополнительных данных пользователя
		sharedGroup.PUT("/users/avatar", users.UploadUserAvatar)          // Загрузка аватара пользователя

		// Маршруты для докторов
		sharedGroup.GET("/doctors", users.GetAllDoctors) // Получение всех докторов
		sharedGroup.GET("/doctors/:id/slots", timeSlots.GetActualTimeSlotsForDoctor)
		sharedGroup.GET("/doctors/:id/all/slots", timeSlots.GetTimeSlotsForDoctor)
		sharedGroup.GET("/doctors/:id/profile", doctors.GetDoctorProfile) // Получение профиля доктора по ID
		sharedGroup.GET("/doctors/filter", doctors.GetFilteredDoctors)    // Получение отфильтрованных докторов

		// Маршруты для записей и расписаний
		sharedGroup.POST("/appointments", appointments.AddAppointment)              // Создание новой записи на прием
		sharedGroup.PUT("/appointments/:id/cancel", appointments.CancelAppointment) // Отмена записи на прием
		sharedGroup.GET("/appointments", appointments.GetAppointmentDetails)        // Получение деталей записи на прием
		sharedGroup.GET("/schedules", schedules.GetFilterSchedules)                 // Получение отфильтрованных расписаний
	}
}
func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")
	c.Header("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}

/*docker exec -it f7dd40fb36c6 psql -U postgres -d fitness_api
 */
