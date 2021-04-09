package controllers

import (
	"assessmentManager/app/models/entities"
	"assessmentManager/app/models/providers/assessment_provider"
	"assessmentManager/app/models/providers/candidate_provider"
	"assessmentManager/app/models/providers/employee_provider"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/revel/revel"
)

//Структура контроллера ассессментов
type CAssessment struct {
	*revel.Controller
	provider          *assessment_provider.PAssessment
	providerEmployee  *employee_provider.PEmployee
	providerCandidate *candidate_provider.PCandidate
}

//Init интерцептор контроллера CAssessment
func (c *CAssessment) Init() revel.Result {
	//Инициализация провайдера
	c.provider = new(assessment_provider.PAssessment)
	err := c.provider.Init()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}

	c.providerEmployee = new(employee_provider.PEmployee)
	err = c.providerEmployee.Init()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}

	c.providerCandidate = new(candidate_provider.PCandidate)
	err = c.providerCandidate.Init()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}

	return nil
}

//Получение всех ассессментов
func (c *CAssessment) GetAll() revel.Result {
	//Получение ассессментов
	assessments, err := c.provider.GetAllAssessments()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(assessments))
}

//Получение ассессмента по ID
func (c *CAssessment) GetByID(id int64) revel.Result {
	//Получение ассессмента
	assessment, err := c.provider.GetAssessmentByID(id)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(assessment))
}

//Создание ассессмента
func (c *CAssessment) Create() revel.Result {
	var (
		assessment *entities.Assessment //Экземпляр сущности для создания
		err        error                //Ошибка входе выполнения функции
	)
	//Формирование сущности для создания из post параметров
	assessment, err = c.fetchPostAssessment()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Создание сущности
	assessment, err = c.provider.CreateAssessment(assessment)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(assessment))
}

//Изменение ассессмента
func (c *CAssessment) Update() revel.Result {
	var (
		assessment *entities.Assessment //Экземпляр сущности для обновления
		err        error                //Ошибка входе выполнения функции
	)
	//Формирование сущности для обновления из post параметров
	assessment, err = c.fetchPostAssessment()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Обновление сущности
	assessment, err = c.provider.UpdateAssessment(assessment)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(assessment))
}

//Удаление ассессмента
func (c *CAssessment) Delete() revel.Result {
	var (
		assessment *entities.Assessment //Экземпляр сущности для удаления
		err        error                //Ошибка в ходе выполнения функции
	)
	//Формирование сущности для удаления из post параметров
	assessment, err = c.fetchPostAssessment()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Удаление сущности
	err = c.provider.DeleteAssessment(assessment)
	if err != nil {
		c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(nil))
}

//GetEmployees метод получения всех сотрудников на ассессменте
func (c *CAssessment) GetEmployees(id int64) revel.Result {
	var (
		employee  *entities.Employee   //Экземпляр сотрудника
		employees []*entities.Employee //Слайс сотрудников
	)

	employeesID, err := c.provider.GetEmployeesFromAssessment(id)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}

	for _, id := range employeesID {
		employee, err = c.providerEmployee.GetEmployeeByID(id)
		if err != nil {
			continue
		}
		employees = append(employees, employee)
	}

	//Рендер положительного результата
	return c.RenderJSON(Success(employees))
}

//DeleteEmployees удаление всех сотрудников из ассессмента
func (c *CAssessment) DeleteEmployees() revel.Result {
	var (
		assessment *entities.Assessment //Экземляр сущности ассессмента для удаления всех сотрудников из него
		err        error                //Ошибка в ходе выполнения функции
	)
	//Формирование сущности для удаления из post параметров
	assessment, err = c.fetchPostAssessment()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Удаление всех сотрудников из ассессмента
	err = c.provider.DeleteEmployeesFromAssessment(assessment.ID)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}

	//Рендер положительного результата
	return c.RenderJSON(Success(nil))
}

//GetCandidates метод получения всех кандидатов на ассессменте
func (c *CAssessment) GetCandidates(id int64) revel.Result {
	var (
		candidate  *entities.Candidate   //Экземпляр кандидата
		candidates []*entities.Candidate //Слайс кандидатов
	)

	candidatesID, err := c.provider.GetCandidatesFromAssessment(id)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	for _, id := range candidatesID {
		candidate, err = c.providerCandidate.GetCandidateByID(id)
		if err != nil {
			return c.RenderJSON(Failed(err.Error()))
		}
		candidates = append(candidates, candidate)
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(candidates))
}

//DeleteCandidates удаление всех кандидатов из ассессмента
func (c *CAssessment) DeleteCandidates() revel.Result {
	var (
		assessment *entities.Assessment //Экземпляр сущности ассессмента для удаления всех кандидатов из него
		err        error                //Ошибка в ходе выполнения функции
	)
	//Формирование сущности для удаления из post параметров
	assessment, err = c.fetchPostAssessment()
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Удаление всех кандидатов из ассессмента
	err = c.provider.DeleteCandidatesFromAssessment(assessment.ID)
	if err != nil {
		return c.RenderJSON(Failed(err.Error()))
	}
	//Рендер положительного результата
	return c.RenderJSON(Success(nil))
}

//Метод получения сущности из post параметров
func (c *CAssessment) fetchPostAssessment() (*entities.Assessment, error) {
	//Промежуточная структура ассессмента
	type intermediateAssessment struct {
		ID         int64    `json:"ID"`         //Идентификатор
		State      string   `json:"state"`      //Вненшний ключ состояния
		Disposer   int64    `json:"disposer"`   //Внешний ключ распорядителя
		Title      string   `json:"title"`      //Название
		Date       string   `json:"date"`       //Дата
		Employees  []string `json:"employees"`  //Сотрудники
		Candidates []string `json:"candidates"` //Кандидаты
	}

	e := new(entities.Assessment)                  //Экземпляр структуры ассессмента
	interAssessment := new(intermediateAssessment) //Экземпляр промежуточноый структуры ассессмента

	var rawRequest []byte //Байтовое представление тела запроса

	//Получение тела запроса
	rawRequest, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return nil, err
	}

	//Преобразование тела запроса в промежутчоную структуру ассессмента
	err = json.Unmarshal(rawRequest, interAssessment)
	if err != nil {
		return nil, err
	}

	e.ID = interAssessment.ID
	e.State = interAssessment.State
	e.Disposer = interAssessment.Disposer
	e.Title = interAssessment.Title
	e.Date = interAssessment.Date

	//Передача в экземпляр сущности Ассессмента сотрдуников
	for _, employeeIdStr := range interAssessment.Employees {
		employeeID, err := strconv.Atoi(employeeIdStr)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		employee, err := c.providerEmployee.GetEmployeeByID(int64(employeeID))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		e.Employees = append(e.Employees, *employee)
	}
	//Передача в экземпляр сущности Ассессмента кандидатов
	for _, candidateIdStr := range interAssessment.Candidates {
		candidateID, err := strconv.Atoi(candidateIdStr)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		candidate, err := c.providerCandidate.GetCandidateByID(int64(candidateID))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		e.Candidates = append(e.Candidates, *candidate)
	}
	return e, err
}
