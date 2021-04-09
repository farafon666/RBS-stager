package entities

//Assessment структура сущности ассессмента
type Assessment struct {
	ID         int64       `json:"ID"`       //Идентификатор
	State      string      `json:"state"`    //Вненшний ключ состояния
	Disposer   int64       `json:"disposer"` //Внешний ключ распорядителя
	Title      string      `json:"title"`    //Название
	Date       string      `json:"date"`     //Дата
	Employees  []Employee  `json:"-"`        //Сотрудники
	Candidates []Candidate `json:"-"`        //Кандидаты
}
