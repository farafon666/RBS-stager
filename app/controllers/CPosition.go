package controllers

import (
	"assessmentManager/app/models/providers/position_provider"

	"github.com/revel/revel"
)

//Контроллер запросов для сущности "Должность"
type CPosition struct {
	*revel.Controller
	provider *position_provider.PPosition
}

//Инициализация контроллера CPosition
func (c *CPosition) Init() revel.Result {
	var err error

	//Инициализация провайдера
	c.provider = new(position_provider.PPosition)
	err = c.provider.Init()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	return nil
}

//Получение всех должностей
func (c *CPosition) GetAll() revel.Result {
	//Получение должностей
	positions, err := c.provider.GetPositions()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(positions))
}
