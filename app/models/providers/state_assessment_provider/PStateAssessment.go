package state_assessment_provider

import (
	"assessmentManager/app/helpers"
	"assessmentManager/app/models/entities"
	"assessmentManager/app/models/mappers"
	"database/sql"
)

//Провайдер контроллера состояний ассессмента
type PStateAssessment struct {
	stateAssessmentMapper *mappers.MStateAssessment
}

//Инициализация
func (p *PStateAssessment) Init() (err error) {
	var db *sql.DB //Экземпляр подключения к БД

	//Получение экземпляра подключения к БД
	db, err = helpers.GetDBConnection()
	if err != nil {
		return err
	}

	//Инициализация маппера состояний ассессмента
	p.stateAssessmentMapper = new(mappers.MStateAssessment)
	p.stateAssessmentMapper.Init(db)

	return
}

//Метод получения состояний ассессмента
func (p *PStateAssessment) GetStatesAssessments() (ss []*entities.StateAssessment, err error) {
	var (
		sdbts []*mappers.StateAssessmentDBType
		s     *entities.StateAssessment
	)

	//Получение данных состояний ассессментов
	sdbts, err = p.stateAssessmentMapper.SelectAll()
	if err != nil {
		return
	}

	for _, sdbt := range sdbts {
		//Преобразование к типу сущности
		s, err = sdbt.ToType()
		if err != nil {
			return
		}
		ss = append(ss, s)
	}
	return
}
