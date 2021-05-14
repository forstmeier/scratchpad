<template>
  <v-container id="database" class="full-height">
    <v-row>
      <v-col>
        <h2>DATABASE</h2>
      </v-col>
      <v-col>
        <Modal
          v-bind:modalAction="modalAction"
          v-bind:schemaType="schemaType"
          v-bind:schemaTypeOptions="schemaTypeOptions"
        />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <Query v-on:variables="updateVariables($event)" />
      </v-col>
      <v-col>
        <List v-bind:contents="contents" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import API from "@/api/index.js";
import Modal from "@/components/Modal.vue";
import Query from "@/components/Query.vue";
import List from "@/components/List.vue";
import { query as bloodSpecimenQuery } from "@/assets/graphql/bloodSpecimen.js";

export default {
  name: "Database",
  components: {
    Modal,
    Query,
    List,
  },
  data: function () {
    return {
      query: "",
      variables: {},
      contents: [],
      modalAction: "add",
      schemaType: "specimen",
      schemaTypeOptions: ["bloodSpecimen"],
    };
  },
  created: async function () {
    const claims = await this.$auth.getIdTokenClaims();
    const idToken = claims.__raw;

    this.api = new API(
      process.env.VUE_APP_BACKEND_URL,
      idToken,
      process.env.VUE_APP_FOLIVORA_CUSTOM_SERVER_SECRET
    );

    this.query = bloodSpecimenQuery;
    this.variables = {
      specimen: {},
      consent: {},
      donor: {},
    };

    this.runQuery();
  },
  methods: {
    updateVariables(variables) {
      this.variables = variables;
      this.runQuery();
    },
    runQuery() {
      this.api.sendRequest(this.query, this.variables, (response) => {
        this.contents = response.data.data.queryBloodSpecimen;
      });
    },
  },
};
</script>