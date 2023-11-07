package handlers

// import (
// 	"fmt"
// 	"net/http"
// 	"path"
// 	"strconv"
// 	"time"

// 	"github.com/ahmedkhalaf1996/Vuityfy-Go-App/models"
// 	"github.com/gofiber/fiber/v2"
// )

// type PostBody struct {
// 	CreatorId int    `json:"creator_id"  binding:"required"`
// 	PostTitle string `json:"post_title"  binding:"required"`
// 	PostBody  string `json:"post_body"  binding:"required"`
// 	// PostFile  string `json:"post_file" binding:"required"`
// }

// type UpdatePostBody struct {
// 	PostTitle string `json:"post_title"  binding:"required"`
// 	PostBody  string `json:"post_body"  binding:"required"`
// 	// PostFile  string `json:"post_file" binding:"required"`
// }

// type AllPosts struct {
// 	ID            int       `json:"id"`
// 	CreatedAt     time.Time `json:"CreatedAt"`
// 	UpdatedAt     time.Time `json:"UpdatedAt"`
// 	CreatorId     int       `json:"creator_id"`
// 	PostTitle     string    `json:"post_title"`
// 	PostBody      string    `json:"post_body"`
// 	PostFile      string    `json:"post_file"`
// 	IsDirectScore bool      `json:"is_direct_score"`
// 	LikeCount     int       `json:"like_count"`
// 	CommmentCount int       `json:"commment_count"`
// 	CreatorName   string    `json:"creator_name"`
// 	CreatorImage  string    `json:"creator_image"`
// }

// // We have Tow ways to create Posts
// // Manual & Automatic

// // Todo !! Create Auto posts

// // manual || CreatePost
// func AutoMaticCreatePost(c *fiber.Ctx) error {
// 	gameid := c.Locals("gameId").(string)
// 	game := models.Game{}

// 	models.DB.First(&game, "id = ?", gameid)
// 	// create post
// 	post := models.Posts{
// 		CreatorId:     game.UserId,
// 		PostTitle:     "Finshed Game",
// 		PostBody:      "score: " + fmt.Sprint(game.TotalScore) + " | Course Rating: " + fmt.Sprint(game.CourseRating) + " | Slope Rating: " + fmt.Sprint(game.SlopeRating),
// 		IsDirectScore: true,
// 		// PostImage: game.PostImage,
// 	}

// 	models.DB.Create(&post)

// 	// return succuss
// 	return nil

// }

// // manual || CreatePost
// func CreatePost(c *fiber.Ctx) error {
// 	// req := &PostBody{}
// if err := c.BodyParser(req); err != nil {
// 	return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
// 		"message": err.Error(),
// 	})
// }
// 	CreatorId := c.FormValue("creator_id")
// 	PostTitle := c.FormValue("post_title")
// 	PostBody := c.FormValue("post_body")

// 	// file
// 	file, errx := c.FormFile("post_file")
// 	if errx != nil {
// 		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
// 			"status":  http.StatusBadRequest,
// 			"message": "some thing went worng!",
// 			"error": fiber.Map{
// 				"deatils": errx.Error(),
// 			},
// 		})
// 	}
// 	// c.SaveFile(file, "public/"+file.Filename)
// 	extension := path.Ext(file.Filename) //obtain the extension of file
// 	randfilename := "util.RandomString(100)"
// 	filename := randfilename + extension
// 	c.SaveFile(file, "public/"+filename)
// 	// chekc if user exist
// 	user := models.User{}
// 	if err := models.DB.First(&user, "id = ?", CreatorId).Error; err != nil {
// 		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
// 			"status":  http.StatusNotFound,
// 			"message": "some thing went worng!",
// 			"error": fiber.Map{
// 				"deatils": "reocred not found! ",
// 			},
// 		})
// 	}
// 	cid, _ := strconv.Atoi(CreatorId)
// 	// create post
// 	post := models.Posts{
// 		CreatorId: cid,
// 		PostTitle: PostTitle,
// 		PostBody:  PostBody,
// 		PostFile:  filename,
// 	}

// 	err := models.DB.Create(&post).Error

// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": "some thing went worng!",
// 			"error": fiber.Map{
// 				"deatils": "unable to create post , internal server error",
// 			},
// 		})
// 	}

// 	// return succuss
// 	return c.Status(http.StatusCreated).JSON(&fiber.Map{
// 		"status":  http.StatusCreated,
// 		"message": "created succussfully",
// 		"post":    &post,
// 		"error":   fiber.Map{},
// 	})

// }

// // get post by id | GetPostByid/:id
// func GetPostByid(c *fiber.Ctx) error {
// 	// check if post exist
// 	var postid = c.Params("id")
// 	var post models.Posts
// 	if err := models.DB.First(&post, "id = ?", postid).Error; err != nil {
// 		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
// 			"status":  http.StatusNotFound,
// 			"message": "some thing went worng!",
// 			"error": fiber.Map{
// 				"deatils": "Recored Not Found!",
// 			},
// 		})
// 	}

