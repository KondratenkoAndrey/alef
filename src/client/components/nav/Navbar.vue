<template>
  <v-app-bar app>
    <v-app-bar-nav-icon
      class="hide-sm-and-up"
      @click="toggleDrawer(drawerState)"
    />
    <v-toolbar-title v-text="title" />
    <v-toolbar-title v-text="drawerState" />
    <v-spacer />
  </v-app-bar>
</template>

<script>
import { mapMutations } from 'vuex'

export default {
  data() {
    return {
      items: [
        {
          icon: 'mdi-apps',
          title: 'Welcome',
          to: '/',
        },
        {
          icon: 'mdi-chart-bubble',
          title: 'Inspire',
          to: '/inspire',
        },
      ],
      title: 'Alef',
    }
  },
  computed: {
    drawerState() {
      return this.$store.state.nav.drawerState
    },
  },
  methods: {
    ...mapMutations({ toggleDrawer: 'nav/toggleDrawer' }),
  },
}
</script>
