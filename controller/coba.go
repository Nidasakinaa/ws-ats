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

func GetAllBiodatas(c *fiber.Ctx) error {
	ps := cek.GetAllBiodata(config.Ulbimongoconn, "Biodata")
	return c.JSON(ps)
}
