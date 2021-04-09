import {Assessment} from '../models/entities/Assessment.js';
import {Candidate}  from '../models/entities/Candidate.js';
import {Employee}   from '../models/entities/Employee.js';
import {Position}   from '../models/entities/Position.js';

//Главное пространство приложения
export default function WorkedPlaceView(assessmentTab,candidateTab,employeeTab){
  return{
    rows:[
      {
        view:"toolbar",
        height:40,        
        css:{"background":"#6a92ff"},
        cols:[
          {view: "label",  label:"assessmentManager", align: "center"}
        ]        
      },
      {
        type:"clean",
        rows:[
          {
            borderless:true, view:"tabbar", id:"tabbar", value:"listView", multiview:true, options:[
              { value:'Ассессменты', id:'assessmentView'},
              { value:'Кандидаты',   id:'candidateView'},
              { value:'Сотрудники',  id:'employeeView'}
            ]
          },              
          {
            cells:[
              assessmentTab.config(),
              candidateTab.config(),
              employeeTab.config()
            ]
          }
        ]
      }
    ]
  };    
};