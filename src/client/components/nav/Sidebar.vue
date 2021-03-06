<template>
  <v-navigation-drawer v-model="drawerState" app>
    <v-list>
      <v-list-item
        v-for="(item, i) in items"
        :key="i"
        :to="item.to"
        router
        exact
      >
        <v-list-item-action>
          <v-icon>{{ item.icon }}</v-icon>
        </v-list-item-action>
        <v-list-item-content>
          <v-list-item-title v-text="item.title" />
        </v-list-item-content>
      </v-list-item>
    </v-list>
    {{ drawerState }}
  </v-navigation-drawer>
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
