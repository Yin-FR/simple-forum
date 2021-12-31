import Vue from 'vue';
import Router from 'vue-router';

import PrimaryPage from "../views/PrimaryPage";
import PostPage from "../views/PostPage";

Vue.use(Router);

export default new Router({
    mode: 'hash',
    routes:[
        {path: '/', name:'home', redirect: '/primary'},
        {path: '/primary', name:'primary', component: PrimaryPage},
        {path: '/post/:postId', name: 'post', component: PostPage},
    ]
});