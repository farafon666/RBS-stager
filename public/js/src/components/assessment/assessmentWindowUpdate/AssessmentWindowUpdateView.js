//Окно редаткирования существующего ассессмента
export function AssessmentWindowUpdateView(){
  return{
    view:"window",
    position:"center",
    head:"Редактирование существующего ассессмента",
    id: "windowUpdateAssessments",
    close: true,
    modal: true,
    body:{
      view:"form",
      id:"formUpdateAssessments",
      width: 1200,
      height: 730,
      elements:[
        {
          view: "datepicker",
          label: "Дата: ",
          value: "2021-02-12T12:26:04.044Z",
          timepicker: true,
          format: "%H:%i %d.%m.%Y",
          labelWidth: 100,
          inputWidth: 300,
          id: "dateAssessmentUpdate",
          name:"date"
        },      
        {
          view:"richselect",
          id:"stateUpdateAssessment",
          name: "state",
          label:"Состояние: ", 
          labelPosition:"top",
          value: 1,
          options:[]
        },
        {
          view:"richselect",
          id:"disposerUpdateAssessment",
          name: "disposer",
          label:"Распорядитель: ", 
          labelPosition:"top"
        },
        {
          view:"dbllist",
          id:"dbllistUpdateEmployees",
          name:"employees",
          list:{ height: 200, scroll: true},  
          labelLeft:"Сотрудники:",
          labelRight:"Выбранные сотрудники:",
        },
        {
          view:"dbllist",
          id:"dbllistUpdateCandidates",
          name:"candidates",
          list:{ height: 200, scroll: true},    
          labelLeft:"Кандидаты:",
          labelRight:"Приглашённые кандидаты:",
        },
        { margin:5, cols:[
          { 
            view:"button",
            id:"AssessmentWindowUpdateCancelButton", 
            value:"Отмена", 
            css:"webix_danger",
          },
          {
            view:"button", 
            id:"AssessmentWindowUpdateConfirmButton", 
            value:"Изменить", 
            css:"webix_primary",
          },          
        ]}
      ],
      rules:{ // имя компонента используется для применения к нему правила
        "date":    webix.rules.isNotEmpty,
        "state":  webix.rules.isNotEmpty,
        "disposer":webix.rules.isNotEmpty,
        "employees":  webix.rules.isNotEmpty,
        "candidates": webix.rules.isNotEmpty
      }
    }
  }
};