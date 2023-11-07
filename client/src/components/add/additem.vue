<!-- <template>
  <div>
    <v-btn color="primary" dark @click="dialog = true">Add New Order</v-btn>
    <v-dialog v-model="dialog" max-width="800px">
      <v-card>
        <v-card-title>
          <span class="headline">Form</span>
        </v-card-title>
        <v-card-text>
          <v-form @submit.prevent="submitForm">
            <v-autocomplete
              v-model="category"
              :items="categories"
              label="Category"
              required
              clearable
              @update:modelValue="showAddNewCategoryButton"
              @change="showAddNewCategoryButton"
            ></v-autocomplete>
            <v-text-field
              v-if="showAddNewCategoryButtonVal"
              label="Your New Category"
              v-model="newcategory"
              :append-icon="newcategory == '' ? 'mdi-close' : 'mdi-plus'"
              @click:append="AddOrCloseCategorie"
            ></v-text-field>

            <v-autocomplete
              v-model="subCategory"
              :items="category?.subCategories"
              label="Sub-Category"
              required
              clearable
              @update:modelValue="showAddNewSubCategoryButton"
              @change="showAddNewSubCategoryButton"
            ></v-autocomplete>
            <v-text-field
              v-if="showAddNewSubCategoryButtonVal"
              label="Your New Sub Category"
              v-model="newsubcategory"
              :append-icon="newsubcategory == '' ? 'mdi-close' : 'mdi-plus'"
              @click:append="AddOrCloseSubCategorie"
            ></v-text-field>

            <v-autocomplete
              v-model="itemDescription"
              :items="subCategory?.itemDescriptions"
              label="item-description"
              required
              clearable
              @update:modelValue="showAddNewDescriptionButton"
              @change="showAddNewDescriptionButton"
            ></v-autocomplete>
            <v-text-field
              v-if="showAddNewDescriptionButtonVal"
              label="Your New item description"
              v-model="newitemdescription"
              :append-icon="newitemdescription == '' ? 'mdi-close' : 'mdi-plus'"
              @click:append="AddOrCloseDescription"
            ></v-text-field>

            <v-text-field
              v-model.number="itemQuantity"
              label="Item Quantity"
              type="number"
              required
            ></v-text-field>

            <v-radio-group
              v-model="itemQuantityType"
              label="Item Quantity Type"
              required
            >
              <v-radio label="LB" value="LB"></v-radio>
              <v-radio label="Piece" value="Piece"></v-radio>
              <v-radio label="Package" value="Package"></v-radio>
              <v-radio label="Case" value="Case"></v-radio>
            </v-radio-group>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" text @click="dialog = false">Cancel</v-btn>
          <v-btn color="primary" @click="submitForm">Submit</v-btn>
        </v-card-actions>
      </v-card>

      <div class="text-center ma-2">
        <v-snackbar v-model="snackbar">
          Form Not Valid Required data!
          <template v-slot:actions>
            <v-btn color="pink" variant="text" @click="snackbar = false">
              Close
            </v-btn>
          </template>
        </v-snackbar>
      </div>
    </v-dialog>
  </div>
</template>

