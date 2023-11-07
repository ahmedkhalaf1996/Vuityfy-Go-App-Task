
<template>
  <v-row>
    <v-col cols="2"> </v-col>

    <v-col cols="9">
      <!-- Right side: List of Items -->
      <v-card class="items-card mm3" v-for="(order, orderindex) in items" :key="orderindex"> <!--- loop here  -->
        <v-card-title>Order User Email :{{  order.email }}</v-card-title>

        <v-list>
          <v-list-item-group>
            <v-list-item v-for="(item, index) in order.order_list" :key="index" >
              <v-list-item-content>
                <v-row align="center">
                  <v-col cols="4">
                    <img :src="item.img" alt="Item Image" class="item-image" />
                  </v-col>
                  <v-col cols="4">
                    <v-list-item-title>{{ item.title }}</v-list-item-title>
                  </v-col>
                  <v-col cols="4" class="item-quantity">
                    <v-btn icon small @click="decreaseQuantity(item.item_id, order.id)">
                      <v-icon>mdi-minus</v-icon>
                    </v-btn>
                    <v-btn class="ma-2" color="purple"
                      >{{ item.item_quantity }} | {{ item.type }}
                    </v-btn>
                    <v-btn icon small @click="increaseQuantity(item.item_id, order.id)">
                      <v-icon>mdi-plus</v-icon>
                    </v-btn>
                  </v-col>
                  <v-col cols="12">
                    <v-btn icon small @click="removeItem(item.item_id, order.id)">
                      <v-icon>mdi-close</v-icon>
                    </v-btn>
                  </v-col>
                </v-row>
              </v-list-item-content>
            </v-list-item>
          </v-list-item-group>
        </v-list>

        <v-row justify="center">
          <v-col cols="4">
            <v-btn color="primary" block @click="addToPickList(order.id)"
              >Add To PickList</v-btn
            >
          </v-col>
        </v-row>
      </v-card>
    </v-col>
  </v-row>
</template>
<script>
// import Additem from "../add/additem.vue";
import { mapActions } from "vuex";

export default {
  name: "ItemsComponent",
  data() {
    return {
      selectedCategories: [],
      selectedSubcategories: [],
      categories: [],
      subcategories: [],
      items: [],
    };
  },
  async created() {
    // this.items = await this.GetOrderdList();
    const {data} = await this.GetOrderdList();
    if(data){
      this.items = data
    console.log("Review I tem List...", this.items);
    }
  },
  computed: {
    // filteredItems()
    filteredItems() {
      if (
        this.selectedCategories.length === 0 ||
        this.selectedSubcategories.length === 0
      ) {
        return [];
      }
      return this.items.filter(
        (item) =>
          this.selectedCategories.includes(item.categories) &&
          this.selectedSubcategories.includes(item.subcategories)
      );
    },
  },
  methods: {
    ...mapActions(["GetOrderdList", "UpdateItemList", "UpdateItemQuntity", "DeleteItem", "AddToPickList"]),
  

    increaseQuantity(itemId, orderid) {
      var  Itemlist = this.items.find(item => item.id === orderid);
      var  gitem = Itemlist.order_list.find(item => item.item_id === itemId)
      gitem.item_quantity++
      console.log("final", this.items)
      console.log("inc item", gitem)
      this.UpdateItemQuntity(gitem)
    },
    decreaseQuantity(itemId, orderid) {
      var  Itemlist = this.items.find(item => item.id == orderid);
      var  gitem = Itemlist.order_list.find(item => item.item_id == itemId)

      if(gitem.item_quantity >1){
        gitem.item_quantity--
        console.log("final", this.items)
        this.UpdateItemQuntity(gitem)
       }
    },
    removeItem(itemId, orderid) {

      var  Itemlist = this.items.find(item => item.id === orderid);
      const  gitem = Itemlist.order_list.find(item => item.item_id == itemId)
      this.DeleteItem(gitem)
      console.log("final", gitem)
      Itemlist.order_list = Itemlist.order_list.filter(i => i.item_id !== itemId)

    },
    addToPickList(id) {
      this.items = this.items.filter(i => i.id !== id);
      console.log('add to pick list order id', id)
      this.AddToPickList(id)
    },
  },
  components: {
    // Additem,
  },
};
</script>
<style scoped>
.mm3 {
  margin: 41px 0px;
}
.sidebar-card {
  margin-right: 16px;
}
.items-card {
  width: 100%;
}
.item-image {
  width: 80px;
  height: 80px;
  object-fit: cover;
}
.item-quantity {
  text-align: center;
}
.confirmation-dialog {
  background-color: #fff;
  border: 1px solid #ccc;
  padding: 20px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}
</style>