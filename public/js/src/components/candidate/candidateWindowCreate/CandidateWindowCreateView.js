//Окно добавления кандидата
export function CandidateWindowCreateView(){
  return{
    view:"window",
    position:"center",
    head:"Добавление нового кандидата",
    id: "windowCreateCandidate",
    close: true,
    modal: true,
    body:{
      view:"form",
      id:"formCreateCandidate",
      width: 1200,
      height: 420,
      elements:[
        { view:"text", label:"Фамилия:",labelWidth:200,labelAlign:"right",   id:"lastnameCreateCandidate",   name:"lastname"},
        { view:"text", label:"Имя:",labelWidth:200,labelAlign:"right",       id:"firstnameCreateCandidate",  name:"firstname"},
        { view:"text", label:"Отчество:",labelWidth:200,labelAlign:"right",  id:"middlenameCreateCandidate", name:"middlename"},
        { view:"text", label:"Контактный телефон:",labelWidth:200,labelAlign:"right",   id:"phonenumberCreateCandidate",name:"phone_number"},
        { view:"text", label:"Электронная почта:",labelWidth:200,labelAlign:"right",    id:"emailCreateCandidate",      name:"email", type:"email"},
        { view:"text", label:"Дата рождения:",labelWidth:200,labelAlign:"right",        id:"birthdateCreateCandidate",  name:"birthdate"},
        {
          view:"richselect",
          id: "genderCreateCandidate",
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
            id:"CandidateWindowCreateCancelButton",
            value:"Отмена", 
            css:"webix_danger",
          },
          {
            view:"button", 
            id:"CandidateWindowCreateConfirmButton",
            value:"Добавить", 
            css:"webix_primary",
          },          
        ]}
      ],
      rules:{ // имя компонента используется для применения к нему правила
        "lastname": webix.rules.isNotEmpty,
        "firstname":webix.rules.isNotEmpty,
        "middlename":webix.rules.isNotEmpty,
        "phone_number": webix.rules.isNotEmpty,
        "email": webix.rules.isNotEmpty,
        "birthdate": webix.rules.isNotEmpty,
        "gender": webix.rules.isNotEmpty
      }
    }
  }
};       