package mappers

import (
	"assessmentManager/app/models/entities"
	"database/sql"

	"github.com/revel/revel"
)

//Тип сущности "Сотрудник" БД
type EmployeeDBType struct {
	Pk_id          int64  //Идентификатор
	Fk_position    int64  //Внешний ключ на должность
	C_lastname     string //Фамилия
	C_firstname    string //Имя
	C_middlename   string //Отчество
	C_phone_number string //Номер телефона
	C_email        string //Электронная почта
	C_birthdate    string //Дата рождения
	C_gender       string //Пол
}

//Функция преобразования типа бд к типу сущности
func (dbt *EmployeeDBType) ToType() (e *entities.Employee, err error) {
	e = new(entities.Employee)

	e.ID = dbt.Pk_id
	e.Lastname = dbt.C_lastname
	e.Firstname = dbt.C_firstname
	e.Middlename = dbt.C_middlename
	e.PhoneNumber = dbt.C_phone_number
	e.Email = dbt.C_email
	e.BirthDate = dbt.C_birthdate
	e.Gender = dbt.C_gender

	return
}

//Функция преобразования типа сущности к типу БД
//допускается, что dbt is nil
func (_ *EmployeeDBType) FromType(e entities.Employee) (dbt *EmployeeDBType, err error) {
	dbt = &EmployeeDBType{
		Pk_id:          e.ID,
		C_lastname:     e.Lastname,
		C_firstname:    e.Firstname,
		C_middlename:   e.Middlename,
		C_phone_number: e.PhoneNumber,
		C_email:        e.Email,
		C_birthdate:    e.BirthDate,
		C_gender:       e.Gender,
	}

	return
}

//Маппер сотрудников
type MEmployee struct {
	db *sql.DB
}

//Инициализация
func (m *MEmployee) Init(db *sql.DB) {
	m.db = db
}

//Получение всех сотрудников
func (m *MEmployee) SelectAll() (es []*EmployeeDBType, err error) {
	var (
		query string    //Строка запроса
		rows  *sql.Rows //Выборка данных
	)

	//Запрос
	query = `
		SELECT
			pk_id,
			fk_position,
			c_lastname,
			c_firstname,
			c_middlename,
			c_phone_number,
			c_email,
			c_birthdate,
			c_gender
		FROM assessment.t_employees
		ORDER BY pk_id;`

	//Выполнение запроса
	rows, err = m.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		revel.AppLog.Errorf("MEmployee.SelectAll : m.db.query, %s\n", err)
		return
	}

	//Обработка строк выборки
	for rows.Next() {
		//Создание экземпляра сущности для считывания строки выборки
		e := new(EmployeeDBType)

		//Считывание строки выборки
		err = rows.Scan(
			&e.Pk_id,          //Идентификатор
			&e.Fk_position,    //Внешний ключ на должность
			&e.C_lastname,     //Фамилия
			&e.C_firstname,    //Имя
			&e.C_middlename,   //Отчество
			&e.C_phone_number, //Номер телефона
			&e.C_email,        //Электронная почта
			&e.C_birthdate,    //Дата рождения
			&e.C_gender,       //Пол
		)
		if err != nil {
			revel.AppLog.Errorf("MEmployee.SelectAll : rows.Scan, %s\n", err)
			continue
		}

		//Добавление сущности в массив
		es = append(es, e)
	}
	return
}

//Получение сотрудника по ID
func (m *MEmployee) SelectByID(id int64) (e *EmployeeDBType, err error) {
	var (
		query string   //Строка запроса
		row   *sql.Row //Выборка данных
	)

	e = new(EmployeeDBType)

	//Запрос
	query = `
		SELECT
			pk_id,
			fk_position,
			c_lastname,
			c_firstname,
			c_middlename,
			c_phone_number,
			c_email,
			c_birthdate,
			c_gender
		FROM assessment.t_employees
		WHERE pk_id = $1
		ORDER BY pk_id;`

	//Выполнение запроса
	row = m.db.QueryRow(query, id)

	//Считывание строки выборки
	err = row.Scan(
		&e.Pk_id,          //Идентификатор
		&e.Fk_position,    //Внешний ключ на должность
		&e.C_lastname,     //Фамилия
		&e.C_firstname,    //Имя
		&e.C_middlename,   //Отчество
		&e.C_phone_number, //Номер телефона
		&e.C_email,        //Электронная почта
		&e.C_birthdate,    //Дата рождения
		&e.C_gender,       //Пол
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

//Добавление сотрудника
func (m *MEmployee) Insert(edbt *EmployeeDBType) (id int64, err error) {
	var (
		query string   //Строка запроса
		row   *sql.Row //Выборка данных
	)

	//Запрос
	query = `
		INSERT INTO assessment.t_employees(
			fk_position,
			c_lastname,
			c_firstname,
			c_middlename,
			c_phone_number,
			c_email,
			c_birthdate,
			c_gender)
		VALUES (
			$1, -- fk_position,
			$2, -- c_lastname,
			$3, -- c_firstname,
			$4, -- c_middlename,
			$5, -- c_phone_number,
			$6, -- c_email,
			$7, -- c_birthdate,
			$8 -- c_gender
		)RETURNING pk_id;	
	`

	//Выполнение запроса
	row = m.db.QueryRow(query,
		edbt.Fk_position,
		edbt.C_lastname,
		edbt.C_firstname,
		edbt.C_middlename,
		edbt.C_phone_number,
		edbt.C_email,
		edbt.C_birthdate,
		edbt.C_gender,
	)

	//Считывание id
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

//Изменение сотрудника
func (m *MEmployee) Update(edbt *EmployeeDBType) (err error) {
	var (
		query string //строка запроса
	)

	//Запрос
	query = `
		UPDATE assessment.t_employees
		SET 
			fk_position = $2,
			c_lastname = $3,
			c_firstname = $4,
			c_middlename = $5,
			c_phone_number = $6,
			c_email = $7,
			c_birthdate = $8,
			c_gender = $9
		WHERE pk_id = $1;
	`

	//Выполнение запроса
	_, err = m.db.Exec(query,
		edbt.Pk_id,
		edbt.Fk_position,
		edbt.C_lastname,
		edbt.C_firstname,
		edbt.C_middlename,
		edbt.C_phone_number,
		edbt.C_email,
		edbt.C_birthdate,
		edbt.C_gender,
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

//Удаление сотрудника
func (m *MEmployee) Delete(edbt *EmployeeDBType) (err error) {
	var (
		query string //Строка запроса
	)

	//Запрос
	query = `
		DELETE FROM assessment.t_employees
		WHERE pk_id = $1;
	`

	//Выполнение запроса
	_, err = m.db.Exec(query, edbt.Pk_id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}
	return
}
