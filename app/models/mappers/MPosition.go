package mappers

import (
	"assessmentManager/app/models/entities"
	"database/sql"

	"github.com/revel/revel"
)

// Тип сущности "Должность" БД
type PositionDBType struct {
	Pk_id  int64  // идентификатор
	C_name string // название должности
}

// ToType функция преобразования типа бд к типу сущности
func (dbt *PositionDBType) ToType() (p *entities.Position, err error) {
	p = new(entities.Position)

	p.ID = dbt.Pk_id
	p.Name = dbt.C_name

	return
}

// FromType функция преобразования типа бд из типа сущности
func (_ *PositionDBType) FromType(p *entities.Position) (dbt *PositionDBType, err error) {
	dbt = &PositionDBType{
		Pk_id:  p.ID,
		C_name: p.Name,
	}

	return
}

// MPosition маппер должностей
type MPosition struct {
	db *sql.DB
}

// Init
func (m *MPosition) Init(db *sql.DB) {
	m.db = db
}

// SelectAll получение всех должностей
func (m *MPosition) SelectAll() (pdbts []*PositionDBType, err error) {
	var (
		query string    // строка запроса
		rows  *sql.Rows // выборка данных
	)

	// запрос
	query = `
		SELECT
			pk_id,
			c_position
		FROM assessment.t_ref_positions
		ORDER BY pk_id;	
	`

	// выполнение запроса
	rows, err = m.db.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}

	// обработка строк выборки
	for rows.Next() {
		// создание экземпляра сущности для считывания строки выборки
		pdbt := new(PositionDBType)

		// считывание строки выборки
		err = rows.Scan(&pdbt.Pk_id, &pdbt.C_name)
		if err != nil {
			revel.AppLog.Errorf("MPosition.SelectAll : rows.Scan, %s\n", err)
			continue
		}

		// добавление сущности в массив
		pdbts = append(pdbts, pdbt)
	}

	return
}

// PositionNameByID получение должности по id
func (m *MPosition) PositionNameByID(id int64) (positionName string, err error) {
	var (
		query string   // строка запроса
		row   *sql.Row // выборка данных
	)

	// запрос
	query = `
		SELECT
			c_position
		FROM assessment.t_ref_positions
		WHERE pk_id = $1;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, id)

	// считывание строки выборки
	err = row.Scan(&positionName)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}

	return
}

//Получение id по должности
func (m *MPosition) IDByPositionName(positionName string) (id int64, err error) {
	var (
		query string   //Строка запроса
		row   *sql.Row //Выборка данных
	)

	// запрос
	query = `
		SELECT
			pk_id
		FROM assessment.t_ref_positions
		WHERE c_position = $1
		ORDER BY pk_id;
	`

	// выполнение запроса
	row = m.db.QueryRow(query, positionName)

	// считывание строки выборки
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
