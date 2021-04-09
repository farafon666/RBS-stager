import {AssessmentWindowCreateView} from './AssessmentWindowCreateView.js';
import employeeModel from '../../../models/EmployeeModel.js';
import candidateModel from '../../../models/CandidateModel.js';
import assessmentModel from '../../../models/AssessmentModel.js';
import {FormatDate} from '../../../../helpers/dateFormatter.js';

export class CAssessmentWindowCreate{
    constructor(){
        this.view;
    };

    init(refreshTable){
        this.view = {
            window:        $$("windowCreateAssessments"),
            confirmButton: $$("AssessmentWindowCreateConfirmButton"),
            cancelButton:  $$("AssessmentWindowCreateCancelButton"),
            form:          $$("formCreateAssessments"),
            formFields: {
                date:       $$('dateCreateAssessment'),
                disposer:   $$('disposerCreateAssessment'),
                employees:  $$('dbllistCreateEmployees'),
                candidates: $$('dbllistCreateCandidates')
            }
        };

        this.refreshTable = refreshTable;
    };

    config(){
        return webix.ui(AssessmentWindowCreateView()); 
    };

    attachEvents(){
        this.view.formFields.employees.attachEvent("onChange", () => {
            let disp = this.view.formFields.disposer.getValue();
            let employees = this.view.formFields.employees.getValue().split(',');
            let strFinal = '';
            employees.forEach((employee) => {
                if (employee != disp){
                    strFinal += employee + ',';
                }
            });
            this.view.formFields.employees.setValue(strFinal);
        });
        this.view.cancelButton.attachEvent("onItemClick", () => {
            this.view.window.hide();
            this.view.form.clear();
            this.view.formFields.employees.setValue([]);
            this.view.formFields.candidates.setValue([]);
        });
        this.view.confirmButton.attachEvent("onItemClick", () => {
            if(this.view.form.validate()){
                assessmentModel.createAssessment(this.fetch()).then(() => {
                    this.view.window.hide();
                    this.view.form.clear();
                    this.view.formFields.employees.setValue([]);
                    this.view.formFields.candidates.setValue([]);
                    this.refreshTable();
                });                
            }else{
                webix.message("Заполните все поля!","error")
            }
        });
        this.view.window.attachEvent("onShow", () => {
            //Загрузка сотрудников
            employeeModel.getEmployees().then((employees) => {
                employees.map((employee) => {
                    employee.id = employee.ID,
                    employee.value = employee.lastname + ' ' + employee.firstname + ' ' + employee.middlename
                });

                //Загрузка данных для выбора распорядителя
                this.view.formFields.disposer.define('options', employees);
                this.view.formFields.disposer.refresh();

                //Загрузка данных для выбора членов комиссии
                this.view.formFields.employees.define('data', employees);
            });

            //Загрузка кандидатов
            candidateModel.getCandidates().then((candidates) => {
                candidates.map((candidate) => {
                    candidate.id = candidate.ID,
                    candidate.value = candidate.lastname + ' ' + candidate.firstname + ' ' + candidate.middlename
                });

                //Загрузка данных для выбора приглашённых кандидатов
                this.view.formFields.candidates.define('data', candidates);
            });
        });
    };
    
    //Метод отображения окна создания ассессментов
    show(){
        this.view.window.show();
    };

    //Метод получения сущности из формы
    fetch(){
        let assessment = this.view.form.getValues();
        assessment.date = FormatDate(assessment.date);
        assessment.title = 'Ассессмент';
        assessment.state = 'Предстоит';
        assessment.employees = assessment.employees.split(',');
        assessment.candidates = assessment.candidates.split(',');
        return assessment;
    };
}