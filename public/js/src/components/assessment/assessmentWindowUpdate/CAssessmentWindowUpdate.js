import {AssessmentWindowUpdateView} from './AssessmentWindowUpdateView.js';
import employeeModel from '../../../models/EmployeeModel.js';
import candidateModel from '../../../models/CandidateModel.js';
import assessmentModel from '../../../models/AssessmentModel.js';
import stateAssessmentModel from '../../../models/StateAssessmentModel.js';
import {FormatDate} from '../../../../helpers/dateFormatter.js';

export class CAssessmentWindowUpdate{
    constructor(){
        this.view;
    };

    init(refreshTable){
        this.view = {
            window:        $$("windowUpdateAssessments"),
            confirmButton: $$("AssessmentWindowUpdateConfirmButton"),
            cancelButton:  $$("AssessmentWindowUpdateCancelButton"),
            form:          $$("formUpdateAssessments"),
            formFields: {
                date:       $$('dateAssessmentUpdate'),
                state:      $$('stateUpdateAssessment'),
                disposer:   $$('disposerUpdateAssessment'),
                employees:  $$('dbllistUpdateEmployees'),
                candidates: $$('dbllistUpdateCandidates')
            }
        };
        this.refreshTable = refreshTable;
    };

    config(){
        return webix.ui(AssessmentWindowUpdateView()); 
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
        //Нажатие на кнопку отмены
        this.view.cancelButton.attachEvent("onItemClick", ()=>{
            this.view.window.hide();
            this.view.form.clear();
            this.view.formFields.employees.setValue([]);
            this.view.formFields.candidates.setValue([]);
        });
        //Нажатие на кнопку подтверждения
        this.view.confirmButton.attachEvent("onItemClick", ()=>{
            if(this.view.form.validate()){
                assessmentModel.delEmployees(this.fetch()).then();
                assessmentModel.delCandidates(this.fetch()).then();

                assessmentModel.updateAssessment(this.fetch()).then(() => {
                    this.view.window.hide();
                    this.view.form.clear();
                    this.view.formFields.employees.setValue([]);
                    this.view.formFields.candidates.setValue([]);
                    this.refreshTable();
                });
            }else{
                webix.message('Заполните все поля','error')
            }
        });
        //Событие открытия окна
        this.view.window.attachEvent("onShow", ()=>{
            //Загрузка состояний ассессментов
            stateAssessmentModel.getStatesAssessment().then((states) => {
                states.map((state) => {
                    state.id = state.state,
                    state.value = state.state
                });

                //Загрузка данных для выбора состояний ассессмента
                this.view.formFields.state.define('options', states);
                this.view.formFields.state.refresh();
            });

            //Загрузка сотрудников
            employeeModel.getEmployees().then((employees) => {
                employees.map((disposer) => {
                    disposer.id = disposer.ID,
                    disposer.value = disposer.lastname + ' ' + disposer.firstname + ' ' + disposer.middlename
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

    //Метод отображения окна
    show(){
        this.view.window.show();
    };

    //Метод парсинга сущности в форму окна
    parse(assessment){
        assessmentModel.getEmployees(assessment.ID).then((employees) => {
            let strEmployee = '';
            employees.forEach((employee) => {
                strEmployee += employee.ID + ','
            });
            this.view.formFields.employees.setValue(strEmployee)
        });
        
        assessmentModel.getCandidates(assessment.ID).then((candidates) => {
            let strCandidate = '';
            candidates.forEach((candidate) => {
                strCandidate += candidate.ID + ','
            });
            this.view.formFields.candidates.setValue(strCandidate)
        });

        this.view.form.setValues(assessment);
    };
    
    //Метод получения сущности из формы окна
    fetch(){
        let assessment = this.view.form.getValues();
        assessment.date = FormatDate(assessment.date);
        assessment.employees = assessment.employees.split(',');
        assessment.candidates = assessment.candidates.split(',');
        return assessment;
    };
}