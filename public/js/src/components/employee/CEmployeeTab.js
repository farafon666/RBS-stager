import {EmployeeTabView} from './EmployeeTabView.js';
import {CEmployeeWindowCreate} from './employeeWindowCreate/CEmployeeWindowCreate.js';
import {CEmployeeWindowUpdate} from './employeeWindowUpdate/CEmployeeWindowUpdate.js';
import {CEmployeeWindowDelete} from './employeeWindowDelete/CEmployeeWindowDelete.js';
import employeeModel from '../../models/EmployeeModel.js';

export class CEmployeeTab{
    constructor(){
        this.view;
    }
    init(){
        this.view = {
            datatable: $$('dataEmployee'),            
            buttons: {
                buttonCreateEmployee: $$('buttonCreateEmployee'),
                buttonUpdateEmployee: $$('buttonUpdateEmployee'),
                buttonDeleteEmployee: $$('buttonDeleteEmployee'),
            }
        };

        //Создание экземпляра класса окна добавления сотрудников
        this.createWindowEmployee = new CEmployeeWindowCreate();
        this.createWindowEmployee.config();
        this.createWindowEmployee.init(() => this.refreshTable());
        this.createWindowEmployee.attachEvents();

        //Создание экземпляра класса окна изменения сотрудника
        this.udpateWindowEmployee = new CEmployeeWindowUpdate(); 
        this.udpateWindowEmployee.config();
        this.udpateWindowEmployee.init(() => this.refreshTable());
        this.udpateWindowEmployee.attachEvents();
        
        //Создание экземпляра класса окна удаления сотрудника
        this.deleteWindowEmployee = new CEmployeeWindowDelete();
        this.deleteWindowEmployee.config();
        this.deleteWindowEmployee.init(() => this.refreshTable());
        this.deleteWindowEmployee.attachEvents();

        this.attachEvents();
    };
    attachEvents(){
        //Загрузка первичных данных в таблицу
        this.refreshTable();
        //Создание нового сотрудника
        this.view.buttons.buttonCreateEmployee.attachEvent('onItemClick',()=>{
            this.createEmployee();
        });
        //Изменение сотрудника
        this.view.buttons.buttonUpdateEmployee.attachEvent('onItemClick',()=>{
            this.updateEmployee();
        });
        //Удаление сотрудника
        this.view.buttons.buttonDeleteEmployee.attachEvent('onItemClick',()=>{
            this.deleteEmployee();
        });        
    };
    config(){
        return EmployeeTabView();
    };
    //Метод обновления таблицы сотрудников
    refreshTable(employees) {
        if (employees) {
            this.view.datatable.clearAll();
            this.view.datatable.parse(employees);
        }else{
            employeeModel.getEmployees().then((employees) => {
                this.view.datatable.clearAll();
                this.view.datatable.parse(employees);
            });
        }
    };

    //Метод создания нового сотрудника
    createEmployee(){
        this.createWindowEmployee.show();
    };

    //Метод изменения данных о сотруднике
    updateEmployee(){
        //Получение выделенного элемента
        let selected = this.view.datatable.getSelectedItem();

        if (!selected) {
            webix.message('Выделите строку!', 'error')
            return
        }

        employeeModel.getEmployeeByID(selected.ID).then((employee) => {
            this.udpateWindowEmployee.parse(employee);
            this.udpateWindowEmployee.show();
        })
    };

    //Метод удаления сотрудника
    deleteEmployee(){
        //Получение выделенного элемента
        let selected = this.view.datatable.getSelectedItem();

        if(!selected){
            webix.message('Выделите строку!', 'error')
            return
        }

        employeeModel.getEmployeeByID(selected.ID).then((employee) => {
            this.deleteWindowEmployee.parse(employee);
            this.deleteWindowEmployee.show();
        });
    };
};