//Класс Кандидат
export class Candidate{
    constructor(ID,lastname,firstname,middlename,birthdate,phone_number,email,gender){
        this.ID = ID;                     //ID
        this.lastname = lastname;         //Фамилия кандидата
        this.firstname = firstname;       //Имя кандидата
        this.middlename = middlename;     //Отчество кандидата
        this.birthdate = birthdate;       //Дата рождения кандидата
        this.phone_number = phone_number; //Номер телефона кандидата
        this.email = email;               //Адрес электронной почты кандидата
        this.gender = gender;             //Пол кандидата
    }
};

