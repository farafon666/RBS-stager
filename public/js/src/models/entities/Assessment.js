//Класс Ассесмент
export class Assessment{
    constructor(ID, name, date, disposer, candidates, employees, state){
        this.ID = ID;                   //ID
        this.name = name;               //Название
        this.date = date;               //Дата и время
        this.disposer = disposer;       //Распорядитель
        this.candidates = candidates;   //Кандидаты приглашённые на ассессмент
        this.employees = employees;     //Сотрудники присутствующие на ассессменте
        this.state = state;             //Состояние ассессмента
    }
};