<script>
import { mapActions } from "vuex";
export default {
  name: "AddItem",
  props: {
    addItemFun: Function,
  },
  data() {
    return {
      dialog: false,
      snackbar: false,
      //
      ItemID: null,
      CatID: null,
      SubCatID: null,

      //
      showAddNewCategoryButtonVal: false,
      showAddNewSubCategoryButtonVal: false,
      showAddNewDescriptionButtonVal: false,

      category: "",
      subCategory: "",
      itemDescription: "",

      newcategory: "",
      newsubcategory: "",
      newitemdescription: "",

      itemQuantity: 1,
      itemQuantityType: "",

      categories: [],
    };
  },
  created() {
    this.GetAllData();
    // You can perform your logic here             this.AddItemToList(item)
  },
  computed: {
    // showAddNewCategoryButton() {
    //   return !this.categories.includes(this.category) && this.category !== '';
    // },
  },
  methods: {
    ...mapActions([
      "CreateNewCategory",
      "CreateNewSubCategory",
      "CreateNewItem",
      "GetAllItemsForAdd",
      "AddItemToList",
    ]),
    async GetAllData() {
      try {
        const { data } = await this.GetAllItemsForAdd();
        console.log("get all data ...", data);
        this.categories = data;
      } catch (error) {
        console.log(error);
      }
    },
    async AddOrCloseCategorie() {

      if (this.newcategory !== "") {
        this.categories = []
        this.showAddNewCategoryButtonVal = false;
        this.categories.push(this.newcategory);
        await this.CreateNewCategory(this.newcategory);
        const { data } = await this.GetAllItemsForAdd();
        this.categories = data;
      }
      this.showAddNewCategoryButtonVal = true;
    },
    async AddOrCloseSubCategorie() {

      if (this.newsubcategory !== "") {
        this.subCategory = []
        this.category?.subCategories.push(this.newsubcategory);
        const t = { title: this.newsubcategory, category_id: this.CatID };
        await this.CreateNewSubCategory(t);        
        const { data } = await this.GetAllItemsForAdd();
        this.categories = data;
        this.showAddNewSubCategoryButtonVal = false;
        // this.dialog = false
      } else {
        this.showAddNewSubCategoryButtonVal = true;

      }
    },

    async AddOrCloseDescription() {
      if (this.newitemdescription !== "") {
        this.categories = []

        this.showAddNewDescriptionButtonVal = false;

        console.log("iqt", this.itemQuantityType);
        if (this.itemQuantityType == "") {
          alert("Please select an item type.");
          return;
        }
        this.subCategory?.itemDescriptions.push(this.newitemdescription);
        const t = {
          title: this.newitemdescription,
          category_id: this.CatID,
          sub_category_id: this.SubCatID,
          type: this.itemQuantityType,
        };

         await this.CreateNewItem(t);
        const { data } = await this.GetAllItemsForAdd();
        this.categories = data;

        // this.dialog = false

      }
      this.showAddNewDescriptionButtonVal = true;

    },

    showAddNewCategoryButton() {
      const category = this.categories?.find(
        (category) => category.title === this.category
      );
      this.CatID = category?.id;
      this.category = category;
      // this.subCategories = category?.subCategories
      // this.itemDescriptions = category?.subCategories?.itemDescription

      if (!category) {
        this.showAddNewCategoryButtonVal = true;
      } else {
        this.showAddNewCategoryButtonVal = false;
      }
    },
    showAddNewSubCategoryButton() {
      const suncategory = this.category?.subCategories?.find(
        (suncategory) => suncategory.title === this.subCategory
      );
      this.SubCatID = suncategory?.id;
      this.subCategory = suncategory;

      if (!suncategory) {
        this.showAddNewSubCategoryButtonVal = true;
      } else {
        this.showAddNewSubCategoryButtonVal = false;
      }
    },

    showAddNewDescriptionButton() {
      const itemdesc = this.subCategory?.itemDescriptions?.find(
        (itemdesc) => itemdesc.title === this.itemDescription
      );
      this.ItemID = itemdesc?.id;
      this.itemDescription = itemdesc;

      if (!itemdesc) {
        this.showAddNewDescriptionButtonVal = true;
      } else {
        console.log("calllled", itemdesc);
        this.itemQuantityType = itemdesc.itemQuantityType;
        this.showAddNewDescriptionButtonVal = false;
      }
    },

    submitForm() {
      // Perform form submission logic here
      // You can access the form field values using this.category, this.subCategory, etc.
      if (
        this.itemQuantity < 0 ||
        this.itemDescription == "" ||
        this.itemQuantityType == ""
      ) {
        this.snackbar = true;
      } else {
        console.log(
          "item descripton",
          this.itemDescription,
          "item quenity",
          this.itemQuantity,
          "item type",
          this.itemQuantityType,
          "item Category",
          this.category
        );

        this.AddItemToList(this.itemDescription);
        console.log("Form submitted!");
        this.dialog = false; // Close the dialog after form submission
        if (typeof this.addItemFun === "function") {
          this.addItemFun();
        }
      }
    },
  },
};
</script> -->


