package mappers

import (
	"assessmentManager/app/models/entities"
	"database/sql"

	"github.com/revel/revel"
)

//Тип сущности "Состояние ассессмента" БД
type StateAssessmentDBType struct {
	Pk_id   int64  //Идентификатор
	C_state string //Сотсояние ассессмента
}

// ToType функциия преобразования бд к типу сущности
func (dbt *StateAssessmentDBType) ToType() (s *entities.StateAssessment, err error) {
	s = new(entities.StateAssessment)

	s.ID = dbt.Pk_id
	s.State = dbt.C_state

	return
}

// FromType функция преобразования типа бд из типа сущности
func (_ *StateAssessmentDBType) FromType(s *entities.StateAssessment) (dbt *StateAssessmentDBType, err error) {
	dbt = &StateAssessmentDBType{
		Pk_id:   s.ID,
		C_state: s.State,
	}

	return
}

// MStateAssessment маппер состояния ассессментов
type MStateAssessment struct {
	db *sql.DB
}

//Init
func (m *MStateAssessment) Init(db *sql.DB) {
	m.db = db
}

// SelectAll получение всех состояний ассессментов
func (m *MStateAssessment) SelectAll() (sdbts []*StateAssessmentDBType, err error) {
	var (
		query string    // строка запроса
		rows  *sql.Rows //выборка данных
	)

	//Запрос
	query = `
		SELECT 
			pk_id,
			c_state
		FROM assessment.t_ref_state_assessments
		ORDER BY pk_id;
	`

	//Выполнение запроса
	rows, err = m.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}

	//Обработка строк выборки
	for rows.Next() {
		//Создание экземпляра сущности для считывания строки выборки
		sdbt := new(StateAssessmentDBType)

		//Считывание строки выборки
		err = rows.Scan(&sdbt.Pk_id, &sdbt.C_state)
		if err != nil {
			revel.AppLog.Errorf("MPosition.SelectAll : rows.Scan, %s\n", err)
			continue
		}

		//Добавление сущности в массив
		sdbts = append(sdbts, sdbt)
	}

	return
}

// AssessmentStateByID получение состояния по id
func (m *MStateAssessment) StateAssessmentByID(id int64) (state string, err error) {
	var (
		query string   //Строка запроса
		row   *sql.Row //Выборка данных
	)

	//Запрос
	query = `
		SELECT 
			c_state
		FROM assessment.t_ref_state_assessments
		WHERE pk_id = $1;	
	`

	//Выполнение запроса
	row = m.db.QueryRow(query, id)

	//Считывание строки выборки
	err = row.Scan(&state)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}

	return
}

// IDByAssessmentState получение id по состоянию ассессмента
func (m *MStateAssessment) IDByStateAssessment(state string) (id int64, err error) {
	var (
		query string   //Строка запроса
		row   *sql.Row //Выборка данных
	)

	//Запрос
	query = `
		SELECT 
			pk_id
		FROM assessment.t_ref_state_assessments
		WHERE c_state = $1;	
	`

	//Выполнение запроса
	row = m.db.QueryRow(query, state)

	//Считывание строки выборки
	err = row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}
	return
}
