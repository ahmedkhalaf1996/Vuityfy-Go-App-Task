import axios from 'axios';

const API = axios.create({ baseURL: 'http://localhost:5000',     withCredentials: true});
API.defaults.headers.post['Content-Type'] = 'application/json';


export const fetchItems = () => {
    const  categories = [
        {
            name: 'Category 1',
            subcategories: [
                { name: 'subcategories 1.1' },
                { name: 'subcategories 1.2' },
                { name: 'subcategories 1.3' }
            ],
            items: [
                { id:4,image: 'path/to/image1.jpg', subcategories:'subcategories 1.1', categories:'Category 1',title: 'Item 1', subtitle: 'Subtitle 1', type:'LB' },
                { id:5,image: 'path/to/image2.jpg',subcategories:'subcategories 1.1',  categories:'Category 1',title: 'Item 2', subtitle: 'Subtitle 2', type:'Piece' },
                { id:6,image: 'path/to/image3.jpg',subcategories:'subcategories 1.2',  categories:'Category 1',title: 'Item 3', subtitle: 'Subtitle 3', type:'Case' },
                { id:7,image: 'path/to/image4.jpg',subcategories:'subcategories 2.1',  categories:'Category 2', title: 'Item 4', subtitle: 'Subtitle 4', type:'Package' }
            ]
            },
            {
            name: 'Category 2',
            subcategories: [
                { name: 'subcategories 2.1' },
                { name: 'subcategories 2.2' },
                { name: 'subcategories 2.3' }
            ],
            items: [
                { id:1,image: 'path/to/image5.jpg',subcategories:'subcategories 1.2',  categories:'Category 1', title: 'Item 5', subtitle: 'Subtitle 5', type:'Piece' },
                { id:2,image: 'path/to/image6.jpg',subcategories:'subcategories 2.2',  categories:'Category 2', title: 'Item 6', subtitle: 'Subtitle 6', type:'LB' },
                { id:3,image: 'path/to/image7.jpg',subcategories:'subcategories 1.1',  categories:'Category 1',title: 'Item 7', subtitle: 'Subtitle 7' , type:'Package'}
            ]
            }
    ]
    return categories;
}

export const fetchallitemsx = ()=>{
    const Items = [
        {id:1,image: 'path/to/image5.jpg', subcategories:'subcategories 2.3',  categories:'Category 2', title: 'Item 5', subtitle: 'Subtitle 5' , type:'Piece'},
        { id:2,image: 'path/to/image6.jpg',subcategories:'subcategories 1.1',  categories:'Category 1', title: 'Item 6', subtitle: 'Subtitle 6' , type:'LB'},
        { id:3,image: 'path/to/image7.jpg',subcategories:'subcategories 2.1',   categories:'Category 2',title: 'Item 7', subtitle: 'Subtitle 7' , type:'Case'},
       { id:4,image: 'path/to/image1.jpg',subcategories:'subcategories 2.1',  categories:'Category 2', title: 'Item 1', subtitle: 'Subtitle 1' , type:'Case'},
       { id:5,image: 'path/to/image2.jpg',subcategories:'subcategories 1.1',  categories:'Category 1', title: 'Item 2', subtitle: 'Subtitle 2' , type:'LB'},
       { id:6,image: 'path/to/image3.jpg',subcategories:'subcategories 1.1',  categories:'Category 1', title: 'Item 3', subtitle: 'Subtitle 3' , type:'Piece'},
       { id:7,image: 'path/to/image4.jpg',subcategories:'subcategories 2.3',  categories:'Category 2', title: 'Item 4', subtitle: 'Subtitle 4' , type:'Case'}
    ]
    return Items;
}

// {id, img, title , sub}

// export const fetchPost = (id) => API.get(`/posts/${id}`);
// export const fetchPosts = (page, id) => API.get(`/posts?page=${page}&id=${id}`);
// export const createPost = (newPost) => API.post('/posts', newPost);
// export const likePost = (id) => API.patch(`/posts/${id}/likePost`);
// export const comment = (value, id) => API.post(`/posts/${id}/commentPost`, { value });
// export const updatePost = (id, updatedPost) => API.patch(`/posts/${id}`, updatedPost);
// export const deletePost = (id) => API.delete(`/posts/${id}`);
// export const fetchPostsUsersBySearch = (searchQuery) => API.get(`/posts/search?searchQuery=${searchQuery}`);



// export const signIn = (formData) => API.post('/user/signin', formData);
// export const signUp = (formData) => API.post('/user/signup', formData);
// export const fetchUserProfile = (id) => API.get(`/user/getUser/${id}`);
// export const getSugUser = (id) => API.get(`/user/getSug?id=${id}`);
// export const UpdateUser = (userData) => API.patch(`/user/Update/${userData._id}`, userData);
// export const following = (id) => API.patch(`/user/${id}/following`);


// -----------------------------------------------------------------


// export const CreateNewCategory = (title) => API.post('/create-category', title)
// export const CreateNewCategory = (title) => axios.post('/items/create-category', title)


export const CreateNewCategory = (data) => API.post('/items/create-category', data)
export const CreateNewSubCategory = (data) => API.post('/items/create-sub-category', data)
export const CreateNewItem = (data) => API.post('/items/create-item', data)
export const GetAllItemsForAdd = ()=> API.get('/items/get-all-items-for-add')

export const fetchallitems = ()=> API.get('/items/getallitems')


// order

export const MakeOrder = (data)=> API.post('/order/create-order', data)
// 	app.Get("/order/get-all-orders", handlers.GetAllOrders)

export const GetAllOrders = () => API.get('/order/get-all-orders')
export const UpdateItemQuntity = (data) => API.post('/order/update-item-order', data)
export const DeleteItem = (data) => API.post('/order/delete-one-item', data)
export const AddToPickList = (id) => API.post('/order/move-to-picklist', id)
export const GetAllDataToPickList = () => API.get('/order/get-picklist')
export const ArcihiveAll = ()=> API.post('/order/arcihive-all')
