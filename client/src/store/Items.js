
// const Items = {
//     state: { isLoading: true,
//post:[], posts: [], SearchResult:[] },
//     getters: {
//      GetPost: (state) => () => {
//       return {...state.post};
//      },
//      GetAllPosts: (state) => () => {
//         return {...state.posts};
//      },
//      GetSearchData: (state) => () => {
//         return {...state.SearchResult};
//      },
//     },
//     mutations: {
//        Post(state, payload){
//           state.post = payload;
//        },
//        Posts(state, payload){
//          state.posts = payload
//        },
//        Search(state, payload){
//         state.SearchResult = payload
//        }
//     },
//     actions: {
//         async getPost(context, id){
//              try {
//                 let {data} =  await api.fetchPost(id);

//                 context.commit('Post',data)
                
//                 return data
//              } catch (error) {
//                 console.log(error)
//              }
//         },
//         async  getPosts(context, page){
//             try {
//                 const user = JSON.parse(localStorage.getItem('profile'));
//                 const userId = user?.result?._id;
//                 if(userId){
//                   const { data } = await api.fetchPosts(page, userId);
//                   context.commit('Posts',data)
//                   return data;
//                 }
            
//               } catch (error) {
//                 console.log(error);
//               }
//         },
//         async  getPostsUsersBySearch(context,serchData){
//             try {
                
//                 const   { data }   = await api.fetchPostsUsersBySearch(serchData);
//                 context.commit('Search',data)
//                 return data;
//               } catch (error) {
//                 console.log(error);
//               }
//         },
//         async  createPost(context, post){
//           try {
//             const { data } = await api.createPost(post);

//             context.commit('Post',data)

//             return data;
//           } catch (error) {
//             return error;
//           }
//         },
//         async  updatePost(context,PostData ){
//           // console.log("post",PostData.id, PostData)
//           const ud = JSON.parse(localStorage.getItem('profile'));
//           const pooStData = 
//             {
//               "title": PostData.title,
//               "message": PostData.message,
//               "creator": ud?.result?._id,
//               "selectedFile": PostData.selectedFile
//             }
          
//           // try {
//               const  post  = await api.updatePost(PostData.id, pooStData);

//               context.commit('Post',post)
//             return post;
//             // } catch (error) {
//             //   console.log(error.message)
//             // }
//         },
//         async  LikePostByUser(context, id){
//             const user = JSON.parse(localStorage.getItem('profile'));

//             const { data } = await api.likePost(id, user?.token);
//             context.commit('Post',data)
//             console.log('data', data)
//         },
//         async  commentPost(context,form ){
//          // console.log(data)
//             const { data } = await api.comment(form.value, form.id);

//             context.commit('Post',data)
//         },
//         async  deletePost (context, id){
//             await await api.deletePost(id);
//            // context.commit('',data)
//         }

//     }
// }
import * as api from '../api/index.js';

const Items = {
    state:{CategoriesItemsList:[], ItemsList:[]},
    getters:{
        GetCategoriesItemsList : (state) => ()=>{
            return {...state.CategoriesItemsList};
        },
        GetAllItemsList: (state) => () =>{
            return {...state.ItemsList}
        }
    },
    mutations: {
        SetCategoriesItemsList(state, payload){
            state.CategoriesItemsList = payload
        },
        SetAllItemsList(state, payload){
            state.ItemsList = payload
        }
    },
    actions: {
        // async GetCategoriesItems(context){
        //     // let {categories} = await api.fetchItems()
        //     let categories =  api.fetchItems()

        //     console.log("store Get Items", categories)

        //     context.commit('SetCategoriesItemsList', categories)
        //     return categories;
        // }, 
        async GetAllItems(context){
            let items = api.fetchallitems()
            // console.log("store all items", items)
            context.commit('SetAllItemsList', items)
            return items
        },
        async AddItemToList(context, payload){
            const items = JSON.parse(localStorage.getItem('items'));            
            if (items) {
              const itemsArray = Object.values(items);
              const item = itemsArray.find((item) => item?.title === payload?.title);
              if (payload?.quantity && item){
               item.quantity = payload.quantity 

              }
              else if (item) {

                item.quantity += 1;
              } else {
                payload.quantity = 1;
                itemsArray.push(payload);
            
              }
              localStorage.setItem('items', JSON.stringify(itemsArray));
            } else {
              payload.quantity = 1;
              localStorage.setItem('items', JSON.stringify([payload]));
            }
        }, 
        
        // async ClearItemList(){
        //   localStorage.clear();
        // },
        async UpdateItemList(context, payload) {
          try {
            const items = JSON.parse(localStorage.getItem('items'));
        
            if (items) {
              const itemsArray = Object.values(items);
              const index = itemsArray.findIndex((item) => item.title === payload.title);
              console.log("store ff >", payload.quantity)

              if (index !== -1) {
                // Check if the quantity is zero
                if (payload.quantity === 0) {
                  // Remove the item with quantity zero
                  console.log("store >", payload.quantity)
                  itemsArray.splice(index, 1);
                } else {
                  // Replace the old item with the new item
                  itemsArray[index] = payload;
                }
              }
        
              localStorage.setItem('items', JSON.stringify(itemsArray));
            }
          } catch (error) {
            console.error('Error updating/removing item:', error);
          }
        },
        
        


        async GetItemsFormLocal(){
            const items = JSON.parse(localStorage.getItem('items'));
            if (items){
                return items
            } else {
                return []
            }
            
        },
        async CreateNewCategory(context,title){
          try {
            const t = {"title":title}
            console.log("d", t)
            const data = await api.CreateNewCategory(t)
            return data;
          } catch (error) {
            console.log(error)
          }
        },
        async CreateNewSubCategory(context,sdata){
          try {
            const t = {"title":sdata.title, "category_id":sdata.category_id}
            console.log("d", t)
            const data = await api.CreateNewSubCategory(t)
            return data;
          } catch (error) {
            console.log(error)
          }
        },
        async CreateNewItem(context,sdata){
          try {
            const t = {"title":sdata.title, "category_id":sdata.category_id, 
             "sub_category_id": sdata.sub_category_id, "type": sdata.type}
            console.log("d", t)
            const data = await api.CreateNewItem(t)
            return data;
          } catch (error) {
            console.log(error)
          }
        },

        // get all data to add item
        async GetAllItemsForAdd(){
          try {
            const {data} = await api.GetAllItemsForAdd();
            return data;
          } catch (error) {
            console.log(error)
          }
        }

    }
}

export default Items;