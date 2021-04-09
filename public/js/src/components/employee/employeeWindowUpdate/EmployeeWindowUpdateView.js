//Окно редаткирования сотрудника
export function EmployeeWindowUpdateView(){
  return{
    view:"window",
    position:"center",
    head:"Редактирование данных о сотруднике",
    id: "windowUpdateEmployee",
    close: true,
    modal: true,
    body:{
      view:"form",
      id:"formUpdateEmployee",
      width: 1200,
      height: 450,
      elements:[
        { view:"text", label:"Фамилия:", labelWidth:200,labelAlign:"right",  id:"lastnameUpdateEmployee",   name:"lastname"},
        { view:"text", label:"Имя:",     labelWidth:200,labelAlign:"right",  id:"firstnameUpdateEmployee",  name:"firstname"},
        { view:"text", label:"Отчество:",labelWidth:200,labelAlign:"right",  id:"middlenameUpdateEmployee", name:"middlename"},
        { 
          view:"richselect", 
          label:"Должность:",labelWidth:200,labelAlign:"right",
          id:"positionUpdateEmployee", name:"position",
          options:[]
        },
        { view:"text", label:"Контактный телефон:",labelWidth:200,labelAlign:"right",  id:"phonenumberUpdateEmployee",name:"phone_number"},
        { view:"text", label:"Электронная почта:",labelWidth:200,labelAlign:"right",   id:"emailUpdateEmployee",      name:"email", type:"email"},
        { view:"text", label:"Дата рождения:",labelWidth:200,labelAlign:"right",       id:"birthdateUpdateEmployee",  name:"birthdate"},
        {
          view:"richselect",
          id: "genderUpdateEmployee",
          name: "gender",
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
            id:"EmployeeWindowUpdateCancelButton", 
            value:"Отмена", 
            css:"webix_danger"
          },
          {
            view:"button",
            id:"EmployeeWindowUpdateConfrimButton",  
            value:"Изменить", 
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