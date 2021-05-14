<template>
  <v-app class="full-height">
    <v-card>
      <v-navigation-drawer
        app
        permanent
        v-if="$auth.isAuthenticated && !$auth.loading"
      >
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title class="title">
              <router-link to="/">Folivora Labs</router-link>
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-divider />
        <v-list dense nav>
          <v-list-item link>
            <v-list-item-content>
              <router-link to="/database">DATABASE</router-link>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>
    </v-card>
    <v-main app class="full-height">
      <v-container fluid class="full-height">
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Ubuntu+Mono:ital,wght@0,400;0,700;1,400;1,700&display=swap");

body,
html {
  height: 100%;
  font-family: "Ubuntu Mono", "sans-serif";
}

#app {
  font-family: "Ubuntu Mono", "sans-serif";
}

.full-height {
  height: 100%;
}
</style>
