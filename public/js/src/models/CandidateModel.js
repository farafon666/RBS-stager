import Model from '../../helpers/model.js'

// CandidateModel объект для работы(CRUD) с данными
class CandidateModel extends Model {
  constructor() {
    super()
  }

  //Получение всех кандидатов
  getCandidates() {
    return this.get('/candidate/all')
  }

  //Получение кандидата по ID
  getCandidateByID(id) {
    return this.get(`/candidate/${id}`)
  }

  //Создание кандидата
  createCandidate(candidate) {
    return this.post('/candidate/create', candidate)
  }

  //Обновление данных о кандидате
  updateCandidate(candidate) {
    return this.post('/candidate/update', candidate)
  }

  //Удаление кандидаты
  deleteCandidate(candidate) {
    return this.post('/candidate/delete', candidate)
  }
}

const candidateModel = new CandidateModel();
export default candidateModel;