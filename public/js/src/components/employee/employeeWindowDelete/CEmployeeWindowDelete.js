import {EmployeeWindowDeleteView} from './EmployeeWindowDeleteView.js';
import employeeModel from '../../../models/EmployeeModel.js';

export class CEmployeeWindowDelete{
    constructor(){
        this.view;
    };

    init(refreshTable){
        this.view = {
            window:         $$("windowDeleteEmployee"),
            buttonConfrim:  $$("EmployeeWindowDeleteConfrimButton"),
            buttonCancel:   $$("EmployeeWindowDeleteCancelButton"),
            form:           $$("formDeleteEmployee")
        };
        this.refreshTable = refreshTable;
    };

    config(){
        return webix.ui(EmployeeWindowDeleteView());
    };

    attachEvents(){
        //Нажатие на кнопку подтверждения
        this.view.buttonConfrim.attachEvent("onItemClick", () => {
            // Удаление сотрудника
            employeeModel.deleteEmployee(this.fetch()).then(() => {
                webix.message('Сотрудник удалён!');
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

    //Метод получения сущности из формы
    fetch(){
        return this.view.form.getValues();
    };

    //Метод загрузки сущности в форму
    parse(employee){
        this.view.form.setValues(employee);
    };
}