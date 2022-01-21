package controllers

import (
	"fis/be/database"
	"fis/be/models"
	"fis/be/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func Login(c *fiber.Ctx) error{
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	db := database.Connect()
	defer db.Close()

	var user models.User
	row, _ := db.Query("SELECT * FROM users "+
		"WHERE email = ?", data["email"])
	if row.Next() {
		row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Email, &user.Password, &user.Role)
	}
	if user.ID == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "not found!",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "password incorrect!",
		})
	}

	token, err := utils.GenerateJwt(strconv.Itoa(int(user.ID)))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "logout successfully!",
	})
}

func GetUserInfo(c *fiber.Ctx) error {
	cookie := c.Cookies("token")
	id, err := utils.ParseJwt(cookie)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message" : "Unauthorized",
		})
	}
	var user models.User
	db := database.Connect()
	defer db.Close()
	row, _ := db.Query("select u.id, u.name, u.phone_number, u.email, r.name "+
		"from users u join roles r on r.id = u.role_id where u.id = ? ", id)
	if row.Next() {
		row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Email, &user.Role)
	}
	return c.JSON(user)
}

func UpdateInfo(c *fiber.Ctx) error {
	cookie := c.Cookies("token")
	id, err := utils.ParseJwt(cookie)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message" : "Unauthorized",
		})
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	db := database.Connect()
	defer db.Close()

	_, err = db.Query("update users set name = ?, phone_number = ?, email = ? "+
		"where id = ?", user.Name, user.PhoneNumber, user.Email, id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if data["password"] != data["password_confirm"] {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "password incorrect!",
		})
	}

	cookie := c.Cookies("token")

	id, _ := utils.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)
	user := models.User{
		ID: int64(userId),
	}
	user.SetPassword(data["password"])

	db := database.Connect()
	defer db.Close()

	_, err := db.Query("update users set password = ? "+
		"where id = ?", user.Password, id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "update password successfully!",
	})
}

