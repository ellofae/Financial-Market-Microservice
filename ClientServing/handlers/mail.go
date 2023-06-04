package handlers

import (
	"github.com/ellofae/Financial-Market-Microservice/ClientServing/data"
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-hclog"
)

type UserData struct {
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Message   string `json:"message"`
}

type MailPart struct {
	log       hclog.Logger
	mailServe *data.MailServer
}

func NewMailService(log hclog.Logger, msrv *data.MailServer) *MailPart {
	m := &MailPart{log: log, mailServe: msrv}
	return m
}

func (m *MailPart) MailServe(c *fiber.Ctx) error {
	m.log.Info("Sending an email to the developer from the user")

	firstName := c.FormValue("FirstNameField")
	lastName := c.FormValue("LastNameField")
	message := c.FormValue("MessageField")

	if firstName == "" || lastName == "" || message == "" {
		m.log.Error("No data was filled into the form sections", "First Name Field", firstName, "Last Name Field", lastName, "Message", message)
		c.Redirect("/contact")
		return c.SendStatus(400)
	}
	m.log.Info("User data to send", "First Name", firstName, "Last Name", lastName, "Message", message)

	err := m.mailServe.SendMail([]byte(message))
	if err != nil {
		m.log.Error("Unable to send an email", "error", err)
		c.Redirect("/contact")
		return c.SendStatus(500)
	}

	c.Render("contact", fiber.Map{})
	return c.SendStatus(fiber.StatusOK)
}
