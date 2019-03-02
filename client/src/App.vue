<template>
  <v-app>
    <v-toolbar app>
      <v-toolbar-title class="headline text-uppercase">
        <router-link to="/" tag="button">Gopham</router-link>
      </v-toolbar-title>
    </v-toolbar>
    <div style="margin-bottom: 5vw"></div>
    <router-view></router-view>
  </v-app>
</template>

<script>
export default {
  name: "App",
  data() {
    return {
      //
    };
  }
};
</script>
