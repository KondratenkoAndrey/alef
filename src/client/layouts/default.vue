<template>
  <v-app>
    <Sidebar />
    <Navbar />

    <v-main>
      <v-container>
        <nuxt />
      </v-container>
    </v-main>

    <Footer />
  </v-app>
</template>

<script>
import Footer from '~/components/nav/Footer'
import Sidebar from '~/components/nav/Sidebar'
import Navbar from '~/components/nav/Navbar'

export default {
  components: {
    Footer,
    Sidebar,
    Navbar,
  },
}
</script>
