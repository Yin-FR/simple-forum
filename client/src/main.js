import Vue from 'vue'
import App from './App.vue'
import axios from "axios";
import VueAxios from "vue-axios";
import Vuex from 'vuex';

Vue.config.productionTip = false
Vue.use(VueAxios, axios);
Vue.use(Vuex);

new Vue({
  render: h => h(App),
}).$mount('#app')