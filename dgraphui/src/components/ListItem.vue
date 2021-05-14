<template>
  <v-container id="list-item">
    <v-card>
      <v-card-title>{{ parentType }}</v-card-title>
      <v-card-text>
        <v-divider />
        <h4>{{ parentType }}</h4>
        <div
          v-for="parentField in parentFields"
          v-bind:key="parentField[0] + parentField[1]"
        >
          {{ parentField[0] }}: {{ parentField[1] }}
        </div>
        <div
          v-for="childrenField in childrenFields"
          v-bind:key="childrenField[0] + childrenField[1]"
        >
          <v-divider />
          <h4>{{ childrenField[0] }}</h4>
          <div
            v-for="(subvalue, subkey) in childrenField[1]"
            v-bind:key="subkey + subvalue"
          >
            {{ subkey }}: {{ subvalue }}
          </div>
        </div>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
export default {
  name: "ListItem",
  props: {
    content: Object,
    parentType: String,
  },
  data: function () {
    return {
      parentFields: [],
      childrenFields: [],
    };
  },
  created: function () {
    const parentFields = Object.entries(this.$props.content).filter(
      (values) => {
        return values[1] === null || typeof values[1] !== "object";
      }
    );
    this.parentFields = parentFields;

    const childrenFields = Object.entries(this.$props.content).filter(
      (values) => {
        return values[1] !== null && typeof values[1] === "object";
      }
    );
    this.childrenFields = childrenFields;
  },
};
</script>