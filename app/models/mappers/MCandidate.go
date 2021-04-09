package mappers

import (
	"assessmentManager/app/models/entities"
	"database/sql"
)

//Тип сущности "Кандидат" БД
type CandidateDBType struct {
	Pk_id          int64  //Идентификатор
	C_lastname     string //Фамилия
	C_firstname    string //Имя
	C_middlename   string //Отчество
	C_phone_number string //Номер телефона
	C_email        string //Электронная почта
	C_birthdate    string //Дата рождения
	C_gender       string //Пол
}

//Функция преобразования типа БД к типу сущности
func (dbt *CandidateDBType) ToType() (с *entities.Candidate, err error) {
	с = new(entities.Candidate)

	с.ID = dbt.Pk_id
	с.Lastname = dbt.C_lastname
	с.Firstname = dbt.C_firstname
	с.Middlename = dbt.C_middlename
	с.PhoneNumber = dbt.C_phone_number
	с.Email = dbt.C_email
	с.BirthDate = dbt.C_birthdate
	с.Gender = dbt.C_gender

	return
}

//Функция преобразования типа сущности к типу БД
//допускается, что dbt is nil
func (_ *CandidateDBType) FromType(с entities.Candidate) (dbt *CandidateDBType, err error) {
	dbt = &CandidateDBType{
		Pk_id:          с.ID,
		C_lastname:     с.Lastname,
		C_firstname:    с.Firstname,
		C_middlename:   с.Middlename,
		C_phone_number: с.PhoneNumber,
		C_email:        с.Email,
		C_birthdate:    с.BirthDate,
		C_gender:       с.Gender,
	}

	return
}

//Маппер кандидатов
type MCandidate struct {
	db *sql.DB
}

//Инициализация
func (m *MCandidate) Init(db *sql.DB) {
	m.db = db
}

//Получение всех кандидатов
func (m *MCandidate) SelectAll() (cs []*CandidateDBType, err error) {
	var (
		query string    //Строка запроса
		rows  *sql.Rows //Выборка данных
	)

	//Запрос
	query = `
		SELECT
			pk_id,
			c_lastname,
			c_firstname,
			c_middlename,
			c_phone_number,
			c_email,
			c_birthdate,
			c_gender
		FROM assessment.t_candidates
		ORDER BY pk_id;`

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
		c := new(CandidateDBType)

		//Считывание строки выборки
		err = rows.Scan(
			&c.Pk_id,          //Идентификатор
			&c.C_lastname,     //Фамилия
			&c.C_firstname,    //Имя
			&c.C_middlename,   //Отчество
			&c.C_phone_number, //Номер телефона
			&c.C_email,        //Электронная почта
			&c.C_birthdate,    //Дата рождения
			&c.C_gender,       //Пол
		)
		if err != nil {
			continue
		}

		//Добавление сущности в массив
		cs = append(cs, c)
	}
	return
}

//Получение Кандидата по ID
func (m MCandidate) SelectByID(id int64) (c *CandidateDBType, err error) {
	var (
		query string   //Строка запроса
		row   *sql.Row //Выборка данных
	)

	c = new(CandidateDBType)

	//Запрос
	query = `
		SELECT
			pk_id,
			c_lastname,
			c_firstname,
			c_middlename,
			c_phone_number,
			c_email,
			c_birthdate,
			c_gender
		FROM assessment.t_candidates
		WHERE pk_id = $1
		ORDER BY pk_id;`

	//Выполнение запроса
	row = m.db.QueryRow(query, id)

	//Считывание строки выборки
	err = row.Scan(
		&c.Pk_id,          //Идентификатор
		&c.C_lastname,     //Фамилия
		&c.C_firstname,    //Имя
		&c.C_middlename,   //Отчество
		&c.C_phone_number, //Номер телефона
		&c.C_email,        //Электронная почта
		&c.C_birthdate,    //Дата рождения
		&c.C_gender,       //Пол
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

//Добавление кандидата
func (m *MCandidate) Insert(cdbt *CandidateDBType) (id int64, err error) {
	var (
		query string   //Строка запроса
		row   *sql.Row //Выборка данных
	)

	//Запрос
	query = `
		INSERT INTO assessment.t_candidates(			
			c_lastname,
			c_firstname,
			c_middlename,
			c_phone_number,
			c_email,
			c_birthdate,
			c_gender)
		VALUES (
			$1, -- c_lastname,
			$2, -- c_firstname,
			$3, -- c_middlename,
			$4, -- c_phone_number,
			$5, -- c_email,
			$6, -- c_birthdate,
			$7  -- c_gender
		)RETURNING pk_id;	
	`

	//Выполнение запроса
	row = m.db.QueryRow(query,
		cdbt.C_lastname,
		cdbt.C_firstname,
		cdbt.C_middlename,
		cdbt.C_phone_number,
		cdbt.C_email,
		cdbt.C_birthdate,
		cdbt.C_gender,
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

//Изменение кандидата
func (m *MCandidate) Update(cdbt *CandidateDBType) (err error) {
	var (
		query string //Строка запроса
	)

	//Запрос
	query = `
		UPDATE assessment.t_candidates
		SET 
			c_lastname = $2,
			c_firstname = $3,
			c_middlename = $4,
			c_phone_number = $5,
			c_email = $6,
			c_birthdate = $7,
			c_gender = $8
		WHERE pk_id = $1;
	`

	//Выполнение запроса
	_, err = m.db.Exec(query,
		cdbt.Pk_id,
		cdbt.C_lastname,
		cdbt.C_firstname,
		cdbt.C_middlename,
		cdbt.C_phone_number,
		cdbt.C_email,
		cdbt.C_birthdate,
		cdbt.C_gender,
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

//Удаление кандидата
func (m *MCandidate) Delete(cdbt *CandidateDBType) (err error) {
	var (
		query string //Строка запроса
	)

	//Запрос
	query = `
		DELETE FROM assessment.t_candidates
		WHERE pk_id = $1;
	`

	//Выполнение запроса
	_, err = m.db.Exec(query, cdbt.Pk_id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}
	return
}
