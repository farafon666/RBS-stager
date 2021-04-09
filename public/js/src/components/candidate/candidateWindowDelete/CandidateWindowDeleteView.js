//Окно удаления кандидата
export function CandidateWindowDeleteView(){
  return{
    view:"window",
    position:"center",
    head:"Удаление существующего кандидата",
    id: "windowDeleteCandidate",
    close: true,
    modal: true,
    body:{
      view:"form",
      id:"formDeleteCandidate",
      width: 480,
      height: 425,
      elements:[
        { view:"label", label:"Вы действительно хотите удалить данного кандидата?", align:"center"},
        { view:"label", label:"",   id:"lastnameDeleteCandidate",     name:'lastname'},
        { view:"label", label:"",   id:"firstnameDeleteCandidate",    name:'firstname'},
        { view:"label", label:"",   id:"middlenameDeleteCandidate",   name:'middlename'},
        { view:"label", label:"",   id:"phone_numberDeleteCandidate", name:'phone_number'},
        { view:"label", label:"",   id:"emailDeleteCandidate",        name:'email'},
        { view:"label", label:"",   id:"birthdateDeleteCandidate",    name:'birthdate'},
        { view:"label", label:"",   id:"genderDeleteCandidate",       name:'gender'},     
        { margin:5, cols:[
          { 
            view:"button", 
            id:"CandidateWindowDeleteCancelButton",
            value:"Отмена", 
            css:"webix_danger"
          },
          {
            view:"button",
            id:"CandidateWindowDeleteConfirmButton", 
            value:"Удалить", 
            css:"webix_primary"
          },          
        ]}
      ]      
    }
  }
};