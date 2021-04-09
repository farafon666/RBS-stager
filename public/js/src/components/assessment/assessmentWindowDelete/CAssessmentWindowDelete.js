import {AssessmentWindowDeleteView} from './AssessmentWindowDeleteView.js';
import assessmentModel from '../../../models/AssessmentModel.js';
import employeeModel from '../../../models/EmployeeModel.js';

export class CAssessmentWindowDelete{
    constructor(){
        this.view;
    };

    init(refreshTable){
        this.view = {
            window:        $$("windowDeleteAssessment"),
            confirmButton: $$("AssessmentWindowDeleteConfirmButton"),
            cancelButton:  $$("AssessmentWindowDeleteCancelButton"),
            form:          $$("formDeleteAssessment"),
            formFields: {
                ID:         $$('IDDeleteAssessment'),
                date:       $$('dateDeleteAssessment'),
                state:      $$('statusDeleteAssessment'),
                disposer:   $$('disposerDeleteAssessment'),
                employees:  $$('employeesDeleteAssessment'),
                candidates: $$('candidatesDeleteAssessment')
            }
        };
        this.refreshTable = refreshTable;
    };

    config(){
        return webix.ui(AssessmentWindowDeleteView()); 
    };

    attachEvents(){
        //Нажатие на кнопку отмены
        this.view.cancelButton.attachEvent("onItemClick", ()=>{
            this.view.window.hide();
            this.view.form.clear();
        });
        //Нажатие на кнопку подтверждения
        this.view.confirmButton.attachEvent("onItemClick", ()=>{
            assessmentModel.deleteAssessment(this.fetch()).then(() => {
                this.view.window.hide();
                this.view.form.clear();
                this.refreshTable();
            });            
        });
        
    };

    //Метод отображения окна
    show(){
        this.view.window.show();
    };

    //Метод парсинга сущности в форму окна
    parse(assessment){
        this.view.formFields.ID.setValue(assessment.ID)
        this.view.formFields.date.setValue(`Дата и время: ${assessment.date}`);
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

    //Метод получения выделеной сущности
    fetch(){
        let assessment = this.view.form.getValues();
        assessment.disposer = 0;
        return assessment;
    }
}