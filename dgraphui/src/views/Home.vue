<template>
  <v-container id="home" class="full-height">
    <v-row>
      <v-col>
        <h2>HOME</h2>
        <div v-if="!$auth.loading">
          <div v-if="$auth.isAuthenticated && !$auth.loading">
            <p>Use the links in the sidebar to work with the app.</p>
          </div>
        </div>
      </v-col>
      <v-col class="d-flex justify-end">
        <v-btn
          v-if="!$auth.isAuthenticated && !$auth.loading"
          v-on:click.prevent="login"
        >
          LOGIN
        </v-btn>
        <v-btn v-if="$auth.isAuthenticated" v-on:click="logout">LOGOUT</v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "Home",
  methods: {
    login: function () {
      this.$auth.loginWithRedirect();
    },
    logout: function () {
      this.$auth.logout();
      this.$router.push({ path: "/" });
    },
  },
};
</script>
