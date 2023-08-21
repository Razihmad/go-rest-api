package controller

import (
	"fmt"
	"log"
	"strconv"
	"time"

	db "github.com/Razihmad/go-rest-api/config"
	"github.com/Razihmad/go-rest-api/model"
	"github.com/gofiber/fiber/v2"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatal(err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Data",
			"error":   true,
			"data":    nil,
		})

	}
	if data["name"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Name is required",
			"error":   true,
			"data":    nil,
		})

	}
	if data["password"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Passcode is required",
			"error":   true,
			"data":    nil,
		})
	}

	cashierModel := model.Cashier{
		Name:      data["name"],
		Email:     data["email"],
		Password:  data["password"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	db.DB.Create(&cashierModel)

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"error":   false,
		"data":    cashierModel,
	})
}

func GetCashiers(c *fiber.Ctx) error {
	var cashiers []model.Cashier
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64
	db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashiers).Count(&count)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"error":   false,
		"data":    cashiers,
	})

}

func GetCashierById(c *fiber.Ctx) error {
	cashierId := c.Params("id")
	fmt.Println(cashierId)
	var cashier model.Cashier
	db.DB.Where("id = ?", cashierId).First(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"error":   false,
		"data":    cashier,
	})

}

func UpdateCashier(c *fiber.Ctx) error {
	cashierId := c.Params("id")
	var cashier model.Cashier
	db.DB.Find(&cashier, cashierId)
	if cashier.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "Cashier not found",
			"error":   true,
			"data":    nil,
		})
	}
	var updateCashier model.Cashier
	err := c.BodyParser(&updateCashier)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid Data",
			"error":   true,
			"data":    nil,
		})

	}
	db.DB.Model(&cashier).Updates(updateCashier)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"error":   false,
		"data":    cashier,
	})

}

func DeleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("id")
	var cashier model.Cashier
	db.DB.Find(&cashier, cashierId)
	if cashier.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Cashier not found",
			"error":   true,
			"data":    nil,
		})
	}
	db.DB.Where("id = ?", cashierId).Delete(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
		"error":   false,
		"data":    nil,
	})

}
