
        <template>
  <v-row>
    <v-col cols="2"> </v-col>

    <v-col cols="8">
      <!-- Right side: List of Items -->
      <MakeOrder v-if="callOrder" :ordercalled="ordercalled" />
      <v-card class="items-card">
        <v-card-title>Items</v-card-title>

        <v-card-actions justify="center">
          <v-col cols="4"><Additem :addItemFun="addItemFun" /></v-col>
        </v-card-actions>
        <v-list>
          <v-list-item-group>
            <v-list-item v-for="(item, index) in items" :key="index">
              <v-list-item-content>
                <v-row align="center">
                  <v-col cols="4">
                    <img :src="item.img" alt="Item Image" class="item-image" />
                  </v-col>
                  <v-col cols="4">
                    <v-list-item-title>{{ item.title }}</v-list-item-title>
                  </v-col>
                  <v-col cols="4" class="item-quantity">
                    <v-btn icon small @click="decreaseQuantity(item)">
                      <v-icon>mdi-minus</v-icon>
                    </v-btn>
                    <v-btn class="ma-2" color="purple"
                      >{{ item.quantity }} | {{ item.type }}
                    </v-btn>
                    <v-btn icon small @click="increaseQuantity(item)">
                      <v-icon>mdi-plus</v-icon>
                    </v-btn>
                  </v-col>
                  <v-col cols="12">
                    <v-btn icon small @click="removeItem(item)">
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
            <v-btn color="primary" block @click="submitOrder"
              >Submit Order</v-btn
            >
          </v-col>
        </v-row>
      </v-card>
    </v-col>
  </v-row>
</template>
<script>
import Additem from "../add/additem.vue";
import MakeOrder from "../order/make.vue";
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
      callOrder: false,
    };
  },
  async created() {
    this.items = await this.GetItemsFormLocal();
    console.log("I tem List...", await this.GetItemsFormLocal());
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
    ...mapActions(["GetItemsFormLocal", "UpdateItemList"]),
    async addItemFun() {
      this.items = await this.GetItemsFormLocal();
    },
    async ordercalled() {
      // this.items = await this.GetItemsFormLocal();

      this.callOrder = false;
      localStorage.clear();
      this.items = []

    },
    increaseQuantity(item) {
      item.quantity++;
      this.UpdateItemList(item);
    },
    decreaseQuantity(item) {
      if (item.quantity > 1) {
        item.quantity--;
        this.UpdateItemList(item);
      }
    },
    removeItem(item) {
      const index = this.items.indexOf(item);
      if (index !== -1) {
        item.quantity = 0;
        this.UpdateItemList(item);
        this.items.splice(index, 1);
      }
    },
    submitOrder() {
      this.callOrder = true;
    },
  },
  components: {
    Additem,
    MakeOrder,
  },
};
</script>
          <style scoped>
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