<template>
  <div>
    <v-btn color="primary" dark @click="dialog = true">Add New Order</v-btn>
    <v-dialog v-model="dialog" max-width="800px">
      <v-card>
        <v-card-title>
          <span class="headline">Form</span>
        </v-card-title>
        <v-card-text>
          <v-form @submit.prevent="submitForm">
            <v-autocomplete
              v-model="selectedCategory"
              :items="categories"
              label="Category"
              required
              clearable
              @update:modelValue="showAddNewCategoryButton"
              @change="showAddNewCategoryButton"
            ></v-autocomplete>
            <v-text-field
              v-if="showAddNewCategoryButtonVal"
              label="Your New Category"
              v-model="newcategory"
              :append-icon="newcategory == '' ? 'mdi-close' : 'mdi-plus'"
              @click:append="AddOrCloseCategorie"
            ></v-text-field>

            <v-autocomplete
              v-model="selectedSubCategory"
              :items="selectedCategory?.subCategories"
              label="Sub-Category"
              required
              clearable
              @update:modelValue="showAddNewSubCategoryButton"
              @change="showAddNewSubCategoryButton"
            ></v-autocomplete>
            <v-text-field
              v-if="showAddNewSubCategoryButtonVal"
              label="Your New Sub Category"
              v-model="newsubcategory"
              :append-icon="newsubcategory == '' ? 'mdi-close' : 'mdi-plus'"
              @click:append="AddOrCloseSubCategorie"
            ></v-text-field>

            <v-autocomplete
              v-model="selectedItemDescription"
              :items="selectedSubCategory?.itemDescriptions"
              label="item-description"
              required
              clearable
              @update:modelValue="showAddNewDescriptionButton"
              @change="showAddNewDescriptionButton"
            ></v-autocomplete>
            <v-text-field
              v-if="showAddNewDescriptionButtonVal"
              label="Your New item description"
              v-model="newitemdescription"
              :append-icon="newitemdescription == '' ? 'mdi-close' : 'mdi-plus'"
              @click:append="AddOrCloseDescription"
            ></v-text-field>

            <v-text-field
              v-model.number="itemQuantity"
              label="Item Quantity"
              type="number"
              required
            ></v-text-field>

            <v-radio-group
              v-model="itemQuantityType"
              label="Item Quantity Type"
              required
            >
              <v-radio label="LB" value="LB"></v-radio>
              <v-radio label="Piece" value="Piece"></v-radio>
              <v-radio label="Package" value="Package"></v-radio>
              <v-radio label="Case" value="Case"></v-radio>
            </v-radio-group>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" text @click="dialog = false">Cancel</v-btn>
          <v-btn color="primary" @click="submitForm">Submit</v-btn>
        </v-card-actions>
      </v-card>

      <div class="text-center ma-2">
        <v-snackbar v-model="snackbar">
          Form Not Valid Required data!
          <template v-slot:actions>
            <v-btn color="pink" variant="text" @click="snackbar = false">
              Close
            </v-btn>
          </template>
        </v-snackbar>
      </div>
    </v-dialog>
  </div>
</template>

