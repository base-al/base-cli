package {{.ModuleName}}

import (
	"github.com/gofiber/fiber/v2"
)

type {{.EntityTitle}}HTTPTransport struct {
	service {{.EntityTitle}}Service
}

func New{{.EntityTitle}}HTTPTransport(service {{.EntityTitle}}Service) *{{.EntityTitle}}HTTPTransport {
	return &{{.EntityTitle}}HTTPTransport{service: service}
}

func (p *{{.EntityTitle}}HTTPTransport) RegisterRoutes(router fiber.Router) {
	router.Get("/", p.Index)        // List all {{.EntityName}}s
	router.Post("/", p.Create)      // Create a new {{.EntityName}}
	router.Get("/:id", p.Read)      // Get a {{.EntityName}} by ID
	router.Put("/:id", p.Update)    // Update a {{.EntityName}} by ID
	router.Delete("/:id", p.Delete) // Delete a {{.EntityName}} by ID
}

func (p *{{.EntityTitle}}HTTPTransport) Index(c *fiber.Ctx) error {
	resp, err := p.service.Index()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to list {{.EntityName}}s"})
	}
	return c.JSON(resp)
}

func (p *{{.EntityTitle}}HTTPTransport) Create(c *fiber.Ctx) error {
	var req Create{{.EntityTitle}}Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	resp, err := p.service.Create(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create {{.EntityName}}"})
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (p *{{.EntityTitle}}HTTPTransport) Read(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	req := Read{{.EntityTitle}}Request{ID: id}
	resp, err := p.service.Read(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "{{.EntityName}} not found"})
	}
	return c.JSON(resp)
}

func (p *{{.EntityTitle}}HTTPTransport) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	var req Update{{.EntityTitle}}Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	req.ID = id
	resp, err := p.service.Update(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update {{.EntityName}}"})
	}
	return c.JSON(resp)
}

func (p *{{.EntityTitle}}HTTPTransport) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
	}
	req := Delete{{.EntityTitle}}Request{ID: id}
	resp, err := p.service.Delete(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete {{.EntityName}}"})
	}
	return c.JSON(resp)
}
