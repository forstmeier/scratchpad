import Vue from 'vue';
import router from './router';
import App from './App.vue';
import { Auth0Plugin } from './auth';
import vuetify from './plugins/vuetify';
import 'vuetify/dist/vuetify.css';


Vue.config.productionTip = false;

Vue.use(Auth0Plugin, {
  domain: process.env.VUE_APP_AUTH0_DOMAIN,
  clientID: process.env.VUE_APP_AUTH0_CLIENT_ID,
  onRedirectCallback: appState => {
    router.push(
      appState && appState.targetUrl
        ? appState.targetUrl
        : window.location.pathname
    );
  }
});

new Vue({
  render: h => h(App),
  vuetify,
  router: router
}).$mount('#app')
