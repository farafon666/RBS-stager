package entities

//Candidate структура сущности кандидата
type Candidate struct {
	ID          int64  `json:"ID"`           // идентификатор
	Lastname    string `json:"lastname"`     // фамилия
	Firstname   string `json:"firstname"`    // имя
	Middlename  string `json:"middlename"`   // отчество
	PhoneNumber string `json:"phone_number"` // номер телефона
	Email       string `json:"email"`        // адрес электронной почты
	BirthDate   string `json:"birthdate"`    // дата рождения
	Gender      string `json:"gender"`       // пол
}
