package {{.ModuleName}}

import "time"

type {{.EntityTitle}}Service interface {
	Index() ([]{{.EntityTitle}}, error)
	Create(req Create{{.EntityTitle}}Request) (Create{{.EntityTitle}}Response, error)
	Read(req Read{{.EntityTitle}}Request) (Read{{.EntityTitle}}Response, error)
	List() ([]{{.EntityTitle}}, error)
	Update(req Update{{.EntityTitle}}Request) (Update{{.EntityTitle}}Response, error)
	Delete(req Delete{{.EntityTitle}}Request) (Delete{{.EntityTitle}}Response, error)
}

type Simple{{.EntityTitle}}Service struct {
	{{.EntityName}}s  []{{.EntityTitle}}
	lastID int
}

func NewSimple{{.EntityTitle}}Service() *Simple{{.EntityTitle}}Service {
	return &Simple{{.EntityTitle}}Service{}
}

// @Summary      	Index
// @Description		Lists all {{.EntityName}}s
// @Tags			{{.EntityTitle}}s
// @Accept			json
// @Produce			json
// @Param			Authorization					header		string			true	"Authorization Key(e.g Bearer key)"
// @Success			200								{array}	List{{.EntityTitle}}Request
// @Router			/{{.EntityName}}s	[GET]
func (s *Simple{{.EntityTitle}}Service) Index() ([]{{.EntityTitle}}, error) {
	return s.{{.EntityName}}s, nil
}

// @Summary      	Create
// @Description		Validates user id and title. If they are up to standard a new {{.EntityName}} will be created. The created {{.EntityName}}s ID will be returned.
// @Tags			{{.EntityTitle}}s
// @Accept			json
// @Produce			json
// @Param			Authorization					header		string			true	"Authorization Key(e.g Bearer key)"
// @Param			Create{{.EntityTitle}}Request					body		Create{{.EntityTitle}}Request	true	"Create{{.EntityTitle}}Request"
// @Success			200								{object}	Create{{.EntityTitle}}Response
// @Router			/{{.EntityName}}s	[POST]
func (s *Simple{{.EntityTitle}}Service) Create(req Create{{.EntityTitle}}Request) (Create{{.EntityTitle}}Response, error) {
	s.lastID++
	new{{.EntityTitle}} := {{.EntityTitle}}{
		ID:        s.lastID,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    "draft",
		AuthorID:  req.AuthorID,
	}
	s.{{.EntityName}}s = append(s.{{.EntityName}}s, new{{.EntityTitle}})
	return Create{{.EntityTitle}}Response{ID: new{{.EntityTitle}}.ID, Status: "created"}, nil
}

// @Summary      	Read
// @Description		Fetches a {{.EntityName}} by ID
// @Tags			{{.EntityTitle}}s
// @Accept			json
// @Produce			json
// @Param			Authorization					header		string			true	"Authorization Key(e.g Bearer key)"
// @Param			id								path		int				true	"{{.EntityTitle}} ID"
// @Success			200								{array}	Read{{.EntityTitle}}Request
// @Router			/{{.EntityName}}s/{id}	[GET]
func (s *Simple{{.EntityTitle}}Service) Read(req Read{{.EntityTitle}}Request) (Read{{.EntityTitle}}Response, error) {
	for _, {{.EntityName}} := range s.{{.EntityName}}s {
		if {{.EntityName}}.ID == req.ID {
			return Read{{.EntityTitle}}Response{ {{.EntityTitle}}: {{.EntityName}} }, nil
		}
	}
	return Read{{.EntityTitle}}Response{}, nil
}

// @Summary      	Update
// @Description		Updates a {{.EntityName}}
// @Tags			{{.EntityTitle}}s
// @Accept			json
// @Produce			json
// @Param			Authorization					header		string			true	"Authorization Key(e.g Bearer key)"
// @Param			id								path			int					true	"{{.EntityTitle}} ID"
// @Param			Update{{.EntityTitle}}Request				body			Update{{.EntityTitle}}Request	true	"Update{{.EntityTitle}}Request"
// @Success			200								{object}		Update{{.EntityTitle}}Response
// @Router			/{{.EntityName}}s/{id}	[PUT]
func (s *Simple{{.EntityTitle}}Service) Update(req Update{{.EntityTitle}}Request) (Update{{.EntityTitle}}Response, error) {
	for i, {{.EntityName}} := range s.{{.EntityName}}s {
		if {{.EntityName}}.ID == req.ID {
			s.{{.EntityName}}s[i] = {{.EntityTitle}}{
				ID:          req.ID,
				Title:       req.Title,
				Content:     req.Content,
				PublishedAt: req.PublishedAt,
				Status:      req.Status,
				AuthorID:    req.AuthorID,
			}
			return Update{{.EntityTitle}}Response{ID: req.ID, Status: "updated"}, nil
		}
	}
	return Update{{.EntityTitle}}Response{}, nil
}

// @Summary      	Delete
// @Description		Deletes a {{.EntityName}} by ID
// @Tags			{{.EntityTitle}}s
// @Accept			json
// @Produce			json
// @Param			Authorization					header		string			true	"Authorization Key(e.g Bearer key)"
// @Param			id								path		int				true	"{{.EntityTitle}} ID"
// @Success			200								{object}	Delete{{.EntityTitle}}Response
// @Router			/{{.EntityName}}s/{id}	[DELETE]
func (s *Simple{{.EntityTitle}}Service) Delete(req Delete{{.EntityTitle}}Request) (Delete{{.EntityTitle}}Response, error) {
	for i, {{.EntityName}} := range s.{{.EntityName}}s {
		if {{.EntityName}}.ID == req.ID {
			s.{{.EntityName}}s = append(s.{{.EntityName}}s[:i], s.{{.EntityName}}s[i+1:]...)
			return Delete{{.EntityTitle}}Response{ID: req.ID, Status: "deleted"}, nil
		}
	}
	return Delete{{.EntityTitle}}Response{}, nil
}
