//Окно добавления сотрудника
export function EmployeeWindowCreateView(){
  return {
    view:"window",
    position:"center",
    head:"Добавление нового сотрудника",
    id: "windowCreateEmployee",
    close: true,
    modal: true,
    body:{
      view:"form",
      id:"formCreateEmployee",
      width: 1200,
      height: 450,
      elements:[
        { view:"text", label:"Фамилия:",labelWidth:200,labelAlign:"right",   id:"lastnameCreateEmployee",   name:"lastname"},
        { view:"text", label:"Имя:",labelWidth:200,labelAlign:"right",       id:"firstnameCreateEmployee",  name:"firstname"},
        { view:"text", label:"Отчество:",labelWidth:200,labelAlign:"right",  id:"middlenameCreateEmployee", name:"middlename"},
        { 
          view:"richselect", 
          label:"Должность:",labelWidth:200,labelAlign:"right", 
          id:"positionCreateEmployee", name:"position",
          options:[]
        },
        { view:"text", label:"Контактный телефон:",labelWidth:200,labelAlign:"right",   id:"phonenumberCreateEmployee",name:"phone_number"},
        { view:"text", label:"Электронная почта:",labelWidth:200,labelAlign:"right",    id:"emailCreateEmployee",      name:"email", type:"email"},
        { view:"text", label:"Дата рождения:",labelWidth:200,labelAlign:"right",        id:"birthdateCreateEmployee",  name:"birthdate"},
        {
          view:"richselect",
          id:"genderCreateEmployee",
          name:"gender",
          label:"Выберите пол: ", 
          labelWidth: 200,
          labelAlign:"right",
          options:[ 
            { id:1, value:"Мужчина"}, 
            { id:2, value:"Женщина"}
          ] 
        },              
        { margin:5, cols:[
          { 
            view:"button", 
            id:"EmployeeWindowCreateCancelButton",
            value:"Отмена", 
            css:"webix_danger"
          },
          {
            view:"button",
            id:"EmployeeWindowCreateConfrimButton", 
            value:"Добавить", 
            css:"webix_primary"
          },          
        ]}
      ],
      rules:{ // имя компонента используется для применения к нему правила
        "lastname": webix.rules.isNotEmpty,
        "firstname":webix.rules.isNotEmpty,
        "middlename":webix.rules.isNotEmpty,
        "position": webix.rules.isNotEmpty,
        "phone_number": webix.rules.isNotEmpty,
        "email": webix.rules.isNotEmpty,
        "birthdate": webix.rules.isNotEmpty,
        "gender": webix.rules.isNotEmpty
      }
    }
  }
};