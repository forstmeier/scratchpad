<template>
  <v-container id="modal">
    <v-dialog v-model="dialog" persistent max-width="600px">
      <template v-slot:activator="{ on, attrs }">
        <v-btn color="primary" dark v-bind="attrs" v-on="on">ADD</v-btn>
      </template>
      <v-card>
        <v-card-title> {{ modalAction }} {{ schemaType }} </v-card-title>
        <v-container>
          <v-row dense>
            <v-col>
              <v-select
                v-if="schemaTypeOptions"
                v-model="schemaTypeSelection"
                v-bind:items="schemaTypeOptions"
                dense
                outlined
                hide-details="auto"
              />
            </v-col>
            <v-col>
              <v-radio-group v-model="inputToggle" row dense>
                <v-radio label="FILE" value="file" />
                <v-radio label="FORM" value="form" />
              </v-radio-group>
            </v-col>
          </v-row>
          <v-row dense v-if="inputToggle === 'form'">
            <ModalItem
              v-for="(formItem, key) in formItems"
              v-bind:key="key"
              v-bind:form="formItem"
              v-bind:type="key"
              v-on:value="updateForm($event)"
            />
          </v-row>
          <v-row dense v-if="inputToggle === 'file'">
            <v-file-input
              label="File input"
              outlined
              dense
              clearable
              v-on:change="chooseFile"
            ></v-file-input>
          </v-row>
        </v-container>

        <v-card-actions>
          <v-spacer />
          <v-btn
            v-if="inputToggle !== ''"
            v-on:click="
              dialog = false;
              handleSubmit(inputToggle);
            "
          >
            Submit
          </v-btn>
          <v-btn v-on:click="dialog = false">Cancel</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import Papa from "papaparse";
import { add as bloodSpecimenAdd } from "@/assets/graphql/bloodSpecimen.js";
import API from "@/api/index.js";
import { bloodSpecimen } from "@/assets/forms";
import ModalItem from "./ModalItem.vue";

function clearFile() {
  let oldInput = document.getElementsByClassName("v-file-input__text")[0];
  let newInput = document.createElement("input");
  newInput.type = oldInput.type;
  newInput.className = oldInput.className;
  oldInput.parentNode.replaceChild(newInput, oldInput);
}

function isID(header) {
  const ids = ["consent", "donor", "test", "result"];

  return ids.includes(header);
}

function isEmpty(value) {
  return value === "" || value === [];
}

export default {
  name: "Modal",
  props: {
    modalAction: String,
    schemaType: String,
    schemaTypeOptions: Array,
  },
  components: {
    ModalItem,
  },
  data: function () {
    return {
      dialog: false,
      inputToggle: "",
      fileData: null,
      formItems: {},
      schemaTypeSelection: this.schemaTypeOptions[0],
    };
  },
  methods: {
    updateForm: function (event) {
      this.formItems[event.name].value = event.value;
    },
    handleSubmit: function (type) {
      if (type === "form") {
        this.submitForm();
      } else if (type === "file") {
        this.submitFile();
      }
    },
    submitForm: function () {
      const orgID = this.$auth.user["https://folivora.io/jwt/claims"].orgID;
      let variables = {
        org: {
          id: orgID,
        },
      };
      Object.keys(this.formItems).forEach((key) => {
        const item = this.formItems[key];

        if (isEmpty(item.value)) {
          return null;
        }

        if (item.query) {
          if (["tests", "results"].includes(key)) {
            const idObjects = item["value"].map((id) => {
              return { ["id"]: id };
            });
            variables[key] = idObjects;
          } else {
            variables[key] = {
              id: item.value,
            };
          }
        } else {
          variables[key] = item.value;
        }
      });

      this.api.sendRequest(
        bloodSpecimenAdd,
        { input: [variables] },
        (response) => {
          if (response.data.errors) {
            alert("error processing form"); // TEMP
            console.log("error:", response.data.errors); // TEMP
          }

          const data = response.data.data;
          console.log("response data:", data); // TEMP
          alert("success processing form"); // TEMP
        }
      );

      this.inputToggle = "";
    },
    chooseFile: async function (file) {
      let fileName = file.name;
      let fileSplit = fileName.split(".");
      const ext = fileSplit[fileSplit.length - 1];

      const orgID = this.$auth.user["https://folivora.io/jwt/claims"].orgID;
      if (ext === "json") {
        const reader = new FileReader();
        reader.onload = (e) => {
          const data = JSON.parse(e.target.result);
          this.fileData = data.map((d) => {
            let newD = Object.assign({}, d);
            for (const key in newD) {
              if (isID(key)) {
                newD[key] = newD[key] ? { id: newD[key] } : null;
              }
            }
            newD.org = { id: orgID };

            return newD;
          });
        };
        reader.readAsText(file);
      } else if (ext === "csv") {
        const parseCSV = (file) => {
          return new Promise((resolve) => {
            Papa.parse(file, {
              header: true,
              dynamicTyping: true,
              skipEmptyLines: true,
              transform: function (value, header) {
                if (!!value && header === "tests") {
                  let tests = value.split(",");
                  return tests.map((testID) => ({ id: testID }));
                }

                if (!!value && isID(header)) {
                  return { id: value };
                }

                return value;
              },
              complete: function (results) {
                resolve(results.data);
              },
              error: function (error) {
                console.log("error:", error); // TEMP
              },
            });
          });
        };

        const data = await parseCSV(file);
        this.fileData = data.map((d) => {
          let newD = Object.assign({}, d);
          newD.org = { id: orgID };
          return newD;
        });
      }
    },
    submitFile: async function () {
      this.api.sendRequest(
        bloodSpecimenAdd,
        { input: this.fileData },
        (response) => {
          if (response.data.errors) {
            alert("error processing json file"); // TEMP
            console.log("error:", response.data.errors); // TEMP
          }

          const data = response.data.data;
          console.log("success:", data); // TEMP
        }
      );

      clearFile();
    },
  },
  created: async function () {
    const claims = await this.$auth.getIdTokenClaims();
    const idToken = claims.__raw;

    this.api = new API(
      process.env.VUE_APP_BACKEND_URL,
      idToken,
      process.env.VUE_APP_FOLIVORA_CUSTOM_SERVER_SECRET
    );

    let queries = [];

    let formItems;
    if (this.schemaTypeSelection === "bloodSpecimen") {
      formItems = bloodSpecimen;
    }

    // build and run queries to populate select input options
    Object.keys(formItems).forEach((key) => {
      const item = formItems[key];
      if (item.query) {
        const aliasQuery = key + ": " + item.query;
        queries.push(aliasQuery);
      }
    });

    const query =
      `query BloodSpecimenPreQuery {\n\t` + queries.join("\n\t") + `\n}`;
    this.api.sendRequest(query, null, (response) => {
      let data = response.data.data;
      Object.keys(data).forEach((key) => {
        const options = data[key].map((d) => d.id);
        formItems[key].options = options;
      });
      this.formItems = formItems;
    });
  },
};
</script>
