package handlers

import (
	"net/http"

	"github.com/ahmedkhalaf1996/Vuityfy-Go-App/models"
	"github.com/gofiber/fiber/v2"
)

type OrderForm struct {
	Email     string          `json:"email"`
	OrderList []OrderItemform `json:"order_list"`
}

type OrderItemform struct {
	ItemID       uint `json:"id"`
	ItemQuantity uint `json:"quantity"`
}

func CreateOrder(c *fiber.Ctx) error {
	var form OrderForm
	if err := c.BodyParser(&form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	order := models.Order{
		Email:      form.Email,
		Status:     false,
		IsArchived: false,
	}

	// Create the order
	if err := models.DB.Create(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create the order",
		})
	}

	for _, itemReq := range form.OrderList {
		var olditem models.Item // main item
		if err := models.DB.First(&olditem, itemReq.ItemID).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve item data",
			})
		}

		//--
		var item models.OrderItem

		item.CategoryId = olditem.CategoryId
		item.SubCategoryId = olditem.SubCategoryId
		item.Img = olditem.Img
		item.Title = olditem.Title
		item.Type = olditem.Type
		item.ItemID = olditem.ID

		item.ItemQuantity = itemReq.ItemQuantity
		item.OrderID = order.ID

		if err := models.DB.Create(&item).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create the item",
			})
		}
	}

	if err := models.DB.Save(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create the order",
		})
	}

	// return c.JSON(order)

	// return succuss
	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  http.StatusCreated,
		"message": "created succussfully",
		"data":    &order,
		"error":   fiber.Map{},
	})

}

// get all orders
type OrderResponse struct {
	ID         uint                `json:"id"`
	Email      string              `json:"email"`
	Status     bool                `json:"status"`
	IsArchived bool                `json:"is_archived"`
	OrderList  []OrderItemResponse `json:"order_list"`
}

type OrderItemResponse struct {
	OrderID       uint   `json:"order_id"`
	ItemID        uint   `json:"item_id"`
	ItemQuantity  uint   `json:"item_quantity"`
	Title         string `json:"title"`
	CategoryID    uint   `json:"category_id"`
	SubCategoryID uint   `json:"sub_category_id"`
	Img           string `json:"img"`
	Type          string `json:"type"`
}

func GetAllOrders(c *fiber.Ctx) error {
	var orders []models.Order
	var orderResponses []OrderResponse

	if err := models.DB.Where("status != ?", true).Find(&orders).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve active orders",
		})
	}

	for _, order := range orders {
		var orderResponse OrderResponse

		orderResponse.ID = order.ID
		orderResponse.Email = order.Email
		orderResponse.Status = order.Status
		orderResponse.IsArchived = order.IsArchived

		var orderItems []models.OrderItem
		if err := models.DB.Where("order_id = ?", order.ID).Find(&orderItems).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve order items",
			})
		}

		var orderItemResponses []OrderItemResponse
		for _, orderItem := range orderItems {
			orderItemResponse := OrderItemResponse{
				OrderID:       orderItem.OrderID,
				ItemID:        orderItem.ItemID,
				ItemQuantity:  orderItem.ItemQuantity,
				Title:         orderItem.Title,
				CategoryID:    orderItem.CategoryId,
				SubCategoryID: orderItem.SubCategoryId,
				Img:           orderItem.Img,
				Type:          orderItem.Type,
			}

			orderItemResponses = append(orderItemResponses, orderItemResponse)
		}

		orderResponse.OrderList = orderItemResponses
		orderResponses = append(orderResponses, orderResponse)
	}

	return c.JSON(orderResponses)
}

// UpdateItemOrder
type UpdateItemOrderForm struct {
	ItemID       uint `json:"item_id"`
	ItemQuantity uint `json:"item_quantity"`
}

func UpdateItemQuntityOrder(c *fiber.Ctx) error {
	var form UpdateItemOrderForm
	if err := c.BodyParser(&form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	item := models.OrderItem{}
	if err := models.DB.First(&item, "id = ?", form.ItemID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	item.ItemQuantity = form.ItemQuantity

	if err := models.DB.Save(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create the order",
		})
	}

	// return succuss
	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  http.StatusCreated,
		"message": "created succussfully",
		"data":    &item,
		"error":   fiber.Map{},
	})

}

