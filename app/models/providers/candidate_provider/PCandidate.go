package candidate_provider

import (
	"assessmentManager/app/helpers"
	"assessmentManager/app/models/entities"
	"assessmentManager/app/models/mappers"
	"database/sql"
)

//Провайдер контроллера кандидатов
type PCandidate struct {
	candidateMapper *mappers.MCandidate
}

//Инициализация
func (p *PCandidate) Init() (err error) {
	var db *sql.DB //Экземпляр подключения к БД

	//Получение экземпляра подключения к БД
	db, err = helpers.GetDBConnection()
	if err != nil {
		return err
	}

	//Инициализация маппера кандидатов
	p.candidateMapper = new(mappers.MCandidate)
	p.candidateMapper.Init(db)

	return
}

//Метод получения всех кандидатов
func (p *PCandidate) GetAllCandidates() (cs []*entities.Candidate, err error) {
	var (
		cdbts []*mappers.CandidateDBType
		c     *entities.Candidate
	)

	//Получение данных кандидатов
	cdbts, err = p.candidateMapper.SelectAll()
	if err != nil {
		return
	}

	for _, cdbt := range cdbts {
		//Преобразование к типу сущности
		c, err = cdbt.ToType()
		if err != nil {
			return
		}

		cs = append(cs, c)
	}
	return
}

//Метод для получения кандидата по ID
func (p *PCandidate) GetCandidateByID(id int64) (c *entities.Candidate, err error) {
	var (
		cdbt *mappers.CandidateDBType
	)

	//Получение данных кандидата
	cdbt, err = p.candidateMapper.SelectByID(id)
	if err != nil {
		return
	}

	//Преобразование типа БД к типу сущности
	c, err = cdbt.ToType()
	if err != nil {
		return
	}

	return
}

//Метод создания кандидата
func (p *PCandidate) CreateCandidate(candidate *entities.Candidate) (c *entities.Candidate, err error) {
	var cdbt *mappers.CandidateDBType

	//Инициализация структур БД из структур сущности
	cdbt, err = cdbt.FromType(*candidate)
	if err != nil {
		return
	}

	//Добавление кандидата
	candidate.ID, err = p.candidateMapper.Insert(cdbt)
	if err != nil {
		return
	}

	return candidate, nil
}

//Метод обновления кандидата
func (p *PCandidate) UpdateCandidate(candidate *entities.Candidate) (c *entities.Candidate, err error) {
	var cdbt *mappers.CandidateDBType

	//Инициализация структуры БД из структуры сущности
	cdbt, err = cdbt.FromType(*candidate)
	if err != nil {
		return
	}

	//Обновление кандидата
	err = p.candidateMapper.Update(cdbt)
	if err != nil {
		return
	}

	return candidate, nil
}

//Метод удаления кандидата
func (p *PCandidate) DeleteCandidate(candidate *entities.Candidate) (err error) {
	var cdbt *mappers.CandidateDBType

	//Инициализация структуры БД из структуры сущности
	cdbt, err = cdbt.FromType(*candidate)
	if err != nil {
		return
	}

	//Удаление кандидата
	err = p.candidateMapper.Delete(cdbt)
	if err != nil {
		return
	}
	return
}
