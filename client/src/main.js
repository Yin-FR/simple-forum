import Vue from 'vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import locale from 'element-ui/lib/locale/lang/en';
import App from './App.vue'
import axios from "axios";
import VueAxios from "vue-axios";
import Vuex from 'vuex';
import router from './router'

Vue.config.productionTip = false;
Vue.use(ElementUI, { locale });
Vue.use(VueAxios, axios);
Vue.use(Vuex);

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
