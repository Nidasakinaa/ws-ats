package controller

import (
	"github.com/Nidasakinaa/beats"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetAllBiodata(c *fiber.Ctx) error {
	ps := beats.GetAllBiodata()
	return c.JSON(ps)
}