<script>
import { mapActions } from "vuex";
export default {
  name: "AddItem",
  props: {
    addItemFun: Function,
  },
  data() {
    return {
      dialog: false,
      snackbar: false,
      ItemID: null,
      CatID: null,
      SubCatID: null,
      showAddNewCategoryButtonVal: false,
      showAddNewSubCategoryButtonVal: false,
      showAddNewDescriptionButtonVal: false,
      selectedCategory: "", // Updated v-model
      selectedSubCategory: "", // Updated v-model
      selectedItemDescription: "", // Updated v-model
      newcategory: "",
      newsubcategory: "",
      newitemdescription: "",
      itemQuantity: 1,
      itemQuantityType: "",
      categories: [],
    };
  },
  created() {
    this.GetAllData();
  },
  methods: {
    ...mapActions([
      "CreateNewCategory",
      "CreateNewSubCategory",
      "CreateNewItem",
      "GetAllItemsForAdd",
      "AddItemToList",
    ]),
    async GetAllData() {
      try {
        const { data } = await this.GetAllItemsForAdd();
        console.log("get all data ...", data);
        this.categories = data;
      } catch (error) {
        console.log(error);
      }
    },
    async AddOrCloseCategorie() {
      console.log("x", this.categories)
      if (this.newcategory !== "") {
        // this.categories = [];
        // this.categories.push(this.newcategory);
        const {data} =  await this.CreateNewCategory(this.newcategory);
        // const { data } = await this.GetAllItemsForAdd();
        this.categories.push(data.data)
        this.CatID = data.data.id
        console.log("this cat", this.categories, "data", data)
        console.log("c cat", this.categories)
        // console.log("selected cat", this.selectedCategory, "new category", this.newcategory)
        // this.selectedCategory = this.newcategory; // Update selectedCategory
        this.showAddNewCategoryButtonVal = false;

      }
      this.showAddNewCategoryButtonVal = true;
    },
    async AddOrCloseSubCategorie() {
      if (this.newsubcategory !== "") {
        // this.subCategory = [];
        // this.selectedCategory?.subCategories?.push(this.newsubcategory);
        const t = { title: this.newsubcategory, category_id: this.CatID };
        const {data} = await this.CreateNewSubCategory(t);
        this.SubCatID = data.data.id
        console.log("data sub cat ", data.data)
        // const { data } = await this.GetAllItemsForAdd();
        this.selectedSubCategory = data.data;
        this.showAddNewSubCategoryButtonVal = false;
        // this.selectedSubCategory = this.newsubcategory; // Update selectedSubCategory
      } else {
        this.showAddNewSubCategoryButtonVal = true;
      }
    },


    async AddOrCloseDescription() {
      if (this.newitemdescription !== "") {
        this.categories = []

        this.showAddNewDescriptionButtonVal = false;

        console.log("iqt", this.itemQuantityType);
        if (this.itemQuantityType == "") {
          alert("Please select an item type.");
          return;
        }
        // this.selectedItemDescription = this.newitemdescription; // Update selectedItemDescription
        // this.subCategory?.itemDescriptions.push(this.newitemdescription);
        const t = {
          title: this.newitemdescription,
          category_id: this.CatID,
          sub_category_id: this.SubCatID,
          type: this.itemQuantityType,
        };
        const {data} = await this.CreateNewItem(t);
        console.log("add item added", data)
        this.selectedItemDescription = data.data

        // const { data } = await this.GetAllItemsForAdd();
       // this.categories = data;

        // this.dialog = false

      }
      this.showAddNewDescriptionButtonVal = true;

    },


    showAddNewCategoryButton() {
      const category = this.categories?.find(
        (category) => category.title === this.selectedCategory
      );
      this.CatID = category?.id;
      this.selectedCategory = category;
      if (!category) {
        this.showAddNewCategoryButtonVal = true;
      } else {
        this.showAddNewCategoryButtonVal = false;
      }
    },
    showAddNewSubCategoryButton() {
      const suncategory = this.selectedCategory?.subCategories?.find(
        (suncategory) => suncategory.title === this.selectedSubCategory
      );
      this.SubCatID = suncategory?.id;
      this.selectedSubCategory = suncategory;
      if (!suncategory) {
        this.showAddNewSubCategoryButtonVal = true;
      } else {
        this.showAddNewSubCategoryButtonVal = false;
      }
    },
    showAddNewDescriptionButton() {
      const itemdesc = this.selectedSubCategory?.itemDescriptions?.find(
        (itemdesc) => itemdesc.title === this.selectedItemDescription
      );
      this.ItemID = itemdesc?.id;
      this.selectedItemDescription = itemdesc;
      if (!itemdesc) {
        this.showAddNewDescriptionButtonVal = true;
      } else {
        console.log("calllled", itemdesc);
        this.itemQuantityType = itemdesc.itemQuantityType;
        this.showAddNewDescriptionButtonVal = false;
      }
    },
    submitForm() {
      // Perform form submission logic here
      // You can access the form field values using this.category, this.subCategory, etc.
      if (
        this.itemQuantity < 0 ||
        this.itemDescription == "" ||
        this.itemQuantityType == ""
      ) {
        this.snackbar = true;
      } else {
        console.log(
          "item descripton",
          this.itemDescription,
          "item quenity",
          this.itemQuantity,
          "item type",
          this.itemQuantityType,
          "item Category",
          this.category
        );

        this.AddItemToList(this.selectedItemDescription);
        console.log("Form submitted!");
        this.dialog = false; // Close the dialog after form submission
        if (typeof this.addItemFun === "function") {
          this.addItemFun();
        }
      }
    },
  },
};
</script>
