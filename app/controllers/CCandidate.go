package controllers

import (
	"assessmentManager/app/models/entities"
	"assessmentManager/app/models/providers/candidate_provider"
	"encoding/json"
	"io/ioutil"

	"github.com/revel/revel"
)

//Структура контроллера кандидатов
type CCandidate struct {
	*revel.Controller
	provider *candidate_provider.PCandidate
}

//Init интерцептор контроллера CCandidate
func (c *CCandidate) Init() revel.Result {
	//Инициализация провайдера
	c.provider = new(candidate_provider.PCandidate)
	err := c.provider.Init()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	return nil
}

//Получение всех кандидатов
func (c *CCandidate) GetAll() revel.Result {
	//Получение кандидатов
	canidates, err := c.provider.GetAllCandidates()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(canidates))
}

//Получение кандидата по ID
func (c *CCandidate) GetByID(id int64) revel.Result {
	//Получение кандидата
	candidate, err := c.provider.GetCandidateByID(id)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(candidate))
}

//Создание кандидата
func (c *CCandidate) Create() revel.Result {
	var (
		candidate *entities.Candidate //Экземпляр сущности для создания
		err       error               //Ошибка входе выполнения функции
	)
	//Формирование сущности для создания из post параметров
	candidate, err = c.fetchPostCandidate()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Создание сущности
	candidate, err = c.provider.CreateCandidate(candidate)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(candidate))
}

//Изменение кандидата
func (c *CCandidate) Update() revel.Result {
	var (
		candidate *entities.Candidate //Экземпляр сущности для обновления
		err       error               //Ошибка в ходе выполнения функции
	)
	//Формирования сущности для обновления из post параметров
	candidate, err = c.fetchPostCandidate()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Обновление сущности
	candidate, err = c.provider.UpdateCandidate(candidate)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(candidate))
}

//Удаление кандидата
func (c *CCandidate) Delete() revel.Result {
	var (
		candidate *entities.Candidate //Экземпляр сущности для удаления
		err       error               //Ошибка в ходе выполнения функции
	)
	//Формирование сущности для удаления из post параметров
	candidate, err = c.fetchPostCandidate()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Удаление сущности
	err = c.provider.DeleteCandidate(candidate)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(nil))
}

//Метод получения сущности из post параметров
func (c *CCandidate) fetchPostCandidate() (e *entities.Candidate, err error) {
	var rawRequest []byte //Байтовое представление тела запроса

	//Получение тела запроса
	rawRequest, err = ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return
	}

	//Преобразование тела запроса в структуру сущности
	err = json.Unmarshal(rawRequest, &e)
	if err != nil {
		return
	}

	return
}
