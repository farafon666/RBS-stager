//Окно проведения ассесмента
export function AssessmentWindowInProcessView(){
    return { 
      view:"window",
      position:"center",
      head:"Проведение ассессмента",
      id: "windowInProcessAssessments",
      close: true,
      modal: true,
      body:{
        view:"form",
        id:"formInProcessAssessments",
        width: 1200,
        height: 730,
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
            view:"dbllist",
            id:"dbllistInProcessCandidates",
            name:"candidates",
            list:{ height: 200, scroll: true},    
            labelLeft:"Кандидаты:",
            labelRight:"Приглашённые кандидаты:",
          },
          { margin:5, cols:[
            { 
              view:"button",
              id:"AssessmentWindowInProcessCancelButton", 
              value:"Отмена", 
              css:"webix_danger",
            },
            {
              view:"button",
              id:"AssessmentWindowInProcessConfirmButton",  
              value:"Добавить", 
              css:"webix_primary",
            },          
          ]}
        ],
        rules:{ // имя компонента используется для применения к нему правила
          "date": webix.rules.isNotEmpty,
          "disposer":webix.rules.isNotEmpty,
          "employees": webix.rules.isNotEmpty,
          "candidates": webix.rules.isNotEmpty
        }
      }
    }
  };