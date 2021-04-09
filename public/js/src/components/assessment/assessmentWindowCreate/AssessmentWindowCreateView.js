//Окно создания нового ассесмента
export function AssessmentWindowCreateView(){
  return { 
    view:"window",
    position:"center",
    head:"Создание нового ассессмента",
    id: "windowCreateAssessments",
    close: true,
    modal: true,
    body:{
      view:"form",
      id:"formCreateAssessments",
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
          id: "dateCreateAssessment",
          name:"date"
        },
        {
          view:"richselect",
          id:"disposerCreateAssessment",
          name: "disposer",
          label:"Выберите распорядителя: ", 
          labelPosition:"top"
        },
        {
          view:"dbllist",
          id:"dbllistCreateEmployees",
          name: "employees",
          list:{ height: 200, scroll: true},  
          labelLeft:"Сотрудники:",
          labelRight:"Выбранные сотрудники:",
        },
        {
          view:"dbllist",
          id:"dbllistCreateCandidates",
          name:"candidates",
          list:{ height: 200, scroll: true},    
          labelLeft:"Кандидаты:",
          labelRight:"Приглашённые кандидаты:",
        },
        { margin:5, cols:[
          { 
            view:"button",
            id:"AssessmentWindowCreateCancelButton", 
            value:"Отмена", 
            css:"webix_danger",
          },
          {
            view:"button",
            id:"AssessmentWindowCreateConfirmButton",  
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