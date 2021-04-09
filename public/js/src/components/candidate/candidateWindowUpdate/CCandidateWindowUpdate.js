import {CandidateWindowUpdateView} from './CandidateWindowUpdateView.js';
import candidateModel from '../../../models/CandidateModel.js';

export class CCandidateWindowUpdate{
    constructor(){
        this.view;
    };

    init(refreshTable){
        this.view = {
            window:        $$("windowUpdateCandidate"),
            buttonConfirm: $$("CandidateWindowUpdateConfirmButton"),
            buttonCancel:  $$("CandidateWindowUpdateCancelButton"),
            form: $$('formUpdateCandidate'),
            formFields: {
                lastname: $$('lastnameUpdateCandidate'),
                firstname: $$('firstnameUpdateCandidate'),
                middlename: $$('middlenameUpdateCandidate'),
                phonenumber: $$('phonenumberUpdateCandidate'),
                email: $$('emailUpdateCandidate'),
                birthdate: $$('birthdateUpdateCandidate'),
                gender: $$('genderUpdateCandidate')
            }
        };
        this.refreshTable = refreshTable;
    };

    config(){
        return webix.ui(CandidateWindowUpdateView()); 
    };

    attachEvents(){
        //Нажатие на кнопку подтверждения
        this.view.buttonConfirm.attachEvent("onItemClick", () => {
            candidateModel.updateCandidate(this.fetch()).then(() => {
                webix.message('Данные о кандидате изменены!', 'success');
                this.view.window.hide();
                this.view.form.clear();
                this.refreshTable();
            });
        });
        //Нажатие на кнопку отмены
        this.view.buttonCancel.attachEvent("onItemClick", () => {
            this.view.window.hide();
            this.view.form.clear();
        });
    };
    //Метод отображения окна
    show(){
        this.view.window.show();
    };
    //Метод получения сущности из форм окна
    fetch(){
        let candidate = this.view.form.getValues();
        if (candidate.gender === 1) {
            candidate.gender = 'Мужчина';
        }else{
            candidate.gender = 'Женщина';
        }
        return candidate;
    };
    //Метод загрузки данных в форму
    parse(candidate){
        if(candidate.gender === 'Мужчина'){
            candidate.gender = 1;
        }else{
            candidate.gender = 2;
        }
        this.view.form.setValues(candidate);

    };
}