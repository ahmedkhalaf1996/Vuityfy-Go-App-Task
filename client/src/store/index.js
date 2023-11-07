import { createStore } from 'vuex'
import Items from './Items'
import Order from './order'


export default createStore({
  
  modules: {
    Items,
    Order
  }
})