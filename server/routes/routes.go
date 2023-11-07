package routes

import (
	"github.com/ahmedkhalaf1996/Vuityfy-Go-App/handlers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// -----------------   POSTS   ---------------//
	// app.Post("/post", handlers.CreatePost)
	// app.Get("/post/:id", handlers.GetPostByid)
	// app.Put("/post/:id", handlers.UpdatePost)
	// app.Delete("/post/:id", handlers.DeletePost)

	// ------------- items --------------- //

	app.Post("/items/create-category", handlers.CreateCategory)
	app.Post("/items/create-sub-category", handlers.CreateSubCategory)
	app.Post("/items/create-item", handlers.CreateItem)
	app.Get("/items/get-all-items-for-add", handlers.GetAllItemsForAdd)
	app.Get("/items/getallitems", handlers.Getallitems)

	// -------------  Order  ------------------------------//
	app.Post("/order/create-order", handlers.CreateOrder)
	app.Get("/order/get-all-orders", handlers.GetAllOrders)

	app.Post("/order/update-item-order", handlers.UpdateItemQuntityOrder)
	app.Post("/order/delete-one-item", handlers.DeleteOneItem)
	app.Post("/order/move-to-picklist", handlers.MoveToPickList)
	app.Get("/order/get-picklist", handlers.GetAllDataToPickList)
	app.Post("/order/arcihive-all", handlers.ArchiveAll)
}
