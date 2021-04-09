import {EmployeeWindowCreateView} from './EmployeeWindowCreateView.js';
import employeeModel from '../../../models/EmployeeModel.js';
import positionModel from '../../../models/PositionModel.js';

export class CEmployeeWindowCreate{
    constructor(){
        this.view;
    };

    init(refreshTable){
        this.view = {
            window:         $$("windowCreateEmployee"),
            buttonConfrim:  $$("EmployeeWindowCreateConfrimButton"),
            buttonCancel:   $$("EmployeeWindowCreateCancelButton"),
            form:           $$("formCreateEmployee"),
            formFields: {
                position: $$('positionCreateEmployee')
            }
        };
        this.refreshTable = refreshTable;
    };

    config(){
        return webix.ui(EmployeeWindowCreateView());
    };
    //Метод инициализации обработчиков событий 
    attachEvents(){
        this.view.buttonConfrim.attachEvent("onItemClick", () => {
            if(this.view.form.validate()){
                employeeModel.createEmployee(this.fetch()).then(() => {
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
                position.id = position.name,
                position.value = position.name
            });

            this.view.formFields.position.define('options', positions);
            this.view.formFields.position.refresh();
        });
    };

    //Метод отображения окна
    show() {
        this.view.window.show();
    };

    //Метод получения сущности из формы
    fetch(){
        let employee = this.view.form.getValues();
        if(employee.gender === 1){
            employee.gender = 'Мужчина';
        }else{
            employee.gender = 'Женщина';
        }
        return employee;
    };
}