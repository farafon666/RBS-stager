package position_provider

import (
	"assessmentManager/app/helpers"
	"assessmentManager/app/models/entities"
	"assessmentManager/app/models/mappers"
	"database/sql"
)

//Провайдер контроллера должностей
type PPosition struct {
	positionMapper *mappers.MPosition
}

//Инициализация
func (p *PPosition) Init() (err error) {
	var db *sql.DB //Экземпляр подключения к БД

	//Получение экземпляра подключения к БД
	db, err = helpers.GetDBConnection()
	if err != nil {
		return err
	}

	//Инициализация маппера должностей
	p.positionMapper = new(mappers.MPosition)
	p.positionMapper.Init(db)

	return
}

//Метод получения должностей
func (p *PPosition) GetPositions() (ps []*entities.Position, err error) {
	var (
		pdbts []*mappers.PositionDBType
		pos   *entities.Position
	)

	//Получение данных должностей
	pdbts, err = p.positionMapper.SelectAll()
	if err != nil {
		return
	}

	for _, pdbt := range pdbts {
		//Преобразование к типу сущности
		pos, err = pdbt.ToType()
		if err != nil {
			return
		}
		ps = append(ps, pos)
	}
	return
}
