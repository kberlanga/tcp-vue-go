<template>
  <v-app>
    <v-app-bar class="app-bar" app elevation="0">
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-logo-dark.png"
          transition="scale-transition"
          width="20"
        />
      </div>
      <v-toolbar-title class="app-bar__title">Server Manager</v-toolbar-title>
      <v-spacer></v-spacer>

      <v-btn icon v-on:click="toggleTheme(!isDark)">
        <v-icon>{{ isDark ? iconThemeLight : iconThemeDark }}</v-icon>
      </v-btn>
    </v-app-bar>
    <v-main>
      <MenuOptions :connections="connections" :stadistics="stadistics" />
    </v-main>
  </v-app>
</template>

<script>
import MenuOptions from "./components/MenuOptions.vue";
import SocketioService from "./services/socketio.service";

export default {
  name: "App",

  components: {
    MenuOptions,
  },

  data: () => ({
    isDark: true,
    iconThemeDark: "mdi-weather-night",
    iconThemeLight: "mdi-brightness-5",
    connections: [],
    stadistics: [],
  }),
  created() {
    SocketioService.setupSocketConnection();
    SocketioService.socket.emit("app", "app connected");
    SocketioService.socket.on("connections", (connections) => {
      this.connections = connections;
    });
    SocketioService.socket.on("stadistics", (stadistics) => {
      this.stadistics = stadistics;
    });
  },
  methods: {
    toggleTheme: function (isDark) {
      this.isDark = isDark;
      this.$vuetify.theme.isDark = this.isDark;
    },
  },
};
</script>

<style>
body {
  font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
    Oxygen, Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
}

.app-bar {
  padding: 0px 10px;
  height: 60px;
}

.app-bar__title {
  display: none;
}

@media screen and (min-width: 768px) {
  .app-bar {
    padding: 0px 90px;
  }
}

@media screen and (min-width: 1024px) {
  .app-bar {
    padding: 0px 174px;
  }

  .app-bar__title {
    display: block;
  }
}

@media screen and (min-width: 1440px) {
  .app-bar {
    padding: 0px 274px;
  }
}
</style>
