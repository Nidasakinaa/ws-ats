package controller

import (
	cek "github.com/Nidasakinaa/beats/module"
	"github.com/Nidasakinaa/ws-ats/config"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetAllBiodata(c *fiber.Ctx) error {
	ps := cek.GetAllBiodata(config.Ulbimongoconn, "biodata")
	return c.JSON(ps)
}
