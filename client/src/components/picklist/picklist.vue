


<template>
  <div>
    <!-- Active Orders -->
    <v-card>
      <v-card-title>Active Orders</v-card-title>
      <v-card-text>
        <v-row>
          <v-col cols="6">
            <!-- Left List: Items Names -->
            <v-list>
              <v-subheader>Items Names</v-subheader>
              <v-list-item
                v-for="(item, index) in activeItemNamesWithNumbers"
                :key="index"
              >
                <v-list-item-content>{{ item }}</v-list-item-content>
              </v-list-item>
            </v-list>
          </v-col>
          <v-col cols="6">
            <!-- Right List: Items Type and Quantity -->
            <v-list>
              <v-subheader>Items Type and Quantity</v-subheader>
              <v-list-item
                v-for="(type, index) in activeItemTypes"
                :key="index"
              >
                <v-list-item-content>
                  {{ type }} - {{ activeItemQuantities[index] }}
                </v-list-item-content>
              </v-list-item>
            </v-list>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>

    <!-- Archived Orders -->
    <v-card>
      <v-card-title>Archived Orders</v-card-title>
      <v-card-text>
        <v-row>
          <v-col cols="6">
            <!-- Left List: Items Names -->
            <v-list>
              <v-subheader>Items Names</v-subheader>
              <v-list-item
                v-for="(item, index) in archivedItemNamesWithNumbers"
                :key="index"
              >
                <v-list-item-content>{{ item }}</v-list-item-content>
              </v-list-item>
            </v-list>
          </v-col>
          <v-col cols="6">
            <!-- Right List: Items Type and Quantity -->
            <v-list>
              <v-subheader>Items Type and Quantity</v-subheader>
              <v-list-item
                v-for="(type, index) in archivedItemTypes"
                :key="index"
              >
                <v-list-item-content>
                  {{ type }} - {{ archivedItemQuantities[index] }}
                </v-list-item-content>
              </v-list-item>
            </v-list>
          </v-col>
        </v-row>
      </v-card-text>
    </v-card>
    <v-btn @click="ArcihiveAll">Done</v-btn>

  </div>
</template>

<script>
import { mapActions } from 'vuex';

export default {
  name: 'picklist-component',
  data() {
    return {
      order: [],
    };
  },

  async created() {
    try {
      const  {data} = await this.GetAllDataToPickList()
      if (data){
        this.order = data 
      } 
    }catch (error){
      console.log(error)
    }

  },

  computed: {
    activeItemNamesWithNumbers() {
      const items = this.filterItems(false)
        .map((order) => order.order_list)
        .flat();

      const itemNames = items.reduce((acc, item) => {
        if (!acc[item.title]) {
          acc[item.title] = 0;
        }
        acc[item.title]++;
        return acc;
      }, {});

      return Object.keys(itemNames).map(
        (name) => `${itemNames[name]} ${name}`
      );
    },
    activeItemTypes() {
      return [...new Set(this.filterItems(false).map((order) =>
        order.order_list.map((item) => item.type)
      ).flat())];
    },
    activeItemQuantities() {
      const items = this.filterItems(false)
        .map((order) => order.order_list)
        .flat();

      const itemQuantities = items.reduce((acc, item) => {
        if (!acc[item.type]) {
          acc[item.type] = 0;
        }
        acc[item.type] += item.item_quantity;
        return acc;
      }, {});

      return this.activeItemTypes.map((type) => itemQuantities[type]);
    },
    archivedItemNamesWithNumbers() {
      const items = this.filterItems(true)
        .map((order) => order.order_list)
        .flat();

      const itemNames = items.reduce((acc, item) => {
        if (!acc[item.title]) {
          acc[item.title] = 0;
        }
        acc[item.title]++;
        return acc;
      }, {});

      return Object.keys(itemNames).map(
        (name) => `${itemNames[name]} ${name}`
      );
    },
    archivedItemTypes() {
      return [...new Set(this.filterItems(true).map((order) =>
        order.order_list.map((item) => item.type)
      ).flat())];
    },
    archivedItemQuantities() {
      const items = this.filterItems(true)
        .map((order) => order.order_list)
        .flat();

      const itemQuantities = items.reduce((acc, item) => {
        if (!acc[item.type]) {
          acc[item.type] = 0;
        }
        acc[item.type] += item.item_quantity;
        return acc;
      }, {});

      return this.archivedItemTypes.map((type) => itemQuantities[type]);
    },
  },

  methods: {
    ...mapActions([ "GetAllDataToPickList", "ArcihiveAllSt"]),

    filterItems(isArchived) {
      return this.order.filter((order) => order.is_archived === isArchived);
    },

    async ArcihiveAll(){
      this.ArcihiveAllSt()
      const  {data} = await this.GetAllDataToPickList()
      this.order = data
      console.log("BackAdmin", data)
    }
  },
};
</script>
