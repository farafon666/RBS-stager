package controllers

import (
	"assessmentManager/app/models/providers/state_assessment_provider"

	"github.com/revel/revel"
)

//Контроллер запросов для сущности "Состояние ассессмента"
type CStateAssessment struct {
	*revel.Controller
	provider *state_assessment_provider.PStateAssessment
}

//Ининицаилазация контроллера CStateAssessment
func (c *CStateAssessment) Init() revel.Result {
	var err error

	//Инициализация провайдера
	c.provider = new(state_assessment_provider.PStateAssessment)
	err = c.provider.Init()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	return nil
}

//Получение всех состояний ассессмента
func (c *CStateAssessment) GetAll() revel.Result {
	//Получение состояний
	statesAssessment, err := c.provider.GetStatesAssessments()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(statesAssessment))
}
