package mappers

import (
	"assessmentManager/app/models/entities"
	"database/sql"
)

//Тип сущности "Ассессмент" БД
type AssessmentDBType struct {
	Pk_id       int64  //Идентификатор
	Fk_state    int64  //Внешний ключ состояния ассессмента
	Fk_disposer int64  //Внешний ключ распорядителя
	C_title     string //Название
	C_date      string //Дата
}

//Функция преобразования типа БД к типу сущности
func (dbt *AssessmentDBType) ToType() (a *entities.Assessment, err error) {
	a = new(entities.Assessment)

	a.ID = dbt.Pk_id
	a.Title = dbt.C_title
	a.Date = dbt.C_date

	return
}

//Функция преобразования типа сущности к типу БД
//допускается, что dbt is nil
func (_ *AssessmentDBType) FromType(a entities.Assessment) (dbt *AssessmentDBType, err error) {
	dbt = &AssessmentDBType{
		Pk_id:   a.ID,
		C_title: a.Title,
		C_date:  a.Date,
	}

	return
}

//Маппер ассессментов
type MAssessment struct {
	db *sql.DB
}

//Инициализация
func (m *MAssessment) Init(db *sql.DB) {
	m.db = db
}

//Получение всех ассессментов
func (m *MAssessment) SelectAll() (as []*AssessmentDBType, err error) {
	var (
		query string    //Строка запроса
		rows  *sql.Rows //Выборка данных
	)

	//Запрос
	query = `
		SELECT
			pk_id,
			fk_state,
			fk_disposer,
			c_title,
			c_date
		FROM assessment.t_assessments
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
		a := new(AssessmentDBType)

		//Считывание строки выборки
		err = rows.Scan(
			&a.Pk_id,       //ID
			&a.Fk_state,    //Внешний ключ на состояние ассессментов
			&a.Fk_disposer, //Внешний ключ распорядителя
			&a.C_title,     //Название
			&a.C_date,      //Дата
		)
		if err != nil {
			continue
		}

		//Добавление сущности в массив
		as = append(as, a)
	}
	return
}

//Получение ассессмента по ID
func (m *MAssessment) SelectByID(id int64) (a *AssessmentDBType, err error) {
	var (
		query string   //Строка запроса
		row   *sql.Row //Выборка данных
	)

	a = new(AssessmentDBType)

	//Запрос
	query = `
		SELECT
			pk_id,
			fk_state,
			fk_disposer,
			c_title,
			c_date
		FROM assessment.t_assessments
		WHERE pk_id = $1
		ORDER BY pk_id;
	`

	//Выполнение запроса
	row = m.db.QueryRow(query, id)

	//Считывание строки выборки
	err = row.Scan(
		&a.Pk_id,       //ID
		&a.Fk_state,    //Внешний ключ на состояние ассессмента
		&a.Fk_disposer, //Внешний ключ распорядителя
		&a.C_title,     //Название
		&a.C_date,      //Дата
	)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}
	return
}

//Добавление ассессмента
func (m *MAssessment) Insert(adbt *AssessmentDBType) (id int64, err error) {
	var (
		query string   //Строка запроса
		row   *sql.Row //Выборка данных
	)

	//Запрос
	query = `
		INSERT INTO assessment.t_assessments(
			fk_state,
			fk_disposer,
			c_title,
			c_date)
		VALUES(
			$1,	-- fk_state
			$2, -- fk_disposer
			$3,	-- c_date
			$4	-- c_title
		)RETURNING pk_id;
	`
	//Выполнение запроса
	row = m.db.QueryRow(query,
		adbt.Fk_state,
		adbt.Fk_disposer,
		adbt.C_title,
		adbt.C_date,
	)

	//Считывание ID
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

//Изменение ассессмента
func (m *MAssessment) Update(adbt *AssessmentDBType) (err error) {
	var (
		query string //Строка запроса
	)

	//Запрос
	query = `
		UPDATE assessment.t_assessments
		SET
			fk_state = $2,
			fk_disposer = $3,	
			c_title = $4,
			c_date = $5
		WHERE pk_id = $1;
	`

	//Выполнение запроса
	_, err = m.db.Exec(query,
		adbt.Pk_id,
		adbt.Fk_state,
		adbt.Fk_disposer,
		adbt.C_title,
		adbt.C_date,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}
	return
}

//Удаление ассессмента
func (m *MAssessment) Delete(adbt *AssessmentDBType) (err error) {
	var (
		query string //Строка запроса
	)
	//Удаление ассессмента из связующей таблицы
	err = m.DeleteEmployeesFromAssessment(adbt.Pk_id)
	if err != nil {
		return
	}
	//Удаление ассессмента из связующей таблицы
	err = m.DeleteCandidatesFromAssessment(adbt.Pk_id)
	if err != nil {
		return
	}

	//Запрос
	query = `
		DELETE FROM assessment.t_assessments
		WHERE pk_id = $1;
	`

	//Выполнение запроса
	_, err = m.db.Exec(query, adbt.Pk_id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}
	return
}

//SelectEmployeesFromAssessment метод получения всех сотрудников с ассессмента
func (m *MAssessment) SelectEmployeesFromAssessment(assessmentID int64) (employeesID []int64, err error) {
	var (
		query string    //Строка запроса
		rows  *sql.Rows //Выборка данных
	)

	//Запрос
	query = `
		SELECT
			fk_employee
		FROM assessment.toc_employees_assessments
		WHERE fk_assessment = $1;
	`
	//Выполнение запроса
	rows, err = m.db.Query(query, assessmentID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}

	//Обработка строк выборки
	for rows.Next() {
		//Создание элемента массива для считывания строки выборки
		var e *int64

		//Считывание строки выборки
		err = rows.Scan(&e)
		if err != nil {
			continue
		}

		//Добавление сущности в массив
		employeesID = append(employeesID, *e)
	}
	//fmt.Println("\n\n\n\n\n\n\n\n\n", rows, "\n\n\n\n\n\n\n\n\n")
	return
}

//InsertEmployeeToAssessment добавление сотрудника в ассессмент
func (m *MAssessment) InsertEmployeeToAssessment(assessmentID, employeeID int64) (err error) {
	var (
		query string // Строка запроса
	)
	//Запрос
	query = `
		INSERT INTO assessment.toc_employees_assessments
			(fk_assessment, fk_employee)
		VALUES 
			($1, $2);
	`
	// Выполнение запроса
	_, err = m.db.Exec(query, assessmentID, employeeID)
	if err != nil {
		return
	}
	return
}

//DeleteEmployeesFromAssessment удаление всех сотрудников из ассессмента
func (m *MAssessment) DeleteEmployeesFromAssessment(assessmentID int64) (err error) {
	var (
		query string //Строка запроса
	)

	//Запрос
	query = `
		DELETE FROM assessment.toc_employees_assessments
		WHERE fk_assessment = $1;
	`
	//Выполнение запроса
	_, err = m.db.Exec(query, assessmentID)
	if err != nil {
		return
	}

	return
}

//SelectCandidatesFromAssessment метод получения всех кандидатов с ассессмента
func (m *MAssessment) SelectCandidatesFromAssessment(assessmentID int64) (candidatesID []int64, err error) {
	var (
		query string    //Строка запроса
		rows  *sql.Rows //Выборка данных
	)

	//Запрос
	query = `
		SELECT
			fk_candidate
		FROM assessment.toc_candidates_assessments
		WHERE fk_assessment = $1;
	`

	//Выполнение запроса
	rows, err = m.db.Query(query, assessmentID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}

	//Обработка строк выборки
	for rows.Next() {
		//Создание элемента массива для считывания строки выборки
		var c *int64

		//Считывание строки выборки
		err = rows.Scan(&c)
		if err != nil {
			continue
		}

		//Добавление сущности в массив
		candidatesID = append(candidatesID, *c)
	}
	return
}

//InsertCandidateToAssessment добавление кандидата в ассессмент
func (m *MAssessment) InsertCandidateToAssessment(candidateID, assessmentID int64) (err error) {
	var (
		query string //Строка запроса
	)

	//Запрос
	query = `
		INSERT INTO assessment.toc_candidates_assessments
			(fk_assessment, fk_candidate)
		VALUES
			($1,$2);
	`

	//Выполнение запроса
	_, err = m.db.Exec(query, assessmentID, candidateID)
	if err != nil {
		return
	}

	return
}

//DeleteCandidatesFromAssessment удаление всех кандидатов из ассессмента
func (m *MAssessment) DeleteCandidatesFromAssessment(assessmentID int64) (err error) {
	var (
		query string //Строка запроса
	)

	//Запрос
	query = `
		DELETE FROM assessment.toc_candidates_assessments
		WHERE fk_assessment = $1;
	`

	//Выполнение запроса
	_, err = m.db.Exec(query, assessmentID)
	if err != nil {
		return
	}

	return
}
