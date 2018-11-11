package actions

import (
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/orest/briefer/analyzers"
	"github.com/orest/briefer/rabbitmq"
)

// BriefHandler is a default handler to serve up
// a brief page.
func BriefHandler(c buffalo.Context) error {
	text := c.Param("text")
	title := c.Param("title")
	analyzer := analyzers.TextAnalyzer{}
	brief := strings.Join(analyzer.Brief(title, text), ", ")
	rabbit := rabbitmq.InitClient()
	rabbit.PublishAsync("projects.brief", brief)
	return c.Render(200, r.JSON(map[string]string{"title": title, "original": text, "brief": brief}))
}
