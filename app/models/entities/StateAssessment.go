package entities

//StateAssessment струкутра сущности сотояния ассессмента
type StateAssessment struct {
	ID    int64  `json:"ID"`    //Идентификатор
	State string `json:"state"` //Состояние ассессмента
}
