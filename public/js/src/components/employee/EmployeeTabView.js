export function EmployeeTabView() {
  return {
    id:"employeeView",                 
    rows: [
      {
        view:"toolbar",
        cols:[
          { 
            label: "Добавить", id:"buttonCreateEmployee",     
            view: "button", height: 36, width: 100, 
            align:"right"
          },
          { 
            label: "Изменить",id:"buttonUpdateEmployee",  
            view: "button", height: 36, width: 100, 
            align:"right",
            disabled: true
          },
          { label: "Удалить", id:"buttonDeleteEmployee",  
            view: "button", height: 36, width: 100, 
            align:"right",
            disabled: true
          }
        ]
      },
      {
        view:"datatable",
        id:"dataEmployee",
        columns:[
          { id:"ID",          header:["ID", {content:"selectFilter"}], width: 30, sort:"text"},
          { id:"lastname",    header:["Фамилия", {content:"selectFilter"}], sort:"text"},
          { id:"firstname",   header:["Имя", {content:"selectFilter"}], sort:"text"},
          { id:"middlename",  header:["Отчество", {content:"selectFilter"}], width:150, sort:"text"},
          { id:"position",    header:["Должность",  {content:"selectFilter"}], width: 250, sort:"text"},
          { id:"birthdate",   header:["Дата рождения", {content:"selectFilter"}], sort:"text"},
          { id:"phone_number",header:["Номер телефона", {content:"selectFilter"}], width: 170},
          { id:"email",       header:["Электронная почта", {content:"selectFilter"}], width: 250},
          { id:"gender",      header:["Пол", {content:"selectFilter"}], fillspace:true, sort:"text"}
        ],
        on: {
          onSelectChange: function(){
            $$("buttonUpdateEmployee").enable();
            $$("buttonDeleteEmployee").enable();                       
          }
        },
        select: true
      }
    ]                      
  }
}