import {CAssessmentTab} from './assessment/CAssessmentTab.js';
import {CCandidateTab}  from './candidate/CCandidateTab.js';
import {CEmployeeTab}   from './employee/CEmployeeTab.js';
import WorkedPlaceView  from './ApplicationView.js';

export class Application{
    constructor(){
        this.view;
        this.assessmentTab = new CAssessmentTab();
        this.candidateTab  = new CCandidateTab();
        this.employeeTab   = new CEmployeeTab();
    };
    init(){
        this.view = {
            tabbar: $$("tabbar"),            
        }; 
        
        this.view.tabbar.setValue("assessmentView");

        //Отрисовка всех конфигов
        this.assessmentTab.config();
        this.candidateTab.config();
        this.employeeTab.config();

        //Инициализация всех конфигов
        this.assessmentTab.init();
        this.candidateTab.init();
        this.employeeTab.init();
    };
    config() {           
        return WorkedPlaceView(this.assessmentTab, this.candidateTab, this.employeeTab);
    };
};
