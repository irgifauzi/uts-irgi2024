// File controller.go
package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	gile "github.com/irgifauzi/bec_uts"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPemain(c *fiber.Ctx) error {
	bola, err := gile.PemainGetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(bola)
}
