
import * as api from '../api/index.js';

const Items = {
    state:{},
    getters:{ },
    mutations: {},
    actions: {
        // async CreateNewCategory(context,title){
        //   try {
        //     const t = {"title":title}
        //     console.log("d", t)
        //     const data = await api.CreateNewCategory(t)
        //     return data;
        //   } catch (error) {
        //     console.log(error)
        //   }
        // },

        async MakeOrder(context, form){
            try {
                const data = await api.MakeOrder(form)
                return data;
            } catch (error) {
                console.log("error", error)
            }
        },
        async GetOrderdList(){
            try {
                const data = await api.GetAllOrders();
                return data;
            } catch (error) {
                console.log("err", error)
            }
        },
        async UpdateItemQuntity(context, form){
            try {
                const data = await api.UpdateItemQuntity(form);
                return data;
            } catch (error) {
                console.log("err", error)
            }
        },
        async DeleteItem(context, form ){
            try {
                const data = await api.DeleteItem(form)
                return data;
            } catch (error) {
                console.log('error', error)
            }
        },
        async AddToPickList(context, id){
            try {
                const d = {"id":id}
                const data = await api.AddToPickList(d)
                return data;
            } catch (error) {
                console.log(error)
            }
        },
        async GetAllDataToPickList(){
            try {
                const data = await api.GetAllDataToPickList()
                return data;
            } catch (error) {
                console.log(error)
            }
        }, 
        async ArcihiveAllSt(){
            try {
                const data = await api.ArcihiveAll()
                return data
            } catch (error) {
                console.log(error)
            }
        }


    }
}

export default Items;