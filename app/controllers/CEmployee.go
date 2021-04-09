package controllers

import (
	"assessmentManager/app/models/entities"
	"assessmentManager/app/models/providers/employee_provider"
	"encoding/json"
	"io/ioutil"

	"github.com/revel/revel"
)

//Структура контроллера сотрудников
type CEmployee struct {
	*revel.Controller
	provider *employee_provider.PEmployee
}

//Init интерцептор контроллера CEmployee
func (c *CEmployee) Init() revel.Result {
	//Инициализация провайдера
	c.provider = new(employee_provider.PEmployee)
	err := c.provider.Init()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	return nil
}

//Получение всех сторудников
func (c *CEmployee) GetAll() revel.Result {
	//Получение сотрудников
	employees, err := c.provider.GetAllEmployees()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(employees))
}

//Получение сотрудника по ID
func (c *CEmployee) GetByID(id int64) revel.Result {
	//Получение сотрудника
	employee, err := c.provider.GetEmployeeByID(id)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(employee))
}

//Создание сотрудника
func (c *CEmployee) Create() revel.Result {
	var (
		employee *entities.Employee //Экземпляр сущности для создания
		err      error              //Ошибка входе выполнения функции
	)
	//Формирование сущности для создания из post параметров
	employee, err = c.fetchPostEmployee()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Создание сущности
	employee, err = c.provider.CreateEmployee(employee)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(employee))
}

//Изменение сотрудника
func (c *CEmployee) Update() revel.Result {
	var (
		employee *entities.Employee //Экземляр сущности для обновления
		err      error              //Ошибка в ходе выполнения функции
	)
	//Формирование сущности для обновления из post параметров
	employee, err = c.fetchPostEmployee()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Обновление сущности
	employee, err = c.provider.UpdateEmployee(employee)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(employee))
}

//Удаление сотрудника
func (c *CEmployee) Delete() revel.Result {
	var (
		employee *entities.Employee //Экземпляр сущности для удаления
		err      error              //Ошибка в ходе выполнения функции
	)
	//Формирование сущности для удаления из post параметров
	employee, err = c.fetchPostEmployee()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Удаление сущности
	err = c.provider.DeleteEmployee(employee)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(nil))
}

//Метод получения сущности из post параметров
func (c *CEmployee) fetchPostEmployee() (e *entities.Employee, err error) {
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
