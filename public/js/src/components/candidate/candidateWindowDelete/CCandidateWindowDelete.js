import {CandidateWindowDeleteView} from './CandidateWindowDeleteView.js';
import candidateModel from '../../../models/CandidateModel.js';

export class CCandidateWindowDelete{
    constructor(){
        this.view;
    };
    init(refreshTable){
        this.view = {
            window:        $$("windowDeleteCandidate"),
            buttonConfirm: $$("CandidateWindowDeleteConfirmButton"),
            buttonCancel:  $$("CandidateWindowDeleteCancelButton"),
            form: $$('formDeleteCandidate'),
            formFields: {
                lastname: $$('lastnameDeleteCandidate'),
                firstname: $$('firstnameDeleteCandidate'),
                middlename: $$('middlenameDeleteCandidate'),
                phonenumber: $$('phonenumberDeleteCandidate'),
                email: $$('emailDeleteCandidate'),
                birthdate: $$('birthdateDeleteCandidate'),
                gender: $$('genderDeleteCandidate')
            }
        };
        this.refreshTable = refreshTable;
    };

    config(){
        return webix.ui(CandidateWindowDeleteView()); 
    };

    attachEvents(){
        //Нажатие на кнопку подтверждения
        this.view.buttonConfirm.attachEvent("onItemClick", () => {
            // Удаление кандидата
            candidateModel.deleteCandidate(this.fetch()).then(() => {
                webix.message('Кандидат удалён!', 'success');
                this.view.window.hide();
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

    //Mетод размещения сущности в форме окна
    parse(values) {
        this.view.form.setValues(values)
    };

    //Метод получения сущности из формы окна
    fetch(){
        return this.view.form.getValues();
    };
}