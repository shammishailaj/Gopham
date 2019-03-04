<template>
  <div style="margin: 0 auto">
    <v-btn class="mybuttons" to="/analysis" large color="green"
      >New Analysis</v-btn
    >
    <v-btn class="mybuttons" large color="primary" @click="retrieveProjectList"
      >Refresh</v-btn
    >
    <div style="margin-bottom: 5%"></div>

    <v-card
      style="width: 100%;"
      v-for="(value, key) in projects"
      v-bind:key="key"
    >
      <v-card-title class="description" primary-title>
        <div style="display: inline-block;">
          <h2 style="width: 400px;" class="headline mb-0">{{ key }}</h2>
          <h4>{{ value }}</h4>
        </div>
      </v-card-title>

      <v-card-actions class="cardbuttons">
        <v-btn flat color="blue" @click="showDetails(key)">Details</v-btn>
        <v-btn flat color="red" @click="deleteProject(key)">Delete</v-btn>
      </v-card-actions>
    </v-card>
  </div>
</template>

<script>
const axios = require("axios");

export default {
  name: "Home",
  data() {
    return {
      projects: {}
    };
  },
  methods: {
    showDetails(pname) {
      let self = this;
      self.$router.push({ name: "details", params: { pid: pname } });
    },
    retrieveProjectList() {
      var $this = this;
      axios
        .get("http://localhost:8000/projectlist")
        .then(function(response) {
          $this.projects = response.data;
        })
        .catch(function(error) {
          console.log(error);
        })
        .then(function() {});
    },
    deleteProject(pname) {
      var $this = this;
      axios
        .post("http://localhost:8000/deleteanalysis", {
          name: pname
        })
        .then(function(response) {
          $this.retrieveProjectList();
          console.log(response);
        })
        .catch(function(error) {
          console.log(error);
        });
    }
  },
  beforeMount() {
    this.retrieveProjectList();
  }
};
</script>

<style scoped>
.description,
.cardbuttons {
  display: inline-block;
  vertical-align: middle;
}
</style>
