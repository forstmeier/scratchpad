import Vue from 'vue';
import createAuth0Client from '@auth0/auth0-spa-js';


const DEFAULT_REDIRECT_CALLBACK = () =>
  window.history.replaceState({}, document.title, window.location.pathname);

let instance;

export const getInstance = () => instance;

export const useAuth0 = ({
  onRedirectCallback = DEFAULT_REDIRECT_CALLBACK,
  redirectUri = window.location.origin,
  ...options
}) => {
  if (instance) return instance;

  instance = new Vue({
    data() {
      return {
        loading: true,
        isAuthenticated: false,
        user: {},
        auth0Client: null,
        popupOpen: false,
        error: null
      };
    },
    methods: {
      loginWithPopup: async function(o) {
        this.popupOpen = true;

        try {
          await this.auth0Client.loginWithPopup(o);
          this.user = await this.auth0Client.getUser();
          this.isAuthenticated = await this.auth0Client.isAuthenticated();
          this.error = null;
        } catch (e) {
          console.error('login with popup error:', e);
          this.error = e;
        } finally {
          this.popupOpen = false;
        }
      },
      handleRedirectCallback: async function() {
        this.loading = true;
        try {
          await this.auth0Client.handleRedirectCallback();
          this.user = await this.auth0Client.getUser();
          this.isAuthenticated = true;
          this.error = null;
        } catch (e) {
          this.error = e;
        } finally {
          this.loading = false;
        }
      },
      loginWithRedirect: function(o) {
        return this.auth0Client.loginWithRedirect(o);
      },
      getIdTokenClaims: function(o) {
        return this.auth0Client.getIdTokenClaims(o);
      },
      getTokenSilently: function(o) {
        return this.auth0Client.getTokenSilently(o);
      },
      getTokenWithPopup: function(o) {
        return this.auth0Client.getTokenWithPopup(o);
      },
      logout: function(o) {
        return this.auth0Client.logout(o);
      }
    },
    created: async function() {
      this.auth0Client = await createAuth0Client({
        domain: options.domain,
        client_id: options.clientID,
        audience: options.audience,
        redirect_uri: redirectUri
      });

      try {
        if (
          window.location.search.includes('code=') &&
          window.location.search.includes('state=')
        ) {
          const { appState } = await this.auth0Client.handleRedirectCallback();
          this.error = null;
          onRedirectCallback(appState);
        }
      } catch (e) {
        this.error = e;
      } finally {
        this.isAuthenticated = await this.auth0Client.isAuthenticated();
        this.user = await this.auth0Client.getUser();
        this.loading = false;
      }
    }
  });

  return instance;
};

export const Auth0Plugin = {
  install(Vue, options) {
    Vue.prototype.$auth = useAuth0(options);
  }
};