// 	var comments []models.Comments
// 	var likes []models.Likes

// 	models.DB.Find(&comments, "post_id = ?", postid)
// 	models.DB.Find(&likes, "post_id = ?", postid)
// 	// Todo | get post comments & Likes
// 	// return the post
// 	return c.Status(http.StatusOK).JSON(&fiber.Map{
// 		"status":    http.StatusOK,
// 		"message":   "success",
// 		"error":     fiber.Map{},
// 		"post":      post,
// 		"commments": comments,
// 		"likes":     likes,
// 	})
// }

// // update post
// func UpdatePost(c *fiber.Ctx) error {
// 	// first, check if the post is exist
// 	post := models.Posts{}
// 	if err := models.DB.First(&post, "id = ?", c.Params("id")).Error; err != nil {
// 		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
// 			"status":  http.StatusNotFound,
// 			"message": "some thing went worng!",
// 			"error": fiber.Map{
// 				"post_id": "post with givend id is not found",
// 			},
// 		})
// 	}

// 	// check if we have the same user
// 	uidFromAuth := c.Locals("userId").(string)
// 	if uidFromAuth != fmt.Sprint(post.CreatorId) {
// 		return c.Status(http.StatusMethodNotAllowed).JSON(&fiber.Map{
// 			"status":  http.StatusMethodNotAllowed,
// 			"message": "some thing went worng!",
// 			"error": fiber.Map{
// 				"deatils": "Access Not Allowed",
// 			},
// 		})
// 	}

// 	// parse the request body
// 	request := &UpdatePostBody{}
// 	if err := c.BodyParser(request); err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
// 			"status":  http.StatusBadRequest,
// 			"message": "some thing went worng!",
// 			"error": fiber.Map{
// 				"deatils": err.Error(),
// 			},
// 		})
// 	}

// 	// update the post
// 	updatedpost := models.Posts{
// 		PostTitle: request.PostTitle,
// 		PostBody:  request.PostBody,
// 	}

// 	err := models.DB.Model(&post).Updates(&updatedpost).Error

// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": "some thing went worng!",
// 			"error": fiber.Map{
// 				"deatils": err.Error(),
// 			},
// 		})
// 	}

// 	return c.Status(http.StatusOK).JSON(&fiber.Map{
// 		"status":  http.StatusOK,
// 		"message": "success",
// 		"error":   fiber.Map{},
// 		"Post":    post,
// 	})
// }

// // Todo
// // When We Remove Post We Should Remove The {Comments & Likes}..
// // delete post --
// func DeletePost(c *fiber.Ctx) error {
// 	// chekc if post exist
// 	var postid = c.Params("id")
// 	if postid == "" {
// 		postid = c.Locals("post_id").(string)
// 	}

// 	post := models.Posts{}
// 	if err := models.DB.First(&post, "id = ?", postid).Error; err != nil {
// 		return c.Status(http.StatusNotFound).JSON(&fiber.Map{
// 			"status":  http.StatusNotFound,
// 			"message": "some thing went worng!",
// 			"error": fiber.Map{
// 				"deatils": "Record not found!",
// 			},
// 		})
// 	}

// 	// check if we have the provided user token is the same that create the post
// 	uidFromAuth := c.Locals("userId").(string)
// 	if uidFromAuth != fmt.Sprint(post.CreatorId) {
// 		return c.Status(http.StatusMethodNotAllowed).JSON(&fiber.Map{
// 			"status":  http.StatusMethodNotAllowed,
// 			"message": "some thing went worng!",
// 			"error": fiber.Map{
// 				"deatils": "Access Not Allowed, provided token is not valid",
// 			},
// 		})
// 	}

// 	// Todo Delete  likes
// 	comments := []models.Comments{}
// 	models.DB.Find(&comments, "post_id = ?", postid)

// 	likes := []models.Likes{}
// 	models.DB.Find(&likes, "post_id = ?", postid)

// 	// delete the post & commentsn
// 	err := models.DB.Delete(&post).Error
// 	err1 := models.DB.Delete(&comments).Error
// 	err2 := models.DB.Delete(&likes).Error

// 	if err != nil && err1 != nil && err2 != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": "some thing went worng!",
// 			"error": fiber.Map{
// 				"deatils": "Can't delete the post",
// 			},
// 		})
// 	}

// 	return c.Status(http.StatusOK).JSON(&fiber.Map{
// 		"status":  http.StatusOK,
// 		"message": "Success",
// 		"error":   fiber.Map{},
// 	})

// }

// // note when we delete should delete the comments likes to

// // get posts for specific user With {pagenation}
// func GetPostsByUserId(c *fiber.Ctx) error {

// 	userid := c.Params("id")
// 	page, _ := strconv.Atoi(c.Params("page"))
// 	pageSize, _ := strconv.Atoi(c.Params("page_size"))

// 	var postsList []models.Posts

