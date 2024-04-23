package {{.ModuleName}}

import (
    "github.com/gofiber/fiber/v2"
)

// Register{{.EntityTitle}}Routes sets up the routing for {{.EntityName}} management.
func Register{{.EntityTitle}}Routes(router fiber.Router, transport *{{.EntityTitle}}HTTPTransport) {
    router.Get("/", transport.Index)       // Handles GET requests to list {{.EntityName}}s
    router.Post("/", transport.Create)     // Handles POST requests to create a new {{.EntityName}}
    router.Get("/:id", transport.Read)     // Handles GET requests to read a {{.EntityName}} by ID
    router.Put("/:id", transport.Update)   // Handles PUT requests to update a {{.EntityName}} by ID
    router.Delete("/:id", transport.Delete) // Handles DELETE requests to delete a {{.EntityName}} by ID
}
