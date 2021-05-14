<template>
  <v-container id="query-filter">
    <v-row>
      <v-col class="filter-label">
        {{ filter.label }}
      </v-col>
      <v-col class="filter-operator">
        <v-select
          v-bind:items="filter.operators"
          item-text="name"
          item-value="value"
          v-on:change="
            $emit('operator', {
              name: filter.type + ':' + filter.field,
              value: $event,
            })
          "
          dense
          outlined
          hide-details="auto"
        />
      </v-col>
      <v-col class="filter-value">
        <v-select
          multiple
          v-bind:items="filter.options"
          v-if="filter.input === 'select'"
          item-text="name"
          item-value="value"
          v-on:change="
            $emit('value', {
              name: filter.type + ':' + filter.field,
              value: $event,
            })
          "
          dense
          outlined
          hide-details="auto"
        />
        <v-text-field
          v-if="filter.input === 'text'"
          v-on:change="
            $emit('value', {
              name: filter.type + ':' + filter.field,
              value: $event,
            })
          "
          dense
          outlined
          hide-details="auto"
        />
        <v-text-field
          v-if="filter.input === 'number'"
          type="number"
          v-on:change="
            $emit('value', {
              name: filter.type + ':' + filter.field,
              value: $event,
            })
          "
          dense
          outlined
          hide-details="auto"
        />
        <v-menu
          v-model="menu"
          v-bind:close-on-content-click="false"
          transition="scale-transition"
          v-if="filter.input === 'date'"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-text-field
              readonly
              v-model="date"
              v-bind="attrs"
              v-on="on"
              dense
              outlined
              hide-details="auto"
            />
          </template>
          <v-date-picker
            v-model="date"
            v-on:input="menu = false"
            v-on:change="
              $emit('value', {
                name: filter.type + ':' + filter.field,
                value: $event,
              })
            "
          />
        </v-menu>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: "QueryFilter",
  props: {
    filter: Object,
  },
  data: function () {
    return {
      date: "",
      menu: false,
    };
  },
};
</script>