<template>
  <v-form class="myform" ref="form" v-model="valid" lazy-validation>
    <v-text-field v-model="name" :rules="nameRules" label="Name" required></v-text-field>
    <v-text-field v-model="directory" :rules="directoryRules" label="Directory Name" required></v-text-field>
    <v-btn :disabled="!valid" @click="submit">submit</v-btn>
  </v-form>
</template>

<script>
const axios = require("axios");
import router from "../router";

export default {
  name: "projectPath",

  data() {
    return {
      valid: false,
      name: "",
      nameRules: [v => !!v || "Name is required"],
      directory: "",
      directoryRules: [v => !!v || "Directory is required"]
    };
  },
  methods: {
    submit() {
      axios
        .post("http://localhost:8000/analysis", {
          name: this.name,
          root: this.directory
        })
        .then(function(response) {
          router.push("/");
          console.log(response);
        })
        .catch(function(error) {
          console.log(error);
        });
    }
  }
};
</script>

<style>
.myform {
  margin: 0 auto;
  width: 30%;
}
</style>
