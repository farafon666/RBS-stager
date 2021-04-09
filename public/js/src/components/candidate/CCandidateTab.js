import {CandidateTabView} from './CandidateTabView.js';
import {CCandidateWindowCreate} from './candidateWindowCreate/CCandidateWindowCreate.js';
import {CCandidateWindowUpdate} from './candidateWindowUpdate/CCandidateWindowUpdate.js';
import {CCandidateWindowDelete} from './candidateWindowDelete/CCandidateWindowDelete.js';
import candidateModel from '../../models/CandidateModel.js';

export class CCandidateTab{
    constructor(){
        this.view;
    }
    init(){
        this.view = {
            datatable: $$('dataCandidate'),            
            buttons: {
                buttonCreateCandidate: $$('buttonCreateCandidate'),
                buttonUpdateCandidate: $$('buttonUpdateCandidate'),
                buttonDeleteCandidate: $$('buttonDeleteCandidate'),
            }
        };

        this.createWindowCandidate = new CCandidateWindowCreate();
        this.createWindowCandidate.config();
        this.createWindowCandidate.init(() => this.refreshTable());
        this.createWindowCandidate.attachEvents();

        this.updateWindowCandidate = new CCandidateWindowUpdate();
        this.updateWindowCandidate.config();
        this.updateWindowCandidate.init(() => this.refreshTable());
        this.updateWindowCandidate.attachEvents();

        this.deleteWindowCandidate = new CCandidateWindowDelete();
        this.deleteWindowCandidate.config();
        this.deleteWindowCandidate.init(() => this.refreshTable());
        this.deleteWindowCandidate.attachEvents();
        
        this.attachEvents();
    };
    attachEvents(){
        //Загрузка начальных данных в таблицу
        this.refreshTable();
        //Создание нового кандидата
        this.view.buttons.buttonCreateCandidate.attachEvent('onItemClick',()=>{
            this.createCandidate();
        });
        //Изменение кандидата
        this.view.buttons.buttonUpdateCandidate.attachEvent('onItemClick',()=>{
            this.updateCandidate();
        });
        //Удаление кандидата
        this.view.buttons.buttonDeleteCandidate.attachEvent('onItemClick',()=>{
            this.deleteCanidate();
        });        
    };
    config(){
        return CandidateTabView();
    };
    //Метод обновления таблицы кандидатов
    refreshTable(candidates){
        if (candidates) {
            this.view.datatable.clearAll();
            this.view.datatable.parse(candidates);
        }else{
            candidateModel.getCandidates().then((candidates) => {
                this.view.datatable.clearAll();
                this.view.datatable.parse(candidates);
            });
        }
    };

    //Метод создания нового кандидата
    createCandidate(){
        this.createWindowCandidate.show();
    };

    //Метод изменения данных о кандидате
    updateCandidate(){
        //Получение выделенного элемента
        let selected = this.view.datatable.getSelectedItem();

        if (!selected) {
            webix.message('Выделите строку!', 'error')
            return
        }

        candidateModel.getCandidateByID(selected.ID).then((candidate) => {
            this.updateWindowCandidate.parse(candidate);
            this.updateWindowCandidate.show();
        })
    };
    
    //Метод удаления данных о кандидате
    deleteCanidate(){
        //Получаем выделенный элемент
        let selected = this.view.datatable.getSelectedItem();

        if (!selected) {
            webix.message('Выделите строку!','error');   
            return
        }

        candidateModel.getCandidateByID(selected.ID).then((candidate) => {
            this.deleteWindowCandidate.parse(candidate);
            this.deleteWindowCandidate.show();
        });
    };
};