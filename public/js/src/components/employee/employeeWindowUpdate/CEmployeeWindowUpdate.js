import {EmployeeWindowUpdateView} from './EmployeeWindowUpdateView.js';
import employeeModel from '../../../models/EmployeeModel.js';
import positionModel from '../../../models/PositionModel.js';

export class CEmployeeWindowUpdate{
    constructor(){
        this.view;
    };

    init(refreshTable){
        this.view = {
            window:         $$("windowUpdateEmployee"),
            buttonConfrim:  $$("EmployeeWindowUpdateConfrimButton"),
            buttonCancel:   $$("EmployeeWindowUpdateCancelButton"),
            form:           $$("formUpdateEmployee"),
            formFields: {
                position: $$('positionUpdateEmployee')
            }
        };
        this.refreshTable = refreshTable;
    };

    config(){
        return webix.ui(EmployeeWindowUpdateView());
    };
    //Метод иниициализации обработчиков событий
    attachEvents(){
        this.view.buttonConfrim.attachEvent("onItemClick", () => {
            if (this.view.form.validate()){
                employeeModel.updateEmployee(this.fetch()).then(() => {
                    webix.message('Данные о сотруднике изменены!', 'success');
                    this.view.window.hide();
                    this.view.form.clear();
                    this.refreshTable();
                });
            }else{
                webix.message("Заполните все поля!","error")
            }
        });
        this.view.buttonCancel.attachEvent("onItemClick", () => {
            this.view.window.hide();
            this.view.form.clear();
        });

        //Подгрузка должностей
        positionModel.getPositions().then((positions) => {
            positions.map((position) => {
                position.ID = position.name,
                position.value = position.name
            });

            this.view.formFields.position.define('options', positions);
            this.view.formFields.position.refresh();
        });
    };

    //Метод отображеения окна
    show(){
        this.view.window.show();
    };
    //Метод получения сущности из форм
    fetch(){
        let employee = this.view.form.getValues();
        if(employee.gender === 1){
            employee.gender = 'Мужчина';
        }else{
            employee.gender = 'Женщина';
        }
        return employee;
    };
    //Метод парсинга данных сущности в форму
    parse(employee){
        if(employee.gender === 'Мужчина'){
            employee.gender = 1;
        }else{
            employee.gender = 2;
        }
        this.view.form.setValues(employee);
    };
}