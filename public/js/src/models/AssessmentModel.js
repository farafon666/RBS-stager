import Model from '../../helpers/model.js'

// AssessmentModel объект для работы(CRUD) с данными
class AssessmentModel extends Model {
  constructor() {
    super()
  }

  //Получение всех ассессментов
  getAssessments() {
    return this.get('/assessment/all')
  }

  //Получение ассессмента по ID
  getAssessmentByID(id) {
    return this.get(`/assessment/${id}`)
  }

  //Создание ассессмента
  createAssessment(assessment) {
    return this.post('/assessment/create', assessment)
  }

  //Обновление данных об ассессменте 
  updateAssessment(assessment) {
    return this.post('/assessment/update', assessment)
  }

  //Удаление ассессмент
  deleteAssessment(assessment) {
    return this.post('/assessment/delete', assessment)
  }

  //Получение всех сотрудников с ассессмента
  getEmployees(id) {
    return this.get(`/assessment/${id}/getEmployees`)
  }

  //Удаление всех сотрудников из ассессмента
  delEmployees(assessment) {
    return this.post('/assessment/delEmployees', assessment)
  }
  
  //Получение всех кандидатов с ассессмента
  getCandidates(id) {
    return this.get(`/assessment/${id}/getCandidates`)
  }

  //Удаление всех кандидатов из ассессмента
  delCandidates(assessment) {
    return this.post('/assessment/delCandidates', assessment)
  }
}

const assessmentModel = new AssessmentModel();
export default assessmentModel;