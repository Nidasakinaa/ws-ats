package controller

import (
	cek "github.com/Nidasakinaa/beats/module"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/Nidasakinaa/ws-ats/config"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetAllBiodatas(c *fiber.Ctx) error {
	ps := cek.GetAllBiodata(config.Ulbimongoconn, "Biodata")
	return c.JSON(ps)
}
