import Model from '../../helpers/model.js';

// EmployeeModel объект для работы(CRUD) с данными
class EmployeeModel extends Model {
  constructor() {
    super()
  }

  //Получение всех сотрудников
  getEmployees() {
    return this.get('/employee/all')
  }

  //Получение сотрудника по ID
  getEmployeeByID(id) {
    return this.get(`/employee/${id}`)
  }

  //Создание сотрудника
  createEmployee(employee) {
    return this.post('/employee/create', employee)
  }

  //Обновление данных о сотруднике
  updateEmployee(employee) {
    return this.post('/employee/update', employee)
  }

  //Удаление сотрудника
  deleteEmployee(employee) {
    return this.post('/employee/delete', employee)
  }
}

const employeeModel = new EmployeeModel();
export default employeeModel;