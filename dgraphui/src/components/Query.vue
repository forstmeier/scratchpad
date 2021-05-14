<template>
  <v-container id="query">
    <v-list style="max-height: 700px" class="overflow-y-auto">
      <h3>SPECIMEN</h3>
      <QueryFilter
        v-for="s in filters.specimen"
        v-bind:key="s.type + ':' + s.field"
        v-bind:filter="s"
        v-on:operator="updateForm($event, 'operator')"
        v-on:value="updateForm($event, 'value')"
      />
      <h4>CONSENT</h4>
      <QueryFilter
        v-for="s in filters.consent"
        v-bind:key="s.type + ':' + s.field"
        v-bind:filter="s"
        v-on:operator="updateForm($event, 'operator')"
        v-on:value="updateForm($event, 'value')"
      />
      <h4>DONOR</h4>
      <QueryFilter
        v-for="s in filters.donor"
        v-bind:key="s.type + ':' + s.field"
        v-bind:filter="s"
        v-on:operator="updateForm($event, 'operator')"
        v-on:value="updateForm($event, 'value')"
      />
    </v-list>
  </v-container>
</template>

<script>
import { specimen, consent, donor } from "@/assets/query/filters.js";
import QueryFilter from "./QueryFilter.vue";

export default {
  name: "Query",
  components: {
    QueryFilter,
  },
  data: function () {
    return {
      filters: {
        specimen: specimen,
        consent: consent,
        donor: donor,
      },
      form: {
        specimen: {},
        bloodSpecimen: {},
        consent: {},
        donor: {},
      },
    };
  },
  methods: {
    updateForm: function (event, key) {
      const name = event.name.split(":");

      const oldFieldObject = this.form[name[0]][name[1]];
      const newFieldValue = checkValue(event.value) ? event.value : null;

      var newFieldObject = {};
      newFieldObject[key] = newFieldValue;
      if (oldFieldObject) {
        newFieldObject = Object.assign(oldFieldObject, newFieldObject);
      }
      this.form[name[0]][name[1]] = newFieldObject;

      this.updateVariables();
    },
    updateVariables: function () {
      var variables = {
        consent: createFilter(Object.assign({}, this.form.consent)),
        donor: createFilter(Object.assign({}, this.form.donor)),
        bloodSpecimen: createFilter(
          Object.assign({}, this.form.specimen, this.form.bloodSpecimen)
        ),
      };
      this.$emit("variables", variables);
    },
  },
};

function checkValue(value) {
  if (typeof value === "string") {
    return value ? true : false;
  } else {
    return value.length > 0 ? true : false;
  }
}

function createFilter(filter) {
  var newFilter = {};

  Object.keys(filter).forEach((key) => {
    const field = filter[key];

    if (!field.operator || !field.value) {
      return;
    }

    var fieldFilter = {};
    const newValue =
      typeof field.value === "string"
        ? field.value
        : "/" + field.value.join("|") + "/";

    fieldFilter[field.operator] = newValue;

    newFilter[key] = fieldFilter;
  });

  return newFilter;
}
</script>
