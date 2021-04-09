package assessment_provider

import (
	"assessmentManager/app/helpers"
	"assessmentManager/app/models/entities"
	"assessmentManager/app/models/mappers"
	"database/sql"
)

//Провайдер контроллера ассессментов
type PAssessment struct {
	assessmentMapper      *mappers.MAssessment
	stateAssessmentMapper *mappers.MStateAssessment
	employeeMapper        *mappers.MEmployee
}

//Инициализация
func (p *PAssessment) Init() (err error) {
	var db *sql.DB //Экземпляр подключения к БД

	//Получение экземпляра подключения к БД
	db, err = helpers.GetDBConnection()
	if err != nil {
		return
	}

	//Инициализация маппера ассессментов
	p.assessmentMapper = new(mappers.MAssessment)
	p.assessmentMapper.Init(db)
	//Инициализация маппера состояний ассессментов
	p.stateAssessmentMapper = new(mappers.MStateAssessment)
	p.stateAssessmentMapper.Init(db)

	return
}

//Метод получения всех ассессментов
func (p *PAssessment) GetAllAssessments() (as []*entities.Assessment, err error) {
	var (
		adbts []*mappers.AssessmentDBType
		a     *entities.Assessment
	)

	//Получение данных ассессментов
	adbts, err = p.assessmentMapper.SelectAll()
	if err != nil {
		return
	}

	for _, adbt := range adbts {
		//Преобразование к типу сущности
		a, err = adbt.ToType()
		if err != nil {
			return
		}

		//Получение значения состояния ассессмента по ключу
		a.State, err = p.stateAssessmentMapper.StateAssessmentByID(adbt.Fk_state)
		if err != nil {
			return
		}

		//Получение внешнего ключа на распорядителя
		a.Disposer = adbt.Fk_disposer

		as = append(as, a)
	}
	return
}

//Метод для получения ассессмента по ID
func (p *PAssessment) GetAssessmentByID(id int64) (a *entities.Assessment, err error) {
	var (
		adbt *mappers.AssessmentDBType
	)

	//Получение данных ассессмента
	adbt, err = p.assessmentMapper.SelectByID(id)
	if err != nil {
		return
	}

	//Преобразование тип БД к типу сущности
	a, err = adbt.ToType()
	if err != nil {
		return
	}

	//Получение значения состояния ассессмента по ключу
	a.State, err = p.stateAssessmentMapper.StateAssessmentByID(adbt.Fk_state)
	if err != nil {
		return
	}

	//Получение внешнего ключа на распорядителя
	a.Disposer = adbt.Fk_disposer

	return
}

//Метод создания ассессмента
func (p *PAssessment) CreateAssessment(assessment *entities.Assessment) (a *entities.Assessment, err error) {
	var adbt *mappers.AssessmentDBType

	//Инициализация структуры БД из структуры сущности
	adbt, err = adbt.FromType(*assessment)
	if err != nil {
		return
	}

	//Получение внешнего ключа на состояние ассессмента
	adbt.Fk_state, err = p.stateAssessmentMapper.IDByStateAssessment(assessment.State)
	if err != nil {
		return
	}

	//Получение внешнего ключа на распорядителя
	adbt.Fk_disposer = assessment.Disposer

	//Добавление ассессмента
	assessment.ID, err = p.assessmentMapper.Insert(adbt)
	if err != nil {
		return
	}

	//Добавление сотрудников в ассессмент
	for _, employee := range assessment.Employees {
		err = p.AddEmployeeToAssessment(employee.ID, assessment.ID)
		if err != nil {
			continue
		}
	}

	//Добавление кандидатов в ассессмент
	for _, candidate := range assessment.Candidates {
		err = p.AddCandidateToAssessment(candidate.ID, assessment.ID)
		if err != nil {
			continue
		}
	}

	return assessment, nil
}

//Метод обновления ассессмента
func (p *PAssessment) UpdateAssessment(assessment *entities.Assessment) (a *entities.Assessment, err error) {
	var adbt *mappers.AssessmentDBType

	//Инициализация структуры БД из структур сущности
	adbt, err = adbt.FromType(*assessment)
	if err != nil {
		return
	}

	//Получение внешнего ключа на состояние ассессмента
	adbt.Fk_state, err = p.stateAssessmentMapper.IDByStateAssessment(assessment.State)
	if err != nil {
		return
	}

	//Получение внешнего ключа на распорядителя
	adbt.Fk_disposer = assessment.Disposer

	//Обновление ассессмента
	err = p.assessmentMapper.Update(adbt)
	if err != nil {
		return
	}

	//Добавление сотрудников в ассессмент
	for _, employee := range assessment.Employees {
		err = p.AddEmployeeToAssessment(employee.ID, assessment.ID)
		if err != nil {
			continue
		}
	}

	//Добавление кандидатов в ассессмент
	for _, candidate := range assessment.Candidates {
		err = p.AddCandidateToAssessment(candidate.ID, assessment.ID)
		if err != nil {
			continue
		}
	}

	return assessment, nil
}

//Метод удаления ассессмента
func (p *PAssessment) DeleteAssessment(assessment *entities.Assessment) (err error) {
	var adbt *mappers.AssessmentDBType

	//Инициализация структуры БД из структуры сущности
	adbt, err = adbt.FromType(*assessment)
	if err != nil {
		return
	}

	//Удаление ассессмента
	err = p.assessmentMapper.Delete(adbt)
	if err != nil {
		return
	}
	return
}

//GetEmployeesFromAssessment метод получения всех сотрудников из ассессмента
func (p *PAssessment) GetEmployeesFromAssessment(assessmentID int64) (employeesID []int64, err error) {
	employeesID, err = p.assessmentMapper.SelectEmployeesFromAssessment(assessmentID)
	if err != nil {
		return
	}
	return
}

//AddEmployeeToAssessment метод добавления сотрудников в ассессмент
func (p *PAssessment) AddEmployeeToAssessment(employeeID, assessmentID int64) (err error) {
	err = p.assessmentMapper.InsertEmployeeToAssessment(assessmentID, employeeID)
	if err != nil {
		return
	}
	return
}

//DeleteEmployeesFromAssessment удаление всех сотрудников из ассессмента
func (p *PAssessment) DeleteEmployeesFromAssessment(assessmentID int64) (err error) {
	err = p.assessmentMapper.DeleteEmployeesFromAssessment(assessmentID)
	if err != nil {
		return
	}
	return
}

//GetCandidatesFromAssessment метод получения всех кандидатов из ассессмента
func (p *PAssessment) GetCandidatesFromAssessment(assessmentID int64) (candidatesID []int64, err error) {
	candidatesID, err = p.assessmentMapper.SelectCandidatesFromAssessment(assessmentID)
	if err != nil {
		return
	}
	return
}

//AddCandidateToAssessment добавление кандидата в ассессмент
func (p *PAssessment) AddCandidateToAssessment(candidateID, assessmentID int64) (err error) {
	err = p.assessmentMapper.InsertCandidateToAssessment(candidateID, assessmentID)
	if err != nil {
		return
	}
	return
}

//DeleteCandidatesFromAssessment удаление всех кандидатов
func (p *PAssessment) DeleteCandidatesFromAssessment(assessmentID int64) (err error) {
	err = p.assessmentMapper.DeleteCandidatesFromAssessment(assessmentID)
	if err != nil {
		return
	}
	return
}
