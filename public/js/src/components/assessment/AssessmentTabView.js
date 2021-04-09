export function AssessmentTabView(){
  return{
    id:"assessmentView",
    rows:[
      {
        view:"toolbar",
        cols:[
          { 
            label: "Добавить", id:"buttonCreateAssessment",     
            view: "button", height: 36, width: 100, 
            align:"right"                  
          },
          { 
            label: "Изменить",id:"buttonUpdateAssessment",  
            view: "button", height: 36, width: 100, 
            align:"right",
            disabled: true    
          },
          { 
            label: "Удалить", id:"buttonDeleteAssessment",  
            view: "button", height: 36, width: 100, 
            align:"right",
            disabled: true
          },
          { 
            label: "Просмотр", id:"buttonViewAssessment",  
            view: "button", height: 36, width: 100, 
            align:"right",
            disabled: true            
          }
        ]
      },
      {
        view:"datatable",
        id:"dataAssessment",
        columns:[
          { id:"ID",     header:["ID", {content:"selectFilter"}], width: 30, sort:"text"},
          { id:"title",  header:["Название", {content:"selectFilter"}]},
          { id:"date",   header:["Дата и время", {content:"selectFilter"}], width: 300, sort:"text"},
          { id:"state",  header:["Состояние", {content:"selectFilter"}], width: 200, fillspace:false, sort:"text"}                            
        ],
        on: {
          onSelectChange: function(){
            $$("buttonUpdateAssessment").enable();  
            $$("buttonDeleteAssessment").enable();                        
            $$("buttonViewAssessment").enable();                  
          }
        },
        select: true      
      }
    ]                      
  }
};