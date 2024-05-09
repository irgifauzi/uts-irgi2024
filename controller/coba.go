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

func GetPemainData(c *fiber.Ctx) error {
	ps := gile.GetAllPemain()
	return c.JSON(ps)
	
}
