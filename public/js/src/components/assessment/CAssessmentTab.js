import {AssessmentTabView} from './AssessmentTabView.js'
import {CAssessmentWindowCreate} from './assessmentWindowCreate/CAssessmentWindowCreate.js';
import {CAssessmentWindowUpdate} from './assessmentWindowUpdate/CAssessmentWindowUpdate.js';
import {CAssessmentWindowDelete} from './assessmentWindowDelete/CAssessmentWindowDelete.js';
import {CAssessmentWindowInfo} from './assessmentWindowInfo/CAssessmentWindowInfo.js';
import assessmentModel from '../../models/AssessmentModel.js';

export class CAssessmentTab{
    constructor(){
        this.view;
    }
    init(){
        this.view = {
            datatable: $$('dataAssessment'),            
            buttons: {
                buttonCreateAssessment: $$('buttonCreateAssessment'),
                buttonUpdateAssessment: $$('buttonUpdateAssessment'),
                buttonDeleteAssessment: $$('buttonDeleteAssessment'),
                buttonViewAssessment:   $$('buttonViewAssessment'),
            }
        };
        //Создание экземпляра класса окна добваления ассессмента
        this.createWindowAssessment = new CAssessmentWindowCreate();
        this.createWindowAssessment.config();
        this.createWindowAssessment.init(() => this.refreshTable());
        this.createWindowAssessment.attachEvents();

        //Создание экземпляра класса окна редактирования ассессмента
        this.updateWindowAssessment = new CAssessmentWindowUpdate();
        this.updateWindowAssessment.config();
        this.updateWindowAssessment.init(() => this.refreshTable());
        this.updateWindowAssessment.attachEvents();

        //Создание экземпляра класса окна удаления ассессмента
        this.deleteWindowAssessment = new CAssessmentWindowDelete();
        this.deleteWindowAssessment.config();
        this.deleteWindowAssessment.init(() => this.refreshTable());
        this.deleteWindowAssessment.attachEvents();

        //Создание экземпляра класса окна просмотра подробной информации об ассессменте
        this.infoWindowAssessment = new CAssessmentWindowInfo();
        this.infoWindowAssessment.config();
        this.infoWindowAssessment.init();
        this.infoWindowAssessment.attachEvents();

        this.attachEvents();
    };
    attachEvents(){
        //Загрузка начальных данных в datatable
        this.refreshTable();
        //Создание нового ассессмента
        this.view.buttons.buttonCreateAssessment.attachEvent('onItemClick',()=>{
            this.createAssessment();
        });
        //Изменение ассессмента
        this.view.buttons.buttonUpdateAssessment.attachEvent('onItemClick',()=>{
            this.updateAssessment();
        });
        //Удаление ассессмента
        this.view.buttons.buttonDeleteAssessment.attachEvent('onItemClick',()=>{
            this.deleteAssessment();
        });
        //Просмотр информации об ассессменте
        this.view.buttons.buttonViewAssessment.attachEvent('onItemClick',()=>{
            this.viewAssessment();
        });        
    };
    config(){
        return AssessmentTabView();
    };
    //Метод обновления таблицы ассессментов
    refreshTable(assessments){
        if (assessments) {
            this.view.datatable.clearAll();
            this.view.datatable.parse(assessments);
        }else{
            assessmentModel.getAssessments().then((assessments) => {
                this.view.datatable.clearAll();
                this.view.datatable.parse(assessments);
            });
        }
    }
    //Метод создания нового ассессмента
    createAssessment(){
        this.createWindowAssessment.show();
    };
    //Метод изменения данных об ассессменте
    updateAssessment(){
        let selected = this.view.datatable.getSelectedItem();

        if (!selected) {
            webix.message('Выделите строку!', 'error')
            return
        }else if(selected.state === 'Заархивирован'){
            webix.message('Этот ассессмент нельзя изменить!', 'error')
            return
        }

        assessmentModel.getAssessmentByID(selected.ID).then((assessment) => {
            this.updateWindowAssessment.parse(assessment);
            this.updateWindowAssessment.show(assessment);
        });
    };
    //Метод удаления ассессмента
    deleteAssessment(){
        let selected = this.view.datatable.getSelectedItem();
        if (!selected) {
            webix.message('Выделите строку!', 'error')
            return
        }else if(selected.state === 'Заархивирован'){
            webix.message('Этот ассессмент нельзя удалить!', 'error')
            return
        }

        assessmentModel.getAssessmentByID(selected.ID).then((assessment) => {
            this.deleteWindowAssessment.parse(assessment);
            this.deleteWindowAssessment.show();
        });
    };
    //Метод просмотра подобной инфоромации об ассессменте
    viewAssessment(){
        let selected = this.view.datatable.getSelectedItem();

        if(!selected){
            webix.message('Выделите строку!','error');
            return
        }

        assessmentModel.getAssessmentByID(selected.ID).then((assessment) => {
            this.infoWindowAssessment.parse(assessment);
            this.infoWindowAssessment.show();
        });
    };
};