import {AssessmentWindowInProcess} from './AssessmentWindowInProcessView.js';

export class CAssessmentWindowInProcess{
    constructor(){
        this.view;
    };

    init(){
        this.view = {
            window: $$('WindowInProcessAssessment')
        }
    };

    config(){
        return webix.ui(AssessmentWindowInProcess());
    };

    attachEvents(){};

    //Метод отображения окна 
    show(){};
    
    //Метод заполнения формы окна данными сущности
    parse(){};

    //Метод получения сущности из формы
    fetch(){};
};
