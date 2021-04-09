package employee_provider

import (
	"assessmentManager/app/helpers"
	"assessmentManager/app/models/entities"
	"assessmentManager/app/models/mappers"
	"database/sql"

	"github.com/revel/revel"
)

//Провайдер контроллера сотрудников
type PEmployee struct {
	employeeMapper *mappers.MEmployee
	positionMapper *mappers.MPosition
}

//Инициализация
func (p *PEmployee) Init() (err error) {
	var db *sql.DB //Экземпляр подключения к БД

	//Получение экземпляра подключения к БД
	db, err = helpers.GetDBConnection()
	if err != nil {
		revel.AppLog.Errorf("PEmployee.Init : helpers.GetDBConnection, %s\n", err)
		return err
	}

	//Инициализация маппера сотрудников
	p.employeeMapper = new(mappers.MEmployee)
	p.employeeMapper.Init(db)
	//Инициализация маппера должностей
	p.positionMapper = new(mappers.MPosition)
	p.positionMapper.Init(db)

	return
}

//Метод получения всех сотрудников
func (p *PEmployee) GetAllEmployees() (es []*entities.Employee, err error) {
	var (
		edbts []*mappers.EmployeeDBType
		e     *entities.Employee
	)

	//Получение данных сотрудников
	edbts, err = p.employeeMapper.SelectAll()
	if err != nil {
		revel.AppLog.Errorf("PEmployee.GetAllEmployees : p.employeeMapper.SelectAll, %s\n", err)
		return
	}

	for _, edbt := range edbts {
		//Преобразование к типу сущности
		e, err = edbt.ToType()
		if err != nil {
			return
		}

		//Получение значения должности по ключу
		e.Position, err = p.positionMapper.PositionNameByID(edbt.Fk_position)
		if err != nil {
			return
		}

		es = append(es, e)
	}
	return
}

//Метод для получения сотрудника по ID
func (p *PEmployee) GetEmployeeByID(id int64) (e *entities.Employee, err error) {
	var (
		edbt *mappers.EmployeeDBType
	)

	//Получение данных сотрудника
	edbt, err = p.employeeMapper.SelectByID(id)
	if err != nil {
		return
	}

	//Преобразование типа БД к типу сущности
	e, err = edbt.ToType()
	if err != nil {
		return
	}

	//Получение значения должности по ключу
	e.Position, err = p.positionMapper.PositionNameByID(edbt.Fk_position)
	if err != nil {
		return
	}

	return
}

//Метод создания сотрудника
func (p *PEmployee) CreateEmployee(employee *entities.Employee) (e *entities.Employee, err error) {
	var edbt *mappers.EmployeeDBType

	//Инициализациия структур БД из структур сущности
	edbt, err = edbt.FromType(*employee)
	if err != nil {
		return
	}

	//Получение внешнего ключа на должность
	edbt.Fk_position, err = p.positionMapper.IDByPositionName(employee.Position)
	if err != nil {
		return
	}

	//Добавление сотрудника
	employee.ID, err = p.employeeMapper.Insert(edbt)
	if err != nil {
		return
	}

	return employee, nil
}

//Метод обновления сотрудника
func (p *PEmployee) UpdateEmployee(employee *entities.Employee) (e *entities.Employee, err error) {
	var edbt *mappers.EmployeeDBType

	//Инициализация структуры БД из структуры сущности
	edbt, err = edbt.FromType(*employee)
	if err != nil {
		return
	}

	//Получение внешнего ключа на должность
	edbt.Fk_position, err = p.positionMapper.IDByPositionName(employee.Position)
	if err != nil {
		return
	}

	//Обновление сотрудника
	err = p.employeeMapper.Update(edbt)
	if err != nil {
		return
	}

	return employee, nil
}

//Метод удаления сотрудника
func (p *PEmployee) DeleteEmployee(employee *entities.Employee) (err error) {
	var edbt *mappers.EmployeeDBType

	//Инициализация структуры БД из структуры сущности
	edbt, err = edbt.FromType(*employee)
	if err != nil {
		return
	}

	//Удаление сотрудника
	err = p.employeeMapper.Delete(edbt)
	if err != nil {
		return
	}
	return
}
