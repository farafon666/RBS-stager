import {CandidateWindowCreateView} from './CandidateWindowCreateView.js';
import candidateModel from '../../../models/CandidateModel.js';

export class CCandidateWindowCreate{
    constructor(){
        this.view;
    };
    init(refreshTable){
        this.view = {
            window:        $$("windowCreateCandidate"),
            buttonConfirm: $$("CandidateWindowCreateConfirmButton"),
            buttonCancel:  $$("CandidateWindowCreateCancelButton"),
            form: $$('formCreateCandidate'),
            formFields: {
                lastname: $$('lastnameCreateCandidate'),
                firstname: $$('firstnameCreateCandidate'),
                middlename: $$('middlenameCreateCandidate'),
                phonenumber: $$('phonenumberCreateCandidate'),
                email: $$('emailCreateCandidate'),
                birthdate: $$('birthdateCreateCandidate'),
                gender: $$('genderCreateCandidate')
            }
        };
        this.refreshTable = refreshTable;
    };
    config(){
        return webix.ui(CandidateWindowCreateView()); 
    };
    attachEvents(){
        this.view.buttonCancel.attachEvent("onItemClick", () => {
            this.view.window.hide();
            this.view.form.clear();
        });
        this.view.buttonConfirm.attachEvent("onItemClick", () => {
            if (this.view.form.validate()){
                candidateModel.createCandidate(this.fetch()).then(() => {
                    webix.message('Кандидат добавлен!','success')
                    this.view.form.clear();
                    this.view.window.hide();
                    this.refreshTable();
                });
            }else{
                webix.message("Заполните все поля!","error")
            }
        });
    };
    show(){  
        this.view.window.show();
    };
    // Метод получения сущности из формы окна
    fetch(){
        let candidate = this.view.form.getValues();
        if (candidate.gender === 1){
            candidate.gender = 'Мужчина';
        }else{
            candidate.gender = 'Женщина';
        }
        return candidate;
        
    }
}