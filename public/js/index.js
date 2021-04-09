import {Application} from "./src/components/Application.js";

webix.ready(() => {
    let app = new Application();
    webix.ui(app.config());
    app.init();
})