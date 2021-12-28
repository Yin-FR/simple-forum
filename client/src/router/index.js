import Vue from 'vue';
import Router from 'vue-router';

import PrimaryPage from "../views/PrimaryPage";
import PostPage from "../views/PostPage";

Vue.use(Router);

export default new Router({
    mode: 'history',
    routes:[
        {path: '/', name:'home', redirect: '/primary'},
        {path: '/primary', name:'primary', component: PrimaryPage},
        {path: '/welcomeM', name: 'welcomeM', component: PostPage},
    ]
});