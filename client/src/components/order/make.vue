<template>
  <v-app>
    <v-container>
      <v-row>
        <v-col cols="12">
          <v-dialog v-model="dialog" max-width="500px">
            <v-card>
              <v-card-title> Enter your email </v-card-title>
              <v-card-text>
                <v-text-field
                  v-model="userEmail"
                  label="Email"
                  outlined
                ></v-text-field>
              </v-card-text>
              <v-card-actions>
                <v-btn @click="submitEmail" color="primary">Submit</v-btn>
                <v-btn @click="closeDialog" color="secondary">Close</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-col>
      </v-row>
    </v-container>
  </v-app>
</template>
  
  <script>
import { mapActions } from 'vuex';
export default {
  name: "make-order",
  props: {
    ordercalled: Function,
  },
  data() {
    return {
      dialog: true,
      userEmail: "",
    };
  },
  methods: {
    //   openDialog() {
    //     this.dialog = true;
    //   },
    ...mapActions(["MakeOrder"]),

    closeDialog() {
      this.dialog = false;
    },
    async submitEmail() {
      console.log("User's email: " + this.userEmail);
      // TODO::
      // clear localsotreage
      const items = JSON.parse(localStorage.getItem("items"));

      const updatedItems = items.map((item) => {
        const updatedItem = {};

        for (const key in item) {
          if (Object.prototype.hasOwnProperty.call(item, key)) {
            updatedItem[key.toLowerCase()] = item[key];
          }
        }

        return updatedItem;

      });


    const Final = updatedItems.map((item) => {
       return { id: item.id, quantity: item.quantity };
    });


    var sendOrder = { "email":this.userEmail, "order_list": Final }
     var {data} = await this.MakeOrder(sendOrder)
     
    console.log("f", data)
      this.ordercalled();
      this.closeDialog();
    },
  },
};
</script>
  

