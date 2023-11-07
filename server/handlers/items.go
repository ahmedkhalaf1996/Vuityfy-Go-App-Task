package handlers

import (
	"fmt"
	"net/http"

	"github.com/ahmedkhalaf1996/Vuityfy-Go-App/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Type          string `json:"type"  binding:"required"`
// Quantity      int `json:"quantity"  binding:"required"`

type CategoriesBody struct {
	Title string `json:"title"  binding:"required"`
}

type CategoryRes struct {
	ID            uint                  `json:"id"`
	Title         string                `json:"title"`
	SubCategories []SubCategoryResponse `json:"subCategories"`
}

// create categores
func CreateCategory(c *fiber.Ctx) error {
	req := &CategoriesBody{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	fmt.Println("data", *req)
	// create category
	cat := models.Categories{
		Title: req.Title,
	}

	err := models.DB.Create(&cat).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "some thing went worng!",
			"error": fiber.Map{
				"deatils": "unable to create Category , internal server error",
			},
		})
	}

	var res = CategoryRes{
		ID:            cat.ID,
		Title:         cat.Title,
		SubCategories: []SubCategoryResponse{},
	}

	// return succuss
	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  http.StatusCreated,
		"message": "created succussfully",
		"data":    res,
		"error":   fiber.Map{},
	})

}

type SubCategory struct {
	Title      string `json:"title"  binding:"required"`
	CategoryId int    `json:"category_id"  binding:"required"`
}

// create subcategory
func CreateSubCategory(c *fiber.Ctx) error {
	req := &SubCategory{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	subcat := models.SubCategory{
		Title:      req.Title,
		CategoryId: uint(req.CategoryId),
	}

	err := models.DB.Create(&subcat).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "some thing went worng!",
			"error": fiber.Map{
				"deatils": "unable to create Sub Category , internal server error",
			},
		})
	}

	// subCategoryResponse = append(subCategoryResponse, SubCategoryResponse{
	// 	ID:               subCategory.ID,
	// 	Title:            subCategory.Title,
	// 	ItemDescriptions: itemDescriptions,
	// })
	var sc = SubCategoryResponse{
		ID:               subcat.ID,
		Title:            subcat.Title,
		ItemDescriptions: []ItemDescription{},
	}

	// return succuss
	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  http.StatusCreated,
		"message": "created succussfully",
		"data":    &sc,
		"error":   fiber.Map{},
	})

}

// create item
type Item struct {
	Title         string `json:"title"  binding:"required"`
	CategoryId    int    `json:"category_id"  binding:"required"`
	SubCategoryId int    `json:"sub_category_id"`
	Type          string `json:"type"`
}

func CreateItem(c *fiber.Ctx) error {
	req := &Item{}
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	item := models.Item{
		Title:         req.Title,
		CategoryId:    uint(req.CategoryId),
		SubCategoryId: uint(req.SubCategoryId),
		Img:           "https://thecaviarco.com/cdn/shop/products/gift-card-gift-card-1_949x.png?v=1508887172",
		Type:          req.Type,
	}

	err := models.DB.Create(&item).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "some thing went worng!",
			"error": fiber.Map{
				"deatils": "unable to create Item , internal server error",
			},
		})
	}

	var itd = ItemDescription{
		ID:    item.ID,
		Title: item.Title,
		Img:   item.Img,
		Type:  item.Type,
	}
	// return succuss
	return c.Status(http.StatusCreated).JSON(&fiber.Map{
		"status":  http.StatusCreated,
		"message": "created succussfully",
		"data":    &itd,
		"error":   fiber.Map{},
	})

}

// func GetAllItemsForAdd(c *fiber.Ctx) error {
// 	return nil
// }

type CategoryResponse struct {
	ID            uint                  `json:"id"`
	Title         string                `json:"title"`
	SubCategories []SubCategoryResponse `json:"subCategories"`
}

type SubCategoryResponse struct {
	ID               uint              `json:"id"`
	Title            string            `json:"title"`
	ItemDescriptions []ItemDescription `json:"itemDescriptions"`
}

type ItemDescription struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Img   string `json:"img"`
	Type  string `json:"type"`
}

func GetAllItemsForAdd(c *fiber.Ctx) error {

	// Fetch all categories from the database
	var categories []models.Categories
	if err := models.DB.Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching categories",
		})
	}

	categoryResponse := make([]CategoryResponse, 0)

	for _, category := range categories {
		subCategories, err := getSubCategoriesAndItems(models.DB, category.ID, category.Title)
		if err != nil {
			return err
		}

		categoryResponse = append(categoryResponse, CategoryResponse{
			ID:            category.ID,
			Title:         category.Title,
			SubCategories: subCategories,
		})
	}

	return c.JSON(fiber.Map{
		"data": categoryResponse,
	})
}

func getSubCategoriesAndItems(db *gorm.DB, categoryID uint, categoryName string) ([]SubCategoryResponse, error) {
	var subCategories []models.SubCategory
	if err := db.Where("category_id = ?", categoryID).Find(&subCategories).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Error fetching sub-categories")
	}

	subCategoryResponse := make([]SubCategoryResponse, 0)

	for _, subCategory := range subCategories {
		itemDescriptions, err := getItemDescriptions(db, categoryID, subCategory.ID)
		if err != nil {
			return nil, err
		}

		subCategoryResponse = append(subCategoryResponse, SubCategoryResponse{
			ID:               subCategory.ID,
			Title:            subCategory.Title,
			ItemDescriptions: itemDescriptions,
		})
	}

	return subCategoryResponse, nil
}

func getItemDescriptions(db *gorm.DB, categoryID, subCategoryID uint) ([]ItemDescription, error) {
	var items []models.Item
	if err := db.Where("category_id = ? AND sub_category_id = ?", categoryID, subCategoryID).Find(&items).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Error fetching item descriptions")
	}

	itemDescriptions := make([]ItemDescription, 0)

	for _, item := range items {
		itemDescriptions = append(itemDescriptions, ItemDescription{
			ID:    item.ID,
			Title: item.Title,
			Img:   item.Img,
			Type:  item.Type,
		})
	}

	return itemDescriptions, nil
}

func Getallitems(c *fiber.Ctx) error {
	var allitems []models.Item
	if err := models.DB.Find(&allitems).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching allitems",
		})
	}

	return c.JSON(fiber.Map{
		"data": allitems,
	})
}
