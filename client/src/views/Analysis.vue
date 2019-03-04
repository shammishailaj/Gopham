<template>
  <div>
    <div id="loader"></div>
    <v-form
      id="formID"
      class="myform"
      ref="form"
      v-model="valid"
      lazy-validation
    >
      <v-text-field
        v-model="name"
        :rules="nameRules"
        label="Name"
        required
      ></v-text-field>
      <v-text-field
        v-model="directory"
        :rules="directoryRules"
        label="Directory Name"
        required
      ></v-text-field>
      <v-btn :disabled="!valid" @click="submit">submit</v-btn>
    </v-form>
    <h3 style="text-align: center;" id="errorField"></h3>
  </div>
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
      document.getElementById("formID").style.display = "none";
      document.getElementById("loader").style.display = "block";
      document.getElementById("errorField").innerText =
        "Please wait while the project is being analyzed";
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
          document.getElementById("formID").style.display = "block";
          document.getElementById("loader").style.display = "none";
          document.getElementById("errorField").innerText =
            "No such directory in /projects/";
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

#loader {
  margin: 0 auto;
  display: none;
  border: 16px solid #f3f3f3; /* Light grey */
  border-top: 16px solid #3498db; /* Blue */
  border-radius: 50%;
  width: 120px;
  height: 120px;
  animation: spin 2s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
