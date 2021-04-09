import Model from '../../helpers/model.js';

//StateAssessmentModel.js объект для работы (CRUD) с данными
class StateAssessmentModel extends Model {
    constructor() {
        super()
    }

    //Получение всех состояний
    getStatesAssessment(){
        return this.get('stateAssessment/all')
    }
}

const stateAssessmentModel = new StateAssessmentModel();
export default stateAssessmentModel;