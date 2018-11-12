package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/orest/briefer/rabbitmq"
)

// BriefHandler is a default handler to serve up
// a brief page.
func BriefHandler(c buffalo.Context) error {
	id := c.Param("id")
	description := c.Param("description")
	title := c.Param("title")
	producer := rabbitmq.ProjectProducer{Id: id, Title: title, Description: description}
	producer.Push()
	return c.Render(200, r.JSON(map[string]string{"id": id, "title": title, "description": description}))
}
