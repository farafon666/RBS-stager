//Окно удаления сотрудника
export function EmployeeWindowDeleteView(){
  return{
    view:"window",
    position:"center",
    head:"Удаление существующего сотрудника",
    id: "windowDeleteEmployee",
    close: true,
    modal: true,
    body:{
      view:"form",
      id:"formDeleteEmployee",
      width: 480,
      height: 470,
      elements:[
        { view:"label", label:"Вы действительно хотите удалить данного сотрудника?", align:"center"},
        { view:"label", label:"",   id:"lastnameDeleteEmployee",    name:'lastname'},
        { view:"label", label:"",   id:"firstnameDeleteEmployee",   name:'firstname'},
        { view:"label", label:"",   id:"middlenameDeleteEmployee",  name:'middlename'},
        { view:"label", label:"",   id:"positionDeleteEmployee",    name:'position'},
        { view:"label", label:"",   id:"phonenumberDeleteEmployee", name:'phone_number'},
        { view:"label", label:"",   id:"emailDeleteEmployee",       name:'email'},
        { view:"label", label:"",   id:"birthdateDeleteEmployee",   name:'birthdate'},
        { view:"label", label:"",   id:"genderDeleteEmployee",      name:'gender'},     
        { margin:5, cols:[
          { 
            view:"button",
            id:"EmployeeWindowDeleteCancelButton", 
            value:"Отмена", 
            css:"webix_danger"
          },
          {
            view:"button",
            id:"EmployeeWindowDeleteConfrimButton", 
            value:"Удалить", 
            css:"webix_primary"
          },          
        ]}
      ]      
    }
  }
};