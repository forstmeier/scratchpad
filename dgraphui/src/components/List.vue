<template>
  <v-container id="list">
    <v-list style="max-height: 700px" class="overflow-y-auto">
      <ListItem
        v-for="content in contents"
        v-bind:key="content.id"
        :content="content"
        :parentType="'specimen'"
      />
    </v-list>
    <v-container class="empty-list" v-if="contents && contents.length === 0">
      <v-banner outlined rounded>NO SPECIMENS FOR CURRENT FILTER</v-banner>
    </v-container>
  </v-container>
</template>

<script>
import ListItem from "@/components/ListItem.vue";

export default {
  name: "List",
  props: {
    contents: Array,
  },
  components: {
    ListItem,
  },
};

// outline:
// [ ] infinite scroll
// - [ ] https://www.digitalocean.com/community/tutorials/vuejs-implementing-infinite-scroll
</script>