type deleteOneItem struct {
	ItemID uint `json:"item_id"`
}

// err := models.DB.Delete(&post).Error
func DeleteOneItem(c *fiber.Ctx) error {
	var form deleteOneItem
	if err := c.BodyParser(&form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	item := models.OrderItem{}
	if err := models.DB.First(&item, "id = ?", form.ItemID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Delete(&item).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// return succuss
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  http.StatusOK,
		"message": "deleted succussfully",
		"data":    "ok",
		"error":   fiber.Map{},
	})
}

// MoveToPickList

type formd struct {
	ID uint `json:"id"`
}

func MoveToPickList(c *fiber.Ctx) error {
	var form formd
	if err := c.BodyParser(&form); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	order := models.Order{}
	if err := models.DB.First(&order, "id = ?", form.ID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// item.ItemQuantity = form.ItemQuantity
	order.Status = true

	if err := models.DB.Save(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create the order",
		})
	}

	// return succuss
	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  http.StatusCreated,
		"message": "created succussfully",
		"data":    &order,
		"error":   fiber.Map{},
	})

}

// ----------------------------------------
// ----------------------------------------
// ----------------------------------------

// GetAllDataToPickList
type OrderResponsePic struct {
	ID         uint                   `json:"id"`
	Email      string                 `json:"email"`
	Status     bool                   `json:"status"`
	IsArchived bool                   `json:"is_archived"`
	OrderList  []OrderItemResponsePic `json:"order_list"`
}

type OrderItemResponsePic struct {
	OrderID       uint   `json:"order_id"`
	ItemID        uint   `json:"item_id"`
	ItemQuantity  uint   `json:"item_quantity"`
	Title         string `json:"title"`
	CategoryID    uint   `json:"category_id"`
	SubCategoryID uint   `json:"sub_category_id"`
	Img           string `json:"img"`
	Type          string `json:"type"`
}

func GetAllDataToPickList(c *fiber.Ctx) error {
	var orders []models.Order
	var orderResponses []OrderResponsePic

	if err := models.DB.Where("status = ?", true).Find(&orders).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve active orders",
		})
	}

	for _, order := range orders {
		var orderResponse OrderResponsePic

		orderResponse.ID = order.ID
		orderResponse.Email = order.Email
		orderResponse.Status = order.Status
		orderResponse.IsArchived = order.IsArchived

		var orderItems []models.OrderItem
		if err := models.DB.Where("order_id = ?", order.ID).Find(&orderItems).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve order items",
			})
		}

		var orderItemResponses []OrderItemResponsePic
		for _, orderItem := range orderItems {
			orderItemResponse := OrderItemResponsePic{
				OrderID:       orderItem.OrderID,
				ItemID:        orderItem.ItemID,
				ItemQuantity:  orderItem.ItemQuantity,
				Title:         orderItem.Title,
				CategoryID:    orderItem.CategoryId,
				SubCategoryID: orderItem.SubCategoryId,
				Img:           orderItem.Img,
				Type:          orderItem.Type,
			}

			orderItemResponses = append(orderItemResponses, orderItemResponse)
		}

		orderResponse.OrderList = orderItemResponses
		orderResponses = append(orderResponses, orderResponse)
	}

	return c.JSON(orderResponses)
}

// ArcihiveAll
func ArchiveAll(c *fiber.Ctx) error {
	// Fetch all orders with status = true
	var orders []models.Order
	if err := models.DB.Where("status = ? AND is_archived = ?", true, false).Find(&orders).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No orders with status true found",
		})
	}

	// Update each order's IsArchived field to true
	for _, order := range orders {
		order.IsArchived = true
		if err := models.DB.Save(&order).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update order",
			})
		}
	}

	// Return success
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Archived successfully",
		"data":    orders,
		"error":   nil,
	})
}
