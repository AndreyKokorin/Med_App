package doctors

import (
	"awesomeProject/internal/database"
	doctorRep "awesomeProject/internal/repositories/doctor"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
)

// GetFilteredDoctors
// @Summary Получение списка докторов с фильтрацией
// @Description Возвращает список докторов, отфильтрованных по указанным параметрам
// @Tags doctors
// @Security ApiKeyAuth
// @Produce json
// @Param specialty query string false "Специальность доктора"
// @Param experience query int false "Опыт работы (лет)"
// @Param languages query string false "Языки, на которых говорит доктор"
// @Param gender query string false "Пол доктора"
// @Param min_age query int false "Минимальный возраст доктора"
// @Param max_age query int false "Максимальный возраст доктора"
// @Success 200 {array} models.User "Список отфильтрованных докторов"
// @Failure 400 {object} map[string]string "Ошибка валидации"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /shared/doctors/filter [get]
func GetFilteredDoctors(ctx *gin.Context) {
	// Парсим параметры запроса
	queryParams := ctx.Request.URL.Query()

	filters, err := getFilterParameters(queryParams)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to get doctors: %v", err)})
		return
	}

	// Получаем отфильтрованных докторов
	doctors, err := doctorRep.GetFilteredDoctors(database.DB, filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get doctors: %v", err)})
		return
	}

	// Возвращаем результат в формате JSON
	ctx.JSON(http.StatusOK, doctors)
}

func getFilterParameters(queryParams url.Values) (map[string]interface{}, error) {
	filters := make(map[string]interface{})

	// Добавляем параметры фильтрации
	if specialty := queryParams.Get("specialty"); specialty != "" {
		filters["specialty"] = specialty
	}

	if experience := queryParams.Get("experience"); experience != "" {
		exp, err := strconv.Atoi(experience)
		if err != nil {
			return nil, fmt.Errorf("invalid experience value")
		}
		filters["experience"] = exp
	}

	if languages := queryParams.Get("languages"); languages != "" {
		filters["languages"] = languages
	}

	if gender := queryParams.Get("gender"); gender != "" {
		filters["gender"] = gender
	}

	if minAge := queryParams.Get("min_age"); minAge != "" {
		min, err := strconv.Atoi(minAge)
		if err != nil {
			return nil, fmt.Errorf("invalid min_age value")
		}
		filters["min_age"] = min
	}

	if maxAge := queryParams.Get("max_age"); maxAge != "" {
		max, err := strconv.Atoi(maxAge)
		if err != nil {
			return nil, fmt.Errorf("invalid max_age value")
		}
		filters["max_age"] = max
	}

	return filters, nil
}