// 	// models.DB.Where("creator_id = ?", userid).Find(&postsList)
// 	models.DB.Scopes(Paginate(page, pageSize)).Where("creator_id = ?", userid).Find(&postsList)

// 	// get user
// 	creator := models.User{}
// 	models.DB.First(&creator, "id = ?", userid)

// 	// loop
// 	var upPost = AllPosts{}
// 	var Afinal []interface{}
// 	// var Allff &[]upPost

// 	for _, el := range postsList {

// 		upPost.ID = int(el.ID)
// 		upPost.CreatedAt = el.CreatedAt
// 		upPost.UpdatedAt = el.UpdatedAt
// 		upPost.CreatorId = el.CreatorId
// 		upPost.PostTitle = el.PostTitle
// 		upPost.PostBody = el.PostBody
// 		upPost.PostFile = el.PostFile
// 		upPost.IsDirectScore = el.IsDirectScore
// 		upPost.CreatorName = creator.Name
// 		upPost.CreatorImage = creator.Savg
// 		// likes
// 		var likesList []models.Likes

// 		likecount := models.DB.Where("post_id = ?", el.ID).Find(&likesList).RowsAffected

// 		upPost.LikeCount = int(likecount)
// 		// comments
// 		var commentsList []models.Comments

// 		commentscount := models.DB.Where("post_id = ?", el.ID).Find(&commentsList).RowsAffected
// 		upPost.CommmentCount = int(commentscount)

// 		// Afinal = append(Afinal, upPost)
// 		Afinal = append([]interface{}{upPost}, Afinal...)

// 	}

// 	return c.Status(http.StatusOK).JSON(&fiber.Map{
// 		"status":    http.StatusOK,
// 		"message":   "success",
// 		"post_list": Afinal,
// 		"error":     fiber.Map{},
// 	})
// }

// // .Order("id ESC")
// // get posts for specific user With {pagenation}
// func GetPostsByUserIdpublic(c *fiber.Ctx) error {

// 	// userid := c.Params("id")
// 	page, _ := strconv.Atoi(c.Params("page"))
// 	pageSize, _ := strconv.Atoi(c.Params("page_size"))

// 	var postsList []models.Posts

// 	// models.DB.Where("creator_id = ?", userid).Find(&postsList)
// 	models.DB.Scopes(Paginate(page, pageSize)).Find(&postsList)

// 	// get user
// 	// creator := models.User{}
// 	// models.DB.First(&creator, "id = ?", userid)

// 	// loop
// 	var upPost = AllPosts{}
// 	var Afinal []interface{}
// 	// var Allff &[]upPost

// 	for _, el := range postsList {

// 		upPost.ID = int(el.ID)
// 		upPost.CreatedAt = el.CreatedAt
// 		upPost.UpdatedAt = el.UpdatedAt
// 		upPost.CreatorId = el.CreatorId
// 		upPost.PostTitle = el.PostTitle
// 		upPost.PostBody = el.PostBody
// 		upPost.PostFile = el.PostFile
// 		upPost.IsDirectScore = el.IsDirectScore

// 		creator := models.User{}
// 		models.DB.First(&creator, "id = ?", el.CreatorId)

// 		upPost.CreatorName = creator.Name
// 		upPost.CreatorImage = creator.Savg
// 		// likes
// 		var likesList []models.Likes

// 		likecount := models.DB.Where("post_id = ?", el.ID).Find(&likesList).RowsAffected

// 		upPost.LikeCount = int(likecount)
// 		// comments
// 		var commentsList []models.Comments

// 		commentscount := models.DB.Where("post_id = ?", el.ID).Find(&commentsList).RowsAffected
// 		upPost.CommmentCount = int(commentscount)

// 		// data = append(data, "Append Item")
// 		// data = append([]string{"Prepend Item"}, data...)

// 		// Afinal = append(Afinal, upPost)
// 		Afinal = append([]interface{}{upPost}, Afinal...)
// 	}

// 	return c.Status(http.StatusOK).JSON(&fiber.Map{
// 		"status":    http.StatusOK,
// 		"message":   "success",
// 		"post_list": Afinal,
// 		"error":     fiber.Map{},
// 	})
// }

// // /
// func DeleteUserPosts(c *fiber.Ctx) bool {
// 	userid := c.Locals("userId")

// 	var postsList []models.Posts

// 	models.DB.Where("creator_id = ? ", userid).Find(&postsList)

// 	// loop
// 	for _, el := range postsList {
// 		// likes
// 		var likes []models.Likes
// 		models.DB.Where("post_id = ?", el.ID).Find(&likes)
// 		go models.DB.Delete(&likes)

// 		// comments
// 		var comments []models.Comments
// 		models.DB.Where("post_id = ?", el.ID).Find(&comments)
// 		go models.DB.Delete(&comments)
// 	}

// 	go models.DB.Delete(&postsList)

// 	return true
// }

// // Todo ..
// // Get Frends&otherusers posts (feed)

// // [PostSys] activity feed page: where you see the (posts) scores and games of your friends
