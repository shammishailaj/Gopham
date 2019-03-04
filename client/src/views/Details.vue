<template>
  <div>
    <!-- Left half of screen with navigation and additional Info -->
    <div style="margin-top:4vw; margin-left: 1vw" class="split left">
      <v-navigation-drawer permanent>
        <v-toolbar flat>
          <v-list>
            <v-list-tile-title class="title">
              {{ pid }}
              <br />
            </v-list-tile-title>
          </v-list>
        </v-toolbar>
        <p>SLoC: {{ totalLoC }}</p>
        <v-divider></v-divider>

        <v-list dense class="pt-0">
          <v-list-tile
            v-for="(value, key) in metrics"
            :key="key"
            @click="showmetric(key)"
          >
            <v-list-tile-content>
              <v-list-tile-title>{{ value.title }}</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-navigation-drawer>
    </div>
    <!-- Right Half of Screen with Content -->
    <div style="margin-top:4vw" class="split right">
      <v-card style="margin-bottom: 5vw">
        <v-card-title>
          {{ tabletitle }}
          <v-spacer></v-spacer>
          <v-text-field
            v-model="search"
            append-icon="search"
            label="Search"
            single-line
            hide-details
          ></v-text-field>
        </v-card-title>
        <v-data-table
          :headers="headers"
          :items="tablevalues"
          :search="search"
          :rows-per-page-items="[50, 100, 200]"
        >
          <template slot="items" slot-scope="props">
            <td>{{ props.item.keyname }}</td>
            <td class="text-xs-right">
              {{ props.item.count }}
            </td>
          </template>
          <v-alert slot="no-results" :value="true" color="error" icon="warning">
            Your search for "{{ search }}" found no results.
          </v-alert>
        </v-data-table>
      </v-card>
    </div>
  </div>
</template>

<script>
const axios = require("axios");

export default {
  props: ["pid"],
  data() {
    return {
      totalLoC: 0,
      search: "",
      tabletitle: "Efferent Couplings (File)",
      headers: [
        { text: "File", sortable: false, value: "keyname" },
        { text: "#", value: "count" }
      ],
      tablevalues: [],
      metrics: {
        efCouplingsFile: {
          title: "Efferent Couplings (File)",
          left: "File",
          right: "#",
          content: {}
        },
        packageEFCount: {
          title: "Efferent Couplings (Package)",
          left: "Package",
          right: "#",
          content: {}
        },
        afCouplingsAllPackage: {
          title: "Afferent Couplings (All Packages)",
          left: "Package",
          right: "#",
          content: {}
        },
        afCouplingsProjectPackage: {
          title: "Afferent Couplings (Project only Packages)",
          left: "Package",
          right: "#",
          content: {}
        },
        fileLoC: {
          title: "Source Lines of Code (File)",
          left: "File",
          right: "LoC",
          content: {}
        },
        packageLoC: {
          title: "Source Lines of Code (Package)",
          left: "Package",
          right: "LoC",
          content: {}
        },
        fileFunctionCount: {
          title: "Number of Functions (File)",
          left: "File",
          right: "#",
          content: {}
        },
        packageFunctionCount: {
          title: "Number of Functions (Package)",
          left: "Package",
          right: "#",
          content: {}
        }
      }
    };
  },
  methods: {
    showmetric(metricname) {
      this.tabletitle = this.metrics[metricname].title;
      this.headers[0].text = this.metrics[metricname].left;
      this.headers[1].text = this.metrics[metricname].right;
      var tmptableheaders = [];
      for (var key in this.metrics[metricname].content) {
        tmptableheaders.push({
          keyname: key,
          count: this.metrics[metricname].content[key]
        });
      }
      this.tablevalues = tmptableheaders;
    },
    getMetric(metric) {
      var self = this;
      axios
        .get("http://localhost:8000/getanalysis/" + self.pid + "/" + metric, {})
        .then(function(response) {
          self.metrics[metric].content = response.data;
          if (metric === "efCouplingsFile") {
            self.showmetric(metric);
          }
        })
        .catch(function(error) {
          console.log(error);
        });
    }
  },
  created() {
    // fetches all the metrics and stores them
    var self = this;
    axios
      .get("http://localhost:8000/getanalysis/" + self.pid + "/totalLoC", {})
      .then(function(response) {
        self.totalLoC = response.data;
      })
      .catch(function(error) {
        console.log(error);
      });
    for (var key in this.metrics) {
      this.getMetric(key);
    }
  }
};
</script>

<style scoped>
.split {
  height: 100%;
  position: fixed;
  z-index: 1;
  top: 0;
  overflow-x: hidden;
  padding-top: 20px;
}
.left {
  left: 0;
  width: 20%;
}
.right {
  right: 0;
  width: 80%;
}
</style>
