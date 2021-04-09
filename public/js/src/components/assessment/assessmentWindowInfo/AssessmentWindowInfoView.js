//Окно просмотра существующего ассессмента
export function AssessmentWindowInfoView(){
  return{
    view:"window",
    position:"center",
    head:"Просмотр информации об ассессменте",
    id: "windowViewAssessment",
    close: true,
    modal: true,
    body:{
      view:"form",
      id:"formViewAssessment",
      width: 1200,
      autoheight: true,
      elements:[      
        { view:"label", label:"Дата: ",            id:"dateViewAssessment", name: "date"},
        { view:"label", label:"Состояние: ",       id:"statusViewAssessment", name: "state"},
        { view:"label", label:"Распорядитель: ",   id:"disposerViewAssessment", name: "disposer"},
        { view:"label", label:"Присутствующие сотрудники: "},
        {
          view:"datatable",
          id:"employeesViewAssessment",
          name:"employees",
          height:150,
          scroll:true,
          columns:[
            { id:"ID",          header:"ID", width:30, sort:"text"},
            { id:"lastname",    header:"Фамилия", sort:"text"},
            { id:"firstname",   header:"Имя", sort:"text"},
            { id:"middlename",  header:"Отчество", sort:"text"},
            { id:"position",    header:"Должность", sort:"text", fillspace: true},
            { id:"birthdate",   header:"Дата рождения", sort:"text"},
            { id:"phone_number",header:"Номер телефона",width:150},
            { id:"email",       header:"Электронная почта",width:150},
            { id:"gender",      header:"Пол", sort:"text"}
          ]
        },
        { view:"label", label:"Присутствующие кандидаты: "},
        {
          view:"datatable",
          id:"candidatesViewAssessment",
          name:"candidates",
          height:150,
          scroll: true,
          columns:[
            { id:"ID",          header:"ID", width: 30, sort:"text"},
            { id:"lastname",    header:"Фамилия", sort:"text"},
            { id:"firstname",   header:"Имя", sort:"text"},
            { id:"middlename",  header:"Отчество", sort:"text"},
            { id:"birthdate",   header:"Дата рождения", sort:"text"},
            { id:"phone_number",header:"Номер телефона",width:150},
            { id:"email",       header:"Электронная почта",width:150},
            { id:"gender",      header:"Пол", fillspace:true, sort:"text"}
          ]        
        },
        { margin:5, cols:[
          { 
            view:"button",
            id:"AssessmentWindowInfoCancelButton", 
            value:"Выйти", 
            css:"webix_danger"
          }        
        ]}
      ]      
    }
  }
};