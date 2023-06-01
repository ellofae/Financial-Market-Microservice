package handlers

import "github.com/gofiber/fiber/v2"

func (g *GetRouter) TeamPage(c *fiber.Ctx) error {
	c.GetRespHeader("Content-Type")

	g.log.Info("Sending page with currency exchange rates to the client's request", "request's URL", c.Path)

	c.Render("team", fiber.Map{})
	return c.SendStatus(fiber.StatusOK)
}
