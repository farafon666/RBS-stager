//Окно редактирования кандидата
export function CandidateWindowUpdateView(){
  return{
    view:"window",
    position:"center",
    head:"Редактирование данных о кандидате",
    id: "windowUpdateCandidate",
    close: true,
    modal: true,
    body:{
      view:"form",
      id:"formUpdateCandidate",
      width: 1200,
      height: 420,
      elements:[
        { view:"text",  label:"Фамилия:",labelWidth:200,labelAlign:"right",   id:"lastnameUpdateCandidate",   name:"lastname"},
        { view:"text", label:"Имя:",labelWidth:200,labelAlign:"right",       id:"firstnameUpdateCandidate",  name:"firstname"},
        { view:"text", label:"Отчество:",labelWidth:200,labelAlign:"right",  id:"middlenameUpdateCandidate", name:"middlename"},
        { view:"text", label:"Контактный телефон:",labelWidth:200,labelAlign:"right",   id:"phonenumberUpdateCandidate",name:"phone_number"},
        { view:"text", label:"Электронная почта:",labelWidth:200,labelAlign:"right",    id:"emailUpdateCandidate",      name:"email", type:"email"},
        { view:"text", label:"Дата рождения:",labelWidth:200,labelAlign:"right",        id:"birthdateUpdateCandidate",  name:"birthdate"},
        {
          view:"richselect",
          id: "genderUpdateCandidate",
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
            id:"CandidateWindowUpdateCancelButton", 
            value:"Отмена", 
            css:"webix_danger"
          },
          {
            view:"button",
            id:"CandidateWindowUpdateConfirmButton", 
            value:"Изменить", 
            css:"webix_primary"
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