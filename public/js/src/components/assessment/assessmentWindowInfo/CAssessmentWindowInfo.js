import {AssessmentWindowInfoView} from './AssessmentWindowInfoView.js';
import assessmentModel from './../../../models/AssessmentModel.js';
import employeeModel from '../../../models/EmployeeModel.js'

export class CAssessmentWindowInfo{
    constructor(){
        this.view;
    };

    init(){
        this.view = {
            window:        $$("windowViewAssessment"),
            cancelButton:  $$("AssessmentWindowInfoCancelButton"),
            form:          $$("formViewAssessment"),
            formFields: {
                date:       $$('dateViewAssessment'),
                state:      $$('statusViewAssessment'),
                disposer:   $$('disposerViewAssessment'),
                employees:  $$('employeesViewAssessment'),
                candidates: $$('candidatesViewAssessment')
            }
        }
    };

    config(){
        return webix.ui(AssessmentWindowInfoView()); 
    };

    attachEvents(){
        this.view.cancelButton.attachEvent("onItemClick", ()=>{
            this.view.window.hide();
            this.view.form.clear();
        });
    };

    //Метод отображения окна
    show(){
        this.view.window.show();
    };

    //Метод парсинга сущности в форму 
    parse(assessment){
        this.view.formFields.date.setValue(`Дата: ${assessment.date}`);
        this.view.formFields.state.setValue(`Состояние: ${assessment.state}`);
        //Получение распорядителя
        employeeModel.getEmployeeByID(assessment.disposer).then((employee) => {
            this.view.formFields.disposer.setValue(`Распорядитель: ${employee.lastname} ${employee.firstname} ${employee.middlename}`);
        });
        //Получение сотрудников
        assessmentModel.getEmployees(assessment.ID).then((employees) => {
            this.view.formFields.employees.clearAll();
            this.view.formFields.employees.parse(employees);
        });
        //Получение кандидатов
        assessmentModel.getCandidates(assessment.ID).then((candidates) => {
            this.view.formFields.employees.clearAll();
            this.view.formFields.candidates.parse(candidates);
        });
    };
}