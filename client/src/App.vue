<template>
  <v-app>
    <v-main>
      <Navbar/>
      <router-view />

    </v-main>
  </v-app>
</template>

<script>
import Navbar from './views/Navbar.vue'

export default {
  name: 'App',

  components: {
    Navbar,
  },

  data: () => ({
    //
  }),
}
</script>
