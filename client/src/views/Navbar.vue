<template>
      <v-app-bar app color="primary">
      <v-toolbar-title>My App</v-toolbar-title>
      <v-btn  class="mr-2" to="/"><v-icon left>mdi-home</v-icon> Home </v-btn>
      <v-btn class="ma-2" to="/itemslist" >
        <v-icon start icon="mdi-dialpad"></v-icon>Items {{ itemsnum }}</v-btn>
    <v-btn class="ma-2" to="/review" >
        <v-icon start icon="mdi-dialpad"></v-icon>Review
    </v-btn>
    <v-btn class="ma-2" to="/picklist" >
        <v-icon start icon="mdi-dialpad"></v-icon>Picklist
    </v-btn>
    <v-btn class="ma-2" to="/backadmin" >
        <v-icon start icon="mdi-dialpad"></v-icon>BackAdmin
    </v-btn>
    </v-app-bar>
</template>

<script>
export default {
    name:'Navbar-component',
    data(){
        return {
            itemsnum: 5
        }
    }
}
</script>

<style>

</style>