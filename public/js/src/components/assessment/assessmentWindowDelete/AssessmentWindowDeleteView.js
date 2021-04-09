//Окно удаления существующего ассессмента
export function AssessmentWindowDeleteView(){
    return{
        view:"window",
        position:"center",
        head:"Удаление ассессмента",
        id: "windowDeleteAssessment",
        close: true,
        modal: true,
        body:{
        view:"form",
        id:"formDeleteAssessment",
        width: 1200,
        autoheight: true,
        elements:[
            { view:"label", lable:"",                  id:"IDDeleteAssessment",         name:'ID'},      
            { view:"label", label:"Дата: ",            id:"dateDeleteAssessment",       name:'date'},
            { view:"label", label:"Состояние: ",       id:"statusDeleteAssessment",     name:'state'},
            { view:"label", label:"Распорядитель: ",   id:"disposerDeleteAssessment",   name:'disposer'},
            { view:"label", label:"Присутствующие сотрудники: "},
            {
            view:"datatable",
            id:"employeesDeleteAssessment",
            name:"employees",
            height:150,
            scroll:true,
            columns:[
                { id:"ID",          header:"ID", width:30, sort:"text"},
                { id:"lastname",    header:"Фамилия", sort:"text"},
                { id:"firstname",   header:"Имя", sort:"text"},
                { id:"middlename",  header:"Отчество", sort:"text"},
                { id:"position",    header:"Должность", sort:"text", fillspace:true},
                { id:"birthdate",   header:"Дата рождения", sort:"text"},
                { id:"phone_number",header:"Номер телефона",width:150},
                { id:"email",       header:"Электронная почта",width:150},
                { id:"gender",      header:"Пол", sort:"text"}
            ]
            },
            { view:"label", label:"Присутствующие кандидаты: "},
            {
            view:"datatable",
            id:"candidatesDeleteAssessment",
            name:"candidates",
            height:150,
            scroll: true,
            columns:[
                { id:"ID",          header:"ID", width:30, sort:"text"},
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
                id:"AssessmentWindowDeleteCancelButton", 
                value:"Выйти", 
                css:"webix_danger"
            },
            {
                view:"button",
                id:"AssessmentWindowDeleteConfirmButton",
                value:"Подтвердить",
                css:"webix_primary"
            }        
            ]}
        ]      
        }
    }
};