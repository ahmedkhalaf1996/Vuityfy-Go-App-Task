
<template>
    <div class="container">
        <div class="categories">
        <h2>Categories</h2>
        <v-list>
            <v-list-item v-for="(category, index) in categories" :key="index">
            <v-list-item-content>
                <v-list-item-title @click="selectCategory(category)">{{ category.title }}</v-list-item-title>
                    <v-list-item-subtitle v-if="category.subCategories.length > 0" 
                        @click="category.expanded =!category.expanded">
                    {{ category.expanded ? '▼' : '►' }}
                    </v-list-item-subtitle>
                    <v-list-item-subtitle v-if="category.expanded">
                        <v-list-item-subtitle  v-for="(subcategory, subIndex) in category.subCategories" :key="subIndex">
                        {{ subcategory.title  }}
                    </v-list-item-subtitle>
                    </v-list-item-subtitle>
 
            </v-list-item-content>
            </v-list-item>
        </v-list>
        </div>
    
        <div class="items" v-if="selectedCategory?.items">
        <h2>Items</h2>
        <v-row>
            <v-col v-for="(item, index) in selectedCategoryItems" :key="index" cols="4">
            <v-card>
                <v-card-text>
                <v-img :src="item.image" height="200"></v-img>
                <v-card-title>{{ item.title }}</v-card-title>
                <v-card-subtitle>{{ item.subtitle }}</v-card-subtitle>
                </v-card-text>
                <v-card-actions>
                <v-btn @click="addItemToCart(item)">Add to Cart</v-btn>
                </v-card-actions>
            </v-card>
            </v-col>
            
        </v-row>
        </div>
        <div class="items" v-else>
        <h2>Items</h2>
        <v-row>
            <v-col v-for="(item, index) in selectedCategory" :key="index" cols="4">
            <v-card>
                <v-card-text>
                <v-img :src="item.img" height="200"></v-img>
                <v-card-title>{{ item.title }}</v-card-title>
                </v-card-text>
                <v-card-actions>
                <v-btn @click="addItemToCart(item)">Add to Cart</v-btn>
                </v-card-actions>
            </v-card>
            </v-col>
            
        </v-row>
        </div>
    </div>
  </template>
    
<script>
import { mapActions } from 'vuex';

 export default {
    name:"Home-component",
    data() {
        return {
        categories: [
       
        ],
        selectedCategory: [] // selected items by category
        };
    },
    computed: {
        
        // selectedCategoryItems() {
        // return this.selectedCategory ? this.selectedCategory.items : [];
        // }
    },
    mounted(){
        this.getCategoriesItems()
        this.getAllItems()
    },
    methods: {
        ...mapActions(['GetAllItemsForAdd','GetAllItems', 'AddItemToList']),
        selectCategory(category) {
            this.selectedCategory = []

            for (let i = 0; i < category?.subCategories.length; i++) {
                const element = category?.subCategories[i];
                console.log("el", element)
                for (let x = 0; x < element.itemDescriptions.length; x++) {
                    const el = element.itemDescriptions[x];
                    this.selectedCategory.push(el)
                } 
            }
        },
        addItemToCart(item) {
            console.log(item)
            this.AddItemToList(item)
        // Handle adding the item to the cart
        },
        async getCategoriesItems(){
            try {
                
            let {data} = await this.GetAllItemsForAdd()
            this.categories = data
            console.log("d", data)

            } catch (error) {
             console.log(error)    
            }

        },
        async getAllItems(){
            try {
                let {data} = await this.GetAllItems()
            this.selectedCategory = data?.data
            console.log('allitems', data)
            } catch (error) {
                console.log(error)
            }
        }

    }
    };
</script>
    
<style>
    .container {
    display: flex;
    }

    .categories {
    flex: 0 0 200px;
    padding-right: 20px;
    }

    .items {
    flex: 1;
    }

    .v-list-item-title {
        cursor: pointer;
    }

    .v-list-item-subtitle {
        cursor: pointer;
    }

</style>