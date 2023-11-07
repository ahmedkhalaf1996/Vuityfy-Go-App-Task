
<template>
  <div class="item-list">
    <v-row>
      <v-col cols="12">
        <v-card class="item-card" v-for="(item, index) in items" :key="index">
          <v-card-text>
            <v-row align="center">
              <v-col cols="2">
                <v-img class="item-image" :src="item.img"></v-img>
              </v-col>
              <v-col cols="10">
                <v-row align="center">
                  <v-col cols="12">
                    <span class="item-title">{{ item.title }}</span>
                    <i> Located At: </i>
                    <span class="item-location">{{ item?.location }}</span>
                  </v-col>
                </v-row>
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>
    
    <script>
import { mapActions } from "vuex";
export default {
  name: "backadmin-component",
  data() {
    return {
      items: [],
    };
  },
  methods: {
    ...mapActions(["GetAllDataToPickList"]),
  },
  async created() {
   try {
    const { data } = await this.GetAllDataToPickList();
    //console.log("d", data);
    if(data){
      
    /// =---------
    const formattedItems = [];

// Process the data from the store, filtering by "status" and "is_archived"
for (const x of data) {
  if (x.status === true && x.is_archived === true) {
    for (const order of x.order_list) {
      const { title, img } = order;
      // Generate a random character for the location
      const location = String.fromCharCode(
        65 + Math.floor(Math.random() * 26)
      );

      // Check if the title already exists in the formattedItems array
      const existingItem = formattedItems.find(
        (item) => item.title === title
      );

      // If it doesn't exist, add it to the array
      if (!existingItem) {
        formattedItems.push({ title, img, location });
      }
    }
  }
}
console.log("foramteddata", formattedItems);
this.items = formattedItems;
    }
   } catch (error) {
    console.log(error)
   }

    //// -------
  },
};
</script>
    
    <style scoped>
.item-list {
  width: 80%;
  margin: 0 auto;
}

.item-card {
  margin-bottom: 10px;
  border: 1px solid #ccc;
}

.item-image {
  width: 100%;
  max-width: 80px;
  height: auto;
}

.item-location {
  margin-bottom: 5px;
}

.item-title {
  font-weight: bold;
}
</style>          