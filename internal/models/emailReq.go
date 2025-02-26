package models

type EmailRequest struct {
	From     From   `json:"from"`     // Отправитель
	To       []To   `json:"to"`       // Получатели (массив)
	Subject  string `json:"subject"`  // Тема письма
	Text     string `json:"text"`     // Текст письма
	Category string `json:"category"` // Категория
}

type From struct {
	Email string `json:"email"` // Email отправителя
	Name  string `json:"name"`  // Имя отправителя
}

// To - структура для получателя
type To struct {
	Email string `json:"email"` // Email получателя
}
