import Model from '../../helpers/model.js';

//PositionModel объект для работы (CRUD) с данными
class PositionModel extends Model {
    constructor() {
        super()
    }

    //Получение должностей
    getPositions(){
        return this.get('/position/all')
    }
}

const positionModel = new PositionModel();
export default positionModel;