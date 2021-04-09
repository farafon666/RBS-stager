//Класс Сотрудник
export class Employee{
    constructor(ID,lastname,firstname,middlename,position,birthdate,phone_number,email,gender){
        this.ID = ID;                     //ID
        this.lastname = lastname;         //Фамилия сотрудника
        this.firstname = firstname;       //Имя сотрудника
        this.middlename = middlename;     //Отчество сотрудника
        this.position = position;         //ID должности
        this.birthdate = birthdate;       //Дата рождения сотрудника
        this.phone_number = phone_number; //Номер телефона сотрудника
        this.email = email;               //Адрес электронной почты сотрудника
        this.gender = gender;             //Пол сотрудника
    }
};