<template>
  <v-container id="modal-item">
    <v-row>
      <v-col class="form-name">
        <div>
          {{ form.label }}
        </div>
      </v-col>
      <v-col class="form-value">
        <v-select
          v-if="form.input === 'select'"
          v-bind:items="form.options"
          item-text="name"
          item-value="value"
          v-on:change="
            $emit('value', {
              name: type,
              value: $event,
            })
          "
          dense
          outlined
          hide-details="auto"
        />
        <v-select
          v-if="form.input === 'multiple'"
          v-bind:items="form.options"
          item-text="name"
          item-value="value"
          v-on:change="
            $emit('value', {
              name: type,
              value: $event,
            })
          "
          multiple
          dense
          outlined
          hide-details="auto"
        />
        <v-text-field
          v-if="form.input === 'text'"
          v-on:change="
            $emit('value', {
              name: type,
              value: $event,
            })
          "
          dense
          outlined
          hide-details="auto"
        />
        <v-text-field
          v-if="form.input === 'number'"
          type="number"
          v-on:change="
            $emit('value', {
              name: type,
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
          v-if="form.input === 'date'"
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
                name: type,
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
  name: "ModalItem",
  props: {
    form: Object,
    type: String,
  },
  data: function () {
    return {
      date: "",
      menu: false,
    };
  },
};
</script>